package mycheck

import (
	"errors"
	"fmt"
	"unicode"
)

// структура с перечнем ошибок
type myError struct {
	numbers error
	spaces  error
	length  error
}

// метод Error - получение строки из структуры с ошибками
// в предопределенном порядке: числа, длина, пробелы разделителем ";"
func (M myError) String() error {

	var s []string
	sep := ";"

	if M.numbers != nil {
		s = append(s, M.numbers.Error())
	}

	if M.length != nil {
		s = append(s, M.length.Error())
	}

	if M.spaces != nil {
		s = append(s, M.spaces.Error())
	}

	result := ""
	j := ""

	for _, v := range s {
		result = result + j + v
		j = sep
	}

	if j == "" {
		return nil
	}

	return errors.New(result)
}

// проверка строки на соотвествие требованиям
// не должно быть чисел
// длина строки не должна превышать 20 символов
// должно быть ровно 2 пробела
func MyCheck(input string) error {

	var result myError
	r := []rune(input)
	s := 0
	n := 0

	for _, i := range r {
		// подсчет пробелов
		if string(i) == " " {
			s++
		}

		// проверка на то, что символ является числом
		if unicode.IsNumber(i) {
			n++
		}
	}

	// проверка наличия цифр
	if n > 0 {
		result.numbers = fmt.Errorf("found numbers")
	}

	// проверка наличия двух пробелов
	if s != 2 {
		result.spaces = fmt.Errorf("no two spaces")
	}

	// проверка общей длины строки
	if len(r) >= 20 {
		result.length = fmt.Errorf("line is too long")
	}

	return result.String()
}
