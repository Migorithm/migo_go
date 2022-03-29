package main

import (
	"fmt"
	"log"

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

	fmt.Println("Press any key on the keyboard. Press ESC to quit.")
	for {
		char, key, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if key != 0 {
			fmt.Println("You pressed", char, key)
		} else {
			fmt.Println("You pressed", char)
		}

		if key == keyboard.KeyEsc {
			break
		}

	}
	fmt.Println("Program exiting...")
}

// 4 whatever follows the defer keyword won't execute immediately.
// Instead, it will execute as soon as the current function ends.
// Plus, this is the way you define anonymous function.

// 5 as soon as the main function finishes, it will close the keyboard.
// It might throw an error, but we ignore it.
