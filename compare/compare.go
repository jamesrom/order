package compare

// SimpleOrdered is a union of all the types where < is transitive.
type SimpleOrdered interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~string
}

// SimpleOrdered is a union of all the types where < is NOT transitive. A NaN
// check can be used in conjunction to ensure transitivity. This is what
// `LessFloat` and `LessOrdered` does.
type FloatOrdered interface {
	~float32 | ~float64
}

// Ordered is the union of SimpleOrdered and FloatOrdered.
type Ordered interface {
	SimpleOrdered | FloatOrdered
}

// Less must describe a transitive ordering:
// - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
// - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
// Not that the < operator is not transitive for float types, and special care
// should be taken to order these values correctly. See 'LessFloat'.
type LessFunc[T any] func(T, T) bool

// IsNaN is a generic implementation of math.IsNaN
func IsNaN(f any) bool {
	return f != f
}

// LessSimple is the LessFunc for the < operator. For correctness, a type constraint
// called SimpleOrdered is defined for types that the < operator can be used for sorting.
// Notably, this excludes float64 and float32, for those consider using LessFloat or Less.
func LessSimple[T SimpleOrdered](a, b T) bool {
	return a < b
}

// LessFloat is the correct less function for float32 and float64 types.
func LessFloat[T FloatOrdered](a, b T) bool {
	return a < b || (IsNaN(a) && IsNaN(b))
}

// Less is the minimum 'correct' function for ordering of both Simple and Float types.
// Use this where you cannot or do not want to constrain the type parameter more than
// necessary.
func Less[T Ordered](a, b T) bool {
	return a < b || (IsNaN(a) && IsNaN(b))
}
