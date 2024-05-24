// Package leap contains functionality for determining whether a year is a leap year.
package leap

// IsLeapYear takes an int and returns a bool indicating whether the int represents a leap year.
func IsLeapYear(year int) bool {
	return year%4 == 0 && (!(year%100 == 0) || year%400 == 0)
}
