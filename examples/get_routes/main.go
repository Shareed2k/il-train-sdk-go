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

	// get routes list
	res, err := client.GetRoutesWithContext(ctx, &api.GetRoutesInput{
		Origin:      4100,                          // Bney Brak
		Dastination: 4680,                          // Yosef Tal
		Date:        time.Now().Format("20060102"), // "20210728",
		Hour:        time.Now().Format("1504"),     // "1600",
	})
	if err != nil {
		log.Fatal(err)
	}

	// get stations list
	resStations, err := client.GetStationWithContext(ctx)
	if err != nil {
		log.Fatal(err)
	}

	stations := make(map[int64]string, len(resStations.Data.CustomPropertys))
	for _, station := range resStations.Data.CustomPropertys {
		stations[station.ID] = station.Eng[0]
	}

	for _, route := range res.Data.Routes {
		for _, train := range route.Train {
			if train.DirectTrain {
				fmt.Printf("From Station `%s` => To Station `%s`:\n\t A %s => D %s\n\n", stations[train.OrignStation], stations[train.DestinationStation], train.DepartureTime, train.ArrivalTime)
			}
		}
	}
}
