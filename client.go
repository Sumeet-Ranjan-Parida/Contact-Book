package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Sumeet-Ranjan-Parida/ContactBook/proto"
	"github.com/gin-gonic/gin"

	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":4040", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not connect: %s", err)
	}

	defer conn.Close()

	client := proto.NewContactClient(conn)

	r := gin.Default()
	r.GET("/health", health)
	r.GET("/addcontact/:name/:number", func(ctx *gin.Context) {

		name := ctx.Param("name")

		number, err := strconv.ParseUint(ctx.Param("number"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter"})
		}

		req := &proto.Request{Name: string(name), Number: int64(number)}

		if response, err := client.Getcontact(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Cname, response.Cnumber),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	r.Run()
}

func health(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{"status": 200, "data": "Testing api", "alive": true})
}
