package alur

import (
    "strings"
    "github.com/eaciit/toolkit"
    "errors"
)

type RequestState string

const(
    RequestDraft RequestState = "Draft"
    RequestRunning RequestState = "Running"
    RequestPending RequestState = "Pending"
    RequestRejected RequestState = "Rejected"
    RequestCancelled RequestState = "Cancelled"
    RequestClosed RequestState = "Closed"
)

type Request struct{
    RouteID string
    State RequestState
    
    _steps map[string]*RequestStep
    _currentSteps []*RequestStep
    _data toolkit.M
}

func NewRequest(route *Route, userId string) *Request{
    q := new(Request)
    q._data = toolkit.M{}
    return q
}

func (r *Request) CurrentSteps() []*RequestStep{
    return r._currentSteps
}

func (r *Request) initSteps(){
    r._steps = map[string]*RequestStep{}
}

func (r *Request) Step(stepId string) *RequestStep{
    r.initSteps()
    return r._steps[stepId]
}

func (r *Request) CurrentStep(stepId string) *RequestStep{
    lowerStepId := strings.ToLower(stepId)
    for _, s := range r._currentSteps{
        if strings.ToLower(s.StepID)==lowerStepId{
            return s
        }
    }
    return nil
}

func (r *Request) Data() toolkit.M{
    if r._data==nil{
        r._data=toolkit.M{}
    }
    return r._data
}

func (r *Request) initStepsFromRoute(){
}

func (r *Request) Start() error{
    r.initStepsFromRoute()
    
    for _, s := range r._steps{
        if len(s.Require)==0 && len(s.RequireReject)==0 && s.AutoStart{
            r._currentSteps =append(r._currentSteps, s)
        }
    }
    
    if len(r._currentSteps)==0{
        return errors.New("Request.Start: No active steps")
    }
    return nil
}

func (r *Request) ReOpen(){
}

func (r *Request) Stop(state RequestState, userid string, message string){
    r.State = state
}

func (r *Request) StepById(id string) *RouteStep{
    return nil
}