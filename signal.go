package signal

import (
	"context"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

type MySignal struct {
	Type           string `json:"type"`
	Message        string `json:"message"`
	ExtendDuration int64  `json:"duration"`
}

// Workflow is a sample workflow definition.
func Workflow(ctx workflow.Context, name string) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("Signal workflow started", "name", name)

	signalVal, err := waitForSignal(ctx, 10)
	if err != nil {
		return "", err
	}

	var result string
	err = workflow.ExecuteActivity(ctx, Activity, signalVal).Get(ctx, &result)
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

func waitForSignal(ctx workflow.Context, extend time.Duration) (string, error) {
	var signalVal MySignal
	signalChan := workflow.GetSignalChannel(ctx, "your-signal-name")
	selector := workflow.NewSelector(ctx)
	selector.AddReceive(signalChan, func(channel workflow.ReceiveChannel, more bool) {
		channel.Receive(ctx, &signalVal)
	})
	timerFuture := workflow.NewTimer(ctx, time.Second*extend)
	selector.AddFuture(timerFuture, func(future workflow.Future) {
		signalVal.Message = "World"
	})
	selector.Select(ctx)

	switch signalVal.Type {
	case "cancel":
		return "", temporal.NewCanceledError()
	case "extend":
		return waitForSignal(ctx, time.Duration(signalVal.ExtendDuration))
	}
	return signalVal.Message, nil
}
