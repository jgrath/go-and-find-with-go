package util

import "time"

type SystemProperty struct {
	Name             string    `json:"property-name"`
	Value            string    `json:"property-value"`
	DefaultValue     string    `json:"DefaultValue"`
	Description      string    `json:"Description"`
	DataType         string    `json:"DataType"`
	Enabled          bool      `json:"Enabled"`
	ActiveFromDate   time.Time `json:"ActiveFromDate"`
	GroupCode        string    `json:"GroupCode"`
	GroupName        string    `json:"GroupName"`
	GroupDescription string    `json:"GroupDescription"`
}
