package main

import (
	"fmt"

	"github.com/brentgroves/replib/mtbf"
	"github.com/brentgroves/replib/trlbal"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())

	message2 := trlbal.Hello("jojo")
	fmt.Println(message2)

	message3 := mtbf.Hello("bobo")
	fmt.Println(message3)

	message4 := trlbal.RunScript("test.sh")
	fmt.Println(message4)

}
