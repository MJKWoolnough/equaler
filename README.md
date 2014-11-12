# equaler
--
    import "github.com/MJKWoolnough/equaler"

Package equaler gives a simple interface to compare objects of arbitrary types.

## Usage

```go
var (
	EFalse eFalse
	EThis  eThis
)
```

#### type Equaler

```go
type Equaler interface {
	Equal(Equaler) bool
}
```
