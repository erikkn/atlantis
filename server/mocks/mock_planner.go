// Automatically generated by pegomock. DO NOT EDIT!
// Source: github.com/hootsuite/atlantis/server (interfaces: Planner)

package mocks

import (
	server "github.com/hootsuite/atlantis/server"
	pegomock "github.com/petergtz/pegomock"
	"reflect"
)

type MockPlanner struct {
	fail func(message string, callerSkip ...int)
}

func NewMockPlanner() *MockPlanner {
	return &MockPlanner{fail: pegomock.GlobalFailHandler}
}

func (mock *MockPlanner) Execute(ctx *server.CommandContext) {
	params := []pegomock.Param{ctx}
	pegomock.GetGenericMockFrom(mock).Invoke("Execute", params, []reflect.Type{})
}

func (mock *MockPlanner) SetLockURL(_param0 func(string) string) {
	params := []pegomock.Param{_param0}
	pegomock.GetGenericMockFrom(mock).Invoke("SetLockURL", params, []reflect.Type{})
}

func (mock *MockPlanner) VerifyWasCalledOnce() *VerifierPlanner {
	return &VerifierPlanner{mock, pegomock.Times(1), nil}
}

func (mock *MockPlanner) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierPlanner {
	return &VerifierPlanner{mock, invocationCountMatcher, nil}
}

func (mock *MockPlanner) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierPlanner {
	return &VerifierPlanner{mock, invocationCountMatcher, inOrderContext}
}

type VerifierPlanner struct {
	mock                   *MockPlanner
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
}

func (verifier *VerifierPlanner) Execute(ctx *server.CommandContext) *Planner_Execute_OngoingVerification {
	params := []pegomock.Param{ctx}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Execute", params)
	return &Planner_Execute_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type Planner_Execute_OngoingVerification struct {
	mock              *MockPlanner
	methodInvocations []pegomock.MethodInvocation
}

func (c *Planner_Execute_OngoingVerification) GetCapturedArguments() *server.CommandContext {
	ctx := c.GetAllCapturedArguments()
	return ctx[len(ctx)-1]
}

func (c *Planner_Execute_OngoingVerification) GetAllCapturedArguments() (_param0 []*server.CommandContext) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]*server.CommandContext, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(*server.CommandContext)
		}
	}
	return
}

func (verifier *VerifierPlanner) SetLockURL(_param0 func(string) string) *Planner_SetLockURL_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "SetLockURL", params)
	return &Planner_SetLockURL_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type Planner_SetLockURL_OngoingVerification struct {
	mock              *MockPlanner
	methodInvocations []pegomock.MethodInvocation
}

func (c *Planner_SetLockURL_OngoingVerification) GetCapturedArguments() func(string) string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *Planner_SetLockURL_OngoingVerification) GetAllCapturedArguments() (_param0 []func(string) string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]func(string) string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(func(string) string)
		}
	}
	return
}
