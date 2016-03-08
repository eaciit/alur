package alur

import (
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
    Route Route
    State RequestState
    CurrentSteps []*RequestStep
    
    _data toolkit.M
}

func NewRequest(route *Route, userId string) *Request{
    q := new(Request)
    q._data = toolkit.M{}
    return q
}

func (r *Request) Data() toolkit.M{
    if r._data==nil{
        r._data=toolkit.M{}
    }
    return r._data
}

func (r *Request) Start() error{
    return errors.New("Request.Start: No active steps")
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