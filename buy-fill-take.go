package main

import "fmt"

var water, milk, beans, cups, money int = 400, 540, 120, 9, 550

func printCoffee() {
	fmt.Println("The coffee machine has:")
	fmt.Println(water, "of water")
	fmt.Println(milk, "of milk")
	fmt.Println(beans, "of coffee beans")
	fmt.Println(cups, "of disposable cups")
	fmt.Printf("$%v of money", money)
}

func fillCoffee() {
	var a, b, c, d int
	fmt.Println("Write how many ml of water do you want to add:")
	fmt.Scan(&a)
	water += a
	fmt.Println("Write how many ml of milk do you want to add:")
	fmt.Scan(&b)
	milk += b
	fmt.Println("Write how many grams of coffee beans do you want to add:")
	fmt.Scan(&c)
	beans += c
	fmt.Println("Write how many disposable cups of coffee do you want to add:")
	fmt.Scan(&d)
	cups += d
	printCoffee()
}
func takeCoffee() {
	fmt.Println("I gave you $", money)
	money = 0
	printCoffee()
}

func buyCoffee() {
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu:")
	var coffee int
	fmt.Scan(&coffee)
	switch coffee {
	case 1:
		if water >= 250 && beans >= 16 && cups >= 1 {
			fmt.Println("I have enough resources, making you a coffee!")
			water -= 250
			beans -= 16
			cups -= 1
			money += 4
			printCoffee()
		} else {
			if water < 250 {
				fmt.Println("Sorry, not enough water!")
			} else if beans < 16 {
				fmt.Println("Sorry, not enough coffee beans!")
			} else if cups < 1 {
				fmt.Println("Sorry, not enough disposable cups!")
			}
		}
	case 2:
		if water >= 350 && milk >= 75 && beans >= 20 && cups >= 1 {
			fmt.Println("I have enough resources, making you a coffee!")
			water -= 350
			milk -= 75
			beans -= 20
			cups -= 1
			money += 7
			printCoffee()
		} else {
			if water < 350 {
				fmt.Println("Sorry, not enough water!")
			} else if milk < 75 {
				fmt.Println("Sorry, not enough milk!")
			} else if beans < 20 {
				fmt.Println("Sorry, not enough coffee beans!")
			} else if cups < 1 {
				fmt.Println("Sorry, not enough disposable cups!")
			}
		}
	case 3:
		if water >= 200 && milk >= 100 && beans >= 12 && cups >= 1 {
			fmt.Println("I have enough resources, making you a coffee!")
			water -= 200
			milk -= 100
			beans -= 12
			cups -= 1
			money += 6
			printCoffee()
		} else {
			if water < 200 {
				fmt.Println("Sorry, not enough water!")
			} else if milk < 100 {
				fmt.Println("Sorry, not enough milk!")
			} else if beans < 12 {
				fmt.Println("Sorry, not enough coffee beans!")
			} else if cups < 1 {
				fmt.Println("Sorry, not enough disposable cups!")
			}
		}
	}
}
