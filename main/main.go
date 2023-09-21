package main

import (
	"fmt"

	"github.com/brentgroves/greetings"

	"github.com/brentgroves/replib/trlbal"

	"github.com/brentgroves/replib/mtbf"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)

	message2 := trlbal.Hello("jojo")
	fmt.Println(message2)

	message3 := mtbf.Hello("bobo")
	fmt.Println(message3)

	message4 := trlbal.RunScript("firefox")
	fmt.Println(message4)

}
