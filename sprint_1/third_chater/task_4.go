package main

import "fmt"

func main() {
	var age int
	var generation string
	const currentYear = 2025

	fmt.Print("Введите возраст: ")
	fmt.Scan(&age)

	switch birth := currentYear - age; {
	case birth >= 1946 && birth <= 1964:
		generation = "Привет бумер!"
	case birth >= 1965 && birth <= 1980:
		generation = "Привет, Х"
	case birth >= 1981 && birth <= 1996:
		generation = "Привет, Миллениал!"
	case birth >= 1997 && birth <= 2012:
		generation = "Привет, Зумер!"
	case birth >= 2013 && birth <= 2025:
		generation = "Привет, Альфа!"
	default:
		generation = "Я хз кто ты"
	}

	fmt.Printf("Тебе %d лет, %s\n", age, generation)
}
