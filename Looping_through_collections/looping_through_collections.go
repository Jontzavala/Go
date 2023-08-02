package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Please select an option")
	fmt.Println("1) Print menu")
	in :=bufio.NewReader(os.Stdin)
	choice, _ := in.ReadString('\n')
	choice = strings.TrimSpace(choice) // we don't know what to do with this yet.

	type menuItem struct {
		name string
		prices map[string]float64
	}

	menu := []menuItem{
		{name: "Coffee", prices: map[string]float64{"small": 1.65, "medium": 1.80, "large": 1.95}},
		{name: "Espresso", prices: map[string]float64{"small": 1.90, "medium": 2.25, "large": 2.55}},
	}

	for _, item := range menu {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, price := range item.prices {
			/* "\t" is a tab character, "%10s" is a formatting verb that prints out a string using 10 characters
			"%10.2f" is also a formatting verb, it prints out with a 10 character column and two decimal percisions and
			f means it's expecting a float */
			fmt.Printf("\t%10s%10.2f\n", size, price)
		}
	}
}