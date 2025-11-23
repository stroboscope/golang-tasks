//go:build !solution

package varjoin


func Join(sep string, args ...string) string {

	result := ""
	s := ""

	for _, v := range args {
		result = result + s + v
		s = sep
	}
	return result
}
