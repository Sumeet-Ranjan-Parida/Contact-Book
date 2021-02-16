package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Sumeet-Ranjan-Parida/ContactBook/proto"
	"github.com/gin-gonic/gin"

	"google.golang.org/grpc"
)

type contactapi struct {
	name string
	phno int
}

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
	// r.GET("/view", view)
	r.GET("/addcontact/:name/:number", func(ctx *gin.Context) {

		name := ctx.Param("name")

		number, err := strconv.ParseUint(ctx.Param("number"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter"})
			return
		}

		req := &proto.Request{Name: string(name), Number: int64(number)}

		if response, err := client.Getcontact(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Result),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	r.GET("/delete/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		db, err := sql.Open("mysql", "root:sumeet@tcp(127.0.0.1:3306)/contactbook")
		if err != nil {
			panic(err.Error())
		}

		defer db.Close()

		del, err := db.Query("DELETE FROM contacts WHERE name=?", name)
		if err != nil {
			panic(err.Error())
		}

		defer del.Close()

		ctx.JSON(http.StatusOK, gin.H{"data": name + " has been successfully deleted"})
	})

	r.Run()
}

func health(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{"status": 200, "data": "Testing API", "alive": true})
}

// func view(g *gin.Context) {

// 	db, err := sql.Open("mysql", "root:sumeet@tcp(127.0.0.1:3306)/contactbook")
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	defer db.Close()

// 	var contacts []*contactapi

// 	rows, err := db.Query("SELECT name, phno FROM contacts")
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	for rows.Next() {
// 		c := new(contactapi)
// 		err := rows.Scan(c.name, c.phno)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		contacts = append(contacts, c)
// 	}
// 	g.JSON(http.StatusOK, contacts)

// }
