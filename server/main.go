package main

import (
	"context"
	pb "github.com/ronexlemon/Proto/services/genproto"
	"google.golang.org/grpc"
)

type Server struct{
	pb.UnimplementedCoffeShopServer
}

func (s*Server) GetMenu(menuRequest *pb.MenuRequest,srv  pb.CoffeShop_GetMenuServer ) error {
	items := []*pb.Item{
		&pb.Item{
			Id: "1",
			Name: "Cappuccino",
			Price: 25,
			Quantity: 10,},
			&pb.Item{
				Id: "2",
				Name: "Black Tea",
				Price: 100,
				Quantity: 15,},
				&pb.Item{
					Id: "3",
					Name: "Black Coffee",
					Price: 5,
					Quantity: 8,},
	}

	for i,_:=range items{
		srv.Send(&pb.Menu{Item: items[0:i+1]})
	}
	return nil

	
}
func (s *Server) PlaceOrder(ctx context.Context,order *pb.Order) (*pb.Receipt, error) {
	return &pb.Receipt{Id: "123458A",},nil

	
}
func (s *Server) GetOrderStatus(ctx context.Context,receipt *pb.Receipt) (*pb.OrderStatus, error) {

	return &pb.OrderStatus{OrderId: receipt.Id,
		Status: "IN PROGRESS",}  ,nil

	
}