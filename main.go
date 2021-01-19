package main

import (
	"net/http"
	"fmt"
	"log"
)

func main(){
	log.Print("Iniciando servidor PORT:3000")
	server := NewServer(":3000")
	server.Handle("GET", "/", HandlerRoot)
	server.Handle("GET", "/getTicket", GetTicket)
	server.Handle("GET", "/getAllTickets", GetAllTickets)
	server.Handle("POST", "/setTicket", SetTicket)
	server.Handle("PUT", "/updateTicket", UpdateTicket)
	server.Handle("DELETE", "/deleteTicket", DeleteTicket)
	server.Listen()
}

func HandlerRoot(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hola Go")
}