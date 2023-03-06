package main

import (
	"context"
	"fmt"
	"log"
	"net"

	livescore "grpcscore2server/pb"

	"google.golang.org/grpc"
)

var matches livescore.List2MatchesResponse

type liveScoreServer struct {
	livescore.UnimplementedScore2ServiceServer
}

func (lss *liveScoreServer) List2Matches(ctx context.Context, req *livescore.List2MatchesRequest) (*livescore.List2MatchesResponse, error) {
	println("Connect List")
	match := &livescore.Match2ScoreResponse{
		Score: "2:2",
		Live:  true,
	}

	matches.Scores = matches.Scores[:0]
	matches.Scores = append(matches.Scores, match)
	return &matches, nil
}

const addr = ":50005"

func main() {
	//create tcp connection on port 50005
	conn, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatal("tcp connection err: ", err.Error())
	}
	//create grpc server
	grpcServer := grpc.NewServer()

	server := liveScoreServer{}
	//register our livescore service
	livescore.RegisterScore2ServiceServer(grpcServer, &server)

	fmt.Println("Starting gRPC server at : ", addr)
	//serve our connection
	if err := grpcServer.Serve(conn); err != nil {
		log.Fatal(err)
	}
}
