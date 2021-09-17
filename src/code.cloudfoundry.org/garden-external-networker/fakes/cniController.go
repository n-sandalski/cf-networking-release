// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/containernetworking/cni/pkg/types"
)

type CNIController struct {
	DownStub        func(string, string) error
	downMutex       sync.RWMutex
	downArgsForCall []struct {
		arg1 string
		arg2 string
	}
	downReturns struct {
		result1 error
	}
	downReturnsOnCall map[int]struct {
		result1 error
	}
	UpStub        func(string, string, map[string]interface{}, map[string]interface{}) (types.Result, error)
	upMutex       sync.RWMutex
	upArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 map[string]interface{}
		arg4 map[string]interface{}
	}
	upReturns struct {
		result1 types.Result
		result2 error
	}
	upReturnsOnCall map[int]struct {
		result1 types.Result
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *CNIController) Down(arg1 string, arg2 string) error {
	fake.downMutex.Lock()
	ret, specificReturn := fake.downReturnsOnCall[len(fake.downArgsForCall)]
	fake.downArgsForCall = append(fake.downArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.DownStub
	fakeReturns := fake.downReturns
	fake.recordInvocation("Down", []interface{}{arg1, arg2})
	fake.downMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *CNIController) DownCallCount() int {
	fake.downMutex.RLock()
	defer fake.downMutex.RUnlock()
	return len(fake.downArgsForCall)
}

func (fake *CNIController) DownCalls(stub func(string, string) error) {
	fake.downMutex.Lock()
	defer fake.downMutex.Unlock()
	fake.DownStub = stub
}

func (fake *CNIController) DownArgsForCall(i int) (string, string) {
	fake.downMutex.RLock()
	defer fake.downMutex.RUnlock()
	argsForCall := fake.downArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *CNIController) DownReturns(result1 error) {
	fake.downMutex.Lock()
	defer fake.downMutex.Unlock()
	fake.DownStub = nil
	fake.downReturns = struct {
		result1 error
	}{result1}
}

func (fake *CNIController) DownReturnsOnCall(i int, result1 error) {
	fake.downMutex.Lock()
	defer fake.downMutex.Unlock()
	fake.DownStub = nil
	if fake.downReturnsOnCall == nil {
		fake.downReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.downReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *CNIController) Up(arg1 string, arg2 string, arg3 map[string]interface{}, arg4 map[string]interface{}) (types.Result, error) {
	fake.upMutex.Lock()
	ret, specificReturn := fake.upReturnsOnCall[len(fake.upArgsForCall)]
	fake.upArgsForCall = append(fake.upArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 map[string]interface{}
		arg4 map[string]interface{}
	}{arg1, arg2, arg3, arg4})
	stub := fake.UpStub
	fakeReturns := fake.upReturns
	fake.recordInvocation("Up", []interface{}{arg1, arg2, arg3, arg4})
	fake.upMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CNIController) UpCallCount() int {
	fake.upMutex.RLock()
	defer fake.upMutex.RUnlock()
	return len(fake.upArgsForCall)
}

func (fake *CNIController) UpCalls(stub func(string, string, map[string]interface{}, map[string]interface{}) (types.Result, error)) {
	fake.upMutex.Lock()
	defer fake.upMutex.Unlock()
	fake.UpStub = stub
}

func (fake *CNIController) UpArgsForCall(i int) (string, string, map[string]interface{}, map[string]interface{}) {
	fake.upMutex.RLock()
	defer fake.upMutex.RUnlock()
	argsForCall := fake.upArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *CNIController) UpReturns(result1 types.Result, result2 error) {
	fake.upMutex.Lock()
	defer fake.upMutex.Unlock()
	fake.UpStub = nil
	fake.upReturns = struct {
		result1 types.Result
		result2 error
	}{result1, result2}
}

func (fake *CNIController) UpReturnsOnCall(i int, result1 types.Result, result2 error) {
	fake.upMutex.Lock()
	defer fake.upMutex.Unlock()
	fake.UpStub = nil
	if fake.upReturnsOnCall == nil {
		fake.upReturnsOnCall = make(map[int]struct {
			result1 types.Result
			result2 error
		})
	}
	fake.upReturnsOnCall[i] = struct {
		result1 types.Result
		result2 error
	}{result1, result2}
}

func (fake *CNIController) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.downMutex.RLock()
	defer fake.downMutex.RUnlock()
	fake.upMutex.RLock()
	defer fake.upMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *CNIController) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}