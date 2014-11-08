// Package equaler gives a simple interface to compare objects of arbitrary types.
package equaler

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
	EFalse eFalse
	EThis  eThis
)
