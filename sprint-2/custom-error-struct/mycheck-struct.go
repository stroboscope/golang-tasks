package mycheck

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

// slice с перечнем ошибок
type myError struct {
	numbers error
	length error
	spaces error
	numbers1 error
	length1 error
	spaces1 error
	numbers2 error
	length2 error
	spaces2 error
	numbers3 error
	length3 error
	spaces3 error
	numbers4 error
	length4 error
	spaces4 error
	numbers5 error
	length5 error
	spaces5 error
}

// метод String - получение строки из слайса с ошибками c разделителем ";"
// ошибки хранятся в предопределенном порядке: наличие чисел, несоответствие длины, наличие 2 пробелов
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
	if M.numbers1 != nil {
		s = append(s, M.numbers1.Error())
	}
	if M.length1 != nil {
		s = append(s, M.length1.Error())
	}
	if M.spaces1 != nil {
		s = append(s, M.spaces1.Error())
	}
	if M.numbers2 != nil {
		s = append(s, M.numbers2.Error())
	}
	if M.length2 != nil {
		s = append(s, M.length2.Error())
	}
	if M.spaces2 != nil {
		s = append(s, M.spaces2.Error())
	}
	if M.numbers3 != nil {
		s = append(s, M.numbers3.Error())
	}
	if M.length3 != nil {
		s = append(s, M.length3.Error())
	}
	if M.spaces3 != nil {
		s = append(s, M.spaces3.Error())
	}
	if M.numbers4 != nil {
		s = append(s, M.numbers4.Error())
	}
	if M.length4 != nil {
		s = append(s, M.length4.Error())
	}
	if M.spaces4 != nil {
		s = append(s, M.spaces4.Error())
	}
	if M.numbers5 != nil {
		s = append(s, M.numbers5.Error())
	}
	if M.length5 != nil {
		s = append(s, M.length5.Error())
	}
	if M.spaces5 != nil {
		s = append(s, M.spaces5.Error())
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

	var result myError
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
		result.numbers = fmt.Errorf("found numbers")
	}

	// проверка общей длины строки
	if len(runes) >= 20 {
		result.length = fmt.Errorf("line is too long")
	}

	// проверка наличия двух пробелов
	if spaces != 2 {
		result.spaces = fmt.Errorf("no two spaces")
	}

	// проверка наличия цифр
	if numbers > 0 {
		result.numbers1 = fmt.Errorf("1found numbers")
	}

	// проверка общей длины строки
	if len(runes) >= 20 {
		result.length1 = fmt.Errorf("1line is too long")
	}

	// проверка наличия двух пробелов
	if spaces != 2 {
		result.spaces1 = fmt.Errorf("1no two spaces")
	}

	// проверка наличия цифр
	if numbers > 0 {
		result.numbers2 = fmt.Errorf("2ound numbers")
	}

	// проверка общей длины строки
	if len(runes) >= 20 {
		result.length2 = fmt.Errorf("2line is too long")
	}

	// проверка наличия двух пробелов
	if spaces != 2 {
		result.spaces2 = fmt.Errorf("2no two spaces")
	}

	// проверка наличия цифр
	if numbers > 0 {
		result.numbers3 = fmt.Errorf("3found numbers")
	}

	// проверка общей длины строки
	if len(runes) >= 20 {
		result.length3 = fmt.Errorf("3line is too long")
	}

	// проверка наличия двух пробелов
	if spaces != 2 {
		result.spaces3 = fmt.Errorf("3no two spaces")
	}

	// проверка наличия цифр
	if numbers > 0 {
		result.numbers4 = fmt.Errorf("4found numbers")
	}

	// проверка общей длины строки
	if len(runes) >= 20 {
		result.length4 = fmt.Errorf("4line is too long")
	}

	// проверка наличия двух пробелов
	if spaces != 2 {
		result.spaces4 = fmt.Errorf("4no two spaces")
	}

	// проверка наличия цифр
	if numbers > 0 {
		result.numbers5 = fmt.Errorf("5found numbers")
	}

	// проверка общей длины строки
	if len(runes) >= 20 {
		result.length5 = fmt.Errorf("5line is too long")
	}

	// проверка наличия двух пробелов
	if spaces != 2 {
		result.spaces5 = fmt.Errorf("5no two spaces")
	}

	return result.String()
}
