package main

import "fmt"

func main() {
	score := 0

	fmt.Println("Get Ready")
	fmt.Println("Ваш счет:", score)

	fmt.Println("Вы пролетели через трубу")
	score += 1
	fmt.Println("Ваш счет:", score)

	score = 11
	quotient := score / 3
	remainder := score % 3

	fmt.Println(score, quotient, remainder)

	text := "hello"
	text += " world"
	fmt.Println(text)

	var number int = 10
	number = 33
	fmt.Println(number)

	var emptyText string
	var emptyNumber int
	var emptyBool bool

	fmt.Println("text:", emptyText)
	fmt.Println("number:", emptyNumber)
	fmt.Println("boolean:", emptyBool)

	var x = 5
	x++
	fmt.Println(x)
}
