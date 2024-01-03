package main

import (
	"fmt"
)

const (
	espressoWater   = 250
	espressoBeans   = 16
	espressoCost    = 4
	latteWater      = 350
	latteMilk       = 75
	latteBeans      = 20
	latteCost       = 7
	cappuccinoWater = 200
	cappuccinoMilk  = 100
	cappuccinoBeans = 12
	cappuccinoCost  = 6
)

func display(water *int, milk *int, beans *int, disposableCups *int, money *int) {
	fmt.Printf("%d ml of water\n", *water)
	fmt.Printf("%d ml of milk\n", *milk)
	fmt.Printf("%d g of coffee beans\n", *beans)
	fmt.Printf("%d disposable cups\n", *disposableCups)
	fmt.Printf("$%d of money\n", *money)
}
func filler(stock *int, stockAdded *int) {
	*stock += *stockAdded
}

func fillMilk(milk *int, milkAdded *int) {
	*milk += *milkAdded
}

func fillBeans(beans *int, beansAdded *int) {
	*beans += *beansAdded
}

func fillCups(cups *int, cupsAdded *int) {
	*cups += *cupsAdded
}
func firstDrink(water *int, beans *int, disposableCups *int, money *int) {
	if *water >= espressoWater && *beans >= espressoBeans && *disposableCups >= 1 {
		*water -= espressoWater
		*beans -= espressoBeans
		*money += espressoCost
		*disposableCups--
		fmt.Println("I have enough resources, making you a coffee!")
	} else if *water < espressoWater {
		fmt.Println("Sorry, not enough water!")
	} else if *beans < espressoBeans {
		fmt.Println("Sorry, not enough beans!")
	} else if *disposableCups < 1 {
		fmt.Println("Sorry, not enough cups!")
	}
}

func secondDrink(water *int, milk *int, beans *int, disposableCups *int, money *int) {
	if *water >= latteWater && *milk >= latteMilk && *beans >= latteBeans && *disposableCups >= 1 {
		*water -= latteWater
		*milk -= latteMilk
		*beans -= latteBeans
		*money += latteCost
		*disposableCups--
		fmt.Println("I have enough resources, making you a coffee!")
	} else if *water < latteWater {
		fmt.Println("Sorry, not enough water!")
	} else if *milk < latteMilk {
		fmt.Println("Sorry, not enough milk!")
	} else if *beans < latteBeans {
		fmt.Println("Sorry, not enough beans!")
	} else if *disposableCups < 1 {
		fmt.Println("Sorry, not enough cups!")
	}
}

func thirdDrink(water *int, milk *int, beans *int, disposableCups *int, money *int) {
	if *water >= cappuccinoWater && *milk >= cappuccinoMilk && *beans >= cappuccinoBeans && *disposableCups >= 1 {
		*water -= cappuccinoWater
		*milk -= cappuccinoMilk
		*beans -= cappuccinoBeans
		*money += cappuccinoCost
		*disposableCups--
		fmt.Println("I have enough resources, making you a coffee!")
	} else if *water < cappuccinoWater {
		fmt.Println("Sorry, not enough water!")
	} else if *milk < cappuccinoMilk {
		fmt.Println("Sorry, not enough milk!")
	} else if *beans < cappuccinoBeans {
		fmt.Println("Sorry, not enough beans!")
	} else if *disposableCups < 1 {
		fmt.Println("Sorry, not enough cups!")
	}
}
func buy(water *int, milk *int, beans *int, disposableCups *int, money *int, exit *bool) {
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino:")
	var drink string
	fmt.Scan(&drink)
	if drink == "back" {
		action(water, milk, beans, disposableCups, money, exit)
	} else if drink == "1" {
		firstDrink(water, beans, disposableCups, money)
	} else if drink == "2" {
		secondDrink(water, milk, beans, disposableCups, money)
	} else if drink == "3" {
		thirdDrink(water, milk, beans, disposableCups, money)
	}
}

func fill(water *int, milk *int, beans *int, disposableCups *int, money *int) {
	fmt.Println("Write how many ml of water you want to add:")
	var waterAdded int
	fmt.Scan(&waterAdded)
	filler(water, &waterAdded)
	fmt.Println("Write how many ml of milk you want to add:")
	var milkAdded int
	fmt.Scan(&milkAdded)
	filler(milk, &milkAdded)
	fmt.Println("Write how many grams of coffee beans you want to add:")
	var beansAdded int
	fmt.Scan(&beansAdded)
	filler(beans, &beansAdded)
	fmt.Println("Write how many disposable cups you want to add:")
	var cupsAdded int
	fmt.Scan(&cupsAdded)
	filler(disposableCups, &cupsAdded)
	fmt.Println("The machine has been refilled successfully!\nHere is the updated stock:")
	display(water, milk, beans, disposableCups, money)
}
func take(water *int, milk *int, beans *int, disposableCups *int, money *int) {
	fmt.Printf("I gave you $%d\n", *money)
	fmt.Println("The machine now has 0 $")
	*money -= *money
}
func action(water *int, milk *int, beans *int, disposableCups *int, money *int, exit *bool) bool {
	fmt.Println("Write action (buy, fill, take, remaining, exit):")
	var action string
	fmt.Scan(&action)
	switch action {
	case "buy":
		buy(water, milk, beans, disposableCups, money, exit)
	case "fill":
		fill(water, milk, beans, disposableCups, money)
	case "take":
		take(water, milk, beans, disposableCups, money)
	case "remaining":
		fmt.Println("The coffee machine has:")
		display(water, milk, beans, disposableCups, money)
	case "exit":
		*exit = true
	}
	return *exit
}
func main() {
	fmt.Println("\tWelcome to your automatic barista!")
	water := 0
	milk := 0
	beans := 0
	disposableCups := 0
	money := 0
	exit := false
	for i := 0; i < 1; {
		if water < 200 && milk < 75 && beans < 12 && disposableCups < 1 {
			fmt.Println("\nThe machine is low on resources, please refill!\n")
		} else if water < 200 {
			fmt.Println("The machine is low on water, please refill!")
		} else if milk < 75 {
			fmt.Println("The machine is low on milk, please refill!")
		} else if beans < 12 {
			fmt.Println("The machine is low on beans, please refill!")
		} else if disposableCups < 1 {
			fmt.Println("The machine is low on cups, please refill!")
		}

		if exit {
			break
		}
		action(&water, &milk, &beans, &disposableCups, &money, &exit)

	}

}
