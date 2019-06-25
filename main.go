package main

import (
	"fmt"
)

type Msg struct {
	Dist   int
	Sender string
}

type Token struct {
	Sender string
}

func main() {
	ch := make(chan interface{}, 10)

	ch <- Token{"P"}
	ch <- Msg{10, "Q"}
	ch <- Token{"R"}

	for v := range ch {
		switch t := v.(type) {
		case Token:
			fmt.Printf("Got a Token %s\n", t.Sender)
		case Msg:
			fmt.Printf("Got a Msg %v\n", t.Dist)
		default:
			fmt.Printf("Other: %v \n", t)
		}

	}

}
