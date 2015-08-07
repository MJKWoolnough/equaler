# equaler
--
    import "github.com/MJKWoolnough/equaler"

Package equaler gives a simple interface to compare objects of arbitrary types.

## Usage

```go
var (
	// EFalse represents an Equaler that never matches, even against itself
	EFalse eFalse
	// EThis represents an Equaler that only and alwats matches itself
	EThis eThis
)
```

#### type Equaler

```go
type Equaler interface {
	Equal(Equaler) bool
}
```

Equaler is a simple interface that is used to determing the equality of two
variables
