package egress

import (
	"fmt"

	"../types"
	"./http"
)

//Egress ...
type Egress struct {
	Type   string
	ContCh chan types.CloseType
	DataCh chan []byte
}

//NewEgress  add new ingress
func (eg *Egress) NewEgress() error {
	fmt.Println("Start adding new Egress method")
	var err error
	if eg.Type == "HTTP" {
		ht := &http.HTTP{}
		err = ht.Connect()
		if err != nil {
			return err
		}
		go func() {
			for {
				data, ok := <-eg.DataCh
				if ok {
					success := ht.Send(data)
					if success != true {
						go func() { eg.ContCh <- types.EGRESS }()
					}
				} else {
					break
				}

			}
		}()
	}
	return err
}
