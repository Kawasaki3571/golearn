package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("%d: %s\n", i, func(i int) string {
			switch true {
			case i % 15 == 0:
				return "fizzbuzz"
			case i % 5 == 0:
				return "buzz"
			case i % 3 == 0:
				return "fizz"
			default:
				return ""
			}
		}(i))
	}
}