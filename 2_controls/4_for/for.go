package main

import (
	"fmt"
)

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i < 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "Even")
		} else {
			fmt.Println(i, "Odd")
		}
	}

	/*for {
		fmt.Println("Infinite loop")
		time.Sleep(time.Second * 2)
	}*/

}
