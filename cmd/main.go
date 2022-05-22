package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/afikrim/protobuf-tutorial/blog"
	handlerBlog "github.com/afikrim/protobuf-tutorial/handlers/blog"
	pbBlog "github.com/afikrim/protobuf-tutorial/pb/blog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/protobuf_tutorial?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Gorm failed:", err)
	}

	db.AutoMigrate(&blog.Blog{})

	blogRepo := blog.NewRepository(db, "blogs")

	blogSvc := blog.NewService(blogRepo)

	blogHandler := handlerBlog.NewHandler(blogSvc)

	// Create a listener on TCP Port
	lis, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Blog service to the server
	pbBlog.RegisterBlogServiceServer(s, blogHandler)
	// Server gRPC server
	log.Println("gRPC server is running on port 8080")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalln("Failed to dial:", err)
	}

	gwmux := runtime.NewServeMux()
	err = pbBlog.RegisterBlogServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Println("gRPC-Gateway server is running on port 8090")
	log.Fatalln(gwServer.ListenAndServe())
}
