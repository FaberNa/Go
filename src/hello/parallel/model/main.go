package model

import (
	"fmt"
	"time"
	"sync"
)

type guestConnection struct {
	ip       string
	userName string
	isAdmin  bool
}

type visitorConnection struct {
	ip       string
	connHour int
}

type notifier interface{
	Notify(*sync.WaitGroup)
}


func GetAllConnections()[] notifier {
	gConn := &guestConnection{ip: "192.168.0.10", userName: "Darth Vader"}
	vConn := &visitorConnection{ip: "192.168.0.11", connHour: time.Now().Hour()}

	return []notifier{gConn, vConn}
}

// one implementation for guestConnection
func (g guestConnection) Notify(group *sync.WaitGroup) {
	fmt.Println("Guest connection from user name:", g.userName)
	group.Done()
}

// and a different implementation for visitorConnection
func (v visitorConnection  ) Notify(group *sync.WaitGroup) {
	fmt.Println("Visitor connected at:", v.connHour)
	group.Done()
}
