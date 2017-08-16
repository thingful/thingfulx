package schema

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSerializeValid(t *testing.T) {
	testcases := []struct {
		name     string
		value    interface{}
		dataType DataType
		expected string
	}{
		{
			"doubleOneDigit",
			23.2,
			DoubleType,
			"23.2",
		},
		{
			"doublePointZero",
			23.0,
			DoubleType,
			"23",
		},
		{
			"doubleListOneDigit",
			[]float64{1.1, 2.22, 3.333, 4.4444, 5.55555},
			DoubleListType,
			"1.1,2.22,3.333,4.4444,5.55555",
		},
		{
			"doubleListPointZero",
			[]float64{1.0, 2.00, 3.0, 4.000000, 5},
			DoubleListType,
			"1,2,3,4,5",
		},
		{
			"doubleListEmpty",
			[]float64{},
			DoubleListType,
			"",
		},
		{
			"integer",
			23,
			IntegerType,
			"23",
		},
		{
			"integerList",
			[]int{1, 2, 3, 4, 5},
			IntegerListType,
			"1,2,3,4,5",
		},
		{
			"integerListEmpty",
			[]int{},
			IntegerListType,
			"",
		},
		{
			"timeDate",
			time.Date(2017, 6, 1, 10, 10, 0, 0, time.UTC),
			DateTimeType,
			"2017-06-01T10:10:00Z",
		},
		{
			"timeDateList",
			[]time.Time{
				time.Date(2017, 6, 1, 10, 10, 0, 0, time.UTC),
				time.Date(2017, 6, 1, 10, 15, 0, 0, time.UTC),
				time.Date(2017, 6, 1, 10, 20, 0, 0, time.UTC),
			},
			DateTimeListType,
			"2017-06-01T10:10:00Z,2017-06-01T10:15:00Z,2017-06-01T10:20:00Z",
		},
		{
			"timeDateListEmpty",
			[]time.Time{},
			DateTimeListType,
			"",
		},
		{
			"timeTime",
			time.Date(2017, 6, 1, 10, 10, 0, 0, time.UTC),
			TimeType,
			"10:10:00",
		},
		{
			"byte",
			[]byte("foo"),
			BytesType,
			"foo",
		},
		{
			"string",
			"foo",
			StringType,
			"foo",
		},
		{
			"stringList",
			[]string{"one", "two", "3.33", "4"},
			StringListType,
			"one,two,3.33,4",
		},
		{
			"stringListComma",
			[]string{"1,1", "two", "three,3", "4"},
			StringListType,
			"\"1,1\",two,\"three,3\",4",
		},
		{
			"stringListEmpty",
			[]string{},
			StringListType,
			"",
		},
		{
			"timeListSingleMember",
			[]time.Time{
				time.Date(2017, 6, 1, 10, 10, 0, 0, time.UTC),
			},
			TimeListType,
			"10:10:00",
		},
		{
			"timeListMultipleMember",
			[]time.Time{
				time.Date(2017, 6, 1, 10, 10, 0, 0, time.UTC),
				time.Date(2017, 6, 1, 10, 15, 0, 0, time.UTC),
				time.Date(2017, 6, 1, 10, 20, 0, 0, time.UTC),
			},
			TimeListType,
			"10:10:00,10:15:00,10:20:00",
		},
		{
			"timeListEmpty",
			[]time.Time{},
			TimeListType,
			"",
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			got, err := Serialize(testcase.value, testcase.dataType)
			assert.Nil(t, err)
			assert.Equal(t, testcase.expected, got)
		})
	}
}

func TestSerializeInvalid(t *testing.T) {
	testcases := []struct {
		name     string
		value    interface{}
		dataType DataType
		message  string
	}{
		{
			"stringToInt",
			"foo",
			IntegerType,
			"cannot type assert value 'foo' to int",
		},
		{
			"floatArrayToIntList",
			[]float64{1.1, 2.2, 3.3, 4.4},
			IntegerListType,
			"cannot type assert value '[1.1 2.2 3.3 4.4]' to []int",
		},
		{
			"stringToDouble",
			"foo",
			DoubleType,
			"cannot type assert value 'foo' to float64",
		},
		{
			"stringArrayToDoubleList",
			[]string{"one", "two"},
			DoubleListType,
			"cannot type assert value '[one two]' to []float64",
		},
		{
			"stringToDateTime",
			"foo",
			DateTimeType,
			"cannot type assert value 'foo' to time.Time",
		},
		{
			"stringArrayToDateTimeList",
			[]string{"one", "two"},
			DateTimeListType,
			"cannot type assert value '[one two]' to []time.Time",
		},
		{
			"stringToTime",
			"foo",
			TimeType,
			"cannot type assert value 'foo' to time.Time",
		},
		{
			"stringToBytes",
			"foo",
			BytesType,
			"cannot type assert value 'foo' to []byte",
		},
		{
			"intToString",
			123,
			StringType,
			"cannot type assert value '123' to string",
		},
		{
			"floatArrayToStringList",
			[]float64{1.1, 2.2, 3.3, 4.4},
			StringListType,
			"cannot type assert value '[1.1 2.2 3.3 4.4]' to []string",
		},
		{
			"stringToTimeList",
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
