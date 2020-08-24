package ingress

import (
	"fmt"

	"../types"
	st "./stomp"
)

//Ingress ...
type Ingress struct {
	Type   string
	ContCh chan types.CloseType
}

//NewIngress  add new ingress
func (ing *Ingress) NewIngress() (chan []byte, error) {
	fmt.Println("Start adding new ingress method")
	rec := make(chan []byte, 0)
	if ing.Type == "STOMP" {
		stomp := &st.STOMP{
			Host:           "192.168.71.71:61613",
			QueueName:      "Test_A",
			Username:       "",
			Password:       "",
			SSL:            false,
			PrefetetchSize: 1,
		}
		err := stomp.Connect()
		if err != nil {
			return nil, err
		}
		rec, err = stomp.Receive(ing.ContCh)
		if err != nil {
			return nil, err
		}
		go func() {
			for {
				chStatus, ok := <-ing.ContCh
				if !ok || chStatus == types.ALL {
					fmt.Println("GET Close Order for Ingress >>>")
					stomp.Close()
				}
			}
		}()
	}

	return rec, nil
}
