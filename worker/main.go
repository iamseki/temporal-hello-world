package main

import (
	"log"

	app "github.com/iamseki/temporal-tutorial"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	// default values will be used: localhost:7233 within default namespace
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create temporal client", err)
	}
	defer c.Close()

	w := worker.New(c, app.GreetingTaskQueue, worker.Options{})
	w.RegisterWorkflow(app.GreetingWorkflow)
	w.RegisterActivity(app.ComposeGreeting)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start worker", err)
	}
}
