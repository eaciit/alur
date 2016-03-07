package alur

import(
    "github.com/eaciit/toolkit"
)

type StepType string

const (
	StepEntry    StepType = "Entry"
	StepAction   StepType = "Action"
	StepApproval StepType = "Approval"
)

type IStep interface{
    Config() toolkit.M
}

type IRequestStep interface{
    Config() toolkit.M
    Run() error
    Request(toolkit.M) error
    Approve(string, toolkit.M) error
    Reject(string, string, toolkit.M) error
}

type RouteStep struct {
	RouteID  string
	StepID   string
	Title    string
	StepType StepType
}
