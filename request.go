package alur

import (
    //"github.com/eaciit/toolkit"
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
    CurrentSteps []*RouteStep
}

func (r *Request) Start(){
}

func (r *Request) ReOpen(){
}

func (r *Request) Stop(state RequestState, userid string, message string){
    r.State = state
}

func (r *Request) StepById(id string) *RouteStep{
    return nil
}