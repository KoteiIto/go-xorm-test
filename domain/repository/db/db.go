package db

type DB interface {
	Transaction(f func(tx Session) (interface{}, error))
}

type Session interface {
	Insert(dto InsertableDto) (int64, error)
	Update(dto UpdatableDto) (int64, error)
	Delete(dto DeletableDto) (int64, error)
}

type InsertableDto interface {
	HasEntity
	ValidatableDto
	IsCreated() bool
	AsCreated()
}

type UpdatableDto interface {
	HasEntity
	HasPrimaryKey
	ValidatableDto
	IsUpdated() bool
	AsUpdated()
	UpdatedColumns() []string
}

type DeletableDto interface {
	HasEntity
	HasPrimaryKey
	IsDeleted() bool
	AsDeleted()
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

type Condition struct {
	Column   string
	Operator string
	Value    interface{}
}
