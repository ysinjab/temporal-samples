package workflows

import (
	"context"
	"fmt"
	"go.temporal.io/sdk/temporal"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

func MoneyTransfer(ctx workflow.Context, fromAccountId string, toAccountId string, referenceId string, amount float64) error {
	rp := &temporal.RetryPolicy{
		InitialInterval:    1 * time.Second,
		BackoffCoefficient: 2,
		MaximumInterval:    100 * time.Second,
		MaximumAttempts:    500,
	}
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
		RetryPolicy:         rp,
	}

	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("MoneyTransfer started")

	err := workflow.ExecuteActivity(ctx, Withdraw, fromAccountId, referenceId, amount).Get(ctx, nil)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return err
	}
	err = workflow.ExecuteActivity(ctx, Deposit, toAccountId, referenceId, amount).Get(ctx, nil)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return err
	}
	logger.Info("MoneyTransfer completed.")
	return nil
}

func Deposit(ctx context.Context, accountId string, referenceId string, amount float64) error {
	logger := activity.GetLogger(ctx)
	logger.Info(fmt.Sprintf("Depositing $%f into account %s. ReferenceId: %s\n", amount, accountId, referenceId))
	return nil
}

func Withdraw(ctx context.Context, accountId string, referenceId string, amount float64) error {
	logger := activity.GetLogger(ctx)
	logger.Info(fmt.Sprintf("nWithdrawing $%f into account %s. ReferenceId: %s\n", amount, accountId, referenceId))
	return nil
}
