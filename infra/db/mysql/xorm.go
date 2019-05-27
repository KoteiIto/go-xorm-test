package mysql

import (
	"context"

	"github.com/KoteiIto/go-xorm-test/domain/repository/db"
	"github.com/go-xorm/xorm"
)

type XormMysqlDB struct {
	engine *xorm.Engine
}

func NewXormMysqlDB(uri string) (*XormMysqlDB, error) {
	engine, err := xorm.NewEngine("mysql", uri)
	if err != nil {
		return nil, err
	}
	engine.ShowSQL(true)
	return &XormMysqlDB{engine: engine}, nil
}

func (db *XormMysqlDB) NewSession() *XormMysqlSession {
	return &XormMysqlSession{session: db.engine.NewSession()}
}

func (db *XormMysqlDB) Truncate(table string) error {
	_, err := db.engine.Exec("truncate table " + table)
	return err
}

func (db *XormMysqlDB) Transaction(f func(tx db.Session) (interface{}, error)) (interface{}, error) {
	s := db.engine.NewSession()
	err := s.Begin()
	if err != nil {
		return nil, err
	}

	v, err := f(&XormMysqlSession{session: s})
	if err != nil {
		s.Rollback()
		return nil, err
	}

	return v, nil
}

type XormMysqlSession struct {
	session *xorm.Session
}

func (sess *XormMysqlSession) Get(ctx context.Context, dto db.CrudDto, conditions ...db.Condition) (bool, error) {
	for _, condition := range conditions {
		sess.session.Where(condition.Column+" "+string(condition.Operator)+" ?", condition.Value)
	}
	e := dto.PEntity()
	has, err := sess.session.Get(e)
	if err != nil {
		return has, err
	}

	return has, nil
}

func (sess *XormMysqlSession) Insert(ctx context.Context, dto db.CrudDto) (int64, error) {
	err := dto.Validate()
	if err != nil {
		return 0, err
	}

	affected, err := sess.session.Insert(dto.PEntity())
	if err != nil {
		return affected, err
	}
	dto.AsCreated()

	return affected, err
}

func (sess *XormMysqlSession) Update(ctx context.Context, dto db.CrudDto) (int64, error) {
	err := dto.Validate()
	if err != nil {
		return 0, err
	}

	cols := dto.UpdatedColumns()
	if len(cols) == 0 {
		// 更新されたcolumnが0なのでスキップします。
		return 0, nil
	}

	sess.session.Cols(cols...)

	keys, vals := dto.PrimaryKeys(), dto.PrimaryKeyValues()
	for i := range keys {
		sess.session.Where(keys[i]+" = ?", vals[i])
	}

	e := dto.PEntity()
	affected, err := sess.session.Update(e)
	if err != nil {
		return affected, err
	}
	dto.AsUpdated()

	return affected, err
}

func (sess *XormMysqlSession) Delete(ctx context.Context, dto db.CrudDto) (int64, error) {
	affected, err := sess.session.Delete(dto.PEntity())
	if err != nil {
		return affected, err
	}
	dto.AsDeleted()

	return affected, err
}

func (sess *XormMysqlSession) Begin() error {
	return sess.session.Begin()
}

func (sess *XormMysqlSession) Commit() error {
	return sess.session.Commit()
}

func (sess *XormMysqlSession) Rollback() error {
	return sess.session.Rollback()
}
