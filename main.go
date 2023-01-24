package main

import (
	"fmt"

	"github.com/Lelo88/desafio-go-bases/internal/tickets"
)

func main() {
	tickets.LoadTicketsCSV()
	
	total, _ := tickets.GetTotalTickets("China")
	horas, _ := tickets.GetMornings("12:41")
	promedio, _ := tickets.AverageDestination("Argentina")

	fmt.Println("Total personas a China:",total)
	fmt.Println("Total de vuelos durante esta hora:",horas)
	fmt.Println("Promedio de horas:",promedio)
}
