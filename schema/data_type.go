package schema

// DataType is a type alias used for describing the type of an observation's
// value
type DataType string

var (
	// IntegerType is our exported const representing integer data types
	IntegerType = DataType(Expand("xsd:integer"))

	// DoubleType is our exported const representing double or float data types
	DoubleType = DataType(Expand("xsd:double"))

	// DateTimeType is our exported const representing date/time values in
	// RFC3339
	DateTimeType = DataType(Expand("xsd:dateTime"))

	// StringType is our exported const representing string values
	StringType = DataType(Expand("xsd:string"))
)
