package tickets

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

//import "encoding/csv"

type Ticket struct {
	Id 				int64
	Name 			string
	Email 			string
	DestinyCountry 	string
	Departure 		string
	Price 			float64
}

var ticket [][]string
var ticketsAmemoria []Ticket

//leer archivo csv y traerlo

func LoadTicketsCSV() {
	archivo := "tickets.csv"
	ar, err := os.Open(archivo)

	if err!=nil{
		log.Fatalf("No se puede abrir el archivo, err es %v", err)
	}

	defer ar.Close()

	archCSV := csv.NewReader(ar)

		
	if err!= nil && err!=io.EOF {
		log.Fatalf("No se puede lerr el archivo, %v", err)
	}
	if err == io.EOF {
		log.Fatalf("No se puede lerr el archivo, %v", err)
	}
		//fmt.Println(row)	
	
	i:=0
	for {
		
		record, err := archCSV.Read()
		if err == io.EOF {
			break
		}
		if err!= nil {
			log.Printf("error: %v", err)
		}

		ticket = append(ticket, record)

		id, _ := strconv.ParseInt(ticket[i][0], 10, 64)
		precio, _:= strconv.ParseFloat(ticket[i][5],64)
		newTicket := Ticket{
			Id:             id,
			Name:           ticket[i][1],
			Email:          ticket[i][2],
			DestinyCountry: ticket[i][3],
			Departure:      ticket[i][4],
			Price:          precio,
		}
		
		newTicket.Save()
		i++

	}

	for _, v := range ticket {
		fmt.Println(v)
	}
	
}

func (t *Ticket) Save(){
	ticketsAmemoria = append(ticketsAmemoria, *t) //como es un slice, utilizo el append cada vez que vaya a agregar un producto
}

// ejemplo 1
func GetTotalTickets(destination string) (int,error) { return 0, nil }

// ejemplo 2
func GetMornings(time string) (int, error) {return 0, nil}

// ejemplo 3
func AverageDestination(destination string, total int) (int,error) {return 0,nil}
