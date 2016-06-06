package slack

import (
	log "github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
)

type Slack struct {
	Client *slack.Client
}

func New(token string) (*Slack, error) {

	a := &Slack{

		Client: slack.New(token),
	}

	_, err := a.Client.GetUsers()
	if err != nil {
		log.Panicf("Auth check failed. Error getting users via Slack API: %v", err)
	}
	return a, err
}

func (a *Slack) GetUserByEmail(email string) (*slack.User, error) {
	user, err := a.Client.GetUserByEmail(email)
	return user, err

}

func (a *Slack) UpdateUserGroup(userGroupName string, user string) error {
	userGroupID, err := a.getUserGroupIdByName(userGroupName)
	_, err = a.Client.UpdateUserGroupMembers(userGroupID, user)
	return err

}

func (a *Slack) getUserGroupIdByName(userGroupName string) (string, error) {
	userGroups, err := a.Client.GetUserGroups(slack.GetUserGroupsOptionIncludeUsers(true))
	if err != nil {
		return "", err
	}

	for _, group := range userGroups {
		if group.Name == userGroupName {
			return group.ID, nil
		}
	}

	return "", nil
}

func (a *Slack) getUserGroupMembers(userGroupName string) ([]string, error) {
	var users []string
	userGroupID, err := a.getUserGroupIdByName(userGroupName)
	if err != nil {
		return users, err
	}

	m, err := a.Client.GetUserGroupMembers(userGroupID)
	if err != nil {
		return users, err
	}

	for _, id := range m {
		users = append(users, id)
	}
	return users, nil
}

func (a *Slack) IsUserMemberOfGroup(user *slack.User, userGroupName string) (bool, error) {

	groupMembers, err := a.getUserGroupMembers(userGroupName)
	if err != nil {
		log.Errorf("Slack user %s not found", user.Name)
	}

	for _, memberID := range groupMembers {
		if memberID == user.ID {
			return true, err
		}
	}
	return false, err
}

func (a *Slack) PostMessage(channel string, message string) error {
	msgParams := slack.NewPostMessageParameters()
	msgParams.AsUser = true
	msgParams.LinkNames = 1

	_, _, err := a.Client.PostMessage(
		channel,
		slack.MsgOptionText(message, false),
		slack.MsgOptionPostMessageParameters(msgParams),
	)
	return err
}
