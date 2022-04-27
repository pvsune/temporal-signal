package main

import (
	"context"
	"flag"
	"log"

	"go.temporal.io/sdk/client"
)

func main() {
	var workflowID, signal string
	flag.StringVar(&workflowID, "w", "temporal_signal_workflowID", "WorkflowID.")
	flag.StringVar(&signal, "s", "World", "Signal data.")
	flag.Parse()

	// The client is a heavyweight object that should be created once per process.
	c, err := client.NewClient(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	err = c.SignalWorkflow(context.Background(), workflowID, "", "your-signal-name", signal)
	if err != nil {
		log.Fatalln("Unable to signal workflow", err)
	}
}
