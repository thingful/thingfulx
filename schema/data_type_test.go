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
			[]float64{1.1, 2.2, 3.3, 4.4, 5.5},
			DoubleListType,
			"1.1,2.2,3.3,4.4,5.5",
		},
		{
			23.0,
			DoubleType,
			"23",
		},
		{
			[]float64{1.0, 2.0, 3.0, 4.0, 5.0},
			DoubleListType,
			"1,2,3,4,5",
		},
		{
			23,
			IntegerType,
			"23",
		},
		{
			[]int{1, 2, 3, 4, 5},
			IntegerListType,
			"1,2,3,4,5",
		},
		{
			time.Date(2017, 6, 1, 10, 10, 0, 0, time.UTC),
			DateTimeType,
			"2017-06-01T10:10:00Z",
		},
		{
			[]time.Time{
				time.Date(2017, 6, 1, 10, 10, 0, 0, time.UTC),
				time.Date(2017, 6, 1, 10, 15, 0, 0, time.UTC),
				time.Date(2017, 6, 1, 10, 20, 0, 0, time.UTC),
			},
			DateTimeListType,
			"2017-06-01T10:10:00Z,2017-06-01T10:15:00Z,2017-06-01T10:20:00Z",
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
		{
			[]string{"one", "two", "three", "4"},
			StringListType,
			"one,two,three,4",
		},
		{
			[]time.Time{
				time.Date(2017, 6, 1, 10, 10, 0, 0, time.UTC),
			},
			TimeListType,
			"10:10:00",
		},
		{
			[]time.Time{
				time.Date(2017, 6, 1, 10, 10, 0, 0, time.UTC),
				time.Date(2017, 6, 1, 10, 15, 0, 0, time.UTC),
				time.Date(2017, 6, 1, 10, 20, 0, 0, time.UTC),
			},
			TimeListType,
			"10:10:00,10:15:00,10:20:00",
		},
		{
			[]time.Time{},
			TimeListType,
			"",
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
			[]float64{1.1, 2.2, 3.3, 4.4},
			IntegerListType,
			"cannot type assert value '[1.1 2.2 3.3 4.4]' to []int",
		},
		{
			"foo",
			DoubleType,
			"cannot type assert value 'foo' to float64",
		},
		{
			[]string{"one", "two"},
			DoubleListType,
			"cannot type assert value '[one two]' to []float64",
		},
		{
			"foo",
			DateTimeType,
			"cannot type assert value 'foo' to time.Time",
		},
		{
			[]string{"one", "two"},
			DateTimeListType,
			"cannot type assert value '[one two]' to []time.Time",
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
		{
			[]float64{1.1, 2.2, 3.3, 4.4},
			StringListType,
			"cannot type assert value '[1.1 2.2 3.3 4.4]' to []string",
		},
		{
			"foo",
			TimeListType,
			"cannot type assert value 'foo' to []time.Time",
		},
	}

	for _, testcase := range testcases {
		_, err := Serialize(testcase.value, testcase.dataType)
		assert.NotNil(t, err)
		assert.Equal(t, testcase.message, err.Error())
	}
}
