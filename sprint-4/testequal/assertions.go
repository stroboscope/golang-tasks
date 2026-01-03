//go:build !solution

package testequal

import (
	"reflect"
)

func compareTypes(a, b interface{}) bool {

	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return false
	}

	a_v := reflect.ValueOf(a)
	b_v := reflect.ValueOf(b)

	if !a_v.IsValid() || !b_v.IsValid() {
		return false
	}

	switch reflect.ValueOf(a).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64:

		if a == b {
			return true
		}

	case reflect.String:
		if a == b {
			return true
		}

	case reflect.Slice:
		if a_v.Len() != b_v.Len() {
			break
		}
		//if a_v.Cap() != b_v.Cap() {
		//	break
		//}
		return reflect.DeepEqual(a, b)

	case reflect.Map:
		if a_v.Len() != b_v.Len() {
			break
		}
		return reflect.DeepEqual(a, b)

	}

	return false
}

func formatError(t T, args ...interface{}) {

	switch len(args) {
	case 0:
		t.Errorf("")
	case 1:
		t.Errorf(args[0].(string))
	default:
		t.Errorf(args[0].(string), args[1:]...)
	}
}

// AssertEqual checks that expected and actual are equal.
//
// Marks caller function as having failed but continues execution.
//
// Returns true iff arguments are equal.
func AssertEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) bool {

	t.Helper()

	if compareTypes(expected, actual) {
		return true
	}

	formatError(t, msgAndArgs...)

	return false
}

// AssertNotEqual checks that expected and actual are not equal.
//
// Marks caller function as having failed but continues execution.
//
// Returns true iff arguments are not equal.
func AssertNotEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) bool {

	t.Helper()

	if !compareTypes(expected, actual) {
		return true
	}

	formatError(t, msgAndArgs...)

	return false
}

// RequireEqual does the same as AssertEqual but fails caller test immediately.
func RequireEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) {

	t.Helper()

	if !compareTypes(expected, actual) {
		formatError(t, msgAndArgs...)
		t.FailNow()
	}

}

// RequireNotEqual does the same as AssertNotEqual but fails caller test immediately.
func RequireNotEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) {

	t.Helper()

	if compareTypes(expected, actual) {
		formatError(t, msgAndArgs...)

		t.FailNow()
	}

}
