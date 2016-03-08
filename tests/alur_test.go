package tests

import (
	"testing"
	"time"

	"github.com/eaciit/alur"
	"github.com/eaciit/toolkit"
)

var (
	r *alur.Route
)

func SkipIfNil(t *testing.T) {
	if r == nil {
		t.Skip()
	}
}

func TestRoute(t *testing.T) {
	r := new(alur.Route)
	r.ID = "wf_simple"
	r.Title = "Simple Workflow"
	r.Enable = true
}

func TestManageStep(t *testing.T) {
	SkipIfNil(t)

	s1 := new(alur.RouteStep)
	s1.StepID = "Start"
	s1.StepType = alur.StepEntry

	s2 := new(alur.RouteStep)
	s2.Require = []string{"Start"}
	s2.StepID = "Validate"
	s2.StepType = alur.StepAction
	s2.Pre = func(ctx *alur.Context) {
		m := ctx.Step.Data.(toolkit.M)
		if m.GetInt("leaveday") <= 0 {
			ctx.Request.Stop(alur.RequestCancelled, "Admin", "Leave Day can't be 0")
		}
	}
	s2.Exec = func(ctx *alur.Context) {
		m := ctx.Step.Data.(toolkit.M)
		dept := m.GetString("department")
		if dept == "it" {
			ctx.Request.StepById("approval").SetApprover(0, "andiek.suncahyo").SetMinimalApprover(0, 1)
		}
	}

	s3 := new(alur.RouteStep)
	s3.Require = []string{"Validate"}
	s3.StepID = "Approval"
	s3.StepType = alur.StepApproval

	s4 := new(alur.RouteStep)
	s4.Require = []string{"Approval"}
	s4.StepID = "Close"
	s4.StepType = alur.StepAction
	s4.Exec = func(ctx *alur.Context) {
		toolkit.Println("Request has been approved")
	}

	s5 := new(alur.RouteStep)
	s5.StepID = "Rejection"
	s5.RequireReject = []string{"Approval"}
	s5.StepType = alur.StepAction
	s5.Exec = func(ctx *alur.Context) {
		toolkit.Println("Request has been rejected")
	}

	r.UpdateSteps([]*alur.RouteStep{s1, s2, s3, s4, s5})
	e := r.Verify()
	if e != nil {
		t.Fatalf("Error verify: %s", e.Error())
	}
}

func TestRequestApprove(t *testing.T) {
	q := alur.NewRequest(r, "user")
	q.Data().Set("leaveday",10).Set("department","Finance")
    toolkit.Println("Data: ", q.Data())
    
    e := q.Start()
    if e!=nil {
        t.Fatal("Fail to start: ", e.Error())    
    }

	t0 := time.Now()
    isOnApproval := false
	for !isOnApproval {
		for _, s := range q.CurrentSteps {
			if s.StepID == "Approval" {
				isOnApproval = true
				s.ApproveReject("admin", alur.Approve, "", nil)
			} else {
				time.Sleep(1 * time.Millisecond)
				if d:=time.Since(t0); d > (time.Duration)(5 * time.Second) {
					t.Fatalf("Error, timeout")
					return
				}
			}
		}
	}
}

func TestRequestReject(t *testing.T) {
	q := alur.NewRequest(r, "user")
	q.Data().Set("leaveday",10).Set("department","Finance")
    toolkit.Println("Data: ", q.Data())
    
    e := q.Start()
    if e!=nil {
        t.Fatal("Fail to start: ", e.Error())    
    }
   
	t0 := time.Now()
    isOnApproval := false
	for !isOnApproval {
		for _, s := range q.CurrentSteps {
			if s.StepID == "Approval" {
				isOnApproval = true
				s.ApproveReject("admin", alur.Reject, "just test", nil)
			} else {
				time.Sleep(1 * time.Millisecond)
				if d:=time.Since(t0); d > (time.Duration)(5 * time.Second) {
					t.Fatalf("Error, timeout")
					return
				}
			}
		}
	}
}
