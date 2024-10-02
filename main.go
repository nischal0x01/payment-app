package main

import (
	"errors"
	"fmt"
	"net/http"
)
func returnsError(password string) error{

	var secretPassword string = "password"
	if password == secretPassword{
		return nil
	}else{
		return errors.New("password mismatch")
	}
}
func main(){
// 	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)
// var err  error= http.ListenAndServe("localhost:3000",nil)
// if err!=nil{
// 	log.Fatal(err)
// }

var err error= returnsError("pasword")
if err!= nil{
fmt.Println(err)}
}

func handleCreatePaymentIntent(http.ResponseWriter, *http.Request){
	fmt.Println("Endpoint called")
}