package session

import (
	"fmt"
	"sync"

	"../types"
)

type (
	//SeesionManager controlls the pool of sessions
	sessionManager struct {
		sessionPool []*Session
	}
)

var instance *sessionManager
var once sync.Once

//GetInstance Create and gety single instance of Session manager
func GetInstance() *sessionManager {
	fmt.Println("Start get  the session manager")
	once.Do(func() {
		instance = &sessionManager{}
	})
	return instance
}

// BuildSession add ssesion to the pool
func (sessM *sessionManager) AddSession() {
	fmt.Println("Start adding the session")
	ContCh := make(chan types.CloseType, 0)
	session := InitSession("HTTP", "STOMP", len(sessM.sessionPool)+1, ContCh)
	sessM.sessionPool = append(sessM.sessionPool, session)
	go sessM.ControlSession(session)

	err := session.Run()
	if err != nil {
		session.ContCh <- types.ALL
		session.Status = types.DOWN
	}

}

//Control the session
func (sessM *sessionManager) ControlSession(session *Session) {
	fmt.Println("Start Controll Session  >>>> ", session.ID)
	for {
		closeType, ok := <-session.ContCh
		fmt.Println("GET Close Message  >>>> ", closeType)

		if ok && closeType == types.EGRESS {
			session.EgFailures++
		}
		fmt.Println("Egress Failures >>>> ", session.EgFailures)

		if session.EgFailures >= 2 {
			if session.Status != types.DOWN {
				fmt.Println("Egress Is Down  >>>> ")
				session.Status = types.DOWN
				session.ContCh <- types.ALL
				session.EgFailures = 0
			}
		}
		if ok && closeType == types.INGRESS {
			session.IngFailures++
		}
		fmt.Println("Ingress Failures >>>> ", session.IngFailures)

		if session.IngFailures >= 5 {
			if session.Status != types.DOWN {
				session.Status = types.DOWN
				session.ContCh <- types.ALL
				session.IngFailures = 0

			}

		}
	}
}

// replace close ch with 3rd tyoe for ctrl ch

//ControlManager ...
func (sessM *sessionManager) ControlSessionPool() {
	for {
		for _, session := range sessM.sessionPool {
			if session.Status == types.DOWN { // down
				println("Get DOWN Status  of session >>>> ", session.ID)
				if session.RecoveryCount == 5 {
					session.Status = types.DELETED // deleted
					sessM.sessionPool = append(sessM.sessionPool[:session.ID], sessM.sessionPool[session.ID+1:]...)
					break

				}
				// start recovery
				session.Run()
				println("Start Recover of session >>>> ", session.ID)
				session.Status = types.UP
				session.RecoveryCount++
			}
		}
	}
}
