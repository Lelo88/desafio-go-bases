package tickets

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
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
	
}

func (t *Ticket) Save(){
	ticketsAmemoria = append(ticketsAmemoria, *t) //como es un slice, utilizo el append cada vez que vaya a agregar un producto
}

// ejemplo 1
func GetTotalTickets(destination string) (int,error) { 
	
	totalPaises := 0

	for _, ticket := range ticketsAmemoria {
		if ticket.DestinyCountry == destination {
			totalPaises ++
		}
	}
	
	return totalPaises, nil 
}

// ejemplo 2
func GetMornings(time string) (int, error) {
	
	horasPersonas := 0
	hora := strings.Split(time, ":")
	horaint,_ := strconv.ParseInt(hora[0], 10, 64)

	for _, ticket := range ticketsAmemoria {
		horaSalida := strings.Split(ticket.Departure, ":")
		horaSalidaInt,_ := strconv.ParseInt(horaSalida[0],10,64)

			if horaint == horaSalidaInt{
				switch{
				case horaSalidaInt >= 0 && horaSalidaInt <= 6:
					horasPersonas++
				case horaSalidaInt >= 7 && horaSalidaInt <= 12:
					horasPersonas++
				case horaSalidaInt >= 13 && horaSalidaInt <=19:
					horasPersonas++
				case horaSalidaInt >= 20 && horaSalidaInt <= 23:
					horasPersonas++
				default:
					return 0, errors.New("no se encuentra el horario especificado")	
				}
			}
	}
	//horaSalida := strings.Split(, ":")

	return horasPersonas, nil
}

// ejemplo 3
func AverageDestination(destination string) (float64,error) {
	
	totalTickets, err := GetTotalTickets(destination)

	if err!= nil {
		return float64(totalTickets), err
	}

	return float64(totalTickets) / float64(len(ticketsAmemoria)), nil
}
