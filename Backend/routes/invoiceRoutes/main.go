package invoiceRoutes

import (
	"dbms/handlers/invoiceHandler"
	"dbms/router"
	"net/http"
)


func NewInvoiceRoutes(router router.RouterInterface){
   router.Post("/create/invoice",invoiceHandler.CreateInvoice)
   router.Get("invoice/:id",func(response http.ResponseWriter, request *http.Request){})
}