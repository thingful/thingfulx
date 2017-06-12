package schema

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSerializeValid(t *testing.T) {
	testcases := []struct {
		value    interface{}
		dataType DataType
		expected string
	}{
		{
			23.2,
			DoubleType,
			"23.2",
		},
		{
			23.0,
			DoubleType,
			"23",
		},
		{
			23,
			IntegerType,
			"23",
		},
		{
			time.Date(2017, 6, 1, 10, 10, 0, 0, time.UTC),
			DateTimeType,
			"2017-06-01T10:10:00Z",
		},
		{
			time.Date(2017, 6, 1, 10, 10, 0, 0, time.UTC),
			TimeType,
			"10:10:00",
		},
		{
			[]byte("foo"),
			BytesType,
			"foo",
		},
		{
			"foo",
			StringType,
			"foo",
		},
	}

	for _, testcase := range testcases {
		got, err := Serialize(testcase.value, testcase.dataType)
		assert.Nil(t, err)
		assert.Equal(t, testcase.expected, got)
	}
}

func TestSerializeInvalid(t *testing.T) {
	testcases := []struct {
		value    interface{}
		dataType DataType
		message  string
	}{
		{
			"foo",
			IntegerType,
			"cannot type assert value 'foo' to int",
		},
		{
			"foo",
			DoubleType,
			"cannot type assert value 'foo' to float64",
		},
		{
			"foo",
			DateTimeType,
			"cannot type assert value 'foo' to time.Time",
		},
		{
			"foo",
			TimeType,
			"cannot type assert value 'foo' to time.Time",
		},
		{
			"foo",
			BytesType,
			"cannot type assert value 'foo' to []byte",
		},
		{
			123,
			StringType,
			"cannot type assert value '123' to string",
		},
	}

	for _, testcase := range testcases {
		_, err := Serialize(testcase.value, testcase.dataType)
		assert.NotNil(t, err)
		assert.Equal(t, testcase.message, err.Error())
	}
}
