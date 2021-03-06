// Code generated by counterfeiter. DO NOT EDIT.
package appfilesfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/cf/appfiles"
)

type FakeCfIgnore struct {
	FileShouldBeIgnoredStub        func(string) bool
	fileShouldBeIgnoredMutex       sync.RWMutex
	fileShouldBeIgnoredArgsForCall []struct {
		arg1 string
	}
	fileShouldBeIgnoredReturns struct {
		result1 bool
	}
	fileShouldBeIgnoredReturnsOnCall map[int]struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCfIgnore) FileShouldBeIgnored(arg1 string) bool {
	fake.fileShouldBeIgnoredMutex.Lock()
	ret, specificReturn := fake.fileShouldBeIgnoredReturnsOnCall[len(fake.fileShouldBeIgnoredArgsForCall)]
	fake.fileShouldBeIgnoredArgsForCall = append(fake.fileShouldBeIgnoredArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("FileShouldBeIgnored", []interface{}{arg1})
	fake.fileShouldBeIgnoredMutex.Unlock()
	if fake.FileShouldBeIgnoredStub != nil {
		return fake.FileShouldBeIgnoredStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.fileShouldBeIgnoredReturns
	return fakeReturns.result1
}

func (fake *FakeCfIgnore) FileShouldBeIgnoredCallCount() int {
	fake.fileShouldBeIgnoredMutex.RLock()
	defer fake.fileShouldBeIgnoredMutex.RUnlock()
	return len(fake.fileShouldBeIgnoredArgsForCall)
}

func (fake *FakeCfIgnore) FileShouldBeIgnoredCalls(stub func(string) bool) {
	fake.fileShouldBeIgnoredMutex.Lock()
	defer fake.fileShouldBeIgnoredMutex.Unlock()
	fake.FileShouldBeIgnoredStub = stub
}

func (fake *FakeCfIgnore) FileShouldBeIgnoredArgsForCall(i int) string {
	fake.fileShouldBeIgnoredMutex.RLock()
	defer fake.fileShouldBeIgnoredMutex.RUnlock()
	argsForCall := fake.fileShouldBeIgnoredArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCfIgnore) FileShouldBeIgnoredReturns(result1 bool) {
	fake.fileShouldBeIgnoredMutex.Lock()
	defer fake.fileShouldBeIgnoredMutex.Unlock()
	fake.FileShouldBeIgnoredStub = nil
	fake.fileShouldBeIgnoredReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeCfIgnore) FileShouldBeIgnoredReturnsOnCall(i int, result1 bool) {
	fake.fileShouldBeIgnoredMutex.Lock()
	defer fake.fileShouldBeIgnoredMutex.Unlock()
	fake.FileShouldBeIgnoredStub = nil
	if fake.fileShouldBeIgnoredReturnsOnCall == nil {
		fake.fileShouldBeIgnoredReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.fileShouldBeIgnoredReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeCfIgnore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.fileShouldBeIgnoredMutex.RLock()
	defer fake.fileShouldBeIgnoredMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCfIgnore) recordInvocation(key string, args []interface{}) {
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

var _ appfiles.CfIgnore = new(FakeCfIgnore)
