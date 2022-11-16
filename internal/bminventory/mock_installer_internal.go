// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/assisted-service/internal/bminventory (interfaces: InstallerInternals)

// Package bminventory is a generated GoMock package.
package bminventory

import (
	context "context"
	io "io"
	reflect "reflect"

	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	gomock "github.com/golang/mock/gomock"
	common "github.com/openshift/assisted-service/internal/common"
	models "github.com/openshift/assisted-service/models"
	installer "github.com/openshift/assisted-service/restapi/operations/installer"
	types "k8s.io/apimachinery/pkg/types"
)

// MockInstallerInternals is a mock of InstallerInternals interface.
type MockInstallerInternals struct {
	ctrl     *gomock.Controller
	recorder *MockInstallerInternalsMockRecorder
}

// MockInstallerInternalsMockRecorder is the mock recorder for MockInstallerInternals.
type MockInstallerInternalsMockRecorder struct {
	mock *MockInstallerInternals
}

// NewMockInstallerInternals creates a new mock instance.
func NewMockInstallerInternals(ctrl *gomock.Controller) *MockInstallerInternals {
	mock := &MockInstallerInternals{ctrl: ctrl}
	mock.recorder = &MockInstallerInternalsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInstallerInternals) EXPECT() *MockInstallerInternalsMockRecorder {
	return m.recorder
}

// AddReleaseImage mocks base method.
func (m *MockInstallerInternals) AddReleaseImage(arg0 context.Context, arg1, arg2, arg3 string, arg4 []string) (*models.ReleaseImage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddReleaseImage", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*models.ReleaseImage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddReleaseImage indicates an expected call of AddReleaseImage.
func (mr *MockInstallerInternalsMockRecorder) AddReleaseImage(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddReleaseImage", reflect.TypeOf((*MockInstallerInternals)(nil).AddReleaseImage), arg0, arg1, arg2, arg3, arg4)
}

// BindHostInternal mocks base method.
func (m *MockInstallerInternals) BindHostInternal(arg0 context.Context, arg1 installer.BindHostParams) (*common.Host, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindHostInternal", arg0, arg1)
	ret0, _ := ret[0].(*common.Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BindHostInternal indicates an expected call of BindHostInternal.
func (mr *MockInstallerInternalsMockRecorder) BindHostInternal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindHostInternal", reflect.TypeOf((*MockInstallerInternals)(nil).BindHostInternal), arg0, arg1)
}

// CancelInstallationInternal mocks base method.
func (m *MockInstallerInternals) CancelInstallationInternal(arg0 context.Context, arg1 installer.V2CancelInstallationParams) (*common.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelInstallationInternal", arg0, arg1)
	ret0, _ := ret[0].(*common.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelInstallationInternal indicates an expected call of CancelInstallationInternal.
func (mr *MockInstallerInternalsMockRecorder) CancelInstallationInternal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelInstallationInternal", reflect.TypeOf((*MockInstallerInternals)(nil).CancelInstallationInternal), arg0, arg1)
}

// DeregisterClusterInternal mocks base method.
func (m *MockInstallerInternals) DeregisterClusterInternal(arg0 context.Context, arg1 *common.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeregisterClusterInternal", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeregisterClusterInternal indicates an expected call of DeregisterClusterInternal.
func (mr *MockInstallerInternalsMockRecorder) DeregisterClusterInternal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeregisterClusterInternal", reflect.TypeOf((*MockInstallerInternals)(nil).DeregisterClusterInternal), arg0, arg1)
}

// DeregisterInfraEnvInternal mocks base method.
func (m *MockInstallerInternals) DeregisterInfraEnvInternal(arg0 context.Context, arg1 installer.DeregisterInfraEnvParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeregisterInfraEnvInternal", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeregisterInfraEnvInternal indicates an expected call of DeregisterInfraEnvInternal.
func (mr *MockInstallerInternalsMockRecorder) DeregisterInfraEnvInternal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeregisterInfraEnvInternal", reflect.TypeOf((*MockInstallerInternals)(nil).DeregisterInfraEnvInternal), arg0, arg1)
}

// GetClusterByKubeKey mocks base method.
func (m *MockInstallerInternals) GetClusterByKubeKey(arg0 types.NamespacedName) (*common.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterByKubeKey", arg0)
	ret0, _ := ret[0].(*common.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterByKubeKey indicates an expected call of GetClusterByKubeKey.
func (mr *MockInstallerInternalsMockRecorder) GetClusterByKubeKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterByKubeKey", reflect.TypeOf((*MockInstallerInternals)(nil).GetClusterByKubeKey), arg0)
}

// GetClusterInternal mocks base method.
func (m *MockInstallerInternals) GetClusterInternal(arg0 context.Context, arg1 installer.V2GetClusterParams) (*common.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterInternal", arg0, arg1)
	ret0, _ := ret[0].(*common.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterInternal indicates an expected call of GetClusterInternal.
func (mr *MockInstallerInternalsMockRecorder) GetClusterInternal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterInternal", reflect.TypeOf((*MockInstallerInternals)(nil).GetClusterInternal), arg0, arg1)
}

// GetClusterSupportedPlatformsInternal mocks base method.
func (m *MockInstallerInternals) GetClusterSupportedPlatformsInternal(arg0 context.Context, arg1 installer.GetClusterSupportedPlatformsParams) (*[]models.PlatformType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterSupportedPlatformsInternal", arg0, arg1)
	ret0, _ := ret[0].(*[]models.PlatformType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterSupportedPlatformsInternal indicates an expected call of GetClusterSupportedPlatformsInternal.
func (mr *MockInstallerInternalsMockRecorder) GetClusterSupportedPlatformsInternal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterSupportedPlatformsInternal", reflect.TypeOf((*MockInstallerInternals)(nil).GetClusterSupportedPlatformsInternal), arg0, arg1)
}

// GetCommonHostInternal mocks base method.
func (m *MockInstallerInternals) GetCommonHostInternal(arg0 context.Context, arg1, arg2 string) (*common.Host, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommonHostInternal", arg0, arg1, arg2)
	ret0, _ := ret[0].(*common.Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommonHostInternal indicates an expected call of GetCommonHostInternal.
func (mr *MockInstallerInternalsMockRecorder) GetCommonHostInternal(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommonHostInternal", reflect.TypeOf((*MockInstallerInternals)(nil).GetCommonHostInternal), arg0, arg1, arg2)
}

// GetCredentialsInternal mocks base method.
func (m *MockInstallerInternals) GetCredentialsInternal(arg0 context.Context, arg1 installer.V2GetCredentialsParams) (*models.Credentials, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCredentialsInternal", arg0, arg1)
	ret0, _ := ret[0].(*models.Credentials)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCredentialsInternal indicates an expected call of GetCredentialsInternal.
func (mr *MockInstallerInternalsMockRecorder) GetCredentialsInternal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCredentialsInternal", reflect.TypeOf((*MockInstallerInternals)(nil).GetCredentialsInternal), arg0, arg1)
}

// GetHostByKubeKey mocks base method.
func (m *MockInstallerInternals) GetHostByKubeKey(arg0 types.NamespacedName) (*common.Host, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHostByKubeKey", arg0)
	ret0, _ := ret[0].(*common.Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHostByKubeKey indicates an expected call of GetHostByKubeKey.
func (mr *MockInstallerInternalsMockRecorder) GetHostByKubeKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostByKubeKey", reflect.TypeOf((*MockInstallerInternals)(nil).GetHostByKubeKey), arg0)
}

// GetInfraEnvByKubeKey mocks base method.
func (m *MockInstallerInternals) GetInfraEnvByKubeKey(arg0 types.NamespacedName) (*common.InfraEnv, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfraEnvByKubeKey", arg0)
	ret0, _ := ret[0].(*common.InfraEnv)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInfraEnvByKubeKey indicates an expected call of GetInfraEnvByKubeKey.
func (mr *MockInstallerInternalsMockRecorder) GetInfraEnvByKubeKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfraEnvByKubeKey", reflect.TypeOf((*MockInstallerInternals)(nil).GetInfraEnvByKubeKey), arg0)
}

// GetInfraEnvHostsInternal mocks base method.
func (m *MockInstallerInternals) GetInfraEnvHostsInternal(arg0 context.Context, arg1 strfmt.UUID) ([]*common.Host, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfraEnvHostsInternal", arg0, arg1)
	ret0, _ := ret[0].([]*common.Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInfraEnvHostsInternal indicates an expected call of GetInfraEnvHostsInternal.
func (mr *MockInstallerInternalsMockRecorder) GetInfraEnvHostsInternal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfraEnvHostsInternal", reflect.TypeOf((*MockInstallerInternals)(nil).GetInfraEnvHostsInternal), arg0, arg1)
}

// GetInfraEnvInternal mocks base method.
func (m *MockInstallerInternals) GetInfraEnvInternal(arg0 context.Context, arg1 installer.GetInfraEnvParams) (*common.InfraEnv, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfraEnvInternal", arg0, arg1)
	ret0, _ := ret[0].(*common.InfraEnv)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInfraEnvInternal indicates an expected call of GetInfraEnvInternal.
func (mr *MockInstallerInternalsMockRecorder) GetInfraEnvInternal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfraEnvInternal", reflect.TypeOf((*MockInstallerInternals)(nil).GetInfraEnvInternal), arg0, arg1)
}

// GetKnownApprovedHosts mocks base method.
func (m *MockInstallerInternals) GetKnownApprovedHosts(arg0 strfmt.UUID) ([]*common.Host, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKnownApprovedHosts", arg0)
	ret0, _ := ret[0].([]*common.Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKnownApprovedHosts indicates an expected call of GetKnownApprovedHosts.
func (mr *MockInstallerInternalsMockRecorder) GetKnownApprovedHosts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKnownApprovedHosts", reflect.TypeOf((*MockInstallerInternals)(nil).GetKnownApprovedHosts), arg0)
}

// GetKnownHostApprovedCounts mocks base method.
func (m *MockInstallerInternals) GetKnownHostApprovedCounts(arg0 strfmt.UUID) (int, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKnownHostApprovedCounts", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetKnownHostApprovedCounts indicates an expected call of GetKnownHostApprovedCounts.
func (mr *MockInstallerInternalsMockRecorder) GetKnownHostApprovedCounts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKnownHostApprovedCounts", reflect.TypeOf((*MockInstallerInternals)(nil).GetKnownHostApprovedCounts), arg0)
}

// HostWithCollectedLogsExists mocks base method.
func (m *MockInstallerInternals) HostWithCollectedLogsExists(arg0 strfmt.UUID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HostWithCollectedLogsExists", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HostWithCollectedLogsExists indicates an expected call of HostWithCollectedLogsExists.
func (mr *MockInstallerInternalsMockRecorder) HostWithCollectedLogsExists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HostWithCollectedLogsExists", reflect.TypeOf((*MockInstallerInternals)(nil).HostWithCollectedLogsExists), arg0)
}

// InstallClusterInternal mocks base method.
func (m *MockInstallerInternals) InstallClusterInternal(arg0 context.Context, arg1 installer.V2InstallClusterParams) (*common.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallClusterInternal", arg0, arg1)
	ret0, _ := ret[0].(*common.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InstallClusterInternal indicates an expected call of InstallClusterInternal.
func (mr *MockInstallerInternalsMockRecorder) InstallClusterInternal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallClusterInternal", reflect.TypeOf((*MockInstallerInternals)(nil).InstallClusterInternal), arg0, arg1)
}

// InstallSingleDay2HostInternal mocks base method.
func (m *MockInstallerInternals) InstallSingleDay2HostInternal(arg0 context.Context, arg1, arg2, arg3 strfmt.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallSingleDay2HostInternal", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallSingleDay2HostInternal indicates an expected call of InstallSingleDay2HostInternal.
func (mr *MockInstallerInternalsMockRecorder) InstallSingleDay2HostInternal(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallSingleDay2HostInternal", reflect.TypeOf((*MockInstallerInternals)(nil).InstallSingleDay2HostInternal), arg0, arg1, arg2, arg3)
}

// RegisterClusterInternal mocks base method.
func (m *MockInstallerInternals) RegisterClusterInternal(arg0 context.Context, arg1 *types.NamespacedName, arg2 installer.V2RegisterClusterParams) (*common.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterClusterInternal", arg0, arg1, arg2)
	ret0, _ := ret[0].(*common.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterClusterInternal indicates an expected call of RegisterClusterInternal.
func (mr *MockInstallerInternalsMockRecorder) RegisterClusterInternal(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterClusterInternal", reflect.TypeOf((*MockInstallerInternals)(nil).RegisterClusterInternal), arg0, arg1, arg2)
}

// RegisterInfraEnvInternal mocks base method.
func (m *MockInstallerInternals) RegisterInfraEnvInternal(arg0 context.Context, arg1 *types.NamespacedName, arg2 installer.RegisterInfraEnvParams) (*common.InfraEnv, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterInfraEnvInternal", arg0, arg1, arg2)
	ret0, _ := ret[0].(*common.InfraEnv)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterInfraEnvInternal indicates an expected call of RegisterInfraEnvInternal.
func (mr *MockInstallerInternalsMockRecorder) RegisterInfraEnvInternal(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterInfraEnvInternal", reflect.TypeOf((*MockInstallerInternals)(nil).RegisterInfraEnvInternal), arg0, arg1, arg2)
}

// TransformClusterToDay2Internal mocks base method.
func (m *MockInstallerInternals) TransformClusterToDay2Internal(arg0 context.Context, arg1 strfmt.UUID) (*common.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransformClusterToDay2Internal", arg0, arg1)
	ret0, _ := ret[0].(*common.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransformClusterToDay2Internal indicates an expected call of TransformClusterToDay2Internal.
func (mr *MockInstallerInternalsMockRecorder) TransformClusterToDay2Internal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransformClusterToDay2Internal", reflect.TypeOf((*MockInstallerInternals)(nil).TransformClusterToDay2Internal), arg0, arg1)
}

// UnbindHostInternal mocks base method.
func (m *MockInstallerInternals) UnbindHostInternal(arg0 context.Context, arg1 installer.UnbindHostParams, arg2 bool) (*common.Host, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnbindHostInternal", arg0, arg1, arg2)
	ret0, _ := ret[0].(*common.Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnbindHostInternal indicates an expected call of UnbindHostInternal.
func (mr *MockInstallerInternalsMockRecorder) UnbindHostInternal(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnbindHostInternal", reflect.TypeOf((*MockInstallerInternals)(nil).UnbindHostInternal), arg0, arg1, arg2)
}

// UpdateClusterInstallConfigInternal mocks base method.
func (m *MockInstallerInternals) UpdateClusterInstallConfigInternal(arg0 context.Context, arg1 installer.V2UpdateClusterInstallConfigParams) (*common.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateClusterInstallConfigInternal", arg0, arg1)
	ret0, _ := ret[0].(*common.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateClusterInstallConfigInternal indicates an expected call of UpdateClusterInstallConfigInternal.
func (mr *MockInstallerInternalsMockRecorder) UpdateClusterInstallConfigInternal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateClusterInstallConfigInternal", reflect.TypeOf((*MockInstallerInternals)(nil).UpdateClusterInstallConfigInternal), arg0, arg1)
}

// UpdateClusterNonInteractive mocks base method.
func (m *MockInstallerInternals) UpdateClusterNonInteractive(arg0 context.Context, arg1 installer.V2UpdateClusterParams) (*common.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateClusterNonInteractive", arg0, arg1)
	ret0, _ := ret[0].(*common.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateClusterNonInteractive indicates an expected call of UpdateClusterNonInteractive.
func (mr *MockInstallerInternalsMockRecorder) UpdateClusterNonInteractive(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateClusterNonInteractive", reflect.TypeOf((*MockInstallerInternals)(nil).UpdateClusterNonInteractive), arg0, arg1)
}

// UpdateHostApprovedInternal mocks base method.
func (m *MockInstallerInternals) UpdateHostApprovedInternal(arg0 context.Context, arg1, arg2 string, arg3 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateHostApprovedInternal", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateHostApprovedInternal indicates an expected call of UpdateHostApprovedInternal.
func (mr *MockInstallerInternalsMockRecorder) UpdateHostApprovedInternal(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHostApprovedInternal", reflect.TypeOf((*MockInstallerInternals)(nil).UpdateHostApprovedInternal), arg0, arg1, arg2, arg3)
}

// UpdateInfraEnvInternal mocks base method.
func (m *MockInstallerInternals) UpdateInfraEnvInternal(arg0 context.Context, arg1 installer.UpdateInfraEnvParams, arg2 *string) (*common.InfraEnv, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInfraEnvInternal", arg0, arg1, arg2)
	ret0, _ := ret[0].(*common.InfraEnv)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateInfraEnvInternal indicates an expected call of UpdateInfraEnvInternal.
func (mr *MockInstallerInternalsMockRecorder) UpdateInfraEnvInternal(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInfraEnvInternal", reflect.TypeOf((*MockInstallerInternals)(nil).UpdateInfraEnvInternal), arg0, arg1, arg2)
}

// V2DeregisterHostInternal mocks base method.
func (m *MockInstallerInternals) V2DeregisterHostInternal(arg0 context.Context, arg1 installer.V2DeregisterHostParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2DeregisterHostInternal", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// V2DeregisterHostInternal indicates an expected call of V2DeregisterHostInternal.
func (mr *MockInstallerInternalsMockRecorder) V2DeregisterHostInternal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2DeregisterHostInternal", reflect.TypeOf((*MockInstallerInternals)(nil).V2DeregisterHostInternal), arg0, arg1)
}

// V2DownloadClusterCredentialsInternal mocks base method.
func (m *MockInstallerInternals) V2DownloadClusterCredentialsInternal(arg0 context.Context, arg1 installer.V2DownloadClusterCredentialsParams) (io.ReadCloser, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2DownloadClusterCredentialsInternal", arg0, arg1)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// V2DownloadClusterCredentialsInternal indicates an expected call of V2DownloadClusterCredentialsInternal.
func (mr *MockInstallerInternalsMockRecorder) V2DownloadClusterCredentialsInternal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2DownloadClusterCredentialsInternal", reflect.TypeOf((*MockInstallerInternals)(nil).V2DownloadClusterCredentialsInternal), arg0, arg1)
}

// V2DownloadClusterFilesInternal mocks base method.
func (m *MockInstallerInternals) V2DownloadClusterFilesInternal(arg0 context.Context, arg1 installer.V2DownloadClusterFilesParams) (io.ReadCloser, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2DownloadClusterFilesInternal", arg0, arg1)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// V2DownloadClusterFilesInternal indicates an expected call of V2DownloadClusterFilesInternal.
func (mr *MockInstallerInternalsMockRecorder) V2DownloadClusterFilesInternal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2DownloadClusterFilesInternal", reflect.TypeOf((*MockInstallerInternals)(nil).V2DownloadClusterFilesInternal), arg0, arg1)
}

// V2EventsSubscribe mocks base method.
func (m *MockInstallerInternals) V2EventsSubscribe(arg0 context.Context, arg1 installer.V2EventsSubscribeParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2EventsSubscribe", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2EventsSubscribe indicates an expected call of V2EventsSubscribe.
func (mr *MockInstallerInternalsMockRecorder) V2EventsSubscribe(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2EventsSubscribe", reflect.TypeOf((*MockInstallerInternals)(nil).V2EventsSubscribe), arg0, arg1)
}

// V2EventsSubscriptionDelete mocks base method.
func (m *MockInstallerInternals) V2EventsSubscriptionDelete(arg0 context.Context, arg1 installer.V2EventsSubscriptionDeleteParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2EventsSubscriptionDelete", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2EventsSubscriptionDelete indicates an expected call of V2EventsSubscriptionDelete.
func (mr *MockInstallerInternalsMockRecorder) V2EventsSubscriptionDelete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2EventsSubscriptionDelete", reflect.TypeOf((*MockInstallerInternals)(nil).V2EventsSubscriptionDelete), arg0, arg1)
}

// V2EventsSubscriptionGet mocks base method.
func (m *MockInstallerInternals) V2EventsSubscriptionGet(arg0 context.Context, arg1 installer.V2EventsSubscriptionGetParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2EventsSubscriptionGet", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2EventsSubscriptionGet indicates an expected call of V2EventsSubscriptionGet.
func (mr *MockInstallerInternalsMockRecorder) V2EventsSubscriptionGet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2EventsSubscriptionGet", reflect.TypeOf((*MockInstallerInternals)(nil).V2EventsSubscriptionGet), arg0, arg1)
}

// V2EventsSubscriptionList mocks base method.
func (m *MockInstallerInternals) V2EventsSubscriptionList(arg0 context.Context, arg1 installer.V2EventsSubscriptionListParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2EventsSubscriptionList", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// V2EventsSubscriptionList indicates an expected call of V2EventsSubscriptionList.
func (mr *MockInstallerInternalsMockRecorder) V2EventsSubscriptionList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2EventsSubscriptionList", reflect.TypeOf((*MockInstallerInternals)(nil).V2EventsSubscriptionList), arg0, arg1)
}

// V2ImportClusterInternal mocks base method.
func (m *MockInstallerInternals) V2ImportClusterInternal(arg0 context.Context, arg1 *types.NamespacedName, arg2 *strfmt.UUID, arg3 installer.V2ImportClusterParams) (*common.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2ImportClusterInternal", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*common.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// V2ImportClusterInternal indicates an expected call of V2ImportClusterInternal.
func (mr *MockInstallerInternalsMockRecorder) V2ImportClusterInternal(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2ImportClusterInternal", reflect.TypeOf((*MockInstallerInternals)(nil).V2ImportClusterInternal), arg0, arg1, arg2, arg3)
}

// V2UpdateHostIgnitionInternal mocks base method.
func (m *MockInstallerInternals) V2UpdateHostIgnitionInternal(arg0 context.Context, arg1 installer.V2UpdateHostIgnitionParams) (*models.Host, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2UpdateHostIgnitionInternal", arg0, arg1)
	ret0, _ := ret[0].(*models.Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// V2UpdateHostIgnitionInternal indicates an expected call of V2UpdateHostIgnitionInternal.
func (mr *MockInstallerInternalsMockRecorder) V2UpdateHostIgnitionInternal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2UpdateHostIgnitionInternal", reflect.TypeOf((*MockInstallerInternals)(nil).V2UpdateHostIgnitionInternal), arg0, arg1)
}

// V2UpdateHostInstallProgressInternal mocks base method.
func (m *MockInstallerInternals) V2UpdateHostInstallProgressInternal(arg0 context.Context, arg1 installer.V2UpdateHostInstallProgressParams) (*common.Host, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2UpdateHostInstallProgressInternal", arg0, arg1)
	ret0, _ := ret[0].(*common.Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// V2UpdateHostInstallProgressInternal indicates an expected call of V2UpdateHostInstallProgressInternal.
func (mr *MockInstallerInternalsMockRecorder) V2UpdateHostInstallProgressInternal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2UpdateHostInstallProgressInternal", reflect.TypeOf((*MockInstallerInternals)(nil).V2UpdateHostInstallProgressInternal), arg0, arg1)
}

// V2UpdateHostInstallerArgsInternal mocks base method.
func (m *MockInstallerInternals) V2UpdateHostInstallerArgsInternal(arg0 context.Context, arg1 installer.V2UpdateHostInstallerArgsParams) (*models.Host, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2UpdateHostInstallerArgsInternal", arg0, arg1)
	ret0, _ := ret[0].(*models.Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// V2UpdateHostInstallerArgsInternal indicates an expected call of V2UpdateHostInstallerArgsInternal.
func (mr *MockInstallerInternalsMockRecorder) V2UpdateHostInstallerArgsInternal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2UpdateHostInstallerArgsInternal", reflect.TypeOf((*MockInstallerInternals)(nil).V2UpdateHostInstallerArgsInternal), arg0, arg1)
}

// V2UpdateHostInternal mocks base method.
func (m *MockInstallerInternals) V2UpdateHostInternal(arg0 context.Context, arg1 installer.V2UpdateHostParams, arg2 Interactivity) (*common.Host, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V2UpdateHostInternal", arg0, arg1, arg2)
	ret0, _ := ret[0].(*common.Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// V2UpdateHostInternal indicates an expected call of V2UpdateHostInternal.
func (mr *MockInstallerInternalsMockRecorder) V2UpdateHostInternal(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V2UpdateHostInternal", reflect.TypeOf((*MockInstallerInternals)(nil).V2UpdateHostInternal), arg0, arg1, arg2)
}

// ValidatePullSecret mocks base method.
func (m *MockInstallerInternals) ValidatePullSecret(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidatePullSecret", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidatePullSecret indicates an expected call of ValidatePullSecret.
func (mr *MockInstallerInternalsMockRecorder) ValidatePullSecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatePullSecret", reflect.TypeOf((*MockInstallerInternals)(nil).ValidatePullSecret), arg0, arg1)
}
