package main

import (
	"go.mongodb.org/mongo-driver/bson"
	"fmt"
	"context"
	"time"
	"net/http"
	"encoding/json"
	"strconv"
)

var Ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

// consulta un Ticket por id
func GetTicket(w http.ResponseWriter, r *http.Request){
	IdGet := r.URL.Query().Get("id")
	idTicket, _ := strconv.Atoi(IdGet)
	
	collection, client := ConectionBD()
	// Ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	defer client.Disconnect(Ctx)

	var data Data
	filter := bson.D{{"id", idTicket}}
	erro := collection.FindOne(context.TODO(), filter).Decode(&data)
	if erro != nil {
		w.Write([]byte("no matches found"))
		return
	}
	response, errJson := data.ToJson()
	if errJson != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}	
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// consulta todos los Ticket
func GetAllTickets(w http.ResponseWriter, r *http.Request){
	// Ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection, client := ConectionBD()
	defer client.Disconnect(Ctx)
	tickets, erro := collection.Find(context.TODO(), bson.M{})
	if erro != nil {
		w.Write([]byte("no matches found"))
		return
	}
	var response []Data
	if erro = tickets.All(Ctx, &response); erro != nil {
		w.Write([]byte("no matches found"))
		return
	}
	defer tickets.Close(Ctx)
	w.Header().Set("Content-Type", "application/json")
	json, _ := json.Marshal(response)
	w.Write(json)
	// DisconnectMongo()
}

//inserta ticket
func SetTicket(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var data Data
	err := decoder.Decode(&data)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
	
	if data.Status == ""{
		w.Write([]byte("Ticket Status is missing"))
		return
	}else
	if data.User == ""{
		w.Write([]byte("Ticket User is missing"))
		return
	}

	// Ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	currentTime := time.Now()
	data.Createdate = currentTime.Format("2006-01-02 15:04:05")
	data.Updatedate = currentTime.Format("2006-01-02 15:04:05")

	collection, client := ConectionBD()

	count, err := collection.CountDocuments(context.Background(), bson.D{})
	data.Id = int(count + 1)

	defer client.Disconnect(Ctx)
	insertResult, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		w.Write([]byte("there was an error saving the ticket"))
		return
	}

	fmt.Println("Found data", insertResult)
	w.Write([]byte("OK"))
}

//actualizar ticket
func UpdateTicket(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var data Data
	err := decoder.Decode(&data)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
	if data.Id == 0{
		w.Write([]byte("Ticket id is missing"))
		return
	}
	if data.Status == ""{
		w.Write([]byte("Ticket Status is missing"))
		return
	}
	
	collection, client := ConectionBD()
	// Ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	defer client.Disconnect(Ctx)
	currentTime := time.Now()
	result, err := collection.UpdateOne(
		Ctx,
		bson.M{"id": data.Id},
		bson.D{
			{"$set", bson.D{
				{"status", data.Status},
				{"updatedate", currentTime.Format("2006-01-02 15:04:05")},
			}},
		},
	)
	if err != nil {
		w.Write([]byte("there was an error updating the ticket"))
		return
	}
	_ = result
	w.Write([]byte("Ok"))
}
//elimina un ticket
func DeleteTicket(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var data Data
	err := decoder.Decode(&data)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
	if data.Id == 0{
		w.Write([]byte("Ticket id is missing"))
		return
	}
	collection, client := ConectionBD()
	defer client.Disconnect(Ctx)
	result, err := collection.DeleteOne(Ctx, bson.M{"id": data.Id})
	if err != nil {
		w.Write([]byte("there was an error delete the ticket"))
		return
	}
	fmt.Printf("DeleteOne removed %v document(s)\n", result.DeletedCount)
	w.Write([]byte("Ok"))
}