package main

import (
	"context"
	"log"

	"github.com/ysinjab/temporal-samples/workflows"

	"go.temporal.io/sdk/client"
)

func main() {
	// The client is a heavyweight object that should be created once per process.
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:        "MoneyTransfer__workflowID",
		TaskQueue: "money-transfer-task-queue",
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, workflows.MoneyTransfer, "1", "2", "N7635", 552.8)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	// Synchronously wait for the workflow completion.
	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}
	log.Println("Workflow result:", result)
}
