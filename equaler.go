// Package equaler gives a simple interface to compare objects of arbitrary types.
package equaler

// Equaler is a simple interface that is used to determing the equality of two
// variables
type Equaler interface {
	Equal(interface{}) bool
}

type eFalse struct{}

func (eFalse) Equal(interface{}) bool {
	return false
}

func (eFalse) String() string {
	return "Equaler - False"
}

type eThis struct{}

func (eThis) Equal(e interface{}) bool {
	_, ok := e.(eThis)
	return ok
}

func (eThis) String() string {
	return "Equaler - This"
}

var (
	// EFalse represents an Equaler that never matches, even against itself
	EFalse Equaler = eFalse{}
	// EThis represents an Equaler that only and always matches itself
	EThis Equaler = eThis{}
)
