package echo

import "net/http"

//GetEchoHandlerInput accepts a handler function and returns an Echo Object.
//Handler function accepts http.ResponseWriter and *http.Request as parameters
func GetEchoHandlerInput(handlr func(http.ResponseWriter, *http.Request)) *EchoObj {
	var echo EchoObj
	echo.Handler = handlr
	return &echo
}

//GetEchoIntentInput accepts echo intents which are a map of intent names along with the functions
//and returns an echo object
func GetEchoIntentInput(intents map[string]func(*EchoIntentReq, *EchoResponse)) *EchoObj {
	var echo EchoObj
	echo.Intents = make(map[string]func(*EchoIntentReq, *EchoResponse), len(intents)) // len given for more efficiency to reduce array generation again
	//TODO check  the feadibility of parallelism
	for i, val := range intents {
		func(i string, val func(*EchoIntentReq, *EchoResponse)) {
			echo.Intents[i] = val
		}(i, val)
	}
	return &echo
}
