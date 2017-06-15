package schema

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)

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

	// TimeListType is our exported const representing a value that is a comma
	// separated list of dateTime instances. Used for transport data values.
	TimeListType = DataType("http://purl.org/iot/vocab/thingful#timeList")

	// XSDTimeFormat is our time format string used for XSD time instances
	XSDTimeFormat = "15:04:05"
)

// Serialize is our helper function that serializes a value of one of the above
// types into a string serialization we return to pomelo.
func Serialize(value interface{}, dataType DataType) (string, error) {
	switch dataType {
	case IntegerType:
		v, ok := value.(int)
		if !ok {
			return "", fmt.Errorf("cannot type assert value '%v' to int", value)
		}
		return strconv.Itoa(v), nil
	case DoubleType:
		v, ok := value.(float64)
		if !ok {
			return "", fmt.Errorf("cannot type assert value '%v' to float64", value)
		}
		return strconv.FormatFloat(v, 'g', -1, 64), nil
	case DateTimeType:
		v, ok := value.(time.Time)
		if !ok {
			return "", fmt.Errorf("cannot type assert value '%v' to time.Time", value)
		}
		return v.Format(time.RFC3339Nano), nil
	case TimeType:
		v, ok := value.(time.Time)
		if !ok {
			return "", fmt.Errorf("cannot type assert value '%v' to time.Time", value)
		}
		return v.Format(XSDTimeFormat), nil
	case BytesType:
		v, ok := value.([]byte)
		if !ok {
			return "", fmt.Errorf("cannot type assert value '%v' to []byte", value)
		}
		return string(v), nil
	case TimeListType:
		v, ok := value.([]time.Time)
		if !ok {
			return "", fmt.Errorf("cannot type assert value '%v' to []time.Time", value)
		}

		var buf bytes.Buffer
		for i, t := range v {
			buf.WriteString(t.Format(XSDTimeFormat))
			if i < (len(v) - 1) {
				buf.WriteString(",")
			}
		}

		return buf.String(), nil
	default:
		v, ok := value.(string)
		if !ok {
			return "", fmt.Errorf("cannot type assert value '%v' to string", value)
		}
		return v, nil
	}
}
