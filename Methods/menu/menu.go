package menu

import (
	"bufio"
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
func (m *menu) add() {
	fmt.Println("Please enter the name of the new item")
	name, _ := in.ReadString('\n')
	// You need to dereferance the pointer here because you're changing the value your pointer is pointing to. Becuase you're receiving a pointer instead of a value.
	*m = append(*m, menuItem{name: strings.TrimSpace(name), prices: make(map[string]float64)})
}

var in = bufio.NewReader(os.Stdin)

// Must be capitalized to be able to share with methods.go
func AddItem() {
	data.add()
}

func Print() {
	data.print()
}