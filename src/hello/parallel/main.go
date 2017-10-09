package main

import("./model"
"sync")

func main() {
	var wg sync.WaitGroup
	notifiers := model.GetAllConnections()
	wg.Add(len(notifiers))
	for _, c := range notifiers {
		go c.Notify(&wg)
	}
	wg.Wait()
}