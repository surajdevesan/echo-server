package echo

import (
	"log"
	"net/http"
)

func panicHandler() { //TODO Better error handling
	if rec := recover(); rec != nil {
		log.Println("Recovered from panic", rec)
	}
}

//GetDefaultEchoResponse sets default EchoResponse
func GetDefaultEchoResponse() *EchoResponse { // Add default when declaring Make it as a reciever
	responseEcho := &EchoResponse{}
	responseEcho.setDefaultEchoResponse()
	return responseEcho
}

//StartServer accpets a port and starts a http server
func (echoApp *EchoObj) StartServer(port string) {
	if echoApp.Handler != nil {
		http.HandleFunc("/", echoApp.Handler)
	} else {
		http.HandleFunc("/", handleEchoReq(echoApp))
	}
	if echoApp.Status.StatusHandler != nil && echoApp.Status.StatusRoute != "" {
		http.HandleFunc(echoApp.Status.StatusRoute, echoApp.Status.StatusHandler)
	} else {
		http.HandleFunc("/status", statusHandler)
	}
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Server startup error", err)
	}
}
