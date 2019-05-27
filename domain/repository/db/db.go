package db

import "context"

type DB interface {
	Transaction(f func(tx Session) (interface{}, error))
	NewSession() Session
}

type Session interface {
	Get(ctx context.Context, dto CrudDto, conditions ...Condition) (bool, error)
	Insert(ctx context.Context, dto CrudDto) (int64, error)
	Update(ctx context.Context, dto CrudDto) (int64, error)
	Delete(ctx context.Context, dto CrudDto) (int64, error)
	Begin() error
	Commit() error
	Rollback() error
}

type CrudDto interface {
	HasEntity
	HasPrimaryKey
	HasOrder
	Validate() error
	Table() string
	CacheKey() string
	IsCreated() bool
	AsCreated()
	IsUpdated() bool
	AsUpdated()
	IsDeleted() bool
	AsDeleted()
	UpdatedColumns() []string
}

type ValidatableDto interface {
	Validate() error
}

type HasEntity interface {
	Entity() interface{}
	PEntity() interface{}
	SetEntity(e interface{})
}

type HasPrimaryKey interface {
	PrimaryKeys() []string
	PrimaryKeyValues() []interface{}
}

type HasCacheKey interface {
	CacheKey() string
}

type HasTable interface {
	Table() string
}

type HasOrder interface {
	SetOrder(o int)
	Order() int
}

type ConditionOperatorType string

const (
	ConditionOperatorEQ  ConditionOperatorType = "="
	ConditionOperatorNEQ ConditionOperatorType = "!="
	ConditionOperatorLT  ConditionOperatorType = "<"
	ConditionOperatorLTE ConditionOperatorType = "<="
	ConditionOperatorGT  ConditionOperatorType = ">"
	ConditionOperatorGTE ConditionOperatorType = ">="
)

type Condition struct {
	Column   string
	Operator ConditionOperatorType
	Value    interface{}
}
