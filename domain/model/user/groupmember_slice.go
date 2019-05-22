// Generated by: gen
// TypeWriter: slice
// Directive: +gen on GroupMember

package user

// GroupMemberSlice is a slice of type GroupMember. Use it where you would use []GroupMember.
type GroupMemberSlice []GroupMember

// Where returns a new GroupMemberSlice whose elements return true for func. See: http://clipperhouse.github.io/gen/#Where
func (rcv GroupMemberSlice) Where(fn func(GroupMember) bool) (result GroupMemberSlice) {
	for _, v := range rcv {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// GroupByInt groups elements into a map keyed by int. See: http://clipperhouse.github.io/gen/#GroupBy
func (rcv GroupMemberSlice) GroupByInt(fn func(GroupMember) int) map[int]GroupMemberSlice {
	result := make(map[int]GroupMemberSlice)
	for _, v := range rcv {
		key := fn(v)
		result[key] = append(result[key], v)
	}
	return result
}

// Any verifies that one or more elements of GroupMemberSlice return true for the passed func. See: http://clipperhouse.github.io/gen/#Any
func (rcv GroupMemberSlice) Any(fn func(GroupMember) bool) bool {
	for _, v := range rcv {
		if fn(v) {
			return true
		}
	}
	return false
}