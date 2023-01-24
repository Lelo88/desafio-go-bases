package main

import (
	"fmt"
	"log"

	"github.com/Lelo88/desafio-go-bases/internal/tickets"
)

func main() {
	tickets.LoadTicketsCSV()
	
	total, _ := tickets.GetTotalTickets("China")
	horas, err := tickets.GetMornings("90:41")
	

	fmt.Println(total)

	if horas == 0{
		log.Println(err)
		return
	}

	
	fmt.Println(horas)
}
