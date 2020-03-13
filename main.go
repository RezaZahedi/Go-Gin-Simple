package main

import (
	"github.com/RezaZahedi/Go-Gin-Simple/rest_service"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	rest_service.InitializeRestService(wg)
	wg.Wait()
}
