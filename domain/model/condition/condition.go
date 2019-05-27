package condition

import "time"

type OperatorType string

const (
	OperatorEQ  OperatorType = "="
	OperatorNEQ OperatorType = "!="
	OperatorLT  OperatorType = "<"
	OperatorLTE OperatorType = "<="
	OperatorGT  OperatorType = ">"
	OperatorGTE OperatorType = ">="
)

type Condition struct {
	Column   string
	Operator OperatorType
	Value    interface{}
}

func (c Condition) Check(val interface{}) bool {
	switch c.Operator {
	case OperatorEQ:
		return val == c.Value
	case OperatorNEQ:
		return val != c.Value
	case OperatorLT:
		switch v := c.Value.(type) {
		case int:
			return val.(int) < v
		case int64:
			return val.(int64) < v
		case int32:
			return val.(int32) < v
		case uint64:
			return val.(uint64) < v
		case uint32:
			return val.(uint32) < v
		case float64:
			return val.(float64) < v
		case float32:
			return val.(float32) < v
		case string:
			return val.(string) < v
		case time.Time:
			return val.(time.Time).UnixNano() < v.UnixNano()
		}
	case OperatorLTE:
		switch v := c.Value.(type) {
		case int:
			return val.(int) <= v
		case int64:
			return val.(int64) <= v
		case int32:
			return val.(int32) <= v
		case uint64:
			return val.(uint64) <= v
		case uint32:
			return val.(uint32) <= v
		case float64:
			return val.(float64) <= v
		case float32:
			return val.(float32) <= v
		case string:
			return val.(string) <= v
		case time.Time:
			return val.(time.Time).UnixNano() < v.UnixNano()
		}
	case OperatorGT:
		switch v := c.Value.(type) {
		case int:
			return val.(int) > v
		case int64:
			return val.(int64) > v
		case int32:
			return val.(int32) > v
		case uint64:
			return val.(uint64) > v
		case uint32:
			return val.(uint32) > v
		case float64:
			return val.(float64) > v
		case float32:
			return val.(float32) > v
		case string:
			return val.(string) > v
		case time.Time:
			return val.(time.Time).UnixNano() > v.UnixNano()
		}
	case OperatorGTE:
		switch v := c.Value.(type) {
		case int:
			return val.(int) >= v
		case int64:
			return val.(int64) >= v
		case int32:
			return val.(int32) >= v
		case uint64:
			return val.(uint64) >= v
		case uint32:
			return val.(uint32) >= v
		case float64:
			return val.(float64) >= v
		case float32:
			return val.(float32) >= v
		case string:
			return val.(string) >= v
		case time.Time:
			return val.(time.Time).UnixNano() >= v.UnixNano()
		}
	}

	return false
}
