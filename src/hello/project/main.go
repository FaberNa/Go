package main

import("./model")

func main() {
	notifiers := model.GetAllConnections()
	for _, c := range notifiers {
		c.Notify()
	}

}