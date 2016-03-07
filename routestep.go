package alur

import(
    "github.com/eaciit/toolkit"
)

type StepType string
type InsertType string

const (
	StepEntry    StepType = "Entry"
	StepAction   StepType = "Action"
	StepApproval StepType = "Approval"

    InsertAfter InsertType = "After"
    InsertBefore InsertType = "Before"
    InsertAt InsertType = "At"
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
    Approvals *ApprovalConfig
    Require []string
    Data interface{}
    
    Pre interface{}
    Exec interface{}
    Post interface{}
}
/*SetApprover Set approver on a step. 
Level is sequence level of approver need to be updated. If level number is bigger than available approvers, command will be ignored.
To insert or update approvers, please use InsertApprover instead
*/
func (s *RouteStep) SetApprover(level int, id ...string) *RouteStep{
    return s.InsertApprover(level, InsertAt, id...)
}

func (s *RouteStep) InsertApprover(level int, insertType InsertType, id ...string) *RouteStep{
    return s
}

func (s *RouteStep) SetMinimalApprover(level int, minimalApproverRequired int) *RouteStep{
    if a:=s.getApprovalByLevel(level); a!=nil{
        a.MinimalApprovers = minimalApproverRequired
    }
    return s
}

func (s *RouteStep) SetApproverById(title string, id ...string) *RouteStep{
    return s.InsertApproverById(title, InsertAt, id...)
}

func (s *RouteStep) InsertApproverById(title string, insertType InsertType, id ...string) *RouteStep{
    return s
}

func (s *RouteStep) SetMinimalApproveryId(title string, minimalApproverRequired int) *RouteStep{
    if a:=s.getApprovalById(title); a!=nil{
        a.MinimalApprovers=minimalApproverRequired
    }
    return s
}

func (s *RouteStep) getApprovalByLevel(level int) *ApprovalLevel{
    return nil
}

func (s *RouteStep) getApprovalById(title string) *ApprovalLevel{
    return nil
}
