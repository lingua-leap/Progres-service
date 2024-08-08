package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	pb "progress-service/generated/progress"
	"progress-service/logs"
	"progress-service/service"
	"progress-service/storage"
	"progress-service/storage/mongodb"
)

func main() {
	logger := logs.InitLogger()

	db, err := mongodb.ConnectMongoDB()
	if err != nil {
		logger.Error("error connecting to mongodb", "error", err)
		log.Fatal(err)
	}

	listner, err := net.Listen("tcp", ":50052")
	if err != nil {
		logger.Error("error listening on tcp port", "error", err)
		log.Fatal(err)
	}

	server := grpc.NewServer()

	progress := service.NewService(logger, storage.NewStorage(db))

	pb.RegisterProgressServiceServer(server, progress)

	logger.Info("starting server", "port", 50052)
	log.Printf("server listening on port %d", 50052)
	log.Fatal(server.Serve(listner))

}
