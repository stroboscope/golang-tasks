package mycheck

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

// slice с перечнем ошибок
type myError []error

// метод String - получение строки из слайса с ошибками c разделителем ";"
// ошибки хранятся в предопределенном порядке: наличие чисел, несоответствие длины, наличие 2 пробелов
func (M myError) String() error {

	var s []string
	sep := ";"

	if M[0] != nil {
		s = append(s, M[0].Error())
	}

	if M[1] != nil {
		s = append(s, M[1].Error())
	}

	if M[2] != nil {
		s = append(s, M[2].Error())
	}

	j := strings.Join(s, sep)

	if j == "" {
		return nil
	}

	return errors.New(j)
}

// проверка строки на соответствие требованиям
// не должно быть чисел
// длина строки не должна превышать 20 символов
// должно быть ровно 2 пробела
func MyCheck(input string) error {

	result := myError{nil, nil, nil}
	runes := []rune(input)
	spaces := 0
	numbers := 0

	for _, i := range runes {
		// подсчет пробелов
		if string(i) == " " {
			spaces++
		}

		// проверка на то, что символ является числом
		if unicode.IsNumber(i) {
			numbers++
		}
	}

	// проверка наличия цифр
	if numbers > 0 {
		result[0] = fmt.Errorf("found numbers")
	}

	// проверка общей длины строки
	if len(runes) >= 20 {
		result[1] = fmt.Errorf("line is too long")
	}

	// проверка наличия двух пробелов
	if spaces != 2 {
		result[2] = fmt.Errorf("no two spaces")
	}

	return result.String()
}
