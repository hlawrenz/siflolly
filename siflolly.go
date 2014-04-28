package main

import (
	//	"html/template"
	//	"io/ioutil"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/hlawrenz/siflolly/noise"
	"log"
	"net/http"
	"os"
	"path"
)

type Context struct {
	Title string
}

//var templates = template.Must(template.ParseFiles("index.html"))

func rootHandler(w http.ResponseWriter, r *http.Request) {

}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		log.Println(err)
		return
	}
	log.Println("Past connection")

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			return
		}
		log.Println(messageType)
		log.Println(p)
		if err = conn.WriteMessage(messageType, p); err != nil {
			return
		}
	}
}

var assetDir = path.Join(os.Getenv("GOPATH"), "src/github.com/hlawrenz/siflolly/assets/")
var echowsDir = path.Join(os.Getenv("GOPATH"), "src/github.com/hlawrenz/siflolly/echows/")
var templateDir = path.Join(os.Getenv("GOPATH"), "src/github.com/hlawrenz/siflolly/templates/")

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(assetDir))))
	http.Handle("/echows/", http.StripPrefix("/echows/", http.FileServer(http.Dir(echowsDir))))
	//http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	fmt.Println(assetDir)
	sb := noise.NewSwitchboard()
	n := noise.Noise{Sb: &sb}
	http.HandleFunc("/echosock", echoHandler)
	http.HandleFunc("/noise", n.ServeWs)
	http.ListenAndServe(":8888", nil)
}
