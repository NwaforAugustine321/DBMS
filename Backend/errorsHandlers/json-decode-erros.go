package errorsHandlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type JsonErrorParser struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Err     string `json:"-"`
}

func (message *JsonErrorParser) PrintJsonError() {
	err := fmt.Sprintf("status %d: err :  %v message : %v ", message.Status, message.Err, message.Message)
	fmt.Println(err)
}

func (message *JsonErrorParser) WriteJsonToStream(io http.ResponseWriter) {

	response, err := json.MarshalIndent(message, "", "")

	if err != nil {
		log.Fatal("json error parser")
	}

	io.Write(response)

}
