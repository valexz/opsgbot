package opsg

import "github.com/opsgenie/opsgenie-go-sdk-v2/schedule"

func (a *Opsg) GetOnCalls(scheduleName string) (*schedule.GetOnCallsResult, error) {
	flat := false

	onCallsList, err := a.Client.GetOnCalls(nil, &schedule.GetOnCallsRequest{
		Flat:                   &flat,
		ScheduleIdentifierType: schedule.Name,
		ScheduleIdentifier:     scheduleName,
	})

	return onCallsList, err

}
