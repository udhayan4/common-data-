package main  

import (  
  "context"  
  "log"  
  "net"  

  pb "github.com/gen3/csoc/proto"  
  "google.golang.org/grpc"  
  "google.golang.org/grpc/credentials"  
)  

type AgentServer struct {  
  pb.UnimplementedAgentServiceServer  
}  

func (s *AgentServer) Register(ctx context.Context, req *pb.AgentInfo) (*pb.Response, error) {  
  log.Printf("Agent %s registered from %s", req.Id, req.Ip)  
  return &pb.Response{Status: "ACK"}, nil  
}  

func main() {  
  creds, _ := credentials.NewServerTLSFromFile("server.crt", "server.key")  
  s := grpc.NewServer(grpc.Creds(creds))  
  pb.RegisterAgentServiceServer(s, &AgentServer{})  
  lis, _ := net.Listen("tcp", ":50051")  
  s.Serve(lis)  
}  