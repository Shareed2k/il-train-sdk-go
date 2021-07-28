package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/shareed2k/il-train-sdk-go/api"
	"github.com/shareed2k/il-train-sdk-go/client"
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

	res, err := client.GetStationWithContext(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, data := range res.Data.CustomPropertys {
		fmt.Println(data.ID, data.Eng)
	}
}
