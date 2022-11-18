package errorsHandlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type DataBaseError struct{
	Message string `json:"message"`
	Status  int    `json:"status"`
	Err     string `json:"-"`
}


func (message *DataBaseError) Error()string {
	err := fmt.Sprintf("status %d: err :  %v message : %v ", message.Status, message.Err, message.Message)
	fmt.Println(err)
	return err
}

func (message *DataBaseError) WriteJsonToStream(io http.ResponseWriter) {

	response, err := json.MarshalIndent(message, "", "")

	if err != nil {
		log.Fatal("json error parser")
	}

	io.WriteHeader(message.Status)
	
	io.Write(response)

}