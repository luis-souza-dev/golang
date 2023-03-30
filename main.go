package main

import (
	"basics/data-structures/basics"
	"basics/data-structures/pointers"
	"fmt"
	"os"

	"github.com/hackebrot/turtle"
)

func main() {
	basics.Basics();
	name := "turtle"
	emoji, ok := turtle.Emojis[name]

	if !ok {
		fmt.Fprintf(os.Stderr, "no emoji found for name: %v\n", name)
		os.Exit(1)
	}

	fmt.Printf("Name: %q\n", emoji.Name)
	fmt.Printf("Char: %s\n", emoji.Char)
	fmt.Printf("Category: %q\n", emoji.Category)
	fmt.Printf("Keywords: %q\n", emoji.Keywords)
	fmt.Println();
	pointers.Pointers();
}