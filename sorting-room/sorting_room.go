package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	var output string = fmt.Sprintf("This is the number %.1f", f)
	return output
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	var describe_number_box_output string = fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
	return describe_number_box_output
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	isFnb, ok := fnb.(FancyNumber)
	if !ok {
		return 0
	}
	number_y, e := strconv.Atoi(isFnb.Value())
	if e == nil {
		return number_y
	}
	return 0
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	if ExtractFancyNumber(fnb) == 0 {
		return fmt.Sprintf("%s", "This is a fancy box containing the number 0.0")
	}
	return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(ExtractFancyNumber(fnb)))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	var default_string_value = "Return to sender"
	switch value := i.(type) {
	case int:
		return DescribeNumber(float64(value))
	case float64:
		return DescribeNumber(value)
	case NumberBox:
		return DescribeNumberBox(value)
	case FancyNumberBox:
		return DescribeFancyNumberBox(value)
	default:
		return default_string_value
	}

}
