package main

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"

	ts "github.com/pvsune/temporal-signal"
)

func main() {
	// The client is a heavyweight object that should be created once per process.
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:        "temporal_signal_workflowID",
		TaskQueue: "temporal-signal",
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, ts.Workflow, "Temporal")
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	// Synchronously wait for the workflow completion.
	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		if temporal.IsCanceledError(err) {
			log.Println("Workflow cancelled")
			return
		}
		log.Fatalln("Unable get workflow result", err)
	}
	log.Println("Workflow result:", result)
}
