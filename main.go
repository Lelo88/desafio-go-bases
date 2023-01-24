package main

import (
	"fmt"

	"github.com/Lelo88/desafio-go-bases/internal/tickets"
)

func main() {
	total, err := tickets.GetTotalTickets("Brazil")
	tickets.LoadTicketsCSV()
	if err!= nil {
		panic(err)
	}

	fmt.Println(total)
}
