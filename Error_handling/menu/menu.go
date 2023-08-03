package menu

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type menuItem struct {
	name   string
	prices map[string]float64
}

type menu []menuItem

func (m menu) print() {
	for _, item := range m {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, cost := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", size, cost)
		}
	}
}

// You need a pointer reciever (*menu) because you will be changing the menu.
func (m *menu) add() error {
	fmt.Println("Please enter the name of the new item")
	name, _ := in.ReadString('\n') // Coffee\n
	name = strings.TrimSpace(name)

	// loops through data to check is the inputed name is the same as any in the data and throws an error if it finds a match.
	for _, item := range data {
		if item.name == name {
			// returns an error using the errors.New() function if a match is found
			return errors.New("menu item already exsists")
		}
	}
	// You need to dereferance the pointer here because you're changing the value your pointer is pointing to. Becuase you're receiving a pointer instead of a value.
	*m = append(*m, menuItem{name: name, prices: make(map[string]float64)})

	// if the an error is never thrown you still have to return something. Since errors are an interface type you can return nil, meaning no error occured.
	return nil
}

var in = bufio.NewReader(os.Stdin)

// Must be capitalized to be able to share with methods.go
func AddItem() error {
	return data.add()
}

func Print() {
	data.print()
}