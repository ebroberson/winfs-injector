// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/hydrator/layermodifier"
	digest "github.com/opencontainers/go-digest"
	oci "github.com/opencontainers/image-spec/specs-go/v1"
)

type OCIDirectory struct {
	AddBlobStub        func(srcPath string, blobDescriptor oci.Descriptor) error
	addBlobMutex       sync.RWMutex
	addBlobArgsForCall []struct {
		srcPath        string
		blobDescriptor oci.Descriptor
	}
	addBlobReturns struct {
		result1 error
	}
	addBlobReturnsOnCall map[int]struct {
		result1 error
	}
	RemoveTopBlobStub        func(sha256 string) error
	removeTopBlobMutex       sync.RWMutex
	removeTopBlobArgsForCall []struct {
		sha256 string
	}
	removeTopBlobReturns struct {
		result1 error
	}
	removeTopBlobReturnsOnCall map[int]struct {
		result1 error
	}
	ClearMetadataStub        func() error
	clearMetadataMutex       sync.RWMutex
	clearMetadataArgsForCall []struct{}
	clearMetadataReturns     struct {
		result1 error
	}
	clearMetadataReturnsOnCall map[int]struct {
		result1 error
	}
	ReadMetadataStub        func() (oci.Manifest, oci.Image, error)
	readMetadataMutex       sync.RWMutex
	readMetadataArgsForCall []struct{}
	readMetadataReturns     struct {
		result1 oci.Manifest
		result2 oci.Image
		result3 error
	}
	readMetadataReturnsOnCall map[int]struct {
		result1 oci.Manifest
		result2 oci.Image
		result3 error
	}
	WriteMetadataStub        func(layers []oci.Descriptor, diffIds []digest.Digest, layerAdded bool) error
	writeMetadataMutex       sync.RWMutex
	writeMetadataArgsForCall []struct {
		layers     []oci.Descriptor
		diffIds    []digest.Digest
		layerAdded bool
	}
	writeMetadataReturns struct {
		result1 error
	}
	writeMetadataReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *OCIDirectory) AddBlob(srcPath string, blobDescriptor oci.Descriptor) error {
	fake.addBlobMutex.Lock()
	ret, specificReturn := fake.addBlobReturnsOnCall[len(fake.addBlobArgsForCall)]
	fake.addBlobArgsForCall = append(fake.addBlobArgsForCall, struct {
		srcPath        string
		blobDescriptor oci.Descriptor
	}{srcPath, blobDescriptor})
	fake.recordInvocation("AddBlob", []interface{}{srcPath, blobDescriptor})
	fake.addBlobMutex.Unlock()
	if fake.AddBlobStub != nil {
		return fake.AddBlobStub(srcPath, blobDescriptor)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.addBlobReturns.result1
}

func (fake *OCIDirectory) AddBlobCallCount() int {
	fake.addBlobMutex.RLock()
	defer fake.addBlobMutex.RUnlock()
	return len(fake.addBlobArgsForCall)
}

func (fake *OCIDirectory) AddBlobArgsForCall(i int) (string, oci.Descriptor) {
	fake.addBlobMutex.RLock()
	defer fake.addBlobMutex.RUnlock()
	return fake.addBlobArgsForCall[i].srcPath, fake.addBlobArgsForCall[i].blobDescriptor
}

func (fake *OCIDirectory) AddBlobReturns(result1 error) {
	fake.AddBlobStub = nil
	fake.addBlobReturns = struct {
		result1 error
	}{result1}
}

func (fake *OCIDirectory) AddBlobReturnsOnCall(i int, result1 error) {
	fake.AddBlobStub = nil
	if fake.addBlobReturnsOnCall == nil {
		fake.addBlobReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addBlobReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *OCIDirectory) RemoveTopBlob(sha256 string) error {
	fake.removeTopBlobMutex.Lock()
	ret, specificReturn := fake.removeTopBlobReturnsOnCall[len(fake.removeTopBlobArgsForCall)]
	fake.removeTopBlobArgsForCall = append(fake.removeTopBlobArgsForCall, struct {
		sha256 string
	}{sha256})
	fake.recordInvocation("RemoveTopBlob", []interface{}{sha256})
	fake.removeTopBlobMutex.Unlock()
	if fake.RemoveTopBlobStub != nil {
		return fake.RemoveTopBlobStub(sha256)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.removeTopBlobReturns.result1
}

func (fake *OCIDirectory) RemoveTopBlobCallCount() int {
	fake.removeTopBlobMutex.RLock()
	defer fake.removeTopBlobMutex.RUnlock()
	return len(fake.removeTopBlobArgsForCall)
}

func (fake *OCIDirectory) RemoveTopBlobArgsForCall(i int) string {
	fake.removeTopBlobMutex.RLock()
	defer fake.removeTopBlobMutex.RUnlock()
	return fake.removeTopBlobArgsForCall[i].sha256
}

func (fake *OCIDirectory) RemoveTopBlobReturns(result1 error) {
	fake.RemoveTopBlobStub = nil
	fake.removeTopBlobReturns = struct {
		result1 error
	}{result1}
}

func (fake *OCIDirectory) RemoveTopBlobReturnsOnCall(i int, result1 error) {
	fake.RemoveTopBlobStub = nil
	if fake.removeTopBlobReturnsOnCall == nil {
		fake.removeTopBlobReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeTopBlobReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *OCIDirectory) ClearMetadata() error {
	fake.clearMetadataMutex.Lock()
	ret, specificReturn := fake.clearMetadataReturnsOnCall[len(fake.clearMetadataArgsForCall)]
	fake.clearMetadataArgsForCall = append(fake.clearMetadataArgsForCall, struct{}{})
	fake.recordInvocation("ClearMetadata", []interface{}{})
	fake.clearMetadataMutex.Unlock()
	if fake.ClearMetadataStub != nil {
		return fake.ClearMetadataStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.clearMetadataReturns.result1
}

func (fake *OCIDirectory) ClearMetadataCallCount() int {
	fake.clearMetadataMutex.RLock()
	defer fake.clearMetadataMutex.RUnlock()
	return len(fake.clearMetadataArgsForCall)
}

func (fake *OCIDirectory) ClearMetadataReturns(result1 error) {
	fake.ClearMetadataStub = nil
	fake.clearMetadataReturns = struct {
		result1 error
	}{result1}
}

func (fake *OCIDirectory) ClearMetadataReturnsOnCall(i int, result1 error) {
	fake.ClearMetadataStub = nil
	if fake.clearMetadataReturnsOnCall == nil {
		fake.clearMetadataReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.clearMetadataReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *OCIDirectory) ReadMetadata() (oci.Manifest, oci.Image, error) {
	fake.readMetadataMutex.Lock()
	ret, specificReturn := fake.readMetadataReturnsOnCall[len(fake.readMetadataArgsForCall)]
	fake.readMetadataArgsForCall = append(fake.readMetadataArgsForCall, struct{}{})
	fake.recordInvocation("ReadMetadata", []interface{}{})
	fake.readMetadataMutex.Unlock()
	if fake.ReadMetadataStub != nil {
		return fake.ReadMetadataStub()
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.readMetadataReturns.result1, fake.readMetadataReturns.result2, fake.readMetadataReturns.result3
}

func (fake *OCIDirectory) ReadMetadataCallCount() int {
	fake.readMetadataMutex.RLock()
	defer fake.readMetadataMutex.RUnlock()
	return len(fake.readMetadataArgsForCall)
}

func (fake *OCIDirectory) ReadMetadataReturns(result1 oci.Manifest, result2 oci.Image, result3 error) {
	fake.ReadMetadataStub = nil
	fake.readMetadataReturns = struct {
		result1 oci.Manifest
		result2 oci.Image
		result3 error
	}{result1, result2, result3}
}

func (fake *OCIDirectory) ReadMetadataReturnsOnCall(i int, result1 oci.Manifest, result2 oci.Image, result3 error) {
	fake.ReadMetadataStub = nil
	if fake.readMetadataReturnsOnCall == nil {
		fake.readMetadataReturnsOnCall = make(map[int]struct {
			result1 oci.Manifest
			result2 oci.Image
			result3 error
		})
	}
	fake.readMetadataReturnsOnCall[i] = struct {
		result1 oci.Manifest
		result2 oci.Image
		result3 error
	}{result1, result2, result3}
}

func (fake *OCIDirectory) WriteMetadata(layers []oci.Descriptor, diffIds []digest.Digest, layerAdded bool) error {
	var layersCopy []oci.Descriptor
	if layers != nil {
		layersCopy = make([]oci.Descriptor, len(layers))
		copy(layersCopy, layers)
	}
	var diffIdsCopy []digest.Digest
	if diffIds != nil {
		diffIdsCopy = make([]digest.Digest, len(diffIds))
		copy(diffIdsCopy, diffIds)
	}
	fake.writeMetadataMutex.Lock()
	ret, specificReturn := fake.writeMetadataReturnsOnCall[len(fake.writeMetadataArgsForCall)]
	fake.writeMetadataArgsForCall = append(fake.writeMetadataArgsForCall, struct {
		layers     []oci.Descriptor
		diffIds    []digest.Digest
		layerAdded bool
	}{layersCopy, diffIdsCopy, layerAdded})
	fake.recordInvocation("WriteMetadata", []interface{}{layersCopy, diffIdsCopy, layerAdded})
	fake.writeMetadataMutex.Unlock()
	if fake.WriteMetadataStub != nil {
		return fake.WriteMetadataStub(layers, diffIds, layerAdded)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.writeMetadataReturns.result1
}

func (fake *OCIDirectory) WriteMetadataCallCount() int {
	fake.writeMetadataMutex.RLock()
	defer fake.writeMetadataMutex.RUnlock()
	return len(fake.writeMetadataArgsForCall)
}

func (fake *OCIDirectory) WriteMetadataArgsForCall(i int) ([]oci.Descriptor, []digest.Digest, bool) {
	fake.writeMetadataMutex.RLock()
	defer fake.writeMetadataMutex.RUnlock()
	return fake.writeMetadataArgsForCall[i].layers, fake.writeMetadataArgsForCall[i].diffIds, fake.writeMetadataArgsForCall[i].layerAdded

}

func (fake *OCIDirectory) WriteMetadataReturns(result1 error) {
	fake.WriteMetadataStub = nil
	fake.writeMetadataReturns = struct {
		result1 error
	}{result1}
}

func (fake *OCIDirectory) WriteMetadataReturnsOnCall(i int, result1 error) {
	fake.WriteMetadataStub = nil
	if fake.writeMetadataReturnsOnCall == nil {
		fake.writeMetadataReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.writeMetadataReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *OCIDirectory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addBlobMutex.RLock()
	defer fake.addBlobMutex.RUnlock()
	fake.removeTopBlobMutex.RLock()
	defer fake.removeTopBlobMutex.RUnlock()
	fake.clearMetadataMutex.RLock()
	defer fake.clearMetadataMutex.RUnlock()
	fake.readMetadataMutex.RLock()
	defer fake.readMetadataMutex.RUnlock()
	fake.writeMetadataMutex.RLock()
	defer fake.writeMetadataMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *OCIDirectory) recordInvocation(key string, args []interface{}) {
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

var _ layermodifier.OCIDirectory = new(OCIDirectory)
