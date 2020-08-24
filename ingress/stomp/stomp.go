package stomp

import (
	"crypto/tls"
	"net"

	"../../types"
	"github.com/go-stomp/stomp"
)

//STOMP ...
type STOMP struct {
	Host           string
	Username       string
	Password       string
	QueueName      string
	SSL            bool
	Conn           *tls.Conn
	NetConn        net.Conn
	Stomp          *stomp.Conn
	PrefetetchSize int16
}

//Connect ...
func (st *STOMP) Connect() error {
	var err error
	if st.SSL == true {
		println("Enter TLS")
		st.Conn, err = tls.Dial("tcp", st.Host, &tls.Config{})
		st.Stomp, err = stomp.Connect(st.Conn, stomp.ConnOpt.HeartBeat(0, 0), stomp.ConnOpt.Login(st.Username, st.Password))
	} else {
		st.NetConn, err = net.Dial("tcp", "192.168.71.71:61613")
		st.Stomp, err = stomp.Connect(st.NetConn, stomp.ConnOpt.HeartBeat(0, 0), stomp.ConnOpt.Login(st.Username, st.Password))

	}

	if err != nil {
		println("ٌٌُERROOROROROROR >>>>>", err)
		return err
	}
	println("STOMP Connected...")
	return nil
}

//Close ...
func (st *STOMP) Close() error {
	println("Closiong is starting .....")
	st.Stomp.Disconnect()
	var err error
	if st.Conn != nil {
		err = st.Conn.Close()
	} else {
		err = st.NetConn.Close()

	}
	println(err)
	return nil
}

//Receive ...
func (st *STOMP) Receive(closeTyoe chan types.CloseType) (chan []byte, error) {

	sub, err := st.Stomp.Subscribe(st.QueueName, stomp.AckClient)
	if err != nil {
		println("ٌٌُERROOROROROROR >>>>>", err)
		return nil, err
	}

	senCh := make(chan []byte)

	go func() {
		for {
			msg, ok := <-sub.C

			if ok {
				println("GET Message >>>>>", string(msg.Body))
				st.Stomp.Ack(msg)
				senCh <- msg.Body

			} else {
				println("NOT OK  >>>>>")
				close(senCh)
				go func() { closeTyoe <- types.INGRESS }()
				st.Close()
				break
			}
		}
	}()
	return senCh, nil
}
