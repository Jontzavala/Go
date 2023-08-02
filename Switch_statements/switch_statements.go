package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	type menuItem struct {
		name   string
		prices map[string]float64
	}

	menu := []menuItem{
		{name: "Coffee", prices: map[string]float64{"small": 1.65, "medium": 1.80, "large": 1.95}},
		{name: "Espresso", prices: map[string]float64{"small": 1.90, "medium": 2.25, "large": 2.55}},
	}

	in := bufio.NewReader(os.Stdin)

loop:

	for {
		fmt.Println("Please select an option")
		fmt.Println("1) Print menu")
		fmt.Println("2) Add item")
		fmt.Println("q) quit")
		choice, _ := in.ReadString('\n')

		switch strings.TrimSpace(choice) {
		case "1":
			for _, item := range menu {
				fmt.Println(item.name)
				fmt.Println(strings.Repeat("-", 10))
				for size, cost := range item.prices {
					/* "\t" is a tab character, "%10s" is a formatting verb that prints out a string using 10 characters
					"%10.2f" is also a formatting verb, it prints out with a 10 character column and two decimal percisions and
					f means it's expecting a float */
					fmt.Printf("\t%10s%10.2f\n", size, cost)
				}
			}
		case "2":
			fmt.Println("Please enter the name of the new item")
			name, _ := in.ReadString('\n')
			// the "make" function allows you to create an empty map so you can add the prices sometime in the future
			menu = append(menu, menuItem{name: name, prices: make(map[string]float64)})
		
		case "q":
			/* "loop:" is a label its a valid Go identifier. You can use that label aftr break so it will break out of the statement that follows the label above.
			The statement that follows the loop label above is the for loop. Without the label our break statement only would have broken out of the switch statement*/
			break loop
		default:
			fmt.Println("Unknown option")
		}
	}
}
