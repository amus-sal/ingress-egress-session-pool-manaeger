package session

import (
	"../egress"
	"../ingress"
	"../types"
)

type (
	//Session is a combine of egress and ingress connection
	Session struct {
		IngSession    *ingress.Ingress
		Egress        *egress.Egress
		CommChannel   chan []byte
		EgType        string
		IngType       string
		ID            int
		ContCh        chan types.CloseType
		Status        types.STATUS // 0 down 1 up 2 recoverd
		EgFailures    int8
		IngFailures   int8
		RecoveryCount int8
	}
)

//InitSession the sessions
func InitSession(EgType string, IngType string, id int, ContCh chan types.CloseType) *Session {
	return &Session{
		EgType:        EgType,
		IngType:       IngType,
		ID:            id,
		Status:        types.UP,
		RecoveryCount: 0,
		ContCh:        ContCh,
	}
}

//Run the sessions
func (se *Session) Run() error {
	var err error
	se.IngSession = &ingress.Ingress{Type: se.IngType, ContCh: se.ContCh}
	se.CommChannel, err = se.IngSession.NewIngress(sendFunc)

	se.Egress = &egress.Egress{Type: se.EgType, DataCh: se.CommChannel, ContCh: se.ContCh}
	sendFunc, err = se.Egress.NewEgress()
	if err != nil {
		println("Error with starting the Session", err)
		return err
	}
	return nil
}
