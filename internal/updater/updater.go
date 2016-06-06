package updater

import (
	"github.com/valexz/opsgbot/internal/config"
	"github.com/valexz/opsgbot/internal/opsg"
	"github.com/valexz/opsgbot/internal/slack"
	"sync"
	"time"
)

type Updater struct {
	Wg         *sync.WaitGroup
	Slack      *slack.Slack
	OpsG       *opsg.Opsg
	LastUpdate time.Time
}

func New() (*Updater, error) {
	u := Updater{}
	u.Wg = &sync.WaitGroup{}

	var err error
	u.Slack, err = slack.New(config.Config.ApiKeys.Slack)
	if err != nil {
		return &u, err
	}

	u.OpsG, err = opsg.New(config.Config.ApiKeys.Opsgenie)
	if err != nil {
		return &u, err
	}
	return &u, err

}
