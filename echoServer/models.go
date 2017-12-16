package echo

import (
	"net/http"
)

type EchoHttpBody struct {
	Request RequestBody
	Session struct {
		User struct {
			UserID string
		}
	}
}
type RequestBody struct {
	Type   string
	Intent IntentBody
}
type IntentBody struct {
	Name  string
	Slots map[string]string
}

type EchoObj struct {
	Handler func(http.ResponseWriter, *http.Request)
	Intents map[string]func(*EchoIntentReq, *EchoResponse)
	Status
}

type Status struct {
	StatusRoute   string
	StatusHandler func(w http.ResponseWriter, r *http.Request)
}

type EchoInterface interface {
	Handler(http.ResponseWriter, *http.Request)
}

type EchoIntentReq struct {
	Slots  map[string]string
	UserID string
}
type EchoResponse struct {
	Version  string       `json:"version"`
	Response ResponseEcho `json:"response"`
}

type ResponseEcho struct {
	OutputSpeech     Output `json:"outputSpeech"`
	ShouldEndSession bool   `json:"ShouldEndSession"`
}
type Output struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func (responseEcho *EchoResponse) setDefaultEchoResponse() { // Add default when declaring Make it as a reciever
	responseEcho = &EchoResponse{
		Response: ResponseEcho{
			OutputSpeech: Output{
				Text: "Ok",
				Type: "PlainText",
			},
			ShouldEndSession: true,
		},
		Version: "1.0",
	}
	return
}
