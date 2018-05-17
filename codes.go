package errors

/*
Code defines an error code type.
*/
type Code int

/*
Metadata contains metadata that can be associated with an error code.
*/
type Metadata struct {
	Ext        string
	Int        string
	HTTPStatus int
}

/*
Codes contains a map of error codes to metadata
*/
var Codes = map[Code]Metadata{}

func init() {
	Codes[ErrUnspecified] = Metadata{"An unknown error occurred", "Error code unspecified", 500}
	Codes[ErrUnknown] = Metadata{"An unknown error occurred", "An unknown error occurred", 500}
}

const (
	// ErrUnspecified - 0: The error code was unspecified
	ErrUnspecified Code = iota
	// ErrUnknown - 1: An unknown error occurred.
	ErrUnknown
)