package opsg

import (
	ctx "context"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/opsgenie/opsgenie-go-sdk-v2/schedule"
	log "github.com/sirupsen/logrus"
)

type Opsg struct {
	Client *schedule.Client
	config *client.Config
}

func New(apikey string) (*Opsg, error) {
	var err error

	a := Opsg{}
	a.config = &client.Config{
		ApiKey: apikey,
	}

	a.Client, err = schedule.NewClient(a.config)

	expandListRequest := false

	//Check authentication with token
	_, err = a.Client.List(ctx.Background(), &schedule.ListRequest{
		Expand: &expandListRequest,
	})

	log.Info("Authenticated with OpsGenie")

	return &a, err

}
