package tokens

type Servicer interface {
	Validate(header Header) error
}