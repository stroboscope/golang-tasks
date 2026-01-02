//go:build !solution

package testequal

import (
	"fmt"
)

// AssertEqual checks that expected and actual are equal.
//
// Marks caller function as having failed but continues execution.
//
// Returns true iff arguments are equal.
// func AssertEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
func AssertEqual[T comparable](t any, expected, actual T, msgAndArgs ...any) bool {

	if expected == actual {
		return true
	}

	e := ""
	if len(msgAndArgs) > 1 {
		e = fmt.Sprintf(msgAndArgs[0].(string), msgAndArgs[1:]...)
		panic(e)
	} else {
		e = fmt.Sprint(msgAndArgs[0].(string))
	}
	fmt.Println(e)
	return false
}

// AssertNotEqual checks that expected and actual are not equal.
//
// Marks caller function as having failed but continues execution.
//
// Returns true iff arguments are not equal.
// func AssertNotEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
func AssertNotEqual[T comparable](t any, expected, actual T, msgAndArgs ...any) bool {

	if expected != actual {
		return true
	}

	e := ""
	if len(msgAndArgs) > 1 {
		e = fmt.Sprintf(msgAndArgs[0].(string), msgAndArgs[1:]...)
		panic(e)
	} else {
		e = fmt.Sprint(msgAndArgs[0].(string))
	}
	fmt.Println(e)
	return false
}

// RequireEqual does the same as AssertEqual but fails caller test immediately.
// func RequireEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) {
func RequireEqual[T comparable](t any, expected, actual T, msgAndArgs ...any) bool {

	if expected == actual {
		return true
	}

	e := ""
	if len(msgAndArgs) > 1 {
		e = fmt.Sprintf(msgAndArgs[0].(string), msgAndArgs[1:]...)
		panic(e)
	} else {
		e = fmt.Sprint(msgAndArgs[0].(string))
	}
	panic(e)
}

// RequireNotEqual does the same as AssertNotEqual but fails caller test immediately.
// func RequireNotEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) {
func RequireNotEqual[T comparable](t any, expected, actual T, msgAndArgs ...any) bool {

	if expected != actual {
		return true
	}

	e := ""
	if len(msgAndArgs) > 1 {
		e = fmt.Sprintf(msgAndArgs[0].(string), msgAndArgs[1:]...)
		panic(e)
	} else {
		e = fmt.Sprint(msgAndArgs[0].(string))
	}
	panic(e)
}
