// Code generated by counterfeiter. DO NOT EDIT.
package commonfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/pluginaction"
	"code.cloudfoundry.org/cli/api/plugin"
	"code.cloudfoundry.org/cli/command/common"
	"code.cloudfoundry.org/cli/util/configv3"
)

type FakeInstallPluginActor struct {
	CreateExecutableCopyStub        func(string, string) (string, error)
	createExecutableCopyMutex       sync.RWMutex
	createExecutableCopyArgsForCall []struct {
		arg1 string
		arg2 string
	}
	createExecutableCopyReturns struct {
		result1 string
		result2 error
	}
	createExecutableCopyReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	DownloadExecutableBinaryFromURLStub        func(string, string, plugin.ProxyReader) (string, error)
	downloadExecutableBinaryFromURLMutex       sync.RWMutex
	downloadExecutableBinaryFromURLArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 plugin.ProxyReader
	}
	downloadExecutableBinaryFromURLReturns struct {
		result1 string
		result2 error
	}
	downloadExecutableBinaryFromURLReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	FileExistsStub        func(string) bool
	fileExistsMutex       sync.RWMutex
	fileExistsArgsForCall []struct {
		arg1 string
	}
	fileExistsReturns struct {
		result1 bool
	}
	fileExistsReturnsOnCall map[int]struct {
		result1 bool
	}
	GetAndValidatePluginStub        func(pluginaction.PluginMetadata, pluginaction.CommandList, string) (configv3.Plugin, error)
	getAndValidatePluginMutex       sync.RWMutex
	getAndValidatePluginArgsForCall []struct {
		arg1 pluginaction.PluginMetadata
		arg2 pluginaction.CommandList
		arg3 string
	}
	getAndValidatePluginReturns struct {
		result1 configv3.Plugin
		result2 error
	}
	getAndValidatePluginReturnsOnCall map[int]struct {
		result1 configv3.Plugin
		result2 error
	}
	GetPlatformStringStub        func(string, string) string
	getPlatformStringMutex       sync.RWMutex
	getPlatformStringArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getPlatformStringReturns struct {
		result1 string
	}
	getPlatformStringReturnsOnCall map[int]struct {
		result1 string
	}
	GetPluginInfoFromRepositoriesForPlatformStub        func(string, []configv3.PluginRepository, string) (pluginaction.PluginInfo, []string, error)
	getPluginInfoFromRepositoriesForPlatformMutex       sync.RWMutex
	getPluginInfoFromRepositoriesForPlatformArgsForCall []struct {
		arg1 string
		arg2 []configv3.PluginRepository
		arg3 string
	}
	getPluginInfoFromRepositoriesForPlatformReturns struct {
		result1 pluginaction.PluginInfo
		result2 []string
		result3 error
	}
	getPluginInfoFromRepositoriesForPlatformReturnsOnCall map[int]struct {
		result1 pluginaction.PluginInfo
		result2 []string
		result3 error
	}
	GetPluginRepositoryStub        func(string) (configv3.PluginRepository, error)
	getPluginRepositoryMutex       sync.RWMutex
	getPluginRepositoryArgsForCall []struct {
		arg1 string
	}
	getPluginRepositoryReturns struct {
		result1 configv3.PluginRepository
		result2 error
	}
	getPluginRepositoryReturnsOnCall map[int]struct {
		result1 configv3.PluginRepository
		result2 error
	}
	InstallPluginFromPathStub        func(string, configv3.Plugin) error
	installPluginFromPathMutex       sync.RWMutex
	installPluginFromPathArgsForCall []struct {
		arg1 string
		arg2 configv3.Plugin
	}
	installPluginFromPathReturns struct {
		result1 error
	}
	installPluginFromPathReturnsOnCall map[int]struct {
		result1 error
	}
	UninstallPluginStub        func(pluginaction.PluginUninstaller, string) error
	uninstallPluginMutex       sync.RWMutex
	uninstallPluginArgsForCall []struct {
		arg1 pluginaction.PluginUninstaller
		arg2 string
	}
	uninstallPluginReturns struct {
		result1 error
	}
	uninstallPluginReturnsOnCall map[int]struct {
		result1 error
	}
	ValidateFileChecksumStub        func(string, string) bool
	validateFileChecksumMutex       sync.RWMutex
	validateFileChecksumArgsForCall []struct {
		arg1 string
		arg2 string
	}
	validateFileChecksumReturns struct {
		result1 bool
	}
	validateFileChecksumReturnsOnCall map[int]struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInstallPluginActor) CreateExecutableCopy(arg1 string, arg2 string) (string, error) {
	fake.createExecutableCopyMutex.Lock()
	ret, specificReturn := fake.createExecutableCopyReturnsOnCall[len(fake.createExecutableCopyArgsForCall)]
	fake.createExecutableCopyArgsForCall = append(fake.createExecutableCopyArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("CreateExecutableCopy", []interface{}{arg1, arg2})
	fake.createExecutableCopyMutex.Unlock()
	if fake.CreateExecutableCopyStub != nil {
		return fake.CreateExecutableCopyStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createExecutableCopyReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeInstallPluginActor) CreateExecutableCopyCallCount() int {
	fake.createExecutableCopyMutex.RLock()
	defer fake.createExecutableCopyMutex.RUnlock()
	return len(fake.createExecutableCopyArgsForCall)
}

func (fake *FakeInstallPluginActor) CreateExecutableCopyCalls(stub func(string, string) (string, error)) {
	fake.createExecutableCopyMutex.Lock()
	defer fake.createExecutableCopyMutex.Unlock()
	fake.CreateExecutableCopyStub = stub
}

func (fake *FakeInstallPluginActor) CreateExecutableCopyArgsForCall(i int) (string, string) {
	fake.createExecutableCopyMutex.RLock()
	defer fake.createExecutableCopyMutex.RUnlock()
	argsForCall := fake.createExecutableCopyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeInstallPluginActor) CreateExecutableCopyReturns(result1 string, result2 error) {
	fake.createExecutableCopyMutex.Lock()
	defer fake.createExecutableCopyMutex.Unlock()
	fake.CreateExecutableCopyStub = nil
	fake.createExecutableCopyReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeInstallPluginActor) CreateExecutableCopyReturnsOnCall(i int, result1 string, result2 error) {
	fake.createExecutableCopyMutex.Lock()
	defer fake.createExecutableCopyMutex.Unlock()
	fake.CreateExecutableCopyStub = nil
	if fake.createExecutableCopyReturnsOnCall == nil {
		fake.createExecutableCopyReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.createExecutableCopyReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeInstallPluginActor) DownloadExecutableBinaryFromURL(arg1 string, arg2 string, arg3 plugin.ProxyReader) (string, error) {
	fake.downloadExecutableBinaryFromURLMutex.Lock()
	ret, specificReturn := fake.downloadExecutableBinaryFromURLReturnsOnCall[len(fake.downloadExecutableBinaryFromURLArgsForCall)]
	fake.downloadExecutableBinaryFromURLArgsForCall = append(fake.downloadExecutableBinaryFromURLArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 plugin.ProxyReader
	}{arg1, arg2, arg3})
	fake.recordInvocation("DownloadExecutableBinaryFromURL", []interface{}{arg1, arg2, arg3})
	fake.downloadExecutableBinaryFromURLMutex.Unlock()
	if fake.DownloadExecutableBinaryFromURLStub != nil {
		return fake.DownloadExecutableBinaryFromURLStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.downloadExecutableBinaryFromURLReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeInstallPluginActor) DownloadExecutableBinaryFromURLCallCount() int {
	fake.downloadExecutableBinaryFromURLMutex.RLock()
	defer fake.downloadExecutableBinaryFromURLMutex.RUnlock()
	return len(fake.downloadExecutableBinaryFromURLArgsForCall)
}

func (fake *FakeInstallPluginActor) DownloadExecutableBinaryFromURLCalls(stub func(string, string, plugin.ProxyReader) (string, error)) {
	fake.downloadExecutableBinaryFromURLMutex.Lock()
	defer fake.downloadExecutableBinaryFromURLMutex.Unlock()
	fake.DownloadExecutableBinaryFromURLStub = stub
}

func (fake *FakeInstallPluginActor) DownloadExecutableBinaryFromURLArgsForCall(i int) (string, string, plugin.ProxyReader) {
	fake.downloadExecutableBinaryFromURLMutex.RLock()
	defer fake.downloadExecutableBinaryFromURLMutex.RUnlock()
	argsForCall := fake.downloadExecutableBinaryFromURLArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeInstallPluginActor) DownloadExecutableBinaryFromURLReturns(result1 string, result2 error) {
	fake.downloadExecutableBinaryFromURLMutex.Lock()
	defer fake.downloadExecutableBinaryFromURLMutex.Unlock()
	fake.DownloadExecutableBinaryFromURLStub = nil
	fake.downloadExecutableBinaryFromURLReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeInstallPluginActor) DownloadExecutableBinaryFromURLReturnsOnCall(i int, result1 string, result2 error) {
	fake.downloadExecutableBinaryFromURLMutex.Lock()
	defer fake.downloadExecutableBinaryFromURLMutex.Unlock()
	fake.DownloadExecutableBinaryFromURLStub = nil
	if fake.downloadExecutableBinaryFromURLReturnsOnCall == nil {
		fake.downloadExecutableBinaryFromURLReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.downloadExecutableBinaryFromURLReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeInstallPluginActor) FileExists(arg1 string) bool {
	fake.fileExistsMutex.Lock()
	ret, specificReturn := fake.fileExistsReturnsOnCall[len(fake.fileExistsArgsForCall)]
	fake.fileExistsArgsForCall = append(fake.fileExistsArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("FileExists", []interface{}{arg1})
	fake.fileExistsMutex.Unlock()
	if fake.FileExistsStub != nil {
		return fake.FileExistsStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.fileExistsReturns
	return fakeReturns.result1
}

func (fake *FakeInstallPluginActor) FileExistsCallCount() int {
	fake.fileExistsMutex.RLock()
	defer fake.fileExistsMutex.RUnlock()
	return len(fake.fileExistsArgsForCall)
}

func (fake *FakeInstallPluginActor) FileExistsCalls(stub func(string) bool) {
	fake.fileExistsMutex.Lock()
	defer fake.fileExistsMutex.Unlock()
	fake.FileExistsStub = stub
}

func (fake *FakeInstallPluginActor) FileExistsArgsForCall(i int) string {
	fake.fileExistsMutex.RLock()
	defer fake.fileExistsMutex.RUnlock()
	argsForCall := fake.fileExistsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeInstallPluginActor) FileExistsReturns(result1 bool) {
	fake.fileExistsMutex.Lock()
	defer fake.fileExistsMutex.Unlock()
	fake.FileExistsStub = nil
	fake.fileExistsReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeInstallPluginActor) FileExistsReturnsOnCall(i int, result1 bool) {
	fake.fileExistsMutex.Lock()
	defer fake.fileExistsMutex.Unlock()
	fake.FileExistsStub = nil
	if fake.fileExistsReturnsOnCall == nil {
		fake.fileExistsReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.fileExistsReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeInstallPluginActor) GetAndValidatePlugin(arg1 pluginaction.PluginMetadata, arg2 pluginaction.CommandList, arg3 string) (configv3.Plugin, error) {
	fake.getAndValidatePluginMutex.Lock()
	ret, specificReturn := fake.getAndValidatePluginReturnsOnCall[len(fake.getAndValidatePluginArgsForCall)]
	fake.getAndValidatePluginArgsForCall = append(fake.getAndValidatePluginArgsForCall, struct {
		arg1 pluginaction.PluginMetadata
		arg2 pluginaction.CommandList
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("GetAndValidatePlugin", []interface{}{arg1, arg2, arg3})
	fake.getAndValidatePluginMutex.Unlock()
	if fake.GetAndValidatePluginStub != nil {
		return fake.GetAndValidatePluginStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getAndValidatePluginReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeInstallPluginActor) GetAndValidatePluginCallCount() int {
	fake.getAndValidatePluginMutex.RLock()
	defer fake.getAndValidatePluginMutex.RUnlock()
	return len(fake.getAndValidatePluginArgsForCall)
}

func (fake *FakeInstallPluginActor) GetAndValidatePluginCalls(stub func(pluginaction.PluginMetadata, pluginaction.CommandList, string) (configv3.Plugin, error)) {
	fake.getAndValidatePluginMutex.Lock()
	defer fake.getAndValidatePluginMutex.Unlock()
	fake.GetAndValidatePluginStub = stub
}

func (fake *FakeInstallPluginActor) GetAndValidatePluginArgsForCall(i int) (pluginaction.PluginMetadata, pluginaction.CommandList, string) {
	fake.getAndValidatePluginMutex.RLock()
	defer fake.getAndValidatePluginMutex.RUnlock()
	argsForCall := fake.getAndValidatePluginArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeInstallPluginActor) GetAndValidatePluginReturns(result1 configv3.Plugin, result2 error) {
	fake.getAndValidatePluginMutex.Lock()
	defer fake.getAndValidatePluginMutex.Unlock()
	fake.GetAndValidatePluginStub = nil
	fake.getAndValidatePluginReturns = struct {
		result1 configv3.Plugin
		result2 error
	}{result1, result2}
}

func (fake *FakeInstallPluginActor) GetAndValidatePluginReturnsOnCall(i int, result1 configv3.Plugin, result2 error) {
	fake.getAndValidatePluginMutex.Lock()
	defer fake.getAndValidatePluginMutex.Unlock()
	fake.GetAndValidatePluginStub = nil
	if fake.getAndValidatePluginReturnsOnCall == nil {
		fake.getAndValidatePluginReturnsOnCall = make(map[int]struct {
			result1 configv3.Plugin
			result2 error
		})
	}
	fake.getAndValidatePluginReturnsOnCall[i] = struct {
		result1 configv3.Plugin
		result2 error
	}{result1, result2}
}

func (fake *FakeInstallPluginActor) GetPlatformString(arg1 string, arg2 string) string {
	fake.getPlatformStringMutex.Lock()
	ret, specificReturn := fake.getPlatformStringReturnsOnCall[len(fake.getPlatformStringArgsForCall)]
	fake.getPlatformStringArgsForCall = append(fake.getPlatformStringArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetPlatformString", []interface{}{arg1, arg2})
	fake.getPlatformStringMutex.Unlock()
	if fake.GetPlatformStringStub != nil {
		return fake.GetPlatformStringStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getPlatformStringReturns
	return fakeReturns.result1
}

func (fake *FakeInstallPluginActor) GetPlatformStringCallCount() int {
	fake.getPlatformStringMutex.RLock()
	defer fake.getPlatformStringMutex.RUnlock()
	return len(fake.getPlatformStringArgsForCall)
}

func (fake *FakeInstallPluginActor) GetPlatformStringCalls(stub func(string, string) string) {
	fake.getPlatformStringMutex.Lock()
	defer fake.getPlatformStringMutex.Unlock()
	fake.GetPlatformStringStub = stub
}

func (fake *FakeInstallPluginActor) GetPlatformStringArgsForCall(i int) (string, string) {
	fake.getPlatformStringMutex.RLock()
	defer fake.getPlatformStringMutex.RUnlock()
	argsForCall := fake.getPlatformStringArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeInstallPluginActor) GetPlatformStringReturns(result1 string) {
	fake.getPlatformStringMutex.Lock()
	defer fake.getPlatformStringMutex.Unlock()
	fake.GetPlatformStringStub = nil
	fake.getPlatformStringReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeInstallPluginActor) GetPlatformStringReturnsOnCall(i int, result1 string) {
	fake.getPlatformStringMutex.Lock()
	defer fake.getPlatformStringMutex.Unlock()
	fake.GetPlatformStringStub = nil
	if fake.getPlatformStringReturnsOnCall == nil {
		fake.getPlatformStringReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getPlatformStringReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeInstallPluginActor) GetPluginInfoFromRepositoriesForPlatform(arg1 string, arg2 []configv3.PluginRepository, arg3 string) (pluginaction.PluginInfo, []string, error) {
	var arg2Copy []configv3.PluginRepository
	if arg2 != nil {
		arg2Copy = make([]configv3.PluginRepository, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.getPluginInfoFromRepositoriesForPlatformMutex.Lock()
	ret, specificReturn := fake.getPluginInfoFromRepositoriesForPlatformReturnsOnCall[len(fake.getPluginInfoFromRepositoriesForPlatformArgsForCall)]
	fake.getPluginInfoFromRepositoriesForPlatformArgsForCall = append(fake.getPluginInfoFromRepositoriesForPlatformArgsForCall, struct {
		arg1 string
		arg2 []configv3.PluginRepository
		arg3 string
	}{arg1, arg2Copy, arg3})
	fake.recordInvocation("GetPluginInfoFromRepositoriesForPlatform", []interface{}{arg1, arg2Copy, arg3})
	fake.getPluginInfoFromRepositoriesForPlatformMutex.Unlock()
	if fake.GetPluginInfoFromRepositoriesForPlatformStub != nil {
		return fake.GetPluginInfoFromRepositoriesForPlatformStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getPluginInfoFromRepositoriesForPlatformReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeInstallPluginActor) GetPluginInfoFromRepositoriesForPlatformCallCount() int {
	fake.getPluginInfoFromRepositoriesForPlatformMutex.RLock()
	defer fake.getPluginInfoFromRepositoriesForPlatformMutex.RUnlock()
	return len(fake.getPluginInfoFromRepositoriesForPlatformArgsForCall)
}

func (fake *FakeInstallPluginActor) GetPluginInfoFromRepositoriesForPlatformCalls(stub func(string, []configv3.PluginRepository, string) (pluginaction.PluginInfo, []string, error)) {
	fake.getPluginInfoFromRepositoriesForPlatformMutex.Lock()
	defer fake.getPluginInfoFromRepositoriesForPlatformMutex.Unlock()
	fake.GetPluginInfoFromRepositoriesForPlatformStub = stub
}

func (fake *FakeInstallPluginActor) GetPluginInfoFromRepositoriesForPlatformArgsForCall(i int) (string, []configv3.PluginRepository, string) {
	fake.getPluginInfoFromRepositoriesForPlatformMutex.RLock()
	defer fake.getPluginInfoFromRepositoriesForPlatformMutex.RUnlock()
	argsForCall := fake.getPluginInfoFromRepositoriesForPlatformArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeInstallPluginActor) GetPluginInfoFromRepositoriesForPlatformReturns(result1 pluginaction.PluginInfo, result2 []string, result3 error) {
	fake.getPluginInfoFromRepositoriesForPlatformMutex.Lock()
	defer fake.getPluginInfoFromRepositoriesForPlatformMutex.Unlock()
	fake.GetPluginInfoFromRepositoriesForPlatformStub = nil
	fake.getPluginInfoFromRepositoriesForPlatformReturns = struct {
		result1 pluginaction.PluginInfo
		result2 []string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeInstallPluginActor) GetPluginInfoFromRepositoriesForPlatformReturnsOnCall(i int, result1 pluginaction.PluginInfo, result2 []string, result3 error) {
	fake.getPluginInfoFromRepositoriesForPlatformMutex.Lock()
	defer fake.getPluginInfoFromRepositoriesForPlatformMutex.Unlock()
	fake.GetPluginInfoFromRepositoriesForPlatformStub = nil
	if fake.getPluginInfoFromRepositoriesForPlatformReturnsOnCall == nil {
		fake.getPluginInfoFromRepositoriesForPlatformReturnsOnCall = make(map[int]struct {
			result1 pluginaction.PluginInfo
			result2 []string
			result3 error
		})
	}
	fake.getPluginInfoFromRepositoriesForPlatformReturnsOnCall[i] = struct {
		result1 pluginaction.PluginInfo
		result2 []string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeInstallPluginActor) GetPluginRepository(arg1 string) (configv3.PluginRepository, error) {
	fake.getPluginRepositoryMutex.Lock()
	ret, specificReturn := fake.getPluginRepositoryReturnsOnCall[len(fake.getPluginRepositoryArgsForCall)]
	fake.getPluginRepositoryArgsForCall = append(fake.getPluginRepositoryArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetPluginRepository", []interface{}{arg1})
	fake.getPluginRepositoryMutex.Unlock()
	if fake.GetPluginRepositoryStub != nil {
		return fake.GetPluginRepositoryStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getPluginRepositoryReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeInstallPluginActor) GetPluginRepositoryCallCount() int {
	fake.getPluginRepositoryMutex.RLock()
	defer fake.getPluginRepositoryMutex.RUnlock()
	return len(fake.getPluginRepositoryArgsForCall)
}

func (fake *FakeInstallPluginActor) GetPluginRepositoryCalls(stub func(string) (configv3.PluginRepository, error)) {
	fake.getPluginRepositoryMutex.Lock()
	defer fake.getPluginRepositoryMutex.Unlock()
	fake.GetPluginRepositoryStub = stub
}

func (fake *FakeInstallPluginActor) GetPluginRepositoryArgsForCall(i int) string {
	fake.getPluginRepositoryMutex.RLock()
	defer fake.getPluginRepositoryMutex.RUnlock()
	argsForCall := fake.getPluginRepositoryArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeInstallPluginActor) GetPluginRepositoryReturns(result1 configv3.PluginRepository, result2 error) {
	fake.getPluginRepositoryMutex.Lock()
	defer fake.getPluginRepositoryMutex.Unlock()
	fake.GetPluginRepositoryStub = nil
	fake.getPluginRepositoryReturns = struct {
		result1 configv3.PluginRepository
		result2 error
	}{result1, result2}
}

func (fake *FakeInstallPluginActor) GetPluginRepositoryReturnsOnCall(i int, result1 configv3.PluginRepository, result2 error) {
	fake.getPluginRepositoryMutex.Lock()
	defer fake.getPluginRepositoryMutex.Unlock()
	fake.GetPluginRepositoryStub = nil
	if fake.getPluginRepositoryReturnsOnCall == nil {
		fake.getPluginRepositoryReturnsOnCall = make(map[int]struct {
			result1 configv3.PluginRepository
			result2 error
		})
	}
	fake.getPluginRepositoryReturnsOnCall[i] = struct {
		result1 configv3.PluginRepository
		result2 error
	}{result1, result2}
}

func (fake *FakeInstallPluginActor) InstallPluginFromPath(arg1 string, arg2 configv3.Plugin) error {
	fake.installPluginFromPathMutex.Lock()
	ret, specificReturn := fake.installPluginFromPathReturnsOnCall[len(fake.installPluginFromPathArgsForCall)]
	fake.installPluginFromPathArgsForCall = append(fake.installPluginFromPathArgsForCall, struct {
		arg1 string
		arg2 configv3.Plugin
	}{arg1, arg2})
	fake.recordInvocation("InstallPluginFromPath", []interface{}{arg1, arg2})
	fake.installPluginFromPathMutex.Unlock()
	if fake.InstallPluginFromPathStub != nil {
		return fake.InstallPluginFromPathStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.installPluginFromPathReturns
	return fakeReturns.result1
}

func (fake *FakeInstallPluginActor) InstallPluginFromPathCallCount() int {
	fake.installPluginFromPathMutex.RLock()
	defer fake.installPluginFromPathMutex.RUnlock()
	return len(fake.installPluginFromPathArgsForCall)
}

func (fake *FakeInstallPluginActor) InstallPluginFromPathCalls(stub func(string, configv3.Plugin) error) {
	fake.installPluginFromPathMutex.Lock()
	defer fake.installPluginFromPathMutex.Unlock()
	fake.InstallPluginFromPathStub = stub
}

func (fake *FakeInstallPluginActor) InstallPluginFromPathArgsForCall(i int) (string, configv3.Plugin) {
	fake.installPluginFromPathMutex.RLock()
	defer fake.installPluginFromPathMutex.RUnlock()
	argsForCall := fake.installPluginFromPathArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeInstallPluginActor) InstallPluginFromPathReturns(result1 error) {
	fake.installPluginFromPathMutex.Lock()
	defer fake.installPluginFromPathMutex.Unlock()
	fake.InstallPluginFromPathStub = nil
	fake.installPluginFromPathReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstallPluginActor) InstallPluginFromPathReturnsOnCall(i int, result1 error) {
	fake.installPluginFromPathMutex.Lock()
	defer fake.installPluginFromPathMutex.Unlock()
	fake.InstallPluginFromPathStub = nil
	if fake.installPluginFromPathReturnsOnCall == nil {
		fake.installPluginFromPathReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.installPluginFromPathReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstallPluginActor) UninstallPlugin(arg1 pluginaction.PluginUninstaller, arg2 string) error {
	fake.uninstallPluginMutex.Lock()
	ret, specificReturn := fake.uninstallPluginReturnsOnCall[len(fake.uninstallPluginArgsForCall)]
	fake.uninstallPluginArgsForCall = append(fake.uninstallPluginArgsForCall, struct {
		arg1 pluginaction.PluginUninstaller
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("UninstallPlugin", []interface{}{arg1, arg2})
	fake.uninstallPluginMutex.Unlock()
	if fake.UninstallPluginStub != nil {
		return fake.UninstallPluginStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.uninstallPluginReturns
	return fakeReturns.result1
}

func (fake *FakeInstallPluginActor) UninstallPluginCallCount() int {
	fake.uninstallPluginMutex.RLock()
	defer fake.uninstallPluginMutex.RUnlock()
	return len(fake.uninstallPluginArgsForCall)
}

func (fake *FakeInstallPluginActor) UninstallPluginCalls(stub func(pluginaction.PluginUninstaller, string) error) {
	fake.uninstallPluginMutex.Lock()
	defer fake.uninstallPluginMutex.Unlock()
	fake.UninstallPluginStub = stub
}

func (fake *FakeInstallPluginActor) UninstallPluginArgsForCall(i int) (pluginaction.PluginUninstaller, string) {
	fake.uninstallPluginMutex.RLock()
	defer fake.uninstallPluginMutex.RUnlock()
	argsForCall := fake.uninstallPluginArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeInstallPluginActor) UninstallPluginReturns(result1 error) {
	fake.uninstallPluginMutex.Lock()
	defer fake.uninstallPluginMutex.Unlock()
	fake.UninstallPluginStub = nil
	fake.uninstallPluginReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstallPluginActor) UninstallPluginReturnsOnCall(i int, result1 error) {
	fake.uninstallPluginMutex.Lock()
	defer fake.uninstallPluginMutex.Unlock()
	fake.UninstallPluginStub = nil
	if fake.uninstallPluginReturnsOnCall == nil {
		fake.uninstallPluginReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.uninstallPluginReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstallPluginActor) ValidateFileChecksum(arg1 string, arg2 string) bool {
	fake.validateFileChecksumMutex.Lock()
	ret, specificReturn := fake.validateFileChecksumReturnsOnCall[len(fake.validateFileChecksumArgsForCall)]
	fake.validateFileChecksumArgsForCall = append(fake.validateFileChecksumArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("ValidateFileChecksum", []interface{}{arg1, arg2})
	fake.validateFileChecksumMutex.Unlock()
	if fake.ValidateFileChecksumStub != nil {
		return fake.ValidateFileChecksumStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.validateFileChecksumReturns
	return fakeReturns.result1
}

func (fake *FakeInstallPluginActor) ValidateFileChecksumCallCount() int {
	fake.validateFileChecksumMutex.RLock()
	defer fake.validateFileChecksumMutex.RUnlock()
	return len(fake.validateFileChecksumArgsForCall)
}

func (fake *FakeInstallPluginActor) ValidateFileChecksumCalls(stub func(string, string) bool) {
	fake.validateFileChecksumMutex.Lock()
	defer fake.validateFileChecksumMutex.Unlock()
	fake.ValidateFileChecksumStub = stub
}

func (fake *FakeInstallPluginActor) ValidateFileChecksumArgsForCall(i int) (string, string) {
	fake.validateFileChecksumMutex.RLock()
	defer fake.validateFileChecksumMutex.RUnlock()
	argsForCall := fake.validateFileChecksumArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeInstallPluginActor) ValidateFileChecksumReturns(result1 bool) {
	fake.validateFileChecksumMutex.Lock()
	defer fake.validateFileChecksumMutex.Unlock()
	fake.ValidateFileChecksumStub = nil
	fake.validateFileChecksumReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeInstallPluginActor) ValidateFileChecksumReturnsOnCall(i int, result1 bool) {
	fake.validateFileChecksumMutex.Lock()
	defer fake.validateFileChecksumMutex.Unlock()
	fake.ValidateFileChecksumStub = nil
	if fake.validateFileChecksumReturnsOnCall == nil {
		fake.validateFileChecksumReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.validateFileChecksumReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeInstallPluginActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createExecutableCopyMutex.RLock()
	defer fake.createExecutableCopyMutex.RUnlock()
	fake.downloadExecutableBinaryFromURLMutex.RLock()
	defer fake.downloadExecutableBinaryFromURLMutex.RUnlock()
	fake.fileExistsMutex.RLock()
	defer fake.fileExistsMutex.RUnlock()
	fake.getAndValidatePluginMutex.RLock()
	defer fake.getAndValidatePluginMutex.RUnlock()
	fake.getPlatformStringMutex.RLock()
	defer fake.getPlatformStringMutex.RUnlock()
	fake.getPluginInfoFromRepositoriesForPlatformMutex.RLock()
	defer fake.getPluginInfoFromRepositoriesForPlatformMutex.RUnlock()
	fake.getPluginRepositoryMutex.RLock()
	defer fake.getPluginRepositoryMutex.RUnlock()
	fake.installPluginFromPathMutex.RLock()
	defer fake.installPluginFromPathMutex.RUnlock()
	fake.uninstallPluginMutex.RLock()
	defer fake.uninstallPluginMutex.RUnlock()
	fake.validateFileChecksumMutex.RLock()
	defer fake.validateFileChecksumMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeInstallPluginActor) recordInvocation(key string, args []interface{}) {
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

var _ common.InstallPluginActor = new(FakeInstallPluginActor)