# Order
Fast, generic, concurrency-safe constructs for ordering anything.

## Packages

### `compare`
The package `compare` aims to bring a consistent implementation of a `Less` function for integers, floats and strings.

#### The problem with `constraints.Ordered`

With type parameters introduced in Go 1.18, a (currently experimental) package for the standard library has been proposed called `constraints`, which defines the following type constraint:

```go
// Ordered is a constraint that permits any ordered type: any type that supports the operators < <= >= >. If future releases of Go add new ordered types, this constraint will be modified to include them.
type Ordered interface {
	Integer | Float | ~string
}
```

However, to actually order elements whose underlying type is `constraints.Ordered`, one must not rely upon the `<` operator alone. The standard library `sort` interface [describes why](https://pkg.go.dev/sort#Interface):

```go
// Note that floating-point comparison (the < operator on float32 or float64 values)
// is not a transitive ordering when not-a-number (NaN) values are involved.
// See Float64Slice.Less for a correct implementation for floating-point values.
```

This is what `compare` does: provide a consistent interface for ordering these types. To achieve this three type constraints have been defined:

- **compare.SimpleOrdered**<br />
   Use this whenever you only care ordering integers or strings. The LessFunc defined for this type is just the `<` operator.
- **compare.FloatOrdered**<br />
   Use this whenever you only care about ordering floating point types. The LessFunc defined for this type uses an `IsNaN` check to ensure a transitive ordering.
- **compare.Ordered**<br />
   The union of the two above, and also a copy of `constraints.Ordered`. The LessFunc defined for this type uses an `IsNaN` check to ensure a transitive ordering. Use this for whenever you're using `constraints.Ordered` and you don't want to use the more constrained `SimpleOrdered` or `FloatOrdered`.

Along with the associated LessFuncs (respectively):

- `LessSimple`
- `LessFloat`
- `Less`

### `priorityqueue`
See [priorityqueue README.md](priorityqueue/README.md).

### `chansort`
See [chansort README.md](chansort/README.md).
