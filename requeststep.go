package alur

import (
    "github.com/eaciit/toolkit"
)

type RequestStep struct {
	RouteStep

	Signals map[string]interface{}
}

func (r *RequestStep) resetSignal() {
	r.Signals = map[string]interface{}{}
	for _, s := range r.Require {
		r.Signals[s] = nil
	}

	for _, s := range r.RequireReject {
		r.Signals[s] = nil
	}
}

type SignalType int

const (
	SignalAll     SignalType = 0
	SignalRegular SignalType = 1
	SignalReject  SignalType = 2
)

func (r *RequestStep) checkSignal(s SignalType) bool {
    proceed := false
    if s==SignalAll || s==SignalRegular {
        for _, s := range r.Require {
            if sign, has := r.Signals[s]; has && sign!=nil{
                proceed = true
            } else {
                return false
            }
        }
        if proceed{
        return proceed
        }
    }
    
    if s==SignalAll || s==SignalReject{
        for _, s := range r.RequireReject {
            if sign, has := r.Signals[s]; has && sign!=nil{
                proceed = true
            } else {
                return false
            }
        }
        if proceed{
        return proceed
        }
    }
    
    return proceed
}

type ApproveReject int
const (
    Approve ApproveReject = 1
    Reject ApproveReject = 0
)

func (r *RequestStep) Enter(data toolkit.M) error{
    return nil
}

func (r *RequestStep) ApproveReject(userid string, state ApproveReject, reason string, data toolkit.M) error{
    return nil
}
