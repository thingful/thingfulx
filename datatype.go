package thingfulx

import (
	"strconv"
)

// DataType is a type alias for our constant representing the type of a
// channel's value. We use a string rather than iota type int constants because
// this value is also published to end clients.
type DataType string

const (
	// NumberType is a const used to represent a number channel value
	NumberType DataType = "number"

	// StringType is a const used to represent a string channel value
	StringType DataType = "string"

	// BooleanType is a const used to represent a boolean channel value
	BooleanType DataType = "boolean"

	// DateTimeType is a const used to represent a date/time channel value
	DateTimeType DataType = "dateTime"

	// URLType is a const used to represent a URL channel value
	URLType DataType = "url"

	// BinaryType is a const used to represent a binary channel value
	BinaryType DataType = "binary"
)

// VerifyNumberType confirms that a given string value is really a number as we
// will later claim.
func VerifyNumberType(val string) (string, error) {
	_, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return "", err
	}
	return val, nil
}

// VerifyBooleanType confirms that a given string is really a boolean type as
// we will later claim.
func VerifyBooleanType(val string) (string, error) {
	_, err := strconv.ParseBool(val)
	if err != nil {
		return "", err
	}
	return val, nil
}
