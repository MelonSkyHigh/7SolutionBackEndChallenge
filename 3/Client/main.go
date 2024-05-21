package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	pb "MeatCounter/MeatCounter"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", ":8080", "the address to connect to")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	// conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMeatCounterClient(conn)

	// Contact the server and print out its response.
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()

	g := gin.Default()
	g.GET("/beef/summary", func(ctx *gin.Context) {
		r, err := c.MeatCount(ctx, &pb.Empty{})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		// log.Println(r.GetMeatCount())
		var meatCountbag map[string]int
		json.Unmarshal([]byte(r.GetMeatCount()), &meatCountbag)

		ctx.JSON(http.StatusOK, gin.H{
			"beef": meatCountbag,
		})
	})

	if err := g.Run(":4040"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}
