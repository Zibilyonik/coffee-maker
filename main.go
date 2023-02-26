package main

import "fmt"

func inputAmounts() {
	var a, b, c, d int
	fmt.Println("Write how many ml of water the coffee machine has:")
	fmt.Scan(&a)
	fmt.Println("Write how many ml of milk the coffee machine has:")
	fmt.Scan(&b)
	fmt.Println("Write how many grams of coffee beans the coffee machine has:")
	fmt.Scan(&c)
	fmt.Println("Write how many cups of coffee you will need:")
	fmt.Scan(&d)
}

func coffeeCalc(water, milk, beans, cups int) int {
	var coffee int
	if water/200 > milk/50 {
		coffee = milk / 50
	} else {
		coffee = water / 200
	}
	if coffee > beans/15 {
		coffee = beans / 15
	}
	if coffee > cups {
		coffee = cups
	}
	return coffee
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
	var a int
	fmt.Println("Write how many cups of coffee you will need:")
	fmt.Scan(&a)
	ingredientCalc(a)
	coffeeInit()
}
