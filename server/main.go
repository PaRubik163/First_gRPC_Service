package main

import (
	"context"
	"log"
	"net"
	
	pb "myfirstserviceapp/proto"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMathServiceServer
}

func (s *server) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error){
	logrus.WithFields(logrus.Fields{
		"Method" : "Sum",
		"x" : req.GetX(),
		"y" : req.GetY(),
		"Result" : req.X + req.Y,
	}).Info()

	return &pb.SumResponse{Result: req.GetX() + req.GetY()}, nil
}

func (s *server) Minus(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error){
	logrus.WithFields(logrus.Fields{
		"Method" : "Minus",
		"x" : req.GetX(),
		"y" : req.GetY(),
		"Result" : req.X - req.Y,
	}).Info()

	return &pb.SumResponse{Result: req.X - req.Y}, nil
}

func (s *server) Pow(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error){
	logrus.WithFields(logrus.Fields{
		"Method" : "Pow",
		"x" : req.GetX(),
		"y" : req.GetY(),
		"Result" : req.X * req.Y,
	}).Info()

	return &pb.SumResponse{Result: req.GetX() * req.GetY()}, nil
}

func main(){
	lis, err := net.Listen("tcp", ":7060")

	if err != nil{
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterMathServiceServer(s, &server{})

	logrus.Info("Сервер начал работу на 7070 порту")
	if err := s.Serve(lis); err != nil{
		log.Fatal(err)
	}
}