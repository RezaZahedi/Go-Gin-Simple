package rest_service

import (
	"github.com/RezaZahedi/Go-Gin-Simple/fibonacci"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type FibonacciService struct {
	FibonacciCalculator func(int) (string, error)
}

func ProvideFibonacciService() FibonacciService {
	return FibonacciService{FibonacciCalculator: fibonacci.NewFiboGenerator().GenerateNumber}
}


func (f *FibonacciService)GetFibonacciAnswer(c *gin.Context) {
	number, err := strconv.Atoi(c.PostForm("number"))
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if number < 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	output, err := (f.FibonacciCalculator)(number)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, output)
}
