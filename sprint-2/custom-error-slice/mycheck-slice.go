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

	for i := 0; i < len(M); i++ {

		if M[i] != nil {
			s = append(s, M[i].Error())
		}
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

	result := make(myError,18)
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

	// проверка наличия цифр
	if numbers > 0 {
		result[3] = fmt.Errorf("1found numbers")
	}

	// проверка общей длины строки
	if len(runes) >= 20 {
		result[4] = fmt.Errorf("1line is too long")
	}

	// проверка наличия двух пробелов
	if spaces != 2 {
		result[5] = fmt.Errorf("1no two spaces")
	}

	// проверка наличия цифр
	if numbers > 0 {
		result[6] = fmt.Errorf("2found numbers")
	}

	// проверка общей длины строки
	if len(runes) >= 20 {
		result[7] = fmt.Errorf("2line is too long")
	}

	// проверка наличия двух пробелов
	if spaces != 2 {
		result[8] = fmt.Errorf("2no two spaces")
	}

	// проверка наличия цифр
	if numbers > 0 {
		result[9] = fmt.Errorf("3found numbers")
	}

	// проверка общей длины строки
	if len(runes) >= 20 {
		result[10] = fmt.Errorf("3line is too long")
	}

	// проверка наличия двух пробелов
	if spaces != 2 {
		result[11] = fmt.Errorf("3no two spaces")
	}

	// проверка наличия цифр
	if numbers > 0 {
		result[12] = fmt.Errorf("4found numbers")
	}

	// проверка общей длины строки
	if len(runes) >= 20 {
		result[13] = fmt.Errorf("4line is too long")
	}

	// проверка наличия двух пробелов
	if spaces != 2 {
		result[14] = fmt.Errorf("4no two spaces")
	}

	// проверка наличия цифр
	if numbers > 0 {
		result[15] = fmt.Errorf("5found numbers")
	}

	// проверка общей длины строки
	if len(runes) >= 20 {
		result[16] = fmt.Errorf("5line is too long")
	}

	// проверка наличия двух пробелов
	if spaces != 2 {
		result[17] = fmt.Errorf("5no two spaces")
	}

	return result.String()
}
