# equaler
--
    import "github.com/MJKWoolnough/equaler"


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
