package signal

import (
	"context"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

// Workflow is a sample workflow definition.
func Workflow(ctx workflow.Context, name string) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("Signal workflow started", "name", name)

	// Wait to receive signal.
	workflow.Sleep(ctx, time.Second*10)

	var signalVal string = "World"
	signalChan := workflow.GetSignalChannel(ctx, "your-signal-name")
	signalChan.ReceiveAsync(&signalVal)

	var result string
	err := workflow.ExecuteActivity(ctx, Activity, signalVal).Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}

	logger.Info("Signal workflow completed.", "result", result)

	return result, nil
}

func Activity(ctx context.Context, name string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Activity", "name", name)
	return "Hello " + name + "!", nil
}
