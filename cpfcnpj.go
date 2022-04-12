package brdoc

import (
	"bytes"
	"regexp"
	"strconv"
	"unicode"
)

// Regexp pattern for CPF and CNPJ.
var (
	CPFRegexp  = regexp.MustCompile(`^\d{3}\.?\d{3}\.?\d{3}-?\d{2}$`)
	CNPJRegexp = regexp.MustCompile(`^\d{2}\.?\d{3}\.?\d{3}\/?(:?\d{3}[1-9]|\d{2}[1-9]\d|\d[1-9]\d{2}|[1-9]\d{3})-?\d{2}$`)
)

// IsCPF verifies if the given string is a valid CPF document.
func IsCPF(doc string) bool {

	const (
		size = 9
		pos  = 10
	)

	return isCPFOrCNPJ(doc, CPFRegexp, size, pos)
}

// IsCNPJ verifies if the given string is a valid CNPJ document.
func IsCNPJ(doc string) bool {

	const (
		size = 12
		pos  = 5
	)

	return isCPFOrCNPJ(doc, CNPJRegexp, size, pos)
}

// isCPFOrCNPJ generates the digits for a given CPF or CNPJ and compares it with the original digits.
func isCPFOrCNPJ(doc string, pattern *regexp.Regexp, size int, position int) bool {

	if !pattern.MatchString(doc) {
		return false
	}

	cleanNonDigits(&doc)

	// Invalidates documents with all digits equal.
	if allEq(doc) {
		return false
	}

	d := doc[:size]
	digit := calculateDigit(d, position)

	d = d + digit
	digit = calculateDigit(d, position+1)

	return doc == d+digit
}

// cleanNonDigits removes every rune that is not a digit.
func cleanNonDigits(doc *string) {

	buf := bytes.NewBufferString("")
	for _, r := range *doc {
		if unicode.IsDigit(r) {
			buf.WriteRune(r)
		}
	}

	*doc = buf.String()
}

// allEq checks if every rune in a given string is equal.
func allEq(doc string) bool {

	base := doc[0]
	for i := 1; i < len(doc); i++ {
		if base != doc[i] {
			return false
		}
	}

	return true
}

// calculateDigit calculates the next digit for the given document.
func calculateDigit(doc string, position int) string {

	var sum int
	for _, r := range doc {

		sum += toInt(r) * position
		position--

		if position < 2 {
			position = 9
		}
	}

	sum %= 11
	if sum < 2 {
		return "0"
	}

	return strconv.Itoa(11 - sum)
}

// -------------------- 

func validateCPNJ(CNPJ string) bool {

	fmt.Println(CNPJ)


	if CNPJ == "" {
		return false
	}

	if len(CNPJ) != 14 ||
	CNPJ == "00000000000000" ||
	CNPJ == "11111111111111" ||
	CNPJ == "22222222222222" ||
	CNPJ == "33333333333333" ||
	CNPJ == "44444444444444" ||
	CNPJ == "55555555555555" ||
	CNPJ == "66666666666666" ||
	CNPJ == "77777777777777" ||
	CNPJ == "88888888888888" ||
	CNPJ == "99999999999999" {
	return false
	}
	return true
}

// checagem feita apenas pelo checkCNPJ

func stringToIntSliceCNPJ(data string) (res []int) {
    for _, d := range data {
        x, err := strconv.Atoi(string(d))
        if err != nil {
            continue
        }
        res = append(res, x)
    }
    return
}

func verifyCNPJ(data []int, j int, n int) bool{
        
    soma := 0        

    for i := 0; i < n; i++ {
        v := data[i]
        soma += v * j
        
        if j == 2 {
            j = 9
        } else {
            j -= 1
        }
    }

    resto := soma % 11

    v := data[n]
    x := 0

    if resto >= 2 {
        x = 11 - resto
    }

    if  v != x {
        return false
    }

    return true
}

func checkCNPJ(data string) bool {
    return verifyCNPJ(stringToIntSliceCNPJ(data), 5, 12) && verifyCNPJ(stringToIntSliceCNPJ(data), 6, 13)
}

