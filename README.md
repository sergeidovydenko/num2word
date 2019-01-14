num2words Convert numbers to words in Russian language
=========

num2words - Numbers to words converter in Go (Golang)

## Usage

First, import package num2words

```import github.com/sergeidovydenko/num2words```

Convert number(int) to text
```go
  male inclination
  str := num2words.Convert(115, false) // outputs "сто пятнадцать"
  ...
  str := num2words.Convert(2012, false) // outputs "две тысячи двенадцать"
  ...
  str := num2words.Convert(-1, false) // outputs "минус один"
  ...
  female inclination
  str := num2words.Convert(-1, true) // outputs "минус одна"
```

Convert number(float) to currency
```go
  str := num2words.ConvertToCurrency(514, "RUB") // outputs "один рубль одна копейка"
  ...
  str := num2words.ConvertToCurrency(123, "EUR") // outputs "один евро ноль центов"
```
