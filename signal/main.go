package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"

	"go.temporal.io/sdk/client"

	ts "github.com/pvsune/temporal-signal"
)

func main() {
	var signal string
	flag.StringVar(&signal, "s", "{}", "Signal data.")
	flag.Parse()

	// The client is a heavyweight object that should be created once per process.
	c, err := client.NewClient(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	var signalVal = &ts.MySignal{}
	if err := json.Unmarshal([]byte(signal), signalVal); err != nil {
		log.Fatalln("Unable to unmarshal signal value", err)
	}
	err = c.SignalWorkflow(context.Background(), "temporal_signal_workflowID", "", "your-signal-name", signalVal)
	if err != nil {
		log.Fatalln("Unable to signal workflow", err)
	}
}
