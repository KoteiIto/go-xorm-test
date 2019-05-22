// Generated by: xorm

package user

import (
	"fmt"
	"time"
)

//go:generate gen

// User
// +gen slice:"Where,GroupBy[int],Any"
type User struct {

	// Id `json:"id" xorm:"pk autoincr BIGINT(20)"`
	Id int64 `json:"id" xorm:"pk autoincr BIGINT(20)"`
	// Email `json:"email" xorm:"not null unique VARCHAR(255)"`
	Email string `json:"email" xorm:"not null unique VARCHAR(255)"`
	// PasswordDigest `json:"password_digest" xorm:"CHAR(100)"`
	PasswordDigest string `json:"password_digest" xorm:"CHAR(100)"`
	// CreatedAt `json:"created_at" xorm:"not null created DATETIME"`
	CreatedAt time.Time `json:"created_at" xorm:"not null created DATETIME"`
	// UpdatedAt `json:"updated_at" xorm:"not null updated DATETIME"`
	UpdatedAt time.Time `json:"updated_at" xorm:"not null updated DATETIME"`
	// DeletedAt `json:"deleted_at" xorm:"deleted DATETIME"`
	DeletedAt time.Time `json:"deleted_at" xorm:"deleted DATETIME"`

	isCreated bool `json:"-" xorm:"-"`
	isUpdated bool `json:"-" xorm:"-"`
	isDeleted bool `json:"-" xorm:"-"`
}

var (
	_UserTableName   = "user"
	_UserColumnNames = [6]string{"id", "email", "password_digest", "created_at", "updated_at", "deleted_at"}
)

// Table テーブル名を返却します
func (m User) Table() string {
	return _UserTableName
}

// Columns カラム名のスライスを返却します
func (m User) Columns() [6]string {
	return _UserColumnNames
}

// PrimaryKeys 主キー名のスライスを返却します
func (m User) PrimaryKeys() []string {
	return []string{"id"}
}

// CacheKey PrimaryKeyを連結して、必ず一意になるKeyを返却します
func (m User) CacheKey() string {
	return fmt.Sprintf(
		"%v",
		m.Id,
	)
}

func (m User) Validate() error {

	if m.Id < -9223372036854775808 || 9223372036854775807 < m.Id {
		return fmt.Errorf("validation error. invalid column value. column=[%s], value=[%v]", "Id", m.Id)
	}

	if 255 < len(m.Email) {
		return fmt.Errorf("validation error. invalid column value. column=[%s], value=[%v]", "Email", m.Email)
	}

	if len(m.PasswordDigest) != 100 {
		return fmt.Errorf("validation error. invalid column value. column=[%s], value=[%v]", "PasswordDigest", m.PasswordDigest)
	}

	return nil
}

// IsCreated DBに存在しないレコードのモデルの場合はtrueを返却します
func (m User) IsCreated() bool {
	return m.isCreated
}

// IsUpdated DBと差分があるレコードのモデルの場合はtrueを返却します
func (m User) IsUpdated() bool {
	return m.isUpdated
}

// IsDeleted DBには存在するが削除されるレコードのモデルの場合はtrueを返却します
func (m User) IsDeleted() bool {
	return m.isDeleted
}

// AsCreated DBにInsertするレコードのモデルとして設定する
func (m *User) AsCreated() {
	if m != nil {
		m.isCreated = true
	}
}

// AsUpdated DBにUpdateするレコードのモデルとして設定する
func (m *User) AsUpdated() {
	if m != nil {
		m.isUpdated = true
	}
}

// AsDeleted DBにDeleteするレコードのモデルとして設定する
func (m *User) AsDeleted() {
	if m != nil {
		m.isDeleted = true
	}
}

// ToMap Mapに変換します
func (m User) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"Id":             m.Id,
		"Email":          m.Email,
		"PasswordDigest": m.PasswordDigest,
		"CreatedAt":      m.CreatedAt,
		"UpdatedAt":      m.UpdatedAt,
		"DeletedAt":      m.DeletedAt,
	}
}
