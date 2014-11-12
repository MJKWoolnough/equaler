// Package equaler gives a simple interface to compare objects of arbitrary types.
package equaler

// Equaler is a simple interface that is used to determing the equalit of two
// variables
type Equaler interface {
	Equal(Equaler) bool
}

type eFalse struct{}

func (f eFalse) Equal(e Equaler) bool {
	return false
}

func (f eFalse) String() string {
	return "Equaler - False"
}

type eThis struct{}

func (t eThis) Equal(e Equaler) bool {
	if _, ok := e.(*eThis); ok || e == nil {
		return true
	}
	return false
}

func (t eThis) String() string {
	return "Equaler - This"
}

var (
	// EFalse represents an Equaler that never matches, even against itself
	EFalse eFalse
	// EThis represents an Equaler that only and alwats matches itself
	EThis eThis
)
