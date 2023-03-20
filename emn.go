package main

import (
	"fmt"
	"math"
)

func main() {
	var number, digitCount, sum, temp, copy int
	var remainder float64

	// Kullanıcıdan bir sayı girişi alınır
	fmt.Print("Bir sayı girin: ")
	fmt.Scan(&number)

	// Sayının basamak sayısını bulma
	//digitCount = int(math.Log10(float64(number))) + 1
	copy = number
	for copy != 0 {
		copy /= 10
		digitCount += 1
	}
	// Sayının basamaklarını toplama
	temp = number
	for temp > 0 {
		remainder = math.Mod(float64(temp), 10) //temp%10
		sum += int(math.Pow(remainder, float64(digitCount)))
		temp /= 10
		digitCount--
	}

	// Disarium sayı kontrolü
	if number == sum {
		fmt.Println(number, "This is an disarium number.")
	} else {
		fmt.Println(number, "This is not.")
	}
}
