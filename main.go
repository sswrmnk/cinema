package main

import (
	"fmt"

	"github.com/sswrmnk/cinema/movie"
	"github.com/sswrmnk/cinema/ticket"
)

func init() {
	fmt.Println("init : main")
}

func main() {
	movieName := movie.FindName("tt4154796")
	fmt.Println(movieName)
	movie.Review(movieName, 8.4)
	ticket.BuyTicket(movieName)
}
