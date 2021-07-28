# Unofficial Israel Rail SDK

This is a unofficial wrapping of the API of Israeli Rail  in Golang.

Use this SDK for checking your own train schedule, integrating with Alexa, and so on!

## Installing

```sh
go get github.com/Shareed2k/il-train-sdk-go
```

## Usage
```golang
func main() {
	client := api.New()

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

```

```sh
From Station 4100 => To Station 4680:
         A 28/07/2021 17:43:00 => D 28/07/2021 18:26:00

From Station 4100 => To Station 4680:
         A 28/07/2021 18:13:00 => D 28/07/2021 18:56:00

From Station 4100 => To Station 4680:
         A 28/07/2021 19:13:00 => D 28/07/2021 19:56:00
```

## Contributing
We'd love your contributions! Fire up a (tested!) Pull request and we'll be happy to approve it.