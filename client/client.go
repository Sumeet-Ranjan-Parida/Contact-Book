package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := proto.NewAddServiceClient(conn)

	r := gin.Default()

	r.GET("/heathz", heathz)

}

func heathz(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"status": 200, "data": "API Testing", "Alive": true,
	})
}
