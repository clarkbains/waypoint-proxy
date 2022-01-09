package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"os"

	pb "github.com/clarkbains/waypoint-proxy/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)
type WaypointClient struct {
	Client pb.WaypointClient
	connection grpc.ClientConn
	Context context.Context
}
func (w WaypointClient) Close () {
	w.connection.Close()
	w.Context.Done()
}

//Get a Waypoint GRPC Client
func GetGrpcClient () (WaypointClient) {
	ctx := context.TODO()
	token := "foo"
	if v, p := os.LookupEnv("WAYPOINT_TOKEN"); p {
		token = v
	}

	 port := 443
	 if v, p := os.LookupEnv("WAYPOINT_PORT"); p {
	 	convPort, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalln("Cannot convert Port to integer")
		}
		port = convPort
	 }

	host := "waypoint-api.cwdc.cbains.ca"
	if v, p := os.LookupEnv("WAYPOINT_HOST"); p {
		host = v
	}


	creds, err := credentials.NewClientTLSFromFile("root.pem", host)
	if err != nil {
		log.Fatalf("Failed to create TLS credentials %v", err)
	}

	connectionStr := fmt.Sprintf("%s:%d", host, port)
	log.Printf("Dialing %s\n", connectionStr)

	conn, err := grpc.Dial(connectionStr, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Failed to Dial: %v", err)
	}


	ctx = metadata.AppendToOutgoingContext(ctx, "client-api-protocol", "1,1", "authorization", token)

	wp := WaypointClient{
		Client: pb.NewWaypointClient(conn),
		Context: ctx,
		connection: *conn,
	}
	return wp

}
