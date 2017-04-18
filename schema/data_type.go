package schema

// DataType is a type alias used for describing the type of an observation's
// value
type DataType string

var (
	// IntegerType is our exported const representing integer data types
	IntegerType = DataType(Expand("xsd:integer"))

	// DoubleType is our exported const representing double or float data types
	DoubleType = DataType(Expand("xsd:double"))

	// DateTimeType is our exported const representing a date/time values
	DateTimeType = DataType(Expand("xsd:dateTime"))

	// TimeType is our exported const representing a value that is a time
	TimeType = DataType(Expand("xsd:time"))

	// StringType is our exported const representing string values
	StringType = DataType(Expand("xsd:string"))

	// BytesType is our exported const representing an opaque slice of bytes
	BytesType = DataType(Expand("thingful:bytes"))
)
