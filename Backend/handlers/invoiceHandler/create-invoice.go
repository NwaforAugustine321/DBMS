package invoiceHandler

import (
	"dbms/repo/interfaces"
	"encoding/json"
	"net/http"
)


func CreateInvoice (response http.ResponseWriter, request *http.Request){
	var invoicePayload interfaces.Invoices 
	
    err := json.NewDecoder(request.Body).Decode(&invoicePayload)

	if err != nil{
		return 
	}
}