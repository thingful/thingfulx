package thingfulx

import (
	"strconv"
)

// DataType is a type alias for our iota constant below representing the type
// of the channel's value.
type DataType int

const (
	// NumberType is a const used to represent a number channel value
	NumberType DataType = 1 + iota
	// StringType is a const used to represent a string channel value
	StringType
	// BooleanType is a const used to represent a boolean channel value
	BooleanType
	// DateTimeType is a const used to represent a date/time channel value
	DateTimeType
	// URLType is a const used to represent a URL channel value
	URLType
	// BinaryType is a const used to represent a binary channel value
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
