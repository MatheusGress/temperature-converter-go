package main

import "fmt"

func main() {
	// Kelvin to Celsius: C = K - 273
	boilingPointKelvin := 373.15
	boilingPointCelsius := boilingPointKelvin - 273.15

	fmt.Printf("O ponto de ebulição da água em Celsius é: %.2f°C\n", boilingPointCelsius)
}