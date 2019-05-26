package mysql

import (
	"github.com/KoteiIto/go-xorm-test/domain/repository/db"
	"github.com/go-xorm/xorm"
)

type DB struct {
	engine xorm.Engine
}

func NewDB(uri string) (*DB, error) {
	engine, err := xorm.NewEngine("mysql", uri)
	if err != nil {
		return nil, err
	}
	engine.ShowSQL(true)
	return &DB{engine: *engine}, nil
}

func (db *DB) Transaction(f func(tx db.Session) (interface{}, error)) (interface{}, error) {
	s := db.engine.NewSession()
	err := s.Begin()
	if err != nil {
		return nil, err
	}

	v, err := f(&Session{session: *s})
	if err != nil {
		s.Rollback()
		return nil, err
	}

	return v, nil
}

type Session struct {
	session xorm.Session
}

func (sess *Session) Insert(dto db.InsertableDto) (int64, error) {
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

func (sess *Session) Update(dto db.UpdatableDto) (int64, error) {
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

func (sess *Session) Delete(dto db.DeletableDto) (int64, error) {
	affected, err := sess.session.Delete(dto.PEntity())
	if err != nil {
		return affected, err
	}
	dto.AsDeleted()

	return affected, err
}
