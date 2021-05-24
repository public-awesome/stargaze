package types

// Default parameter namespace
const (
	DefaultParamspace string = ModuleName
	DefaultFunder     string = "stars1czlu4tvr3dg3ksuf8zak87eafztr2u004zyh5a"
)

// Parameter store keys
var (
	KeyFunder = []byte("Funder")
)

// NewParams creates a new Params object
func NewParams(
	funder string,
) Params {
	return Params{
		Funder: funder,
	}
}

// DefaultParams defines the parameters for this module
func DefaultParams() Params {
	return NewParams(
		DefaultFunder,
	)
}
