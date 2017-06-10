package schema

// DataType is a type alias used for describing the type of an observation's
// value
type DataType string

const (
	// IntegerType is our exported const representing integer data types
	IntegerType = DataType("http://www.w3.org/2001/XMLSchema#integer")

	// DoubleType is our exported const representing double or float data types
	DoubleType = DataType("http://www.w3.org/2001/XMLSchema#double")

	// DateTimeType is our exported const representing a date/time values
	DateTimeType = DataType("http://www.w3.org/2001/XMLSchema#dateTime")

	// TimeType is our exported const representing a value that is a time
	TimeType = DataType("http://www.w3.org/2001/XMLSchema#time")

	// StringType is our exported const representing string values
	StringType = DataType("http://www.w3.org/2001/XMLSchema#string")

	// BytesType is our exported const representing an opaque slice of bytes
	BytesType = DataType("http://www.w3.org/2001/XMLSchema#bytes")
)
