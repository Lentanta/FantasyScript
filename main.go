package main

import (
	"fmt"

	"github.com/Lentanta/FantasyScript/lexer"
)

func main() {
	code := `
    let five = 5;
    let ten = 10;
    let add = fn(x, y) {
      x + y;
    };
    let result = add(five, ten);
  `
	lx := lexer.New(code)
	fmt.Println(lx)
}
