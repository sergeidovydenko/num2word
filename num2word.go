package num2word

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	zero = "ноль"

	onesFemnine = map[int]string{
		1: "одна",
		2: "две",
		3: "три",
		4: "четыре",
		5: "пять",
		6: "шесть",
		7: "семь",
		8: "восемь",
		9: "девять",
	}

	ones = map[int]string{
		1: "один",
		2: "два",
		3: "три",
		4: "четыре",
		5: "пять",
		6: "шесть",
		7: "семь",
		8: "восемь",
		9: "девять",
	}

	tens = map[int]string{
		0: "десять",
		1: "одиннадцать",
		2: "двенадцать",
		3: "тринадцать",
		4: "четырнадцать",
		5: "пятнадцать",
		6: "шестнадцать",
		7: "семнадцать",
		8: "восемнадцать",
		9: "девятнадцать",
	}

	twenties = map[int]string{
		2: "двадцать",
		3: "тридцать",
		4: "сорок",
		5: "пятьдесят",
		6: "шестьдесят",
		7: "семьдесят",
		8: "восемьдесят",
		9: "девяносто",
	}

	hundreds = map[int]string{
		1: "сто",
		2: "двести",
		3: "триста",
		4: "четыреста",
		5: "пятьсот",
		6: "шестьсот",
		7: "семьсот",
		8: "восемьсот",
		9: "девятьсот",
	}

	thousands = map[int][]string{
		1:  []string{"тысяча", "тысячи", "тысяч"},
		2:  []string{"миллион", "миллиона", "миллионов"},
		3:  []string{"миллиард", "миллиарда", "миллиардов"},
		4:  []string{"триллион", "триллиона", "триллионов"},
		5:  []string{"квадриллион", "квадриллиона", "квадриллионов"},
		6:  []string{"квинтиллион", "квинтиллиона", "квинтиллионов"},
		7:  []string{"секстиллион", "секстиллиона", "секстиллионов"},
		8:  []string{"септиллион", "септиллиона", "септиллионов"},
		9:  []string{"октиллион", "октиллиона", "октиллионов"},
		10: []string{"нониллион", "нониллиона", "нониллионов"},
	}

	currencyForms = map[string][][]string{
		"RUB": [][]string{[]string{"рубль", "рубля", "рублей"}, {"копейка", "копейки", "копеек"}},
		"EUR": [][]string{[]string{"евро", "евро", "евро"}, {"цент", "цента", "центов"}},
		"USD": [][]string{[]string{"доллар", "доллара", "долларов"}, {"цент", "цента", "центов"}},
	}

	negword = "минус"
)

//ConvertToCurrency translate float to word
func ConvertToCurrency(n float64, currency string) string {
	nString := (strconv.FormatFloat(n, 'f', 2, 64))
	splitN := strings.Split(nString, ".")
	intMaj, _ := strconv.Atoi(string(splitN[0]))
	intMin, _ := strconv.Atoi(string(splitN[1][:2]))
	mjC := currencyForms[currency][0]
	miC := currencyForms[currency][1]
	famaleCent := false
	if currency == "RUB" {
		famaleCent = true
	}
	return fmt.Sprintf(
		"%s %s %s %s",
		Convert(intMaj, false),
		morph(intMaj, mjC[0], mjC[1], mjC[2]),
		Convert(intMin, famaleCent),
		morph(intMin, miC[0], miC[1], miC[2]),
	)
}

//Convert translate number to text
func Convert(n int, feminine bool) string {
	if n < 0 {
		return fmt.Sprintf("%s %s", negword, Convert(n*-1, false))
	}
	if n == 0 {
		return zero
	}
	words := []string{}
	chunks := splitbyx(strconv.Itoa(n), 3)
	i := len(chunks)
	for _, x := range chunks {
		i--
		if x == 0 {
			continue
		}
		n1, n2, n3 := getDigits(x)
		if n3 > 0 {
			words = append(words, hundreds[n3])
		}
		if n2 > 1 {
			words = append(words, twenties[n2])
		}
		if n2 == 1 {
			words = append(words, tens[n1])
		} else if n1 > 0 {
			onesC := onesFemnine
			if i == 1 || feminine && i == 0 {
			} else {
				onesC = ones
			}
			words = append(words, onesC[n1])
		}
		if i > 0 {
			words = append(words, pluralize(x, thousands[i]))
		}
	}
	return strings.Join(words, " ")
}

func morph(n int, f1, f2, f5 string) string {
	n = n % 100
	if n > 10 && n < 20 {
		return f5
	}
	n = n % 10
	if n > 1 && n < 5 {
		return f2
	}
	if n == 1 {
		return f1
	}
	return f5
}

func pluralize(n int, forms []string) string {
	var form int
	if n%100 < 10 || n%100 > 20 {
		if n%10 == 1 {
			form = 0
		} else if 5 > n%10 && n%10 > 1 {
			form = 1
		} else {
			form = 2
		}
	} else {
		form = 2
	}
	return forms[form]
}

func splitbyx(n string, x int) []int {
	length := len(n)
	intArray := []int{}
	if length > x {
		start := length % x
		if start > 0 {
			result := n[:start]
			intN, _ := strconv.Atoi(result)
			intArray = append(intArray, intN)
		}
		for i := start; i < length; i += x {
			result := n[i : i+x]
			intN, _ := strconv.Atoi(result)
			intArray = append(intArray, intN)
		}
	} else {
		intN, _ := strconv.Atoi(n)
		intArray = append(intArray, intN)
	}
	return intArray
}

func getDigits(n int) (int, int, int) {
	strArray := []rune(strconv.Itoa(n))
	intArray := make([]int, 3)
	for i, v := range reverse(strArray) {
		intN, _ := strconv.Atoi(string(v))
		intArray[i] = intN
	}
	return intArray[0], intArray[1], intArray[2]
}

func reverse(numbers []rune) []rune {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}
