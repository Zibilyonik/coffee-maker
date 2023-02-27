package main

import "fmt"

func inputAmounts() (int, int) {
	var a, b, c, d int
	fmt.Println("Write how many ml of water the coffee machine has:")
	fmt.Scan(&a)
	fmt.Println("Write how many ml of milk the coffee machine has:")
	fmt.Scan(&b)
	fmt.Println("Write how many grams of coffee beans the coffee machine has:")
	fmt.Scan(&c)
	fmt.Println("Write how many cups of coffee you will need:")
	fmt.Scan(&d)
	return coffeeCalc(a, b, c, d)
}

func coffeeCalc(water, milk, beans, cups int) (int, int) {
	var coffee int
	if water/200 > milk/50 {
		coffee = milk / 50
	} else {
		coffee = water / 200
	}
	if coffee > beans/15 {
		coffee = beans / 15
	}
	return coffee, cups
}

func coffeeInit() {
	fmt.Println("Starting to make a coffee")
	fmt.Println("Grinding coffee beans")
	fmt.Println("Boiling water")
	fmt.Println("Mixing boiled water with crushed coffee beans")
	fmt.Println("Pouring coffee into the cup")
	fmt.Println("Pouring some milk into the cup")
	fmt.Println("Coffee is ready!")
}

func ingredientCalc(a int) {
	fmt.Println("For ", a, " cups of coffee you will need:")
	fmt.Println(a*200, " ml of water")
	fmt.Println(a*50, " ml of milk")
	fmt.Println(a*15, " g of coffee beans")
}

func main() {
	// this is 1-3 stages of the project
	/*
		var coffee, a = inputAmounts()
		ingredientCalc(a)
		if coffee == a {
			fmt.Println("Yes, I can make that amount of coffee")
		} else if coffee > a {
			fmt.Println("Yes, I can make that amount of coffee (and even", coffee-a, "more than that)")
		} else {
			fmt.Println("No, I can make only", coffee, "cups of coffee")
		}
		coffeeInit()
	*/

	// this is 4-6 stages of the project'
	printCoffee()
	for {
		fmt.Println("Write action (buy, fill, take, remaining, exit):")
		var action string
		fmt.Scan(&action)
		switch action {
		case "buy":
			buyCoffee()
		case "fill":
			fillCoffee()
		case "take":
			takeCoffee()
		case "remaining":
			printCoffee()
		case "exit":
			return
		}
	}
}
