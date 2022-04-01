package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open() //1
	if err != nil {        //2
		log.Fatal(err) //3
	}
	defer func() { //4
		_ = keyboard.Close() //5
	}()

	// convert the menu into a map
	coffees := make(map[int]string) //maps are specialized data structure so you can't just say "var coffees = map()"
	coffees[1] = "Cappucino"
	coffees[2] = "Latte"
	coffees[3] = "Americano"
	coffees[4] = "Mocha"
	coffees[5] = "Macchiato"
	coffees[6] = "Espresso"

	fmt.Println("MENU")
	fmt.Println("----")
	fmt.Println("1 - Cappucino")
	fmt.Println("2 - Latte")
	fmt.Println("3 - Americano")
	fmt.Println("4 - Mocha")
	fmt.Println("5 - Macchiato")
	fmt.Println("6 - Espresso")
	fmt.Println("Q - Quit the program") //Instead of ESC, now we are looking Q being pressed

	for {
		char, _, err := keyboard.GetSingleKey() // rune is basically single character(lower level than the type string)
		if err != nil {
			log.Fatal(err)
		}
		if char == 'q' || char == 'Q' { //Note that unlike python, Go distinguishes "" and ''. "" is used for string whereas '' is for rune
			break
		}

		//How can you convert the rune to human-readable character? - By the following!
		t := fmt.Sprintf("You chose %q", char) // This returns a string. fmt.Sprintf replaces %q with char
		fmt.Println(t)

		//What about converting string to integer?
		i, _ := strconv.Atoi(string(char))
		t2 := fmt.Sprintf("You chose %d", i) //placeholder for integer is 'd'
		fmt.Println(t2)

		t3 := fmt.Sprintf("You chose %s", coffees[i])
		fmt.Println(t3)

		//What if you press things that are not numeric? -> it will return 0

	}
	fmt.Println("Program exiting...")
}
