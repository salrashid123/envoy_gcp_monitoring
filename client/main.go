package main

import (
	"fmt"
	"time"

	pubsub "cloud.google.com/go/pubsub"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
)

const (
	projectID = "your_project_id_here"
)

func main() {
	for {
		ctx := context.Background()
		pubsubClient, err := pubsub.NewClient(ctx, projectID)
		if err != nil {
			fmt.Printf("pubsub.NewClient: %v", err)
			return
		}
		defer pubsubClient.Close()

		pit := pubsubClient.Topics(ctx)
		for {
			topic, err := pit.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				fmt.Printf("pubssub.Iterating error: %v", err)
				return
			}
			fmt.Printf("Topic Name: %s\n", topic.ID())
		}
		time.Sleep(1000)
	}
}
