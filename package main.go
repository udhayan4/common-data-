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



import { NextResponse } from 'next/server'  
import { getToken } from 'next-auth/jwt'  

export async function middleware(req: NextRequest) {  
  const token = await getToken({ req, secret: process.env.JWT_SECRET })  
  if (!token || token.role !== 'admin') {  
    return NextResponse.redirect('/unauthorized')  
  }  
  return NextResponse.next()  
}  



# grafana.tf  
resource "grafana_dashboard" "gen3_health" {  
  config_json = templatefile("${path.module}/dashboards/gen3.json", {  
    prometheus_ds = var.prometheus_datasource  
  })  
}  

resource "grafana_data_source" "prometheus" {  
  type = "prometheus"  
  name = "Gen3 Prometheus"  
  url  = "http://prometheus:9090"  
}  