package main

import "fmt"

func main() {
  fmt.Print("Enter the temperature in Farenheit: ")
  var degrees float64
  fmt.Scanf("%f", &degrees)
  
  degrees_in_celsius := (degrees - 32) * 5/9

  fmt.Println(degrees_in_celsius)
}
