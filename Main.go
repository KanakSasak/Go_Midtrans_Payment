package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/veritrans/go-midtrans"
)

var midclient midtrans.Client
var coreGateway midtrans.CoreGateway
var snapGateway midtrans.SnapGateway

var tmpl = template.Must(template.ParseGlob("template/*"))

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type parsetohtml struct {
	Data1 string
	Data2 string
}

func main() {
	setupMidtrans()
	var addr = flag.String("port", ":9999", "The address of the application")
	flag.Parse()
	fmt.Println("Server started on port: ", *addr)

	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/charge", ChargeHandler)
	r.HandleFunc("/notif", NotifHandler).Methods("POST")
	r.HandleFunc("/payment/finish", PayFinHandler)
	r.HandleFunc("/payment/unfinish", PayUnfinHandler)
	r.HandleFunc("/payment/error", PayErrHandler)
	r.HandleFunc("/websocket", Websocket)

	log.Fatal(http.ListenAndServe(":9999", r))

}

func setupMidtrans() {
	midclient = midtrans.NewClient()
	midclient.ServerKey = "SB-Mid-server-iuU8vdc_CQ6Y2L9jaOYCg5_Q"
	midclient.ClientKey = "SB-Mid-client-UQztUNrjlrcueE9C"
	midclient.ApiEnvType = midtrans.Sandbox

	coreGateway = midtrans.CoreGateway{
		Client: midclient,
	}

	snapGateway = midtrans.SnapGateway{
		Client: midclient,
	}
}
