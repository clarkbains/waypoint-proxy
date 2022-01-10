package main

import (
	"fmt"
	"log"
	"sync"
	pb "github.com/clarkbains/waypoint-proxy/proto"
	"github.com/clarkbains/waypoint-proxy/routers/userRouter"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	var wg sync.WaitGroup
	
	go userrouter.Start(&wg, 2000)

	c := GetGrpcClient()

	job,err := c.Client.ListProjects(c.Context, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("fail to List Project: %v", err)
	}

	for _, p := range job.Projects {
		fmt.Println(p.Project)
		proj, err := c.Client.GetProject(c.Context, &pb.GetProjectRequest{Project: p})
		if err != nil {
			log.Fatalf("fail to get project: %s", p)
		}
		req := pb.GetProjectRequest{
			Project: p,
		}
		data, err := c.Client.GetProject(c.Context, &req)
		if err != nil {
			log.Fatalf("Failed to get project %s", proj.Project.Name)
		}
		
		for _, p := range data.Project.Variables {
			fmt.Printf("\t%s: %s\n", p.Name, p.GetStr())
		}
		
		
	}
	wg.Wait()
}

