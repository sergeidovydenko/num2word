num2word Convert numbers to words in Russian language
=========

num2word - Numbers to words converter in Go (Golang)

## Usage

First, import package num2word

```import github.com/sergeidovydenko/num2word```

Convert number(int) to text
```go
  male inclination
  str := num2word.Convert(115, false) // outputs "сто пятнадцать"
  ...
  str := num2word.Convert(2012, false) // outputs "две тысячи двенадцать"
  ...
  str := num2word.Convert(-1, false) // outputs "минус один"
  ...
  female inclination
  str := num2word.Convert(-1, true) // outputs "минус одна"
```

Convert number(float) to currency
```go
  str := num2word.ConvertToCurrency(514, "RUB") // outputs "один рубль одна копейка"
  ...
  str := num2word.ConvertToCurrency(123, "EUR") // outputs "один евро ноль центов"
```
