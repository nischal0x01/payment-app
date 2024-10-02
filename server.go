package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)
	fmt.Println("Listening on localhost:4242...")
	var err  error= http.ListenAndServe("localhost:4242",nil)
	if err!=nil{
	log.Fatal(err)
}

}
func handleCreatePaymentIntent(http.ResponseWriter, *http.Request){
	fmt.Println("Endpoint called")
}