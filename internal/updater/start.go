package updater

import (
	log "github.com/sirupsen/logrus"
	"time"
)

// Start the updater process
func (u *Updater) Start() {
	u.Wg.Add(1)
	go u.run()
}

// Loop for updater
// Will call for new data then call the update function
// Runs on each `updateEvery` interval
const updateEvery time.Duration = time.Minute * 5

func (u *Updater) run() {
	defer u.Wg.Done()

	for {
		start := time.Now().UTC()
		log.Debug("Update starting...")

		u.updateUserGroups()

		u.LastUpdate = time.Now().UTC()

		// slack groups depend on users & schedules
		end := time.Now().UTC()

		log.WithFields(log.Fields{
			"duration": end.Sub(start),
		}).Info("Update complete")

		time.Sleep(updateEvery)
	}
}
