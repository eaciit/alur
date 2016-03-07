package alur

import(
    //"time"
)

type ApprovalConfig struct{
    ID string
    Enable bool
    Approvals []ApprovalLevel
}

type ApprovalLevel struct{
    Approvers []Approver
    Title string
    MinimalApprovers int
    //Timeout *time.Duration
}

type Approver struct{
    UserID string
    Email string
}