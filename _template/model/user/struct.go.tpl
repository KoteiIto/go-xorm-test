// Generated by: xorm

package {{.Models}}

{{$ilen := len .Imports}}
{{if gt $ilen 0}}
import (
	"fmt"
	{{range .Imports}}"{{.}}"{{end}}

	"github.com/KoteiIto/go-xorm-test/domain/model/condition"
)
{{end}}

{{range .Tables}}
// {{Mapper .Name}} {{.Comment}}
// +gen slice:"Where,GroupBy[int],Any"
type {{Mapper .Name}} struct {
{{$table := .}}
{{range .ColumnsSeq}}{{$col := $table.GetColumn .}}	// {{Mapper $col.Name}} {{$col.Comment}}{{Tag $table $col}}
{{Mapper $col.Name}}	{{Type $col}} {{Tag $table $col}}
{{end}}
}

type {{Mapper .Name}}Dto struct {
{{$table := .}}
entity {{Mapper .Name}}
updatedColumnMap map[string]interface{}
order int
isCreated bool
isUpdated bool
isDeleted bool
}

{{range $i,$c := $table.Columns}}
{{if ne (len $c.EnumOptions) 0}}
const (
	{{range $o, $v := $c.EnumOptions}}
	{{Mapper $table.Name}}{{Mapper $c.Name}}{{Mapper $o}} = "{{$o}}"{{end}}
)
{{end}}
{{end}}

const (
{{range $i,$e := $table.Columns}}{{Mapper $table.Name}}Column{{Mapper $e.Name}} = "{{$e.Name}}"
{{end}}
)

var (
	_{{Mapper .Name}}TableName = "{{.Name}}"
	_{{Mapper .Name}}ColumnNames = []string{ {{range $i,$e := $table.Columns}}{{if $i}},{{end}}"{{$e.Name}}"{{end}} }
	_{{Mapper .Name}}PrimaryKeys = []string{ {{range $i,$e := $table.PrimaryKeys}}{{if $i}},{{end}}"{{$e}}"{{end}} }
	{{range $i,$c := $table.Columns}}{{if ne (len $c.EnumOptions) 0}}_{{Mapper $table.Name}}{{Mapper $c.Name}}Enums = []string{ {{range $o, $_ := $c.EnumOptions}}"{{$o}}",{{end}} }{{end}}{{end}}
)

// New{{Mapper .Name}}Dto Dtoを返却します
func New{{Mapper .Name}}Dto(e {{Mapper .Name}}) *{{Mapper .Name}}Dto {
	return &{{Mapper .Name}}Dto {
		entity: e,
		updatedColumnMap: make(map[string]interface{}, {{len $table.Columns}}),
	}
}

// New{{Mapper .Name}}DtoEmpty 空のDtoを返却します
func New{{Mapper .Name}}DtoEmpty() *{{Mapper .Name}}Dto {
	return &{{Mapper .Name}}Dto {
		entity: {{Mapper .Name}}{},
		updatedColumnMap: make(map[string]interface{}, {{len $table.Columns}}),
	}
}

{{range $i,$e := $table.Columns}}
{{if not $e.IsPrimaryKey}}
// Set{{Mapper $e.Name}} setter for {{$e.Name}}
func (m *{{Mapper $table.Name}}Dto) Set{{Mapper $e.Name}}({{Mapper $e.Name}} {{Type $e}}) {
	if m == nil {
		return
	}
	
	if _, ok := m.updatedColumnMap["{{$e.Name}}"]; !ok {
		m.updatedColumnMap["{{$e.Name}}"] = m.entity.{{Mapper $e.Name}}
	}

	m.entity.{{Mapper $e.Name}} = {{Mapper $e.Name}}
}
{{end}}

// Get{{Mapper $e.Name}} getter for {{$e.Name}}
func (m {{Mapper $table.Name}}Dto) Get{{Mapper $e.Name}}() {{Type $e}} {
	return m.entity.{{Mapper $e.Name}}
}
{{end}}

// SetEntity テーブルのエンティティを設定します
func (m *{{Mapper .Name}}Dto) SetEntity(e interface{}) {
	m.entity = (e).({{Mapper .Name}})
}

// Entity テーブルのエンティティを返却します
func (m {{Mapper .Name}}Dto) Entity() interface{} {
	return m.entity
}

// PEntity テーブルのエンティティのポインタを返却します
func (m *{{Mapper .Name}}Dto) PEntity() interface{} {
	return &m.entity
}

// PEntityEmpty テーブルの空のエンティティのポインタを返却します
func (m {{Mapper .Name}}Dto) PEntityEmpty() interface{} {
	return &{{Mapper .Name}}{}
}

// Table テーブル名を返却します
func (m {{Mapper .Name}}Dto) Table () string {
	return _{{Mapper .Name}}TableName
}

// Columns カラム名のスライスを返却します
func (m {{Mapper .Name}}Dto) Columns () []string {
	return _{{Mapper .Name}}ColumnNames
}

// PrimaryKeys 主キー名のスライスを返却します
func (m {{Mapper .Name}}Dto) PrimaryKeys () []string {
	return _{{Mapper .Name}}PrimaryKeys
}

// CacheKey PrimaryKeyの値のスライスを返却します
func (m {{Mapper .Name}}Dto) PrimaryKeyValues () []interface{} {
	return []interface{} {
		{{range $i,$e := $table.PrimaryKeys}}{{$col := $table.GetColumn $e}}{{if $i}},{{end}}m.entity.{{Mapper $col.Name}}{{end}},
	}
}

// CacheKey PrimaryKeyを連結して、必ず一意になるKeyを返却します
func (m {{Mapper .Name}}Dto) CacheKey () string {
	return fmt.Sprintf(
		"{{range $i,$e := $table.PrimaryKeys}}{{if $i}}_{{end}}%v{{end}}", 
		{{range $i,$e := $table.PrimaryKeys}}{{$col := $table.GetColumn $e}}{{if $i}},{{end}}m.entity.{{Mapper $col.Name}}{{end}},
	)
}

func (m {{Mapper .Name}}Dto) UpdatedColumns() []string {
	cols := make([]string, 0, len(m.updatedColumnMap))
	for col, val := range m.updatedColumnMap {
		if val != m.Value(col) {
			cols = append(cols, col)
		}
	}
	return cols
}

func (m {{Mapper .Name}}Dto) Validate() error {
	{{range $i,$e := $table.Columns}}
		{{if and (eq $e.SQLType.Name "BIGINT") }}
			if m.entity.{{Mapper $e.Name}} <  -9223372036854775808 || 9223372036854775807 < m.entity.{{Mapper $e.Name}} {
				return fmt.Errorf("validation error. invalid column value. column=[%s], value=[%v]", "{{Mapper $e.Name}}", m.entity.{{Mapper $e.Name}})
			}
		{{end}}

		{{if and (eq $e.SQLType.Name "INT") (eq $e.SQLType.DefaultLength 11) }}
			if m.entity.{{Mapper $e.Name}} <  -2147483648 || 2147483647 < m.entity.{{Mapper $e.Name}} {
				return fmt.Errorf("validation error. invalid column value. column=[%s], value=[%v]", "{{Mapper $e.Name}}", m.entity.{{Mapper $e.Name}})
			}
		{{end}}

		{{if and (eq $e.SQLType.Name "INT") (eq $e.SQLType.DefaultLength 10) }}
			if m.entity.{{Mapper $e.Name}} <  0 || 4294967295 < m.entity.{{Mapper $e.Name}} {
				return fmt.Errorf("validation error. invalid column value. column=[%s], value=[%v]", "{{Mapper $e.Name}}", m.entity.{{Mapper $e.Name}})
			}
		{{end}}

		{{if and (eq $e.SQLType.Name "MEDIUMINT") (eq $e.SQLType.DefaultLength 9) }}
			if m.entity.{{Mapper $e.Name}} <  -8388608 || 8388607 < m.entity.{{Mapper $e.Name}} {
				return fmt.Errorf("validation error. invalid column value. column=[%s], value=[%v]", "{{Mapper $e.Name}}", m.entity.{{Mapper $e.Name}})
			}
		{{end}}

		{{if and (eq $e.SQLType.Name "MEDIUMINT") (eq $e.SQLType.DefaultLength 8) }}
			if m.entity.{{Mapper $e.Name}} <  0 || 16777215 < m.entity.{{Mapper $e.Name}} {
				return fmt.Errorf("validation error. invalid column value. column=[%s], value=[%v]", "{{Mapper $e.Name}}", m.entity.{{Mapper $e.Name}})
			}
		{{end}}

		{{if and (eq $e.SQLType.Name "SMALLINT") (eq $e.SQLType.DefaultLength 7) }}
			if m.entity.{{Mapper $e.Name}} <  -32768 || 32767 < m.entity.{{Mapper $e.Name}} {
				return fmt.Errorf("validation error. invalid column value. column=[%s], value=[%v]", "{{Mapper $e.Name}}", m.entity.{{Mapper $e.Name}})
			}
		{{end}}

		{{if and (eq $e.SQLType.Name "SMALLINT") (eq $e.SQLType.DefaultLength 6) }}
			if m.entity.{{Mapper $e.Name}} <  0 || 65535 < m.entity.{{Mapper $e.Name}} {
				return fmt.Errorf("validation error. invalid column value. column=[%s], value=[%v]", "{{Mapper $e.Name}}", m.entity.{{Mapper $e.Name}})
			}
		{{end}}

		{{if and (eq $e.SQLType.Name "TINYINT") (eq $e.SQLType.DefaultLength 4) }}
			if m.entity.{{Mapper $e.Name}} <  -128 || 127 < m.entity.{{Mapper $e.Name}} {
				return fmt.Errorf("validation error. invalid column value. column=[%s], value=[%v]", "{{Mapper $e.Name}}", m.entity.{{Mapper $e.Name}})
			}
		{{end}}

		{{if and (eq $e.SQLType.Name "TINYINT") (eq $e.SQLType.DefaultLength 3) }}
			if m.entity.{{Mapper $e.Name}} <  0 || 255 < m.entity.{{Mapper $e.Name}} {
				return fmt.Errorf("validation error. invalid column value. column=[%s], value=[%v]", "{{Mapper $e.Name}}", m.entity.{{Mapper $e.Name}})
			}
		{{end}}

		{{if and (eq $e.SQLType.Name "VARCHAR") }}
			if {{$e.Length }} < len(m.entity.{{Mapper $e.Name}}) {
				return fmt.Errorf("validation error. invalid column value. column=[%s], value=[%v]", "{{Mapper $e.Name}}", m.entity.{{Mapper $e.Name}})
			}
		{{end}}

		{{if and (eq $e.SQLType.Name "CHAR") }}
			if len(m.entity.{{Mapper $e.Name}}) != {{$e.Length }} {
				return fmt.Errorf("validation error. invalid column value. column=[%s], value=[%v]", "{{Mapper $e.Name}}", m.entity.{{Mapper $e.Name}})
			}
		{{end}}

		{{if and (eq $e.SQLType.Name "ENUM") }}
			ok := false
			for _, v := range _{{Mapper $table.Name}}{{Mapper $e.Name}}Enums {
				if m.entity.{{Mapper $e.Name}} == v {
					ok = true
					break
				}
			}
			if !ok {
				return fmt.Errorf("validation error. invalid column value. column=[%s], value=[%v]", "{{Mapper $e.Name}}", m.entity.{{Mapper $e.Name}})
			}
		{{end}}

	{{end}}
	return nil
}

// IsCreated DBに存在しないレコードのモデルの場合はtrueを返却します
func (m {{Mapper .Name}}Dto) IsCreated () bool {
	return m.isCreated
}

// IsUpdated DBと差分があるレコードのモデルの場合はtrueを返却します
func (m {{Mapper .Name}}Dto) IsUpdated () bool {
	return m.isUpdated
}

// IsDeleted DBには存在するが削除されるレコードのモデルの場合はtrueを返却します
func (m {{Mapper .Name}}Dto) IsDeleted () bool {
	return m.isDeleted
}

// AsCreated DBにInsertするレコードのモデルとして設定する
func (m *{{Mapper .Name}}Dto) AsCreated () {
	if m != nil {
		m.isCreated = true
		m.isUpdated = false
		m.isDeleted = false
	}	
}

// AsUpdated DBにUpdateするレコードのモデルとして設定する
func (m *{{Mapper .Name}}Dto) AsUpdated () {
	if m != nil {
		m.isCreated = false
		m.isUpdated = true
		m.isDeleted = false
	}
}

// AsDeleted DBにDeleteするレコードのモデルとして設定する
func (m *{{Mapper .Name}}Dto) AsDeleted () {
	if m != nil {
		m.isCreated = false
		m.isUpdated = false
		m.isDeleted = true
	}
}

// Value カラム名の値を返却します
func (m {{Mapper .Name}}Dto) Value (col string) interface{} {
	switch col {
		{{range $i,$e := $table.Columns}}case {{Mapper $table.Name}}Column{{Mapper $e.Name}}: 
			return m.entity.{{Mapper $e.Name}}
		{{end}}
	}
	return nil
}

// ToMap Mapに変換します
func (m {{Mapper .Name}}Dto) ToMap () map[string]interface{} {
	return map[string]interface{}{ 
		{{range $i,$e := $table.Columns}}"{{Mapper $e.Name}}": m.entity.{{Mapper $e.Name}},
		{{end}}
	}
}

// SetOrder Dtoの更新順序を設定する
func (m *{{Mapper .Name}}Dto) SetOrder (o int) {
	m.order = o
}

// Order Dtoの更新順序を返却する
func (m {{Mapper .Name}}Dto) Order () int {
	return m.order
}

{{range $i,$e := $table.Columns}}
func Gen{{Mapper $table.Name}}{{Mapper $e.Name}}Condition (operator condition.OperatorType, val {{Type $e}}) condition.Condition {
	return condition.Condition {
		Column: {{Mapper $table.Name}}Column{{Mapper $e.Name}},
		Operator: operator,
		Value: val,
	}
}
{{end}}

{{end}}