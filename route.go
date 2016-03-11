package alur

import (
   "github.com/eaciit/orm/v1"
)

type Route struct {
    orm.ModelBase
	ID     string
	Title  string
	Owner  string
	Enable bool
    Port int
    CodePath string
    
    Steps map[string]*RouteStep
}

func (r *Route) TableName() string{
    return "routes"
}

func (r *Route) RecordID() interface{}{
    return r.ID
}

func (r *Route) initStep(){
    if r.Steps==nil {
        r.Steps = map[string]*RouteStep{}
    }
}

func (r *Route) UpdateSteps(steps []*RouteStep){
    r.initStep()
    for _, s := range steps{
        r.Steps[s.StepID]=s
    }
}

func (r *Route) DeleteStep(stepId string){
    r.initStep()
    delete(r.Steps, stepId)
}

func (r *Route) Verify() error{
    return nil
}
