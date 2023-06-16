# OpsgBot

Update your Slack user groups based on your Opsgenie Schedules.

Based on https://github.com/karlkfi/pagerbot,
which is a fork of https://github.com/goodeggs/pagerbot,
which is a fork of https://github.com/YoSmudge/pagerbot.

Provided with API credentials and some configuration, OpsgBot will
automatically update Slack user group membership and post a message to channels
you select informing everyone who's currently on the rotation.

OpsgBot matches Opsgenie users to Slack users by their email addresses,
so your users must have the same email address in Slack as in Opsgenie.
OpsgBot will log warnings for any users it finds in Opsgenie but not in Slack.

# Docker Image
https://hub.docker.com/r/valexz/opsgbot/tags

You MUST mount configuration file to container

# Local Build

Use [goenv](https://github.com/syndbg/goenv) to install dependencies:

`goenv local`

Compile the `opsgbot` binary:

`go build`

You should have a nice `opsbot` binary ready to go.


# Slack Setup

1. [Create a Slack App](https://api.slack.com/apps)
2. Configure the App Scopes under `OAuth & Permissions`:
    - Send messages as opsgbot (chat:write:bot)
    - Post to specific channels in Slack (incoming-webhook)
    - Access basic information about the workspace’s User Groups (usergroups:read)
    - Change user’s User Groups (usergroups:write)
    - Access your workspace’s profile information (users:read)
    - View email addresses of people on this workspace (users:read.email)
3. Install the App and copy the `OAuth Access Token` (requires workspace admin)
4. Save the token `echo "SLACK_TOKEN=<token>" >> .secrets.env`

# Opsgenie Setup

1. [Create a read-only Opsgenie API Key](https://support.atlassian.com/opsgenie/docs/api-key-management/)
2. Enable configuration access
3. Save the key `echo "OPSGENIE_TOKEN=<key>" >> .secrets.env`

# Config

A basic configuration file will look like

```yaml
# Secrets in cluster
api_keys:
   slack: "${SLACK_TOKEN}"
   opsgenie: "${OPSGENIE_TOKEN}"

# It changes responsible users according to desired `schedules` for each Slack group `name` (e.g. @mygroup)
# Fires `update_message` to `channels` on change

schedules: # one or more Opsgenie schedules
   - name: "MySchedule" # Opsgenie schedule name
     slackUserGroup: "myusergroup" # name of the Slack user group to update
     updateMessage:
        # 1st parameter is slack user group, the second parameter is slack username
        message: " On duty for L1: %s"
        channels:
           - "my-slack-channel"
```

This config specifies the use of environment variables which opsgbot will
interpolate at runtime, allowing you to inject secrets and env vars.

Specify the config when launching opsgbot:

```
./opsgbot --config /path/to/config.yml --env-file /path/to/.secrets.env
```

# Deploy

It's recommended to run OpsgBot using a package manager or container platform.

## Docker Image

The included Dockerfile uses a multi-stage build to compile opsgbot and package
it in a small alpine-based Docker image for use in production:

```
# build and tag image
docker build -t valexz/opsgbot:latest .

# run locally in docker
docker run --env-file .secrets.env valexz/opsgbot:latest

# push to a docker image registry
docker push valexz/opsgbot:latest
```
