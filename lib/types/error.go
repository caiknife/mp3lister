package types

type Error string

func (e Error) Error() string {
	return string(e)
}
