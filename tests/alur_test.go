package tests

import(
    "testing"
    "github.com/eaciit/alur"
    "github.com/eaciit/toolkit"
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
        
    s2 := new(alur.RouteStep)
   s2.Require = []string{"Start"}
    s2.StepID = "Validate"
    s2.StepType = alur.StepAction
    s2.Pre = func(ctx alur.Context){
        m := ctx.Step.Data.(toolkit.M)
        if m.GetInt("leaveday") <= 0 {
            ctx.Request.Stop(alur.RequestCancelled, "Admin", "Leave Day can't be 0")
        }
    }
    s2.Exec= func(ctx alur.Context){
        m := ctx.Step.Data.(toolkit.M)
        dept := m.GetString("department")
        if dept=="it" {
            ctx.Request.StepById("approval").SetApprover(0,"andiek.suncahyo").SetMinimalApprover(0,1)
        } 
    }
}