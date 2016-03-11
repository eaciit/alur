package tests

import (
	"testing"
	"time"
	//"bufio"
	"os"
    "errors"

    _ "github.com/eaciit/dbox/dbc/jsons"
	"github.com/eaciit/dbox"
    "github.com/eaciit/alur"
	"github.com/eaciit/toolkit"
    "path/filepath"
)

var (
	r *alur.Route
)

func SkipIfNil(t *testing.T) {
	if r == nil {
		t.Skip()
	}
}

func prepareConnection() error{
    wd, e := os.Getwd()
    if e!=nil {
        return e
    }
    
    wd = filepath.Join(wd, "..", "data")   
    if fi, e := os.Stat(wd); e!=nil {
        return e
    } else if !fi.IsDir(){
        return errors.New(toolkit.Sprintf("%s is not directory", wd))
    }
    
    conn, _ := dbox.NewConnection("jsons", &dbox.ConnectionInfo{Host:wd,})
    e = conn.Connect()
    if e!=nil {
        return errors.New(toolkit.Sprintf("Unable to connect: %s", e.Error()))
    }
    
    alur.SetDb(conn)
    return nil
}

func TestRoute(t *testing.T) {
	r = new(alur.Route)
	r.ID = "wf_simple"
	r.Title = "Simple Workflow"
	r.Enable = true
    
    e := prepareConnection()
    if e!=nil {
        t.Fatalf("Error connection: %s", e.Error())
    }
    
    e = alur.Db().Save(r)
    if e!=nil {
        t.Fatalf("Error save: %s", e.Error())
    }
}

func TestManageStep(t *testing.T) {
	SkipIfNil(t)

	s1 := new(alur.RouteStep)
	s1.StepID = "Start"
    s1.AutoStart = true
	s1.StepType = alur.StepEntry

	s2 := new(alur.RouteStep)
	s2.Require = []string{"Start"}
	s2.StepID = "Validate"
	s2.StepType = alur.StepAction
	/*
    s2.Pre = func(ctx *alur.Context) {
		if ctx.Request.Data().GetInt("leaveday") <= 0 {
			ctx.Request.Stop(alur.RequestCancelled, "Admin", "Leave Day can't be 0")
		}
	}
    
	s2.Exec = func(ctx *alur.Context) {
		dept := ctx.Request.Data().GetString("department")
		if dept == "it" {
			ctx.Request.StepById("approval").SetApprover(0, "andiek.suncahyo").SetMinimalApprover(0, 1)
		}
	}
    */

	s3 := new(alur.RouteStep)
	s3.Require = []string{"Validate"}
	s3.StepID = "Approval"
	s3.StepType = alur.StepApproval

	s4 := new(alur.RouteStep)
	s4.Require = []string{"Approval"}
	s4.StepID = "Close"
	s4.StepType = alur.StepAction
	/*
    s4.Exec = func(ctx *alur.Context) {
		toolkit.Println("Request has been approved")
	}
    */

	s5 := new(alur.RouteStep)
	s5.StepID = "Rejection"
	s5.RequireReject = []string{"Approval"}
	s5.StepType = alur.StepAction
	/*
    s5.Exec = func(ctx *alur.Context) {
		toolkit.Println("Request has been rejected")
	}
    */

	r.UpdateSteps([]*alur.RouteStep{s1, s2, s3, s4, s5})
	e := r.Verify()
	if e != nil {
		t.Fatalf("Error verify: %s", e.Error())
	}
    
    esave := alur.Db().Save(r)
    if esave!=nil {
        t.Fatalf("Error save: %s", esave.Error())
    }
}

func TestRequestApprove(t *testing.T) {
	//--- Get data
	user := "Arief"
	if user == "" {
		t.Fatalf("User can't be empty")
	}

	leaveDay := 5
	if leaveDay == 0 {
		t.Fatalf("Leave day should be > 0")
	}

	q := alur.NewRequest(r, user)
	q.Data().Set("leaveday", leaveDay).Set("verbose", true)
	toolkit.Println("Data: ", q.Data())

	e := q.Start()
	if e != nil {
		t.Fatal("Fail to start: ", e.Error())
	}
	waitForEntry(q, t)
    waitForApproval(q, t, "admin", alur.Approve, "test")
}

func waitForEntry(r *alur.Request, t *testing.T) {
	t0 := time.Now()
	for {
		s := r.CurrentStep("Start")
		if s == nil {
			time.Sleep(1 * time.Millisecond)
			if time.Since(t0) > (5 * time.Second) {
				t.Fatalf("Waiting for Entry Timeout")
				return
			}
		} else {
			e := s.Enter(toolkit.M{}.Set("Department", "IT"))
			if e != nil {
				t.Fatalf("Entry error: %s", e.Error())
				return
			}
			return
		}
	}
}

func waitForApproval(r *alur.Request, t *testing.T, user string, approval alur.ApproveReject, reason string) {
	t0 := time.Now()
	for {
		s := r.CurrentStep("Approval")
		if s == nil {
			time.Sleep(1 * time.Millisecond)
			if time.Since(t0) > (5 * time.Second) {
				t.Fatalf("Waiting for Approval Timeout")
				return
			}
		} else {
			e := s.ApproveReject(user, approval, reason, nil)
			if e != nil {
				t.Fatalf("ApproveReject fails. %s", e.Error())
				return
			}
		}
	}
}

func TestClose(t *testing.T){
    alur.CloseDb()
}