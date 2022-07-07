package main

import (
	"fmt"
)

func main() {
  currencies := map[string]float32{"JPY": 130.2, "EUR": 0.95}
  var dollarAmount  float32
  var currency string

  fmt.Println("Please enter the dollar amount")
  fmt.Scan(&dollarAmount)

if dollarAmount == 0 {
  fmt.Println("Invalid dollar amount.")
} else {
  fmt.Println("Target currency?")
  fmt.Scan(&currency)
  rate, isValid := currencies[currency]

  if isValid {
     fmt.Println("You are now in possession of", rate * dollarAmount, currency)
} else {
  fmt.Println("Currency not available")
}
} }