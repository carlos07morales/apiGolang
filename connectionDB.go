package main

import (
	"fmt"
	"time"
	"context"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var client *mongo.Client

func ConectionBD() (*mongo.Collection, *mongo.Client){
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://apiGo:apiGo@cluster0.vaoqy.mongodb.net/apiGo?retryWrites=true&w=majority"))
    if err != nil {
        log.Fatal(err)
    }
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Conexión exitosa")
	collection := client.Database("apiGo").Collection("tickets")
	// defer client.Disconnect(ctx)
	return collection, client
}

// func DisconnectMongo(){
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer client.Disconnect(ctx)
// 	fmt.Println("Desconectado")
// }