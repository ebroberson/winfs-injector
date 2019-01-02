// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"
)

type Zipper struct {
	ZipStub        func(dir, zipFile string) error
	zipMutex       sync.RWMutex
	zipArgsForCall []struct {
		dir     string
		zipFile string
	}
	zipReturns struct {
		result1 error
	}
	zipReturnsOnCall map[int]struct {
		result1 error
	}
	UnzipStub        func(zipFile, dest string) error
	unzipMutex       sync.RWMutex
	unzipArgsForCall []struct {
		zipFile string
		dest    string
	}
	unzipReturns struct {
		result1 error
	}
	unzipReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Zipper) Zip(dir string, zipFile string) error {
	fake.zipMutex.Lock()
	ret, specificReturn := fake.zipReturnsOnCall[len(fake.zipArgsForCall)]
	fake.zipArgsForCall = append(fake.zipArgsForCall, struct {
		dir     string
		zipFile string
	}{dir, zipFile})
	fake.recordInvocation("Zip", []interface{}{dir, zipFile})
	fake.zipMutex.Unlock()
	if fake.ZipStub != nil {
		return fake.ZipStub(dir, zipFile)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.zipReturns.result1
}

func (fake *Zipper) ZipCallCount() int {
	fake.zipMutex.RLock()
	defer fake.zipMutex.RUnlock()
	return len(fake.zipArgsForCall)
}

func (fake *Zipper) ZipArgsForCall(i int) (string, string) {
	fake.zipMutex.RLock()
	defer fake.zipMutex.RUnlock()
	return fake.zipArgsForCall[i].dir, fake.zipArgsForCall[i].zipFile
}

func (fake *Zipper) ZipReturns(result1 error) {
	fake.ZipStub = nil
	fake.zipReturns = struct {
		result1 error
	}{result1}
}

func (fake *Zipper) ZipReturnsOnCall(i int, result1 error) {
	fake.ZipStub = nil
	if fake.zipReturnsOnCall == nil {
		fake.zipReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.zipReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Zipper) Unzip(zipFile string, dest string) error {
	fake.unzipMutex.Lock()
	ret, specificReturn := fake.unzipReturnsOnCall[len(fake.unzipArgsForCall)]
	fake.unzipArgsForCall = append(fake.unzipArgsForCall, struct {
		zipFile string
		dest    string
	}{zipFile, dest})
	fake.recordInvocation("Unzip", []interface{}{zipFile, dest})
	fake.unzipMutex.Unlock()
	if fake.UnzipStub != nil {
		return fake.UnzipStub(zipFile, dest)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.unzipReturns.result1
}

func (fake *Zipper) UnzipCallCount() int {
	fake.unzipMutex.RLock()
	defer fake.unzipMutex.RUnlock()
	return len(fake.unzipArgsForCall)
}

func (fake *Zipper) UnzipArgsForCall(i int) (string, string) {
	fake.unzipMutex.RLock()
	defer fake.unzipMutex.RUnlock()
	return fake.unzipArgsForCall[i].zipFile, fake.unzipArgsForCall[i].dest
}

func (fake *Zipper) UnzipReturns(result1 error) {
	fake.UnzipStub = nil
	fake.unzipReturns = struct {
		result1 error
	}{result1}
}

func (fake *Zipper) UnzipReturnsOnCall(i int, result1 error) {
	fake.UnzipStub = nil
	if fake.unzipReturnsOnCall == nil {
		fake.unzipReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.unzipReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Zipper) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.zipMutex.RLock()
	defer fake.zipMutex.RUnlock()
	fake.unzipMutex.RLock()
	defer fake.unzipMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Zipper) recordInvocation(key string, args []interface{}) {
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