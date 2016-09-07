package dto

const JiraDateFormat = "2006-01-02T15:04:05.000-0700";

type WorkLogRequest struct {
	Comment          string `json:"comment"`
	Started          string `json:"started"`
	TimeSpentSeconds int `json:"timeSpentSeconds"`
}