package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/ronexlemon/Proto/services/genproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)
func main() {
	start := time.Now()
	conn, err := grpc.NewClient("localhost:9090",grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err !=nil{
		log.Fatalf("Failed to connect to grpc server %v",err)
	}
	defer conn.Close()

	client := pb.NewCoffeShopClient(conn)
	ctx,cancel := context.WithTimeout(context.Background(),time.Second)
	defer cancel()

	menuStream,err := client.GetMenu(ctx,&pb.MenuRequest{})
	if err !=nil{
		log.Fatalf("Failed to get menu %v",err)
	}
	done := make(chan bool)
	var items []*pb.Item

	go func(){
		for{
			resp,err := menuStream.Recv()
			if err == io.EOF{
				done <- true
				return
		}
		if err !=nil{
			log.Fatalf("Failed to get menu %v",err)
		}
		items = append(items,resp.Item...)
		log.Printf("response received %v",resp.Item)
		
	}

	}()
	<-done
	fmt.Println("Items ",items)
	receipt,err := client.PlaceOrder(ctx,&pb.Order{Items: items})
	if err != nil{
		log.Fatalf("Failed to place order %v",err)
	}
	fmt.Println("Receipt ",receipt)

	status,err := client.GetOrderStatus(ctx,receipt)
	if err != nil{
		log.Fatalf("Failed to get order status %v",err)
		}
		fmt.Println("Order status ",status)
		fmt.Println("Time Taken",time.Since(start))
}