package main

import (
	"log"

	"github.com/ysinjab/temporal-samples/workflows"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.NewClient(client.Options{HostPort: client.DefaultHostPort})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}

	defer c.Close()

	w := worker.New(c, "money-transfer-task-queue", worker.Options{})

	w.RegisterWorkflow(workflows.MoneyTransfer)
	w.RegisterActivity(workflows.Deposit)
	w.RegisterActivity(workflows.Withdraw)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}

}
