package http

import (
	"fmt"
	"net/http"
)

//HTTP ...
type HTTP struct {
}

//Connect ...
func (ht *HTTP) Connect() error {
	fmt.Println("No Connection is needed")
	return nil
}

//Send ...
func (ht *HTTP) Send(data []byte) bool {
	fmt.Println("Received New Data  For Egress >>>>>", string(data))
	resp, err := http.Get(string(data))
	if err != nil {
		return false
	}
	fmt.Println("Get Response from HTTP", resp)
	return true
}

//Close ...
func (ht *HTTP) Close() {
	fmt.Println("Connection is closed")
}
