package main

import (
	"fmt"
	"github.com/RezaZahedi/Go-Gin-Simple/rest_service"
	"sync"
)

func main() {
	fmt.Println("hello!!")
	wg := sync.WaitGroup{}
	wg.Add(1)
	rest_service.InitializeRestService(wg)
	wg.Wait()
}
