package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("What would you like me to scream?")
	// bufio.NewReader is a decorator
	in := bufio.NewReader(os.Stdin)
	//Single quotes used to delimint a single character.
	s, _ := in.ReadString('\n')
	//.TrimSpace gets rid of the new line
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)

	fmt.Println(s + "!")
}