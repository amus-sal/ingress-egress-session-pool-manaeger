package main

import (
	"./session"
)

func main() {
	sessionManager := session.GetInstance()
	go sessionManager.AddSession()
	sessionManager.ControlSessionPool()

}
