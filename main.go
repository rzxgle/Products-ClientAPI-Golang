package main

import (
	"apiproducts_client/src/pb/products"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("error on get client. error: ", err)
	}
	defer conn.Close()
	//createProduct(conn)
	FindAllProducts(conn)
}

func createProduct(conn *grpc.ClientConn) {
	newProduct := &products.Product{
		Name:        "Frete Platinum Plus",
		Description: "Produto de impulsionamento de anúncios",
		Price:       1500,
		Quantity:    100,
	}

	productClient := products.NewProductServiceClient(conn)
	newProduct, err := productClient.Create(context.Background(), newProduct)
	if err != nil {
		log.Fatalln("error on create product. error: ", err)
	}

	fmt.Printf("Produto criado. \n %+v\n", *newProduct)
}

func FindAllProducts(conn *grpc.ClientConn) {
	productClient := products.NewProductServiceClient(conn)
	productList, err := productClient.FindAll(context.Background(), &products.Product{})
	if err != nil {
		log.Fatalln("error on list products. error: ", err)
	}

	fmt.Printf("products: %+v\n", productList)
}
