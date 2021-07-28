package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Shareed2k/il-train-sdk-go/api"
	"github.com/Shareed2k/il-train-sdk-go/client"
)

func main() {
	client := api.New(
		api.WithClient(
			client.New(
				client.WithLogger(log.New(os.Stderr, "", log.LstdFlags)),
			),
		),
	)

	// Time out if it takes more than 10 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Always call cancel.

	res, err := client.GetRoutesWithContext(ctx, &api.GetRoutesInput{
		Origin:      4100,                          // Bney Brak
		Dastination: 4680,                          // Yosef Tal
		Date:        time.Now().Format("20060102"), // "20210728",
		Hour:        time.Now().Format("1504"),     // "1600",
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, route := range res.Data.Routes {
		for _, train := range route.Train {
			if train.DirectTrain {
				fmt.Printf("From Station %d => To Station %d:\n\t A %s => D %s\n\n", train.OrignStation, train.DestinationStation, train.DepartureTime, train.ArrivalTime)
			}
		}
	}
}
