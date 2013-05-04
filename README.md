PACKAGE

package equaler
    import "/home/michael/Programming/Go/src/github.com/MJKWoolnough/equaler"


VARIABLES

var (
    EFalse eFalse
    EThis  eThis
)


TYPES

type Equaler interface {
    Equal(Equaler) bool
}


