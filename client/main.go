package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "myfirstserviceapp/proto"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func GetInt() (int64, error) {
	var a int64

	fmt.Printf("Введите число: ")
	_, err := fmt.Scan(&a) 
	if err != nil{
		return 0, err
	}
	return a, nil
}

func main(){
	x, err := GetInt()

	if err != nil{
		log.Fatal(err)
	}

	y, err := GetInt()

	if err != nil{
		log.Fatal(err)
	}

	conn, err := grpc.Dial("localhost:7060", grpc.WithInsecure())

	if err != nil{
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewMathServiceClient(conn)

	ctx, cansel := context.WithTimeout(context.Background(), time.Second)
	defer cansel()

	sumresp, err  := client.Sum(ctx, &pb.SumRequest{X:x, Y:y})

	if err != nil{
		log.Fatal(err)
	}

	minusresp, err := client.Minus(ctx, &pb.SumRequest{X:x, Y:y})

	if err != nil{
		log.Fatal(err)
	}
	
	powresp, err := client.Pow(ctx, &pb.SumRequest{X:x, Y:y})

	if err != nil{
		log.Fatal(err)
	}

	logrus.Info("Ответ от сервера: \n", sumresp.Result, "\n", minusresp.Result, "\n", powresp.Result)
}