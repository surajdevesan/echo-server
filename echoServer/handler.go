package echo

import (
	"encoding/json"
	"log"
	"net/http"
)

func handleEchoReq(echoApp *EchoObj) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer panicHandler()
		var echoBody EchoHttpBody
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&echoBody)
		intentName := echoBody.Request.Intent.Name
		if err != nil {
			log.Panic("Error in Decoding body", err)
		}
		if echoApp.Intents[intentName] != nil { //TODO make it switch
			intentReq := &EchoIntentReq{
				Slots:  echoBody.Request.Intent.Slots,
				UserID: echoBody.Session.User.UserID,
			}
			intentRes := GetDefaultEchoResponse() //Creating echo response object with default values
			echoApp.Intents[intentName](intentReq, intentRes)
			log.Println("Intent Query succesfully executed")
			resJSON, _ := json.Marshal(intentRes)
			w.Write(resJSON)
		}
	})
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	return
}
