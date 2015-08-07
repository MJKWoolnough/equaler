// Package equaler gives a simple interface to compare objects of arbitrary types.
package equaler

// Equaler is a simple interface that is used to determing the equality of two
// variables
type Equaler interface {
	Equal(Equaler) bool
}

type eFalse struct{}

func (eFalse) Equal(Equaler) bool {
	return false
}

func (eFalse) String() string {
	return "Equaler - False"
}

type eThis struct{}

func (eThis) Equal(e Equaler) bool {
	_, ok := e.(eThis)
	return ok
}

func (eThis) String() string {
	return "Equaler - This"
}

var (
	// EFalse represents an Equaler that never matches, even against itself
	EFalse eFalse
	// EThis represents an Equaler that only and alwats matches itself
	EThis eThis
)
