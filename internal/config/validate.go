package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

// Validate the configuration file for sanity
func (c *config) Validate() error {
	if c.ApiKeys.Slack == "" || c.ApiKeys.Opsgenie == "" {
		return fmt.Errorf("You must provide API keys for Slack and Opsgenie")
	}

	if len(c.Schedules) == 0 {
		return fmt.Errorf("You must specify at least one schedule")
	}

	for i, schedule := range c.Schedules {
		if schedule.Name == "" {
			return fmt.Errorf("Must specify schedule name for Schedule %d", i)

		} else if schedule.SlackUserGroup == "" {
			return fmt.Errorf("Must specify group name for slack user group of schedule %s", schedule.Name)

		}
	}

	log.WithFields(log.Fields{"schedules": len(c.Schedules)}).Debug("Loaded config")

	return nil
}
