# Amazon Echo Server in Golang
Note: *This is still in beta*

This package helps you to easily set up an amazon echo server in golang.

## Getting Started

### Prerequisites
* Go version 1.9

### Installing
Install the package with the go get command
```
go get package
```
### Usage
```
	func IntentFunction(req *echo.EchoIntentReq, res *echo.EchoResponse) {
	...
  ...
}
  
  intents := map[string]func(*echo.EchoIntentReq, *echo.EchoResponse){
		"IntentName":       IntentFunction,
	}
	echoIns := echo.GetEchoIntentInput(intents)
	echoIns.StartServer(":8080")
  ```

  **More Detailed explanation will be updated shortly**
