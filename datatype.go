package thingfulx

// DataType is a type alias for our iota constant below representing the type
// of the channel's value.
type DataType int

const (
	NumberType DataType = 1 + iota
	StringType
	BooleanType
	DateTimeType
	URLType
	BinaryType
)

// dataTypes is a slice of strings used in the Stringer representation
var dataTypes = []string{
	"number",
	"string",
	"boolean",
	"dateTime",
	"url",
	"binary",
}

// String returns a string representation of the const.
func (d DataType) String() string {
	return dataTypes[d-1]
}
