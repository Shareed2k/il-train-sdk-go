package api

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/go-querystring/query"
	jsoniter "github.com/json-iterator/go"

	"github.com/Shareed2k/il-train-sdk-go/client"
)

const BaseURL = "https://www.rail.co.il/apiinfo/api/Plan/"

var json = jsoniter.ConfigDefault

type (
	Api struct {
		Client
	}

	GetStationOutput struct {
		MessageType int64
		Message     string
		Data        *StationData `json:"Data"`
	}

	GetRoutesInput struct {
		// OId=4100&TId=4680&Date=20210726&Hour=1800
		Origin      int64  `url:"OId"`
		Dastination int64  `url:"TId"`
		Date        string `url:"Date"`
		Hour        string `url:"Hour"`
	}

	GetRoutesOutput struct {
		MessageType int64
		Message     string
		Data        *RouteData `json:"Data"`
	}

	Route struct {
		IsExchange bool
		EstTime    string
		Train      []*RouteTrain
	}

	RouteTrain struct {
		ArrivalTime, DepartureTime, LineNumber, TrainParvariBenironi             string
		DestinationStation, OrignStation, Platform, DestPlatform, Route, Trainno int64 `json:",string"`
		DirectTrain, Handicap, IsFullTrain, Midnight, ReservedSeat               bool
	}

	RouteData struct {
		BeforeRoutes, Error string
		StartIndex          int
		Details             *RouteDetail
		Routes              []*Route
	}

	RouteDetail struct {
		Destination, Origin int64 `json:",string"`
		Date, Hour, SugKav  string
	}

	StationData struct {
		CustomPropertys []*Station
	}

	Station struct {
		ID                 int64 `json:",string"`
		Heb, Rus, Eng, Arb []string
	}

	Client interface {
		Do(*http.Request) (*http.Response, error)
	}
)

func New(options ...func(*Api)) *Api {
	api := &Api{
		Client: client.New(),
	}

	for _, option := range options {
		option(api)
	}

	return api
}

func (a *Api) GetStationWithContext(ctx context.Context) (*GetStationOutput, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%sGetStations", BaseURL), nil)
	if err != nil {
		return nil, err
	}

	res, err := a.Client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(ioutil.NopCloser(res.Body))
	out := &GetStationOutput{}
	if err := decoder.Decode(out); err != nil {
		return nil, err
	}

	return out, nil
}

func (a *Api) GetRoutesWithContext(ctx context.Context, input *GetRoutesInput) (*GetRoutesOutput, error) {
	v, err := query.Values(input)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%sGetRoutes?%s", BaseURL, v.Encode()), nil)
	if err != nil {
		return nil, err
	}

	res, err := a.Client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(ioutil.NopCloser(res.Body))
	out := &GetRoutesOutput{}
	if err := decoder.Decode(out); err != nil {
		return nil, err
	}

	return out, nil
}

func WithClient(c Client) func(*Api) {
	return func(a *Api) {
		a.Client = c
	}
}
