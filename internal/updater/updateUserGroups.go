package updater

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/valexz/opsgbot/internal/config"
)

// Fetch the users from Pagerduty and slack, and make sure we can match them
// all up. We match Pagerduty users to Slack users based on their email address
func (u *Updater) updateUserGroups() {
	for _, schedule := range config.Config.Schedules {

		var err error
		opsgUsers, err := u.OpsG.GetOnCalls(schedule.Name)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Warning("Failed fetching from  opsgenie on-call user for schedule %s", schedule.Name)
			return
		}
		log.WithFields(log.Fields{
			"users": opsgUsers,
		}).Debug("Fetched opsgenie on-call user for schedule %s", schedule.Name)

		for _, opsgUser := range opsgUsers.OnCallParticipants {

			slackUser, err := u.Slack.GetUserByEmail(opsgUser.Name)
			if err != nil {
				log.Errorf("Failed fetching Slack user via e-mail: %v", err)
			}

			msgText := fmt.Sprintf(schedule.UpdateMessage.Message, "@"+slackUser.Name)

			isCurrentGroupMember, err := u.Slack.IsUserMemberOfGroup(slackUser, schedule.SlackUserGroup)
			if !isCurrentGroupMember {
				err = u.Slack.UpdateUserGroup(schedule.SlackUserGroup, slackUser.ID)
				if err != nil {
					log.WithFields(log.Fields{
						"error":     err,
						"user":      slackUser.Name,
						"schedule":  schedule.Name,
						"userGroup": schedule.SlackUserGroup,
					}).Error("Failed to update Slack user group")
				}
				err = u.Slack.PostMessage(schedule.UpdateMessage.Channels[0], msgText)
				if err != nil {
					log.WithFields(log.Fields{
						"error": err,
					}).Error("Failed to post user group update message")
				}
			}

		}

	}

}
