package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// IndexHandler is a representation of a index
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Please Get Data From Client"))
	tmpl.ExecuteTemplate(w, "index", "")

}

// NotifHandler is a representation of a notif
func NotifHandler(w http.ResponseWriter, r *http.Request) {
	var parsebody JSONNotif

	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	defer r.Body.Close()

	err := json.Unmarshal(body, &parsebody)
	if err != nil {
		panic(err)
	}

	// w.Header().Set("Content-Type", "application/json")
	// //json.NewEncoder(w).Encode(parsebody)
	//
	// conn, err := upgrader.Upgrade(w, r, nil)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// msgType, msg, err := conn.ReadMessage()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	//
	// if string(msg) != "" {
	// 	err = conn.WriteMessage(msgType, []byte("pong"))
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// } else {
	// 	conn.Close()
	// 	fmt.Println(string(msg))
	// 	return
	// }

	fmt.Println(parsebody.StatusMessage)
	fmt.Println(parsebody.TransactionStatus)

}

// ChargeHandler is a representation of a charge
func ChargeHandler(w http.ResponseWriter, r *http.Request) {
	var parsebody JSONBody
	reqBody := r.Body
	token := "Basic " + base64.StdEncoding.EncodeToString([]byte(midclient.ServerKey+":"))
	url := "https://app.sandbox.midtrans.com/snap/v1/transactions"

	spaceClient := http.Client{
		Timeout: time.Second * 60, // Maximum of 2 secs
	}

	jsonBody, err := ioutil.ReadAll(reqBody)
	defer r.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(jsonBody, &parsebody)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	data, err := json.Marshal(parsebody)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBufferString(string(data)))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", token)

	res, err := spaceClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	datajson := JSONReturn{}
	jsonErr := json.Unmarshal(body, &datajson)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	w.Write([]byte(body))

}

// PayFinHandler is a representation of a finish payment
func PayFinHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Payment Finish"))
}

// PayUnfinHandler is a representation of a unfinish payment
func PayUnfinHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Payment Unfinish"))
}

// PayErrHandler is a representation of a error payment
func PayErrHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Payment Error"))
}

// Websocket is a representation of a realtime data
func Websocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	msgType, msg, err := conn.ReadMessage()
	if err != nil {
		fmt.Println(err)
		return
	}
	if string(msg) == "ping" {
		err = conn.WriteMessage(msgType, []byte("pong"))
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		conn.Close()
		fmt.Println(string(msg))
		return
	}
}
