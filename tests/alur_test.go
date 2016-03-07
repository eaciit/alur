package tests

import(
    "testing"
    "github.com/eaciit/alur"
    //"github.com/eaciit/toolkit"
)

var (
    r *alur.Route
)

func SkipIfNil(t *testing.T){
    if r==nil {
        t.Skip()
    }
}

func TestRoute(t *testing.T){
  r := new(alur.Route)
  r.ID = "wf_simple"
  r.Title = "Simple Workflow"
  r.Enable = true
}

func TestManageStep(t *testing.T){
    SkipIfNil(t)
    s1 := new(alur.RouteStep)
    s1.StepID = "Start"
    s1.StepType = alur.StepEntry
}