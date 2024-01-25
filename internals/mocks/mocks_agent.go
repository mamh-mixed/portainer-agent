// Code generated by MockGen. DO NOT EDIT.
// Source: ./agent.go
//
// Generated by this command:
//
//	mockgen -source=./agent.go -destination=./internals/mocks -package mocks
//

// Package mock_agent is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	agent "github.com/portainer/agent"
	libstack "github.com/portainer/portainer/pkg/libstack"
	gomock "go.uber.org/mock/gomock"
)

// MockClusterService is a mock of ClusterService interface.
type MockClusterService struct {
	ctrl     *gomock.Controller
	recorder *MockClusterServiceMockRecorder
}

// MockClusterServiceMockRecorder is the mock recorder for MockClusterService.
type MockClusterServiceMockRecorder struct {
	mock *MockClusterService
}

// NewMockClusterService creates a new mock instance.
func NewMockClusterService(ctrl *gomock.Controller) *MockClusterService {
	mock := &MockClusterService{ctrl: ctrl}
	mock.recorder = &MockClusterServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterService) EXPECT() *MockClusterServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockClusterService) Create(advertiseAddr string, joinAddr []string, probeTimeout, probeInterval time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", advertiseAddr, joinAddr, probeTimeout, probeInterval)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockClusterServiceMockRecorder) Create(advertiseAddr, joinAddr, probeTimeout, probeInterval any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockClusterService)(nil).Create), advertiseAddr, joinAddr, probeTimeout, probeInterval)
}

// GetMemberByNodeName mocks base method.
func (m *MockClusterService) GetMemberByNodeName(nodeName string) *agent.ClusterMember {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMemberByNodeName", nodeName)
	ret0, _ := ret[0].(*agent.ClusterMember)
	return ret0
}

// GetMemberByNodeName indicates an expected call of GetMemberByNodeName.
func (mr *MockClusterServiceMockRecorder) GetMemberByNodeName(nodeName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMemberByNodeName", reflect.TypeOf((*MockClusterService)(nil).GetMemberByNodeName), nodeName)
}

// GetMemberByRole mocks base method.
func (m *MockClusterService) GetMemberByRole(role agent.DockerNodeRole) *agent.ClusterMember {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMemberByRole", role)
	ret0, _ := ret[0].(*agent.ClusterMember)
	return ret0
}

// GetMemberByRole indicates an expected call of GetMemberByRole.
func (mr *MockClusterServiceMockRecorder) GetMemberByRole(role any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMemberByRole", reflect.TypeOf((*MockClusterService)(nil).GetMemberByRole), role)
}

// GetMemberWithEdgeKeySet mocks base method.
func (m *MockClusterService) GetMemberWithEdgeKeySet() *agent.ClusterMember {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMemberWithEdgeKeySet")
	ret0, _ := ret[0].(*agent.ClusterMember)
	return ret0
}

// GetMemberWithEdgeKeySet indicates an expected call of GetMemberWithEdgeKeySet.
func (mr *MockClusterServiceMockRecorder) GetMemberWithEdgeKeySet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMemberWithEdgeKeySet", reflect.TypeOf((*MockClusterService)(nil).GetMemberWithEdgeKeySet))
}

// GetRuntimeConfiguration mocks base method.
func (m *MockClusterService) GetRuntimeConfiguration() *agent.RuntimeConfiguration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRuntimeConfiguration")
	ret0, _ := ret[0].(*agent.RuntimeConfiguration)
	return ret0
}

// GetRuntimeConfiguration indicates an expected call of GetRuntimeConfiguration.
func (mr *MockClusterServiceMockRecorder) GetRuntimeConfiguration() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRuntimeConfiguration", reflect.TypeOf((*MockClusterService)(nil).GetRuntimeConfiguration))
}

// Leave mocks base method.
func (m *MockClusterService) Leave() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Leave")
}

// Leave indicates an expected call of Leave.
func (mr *MockClusterServiceMockRecorder) Leave() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Leave", reflect.TypeOf((*MockClusterService)(nil).Leave))
}

// Members mocks base method.
func (m *MockClusterService) Members() []agent.ClusterMember {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Members")
	ret0, _ := ret[0].([]agent.ClusterMember)
	return ret0
}

// Members indicates an expected call of Members.
func (mr *MockClusterServiceMockRecorder) Members() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Members", reflect.TypeOf((*MockClusterService)(nil).Members))
}

// UpdateRuntimeConfiguration mocks base method.
func (m *MockClusterService) UpdateRuntimeConfiguration(runtimeConfiguration *agent.RuntimeConfiguration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRuntimeConfiguration", runtimeConfiguration)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRuntimeConfiguration indicates an expected call of UpdateRuntimeConfiguration.
func (mr *MockClusterServiceMockRecorder) UpdateRuntimeConfiguration(runtimeConfiguration any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRuntimeConfiguration", reflect.TypeOf((*MockClusterService)(nil).UpdateRuntimeConfiguration), runtimeConfiguration)
}

// MockDigitalSignatureService is a mock of DigitalSignatureService interface.
type MockDigitalSignatureService struct {
	ctrl     *gomock.Controller
	recorder *MockDigitalSignatureServiceMockRecorder
}

// MockDigitalSignatureServiceMockRecorder is the mock recorder for MockDigitalSignatureService.
type MockDigitalSignatureServiceMockRecorder struct {
	mock *MockDigitalSignatureService
}

// NewMockDigitalSignatureService creates a new mock instance.
func NewMockDigitalSignatureService(ctrl *gomock.Controller) *MockDigitalSignatureService {
	mock := &MockDigitalSignatureService{ctrl: ctrl}
	mock.recorder = &MockDigitalSignatureServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDigitalSignatureService) EXPECT() *MockDigitalSignatureServiceMockRecorder {
	return m.recorder
}

// IsAssociated mocks base method.
func (m *MockDigitalSignatureService) IsAssociated() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAssociated")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsAssociated indicates an expected call of IsAssociated.
func (mr *MockDigitalSignatureServiceMockRecorder) IsAssociated() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAssociated", reflect.TypeOf((*MockDigitalSignatureService)(nil).IsAssociated))
}

// VerifySignature mocks base method.
func (m *MockDigitalSignatureService) VerifySignature(signature, key string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifySignature", signature, key)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifySignature indicates an expected call of VerifySignature.
func (mr *MockDigitalSignatureServiceMockRecorder) VerifySignature(signature, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifySignature", reflect.TypeOf((*MockDigitalSignatureService)(nil).VerifySignature), signature, key)
}

// MockDockerInfoService is a mock of DockerInfoService interface.
type MockDockerInfoService struct {
	ctrl     *gomock.Controller
	recorder *MockDockerInfoServiceMockRecorder
}

// MockDockerInfoServiceMockRecorder is the mock recorder for MockDockerInfoService.
type MockDockerInfoServiceMockRecorder struct {
	mock *MockDockerInfoService
}

// NewMockDockerInfoService creates a new mock instance.
func NewMockDockerInfoService(ctrl *gomock.Controller) *MockDockerInfoService {
	mock := &MockDockerInfoService{ctrl: ctrl}
	mock.recorder = &MockDockerInfoServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDockerInfoService) EXPECT() *MockDockerInfoServiceMockRecorder {
	return m.recorder
}

// GetContainerIpFromDockerEngine mocks base method.
func (m *MockDockerInfoService) GetContainerIpFromDockerEngine(containerName string, ignoreNonSwarmNetworks bool) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContainerIpFromDockerEngine", containerName, ignoreNonSwarmNetworks)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContainerIpFromDockerEngine indicates an expected call of GetContainerIpFromDockerEngine.
func (mr *MockDockerInfoServiceMockRecorder) GetContainerIpFromDockerEngine(containerName, ignoreNonSwarmNetworks any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainerIpFromDockerEngine", reflect.TypeOf((*MockDockerInfoService)(nil).GetContainerIpFromDockerEngine), containerName, ignoreNonSwarmNetworks)
}

// GetRuntimeConfigurationFromDockerEngine mocks base method.
func (m *MockDockerInfoService) GetRuntimeConfigurationFromDockerEngine() (*agent.RuntimeConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRuntimeConfigurationFromDockerEngine")
	ret0, _ := ret[0].(*agent.RuntimeConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRuntimeConfigurationFromDockerEngine indicates an expected call of GetRuntimeConfigurationFromDockerEngine.
func (mr *MockDockerInfoServiceMockRecorder) GetRuntimeConfigurationFromDockerEngine() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRuntimeConfigurationFromDockerEngine", reflect.TypeOf((*MockDockerInfoService)(nil).GetRuntimeConfigurationFromDockerEngine))
}

// GetServiceNameFromDockerEngine mocks base method.
func (m *MockDockerInfoService) GetServiceNameFromDockerEngine(containerName string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceNameFromDockerEngine", containerName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServiceNameFromDockerEngine indicates an expected call of GetServiceNameFromDockerEngine.
func (mr *MockDockerInfoServiceMockRecorder) GetServiceNameFromDockerEngine(containerName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceNameFromDockerEngine", reflect.TypeOf((*MockDockerInfoService)(nil).GetServiceNameFromDockerEngine), containerName)
}

// MockDeployer is a mock of Deployer interface.
type MockDeployer struct {
	ctrl     *gomock.Controller
	recorder *MockDeployerMockRecorder
}

// MockDeployerMockRecorder is the mock recorder for MockDeployer.
type MockDeployerMockRecorder struct {
	mock *MockDeployer
}

// NewMockDeployer creates a new mock instance.
func NewMockDeployer(ctrl *gomock.Controller) *MockDeployer {
	mock := &MockDeployer{ctrl: ctrl}
	mock.recorder = &MockDeployerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeployer) EXPECT() *MockDeployerMockRecorder {
	return m.recorder
}

// Deploy mocks base method.
func (m *MockDeployer) Deploy(ctx context.Context, name string, filePaths []string, options agent.DeployOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deploy", ctx, name, filePaths, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deploy indicates an expected call of Deploy.
func (mr *MockDeployerMockRecorder) Deploy(ctx, name, filePaths, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deploy", reflect.TypeOf((*MockDeployer)(nil).Deploy), ctx, name, filePaths, options)
}

// Pull mocks base method.
func (m *MockDeployer) Pull(ctx context.Context, name string, filePaths []string, options agent.PullOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pull", ctx, name, filePaths, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Pull indicates an expected call of Pull.
func (mr *MockDeployerMockRecorder) Pull(ctx, name, filePaths, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pull", reflect.TypeOf((*MockDeployer)(nil).Pull), ctx, name, filePaths, options)
}

// Remove mocks base method.
func (m *MockDeployer) Remove(ctx context.Context, name string, filePaths []string, options agent.RemoveOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", ctx, name, filePaths, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockDeployerMockRecorder) Remove(ctx, name, filePaths, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockDeployer)(nil).Remove), ctx, name, filePaths, options)
}

// Validate mocks base method.
func (m *MockDeployer) Validate(ctx context.Context, name string, filePaths []string, options agent.ValidateOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", ctx, name, filePaths, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockDeployerMockRecorder) Validate(ctx, name, filePaths, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockDeployer)(nil).Validate), ctx, name, filePaths, options)
}

// WaitForStatus mocks base method.
func (m *MockDeployer) WaitForStatus(ctx context.Context, name string, status libstack.Status) <-chan string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForStatus", ctx, name, status)
	ret0, _ := ret[0].(<-chan string)
	return ret0
}

// WaitForStatus indicates an expected call of WaitForStatus.
func (mr *MockDeployerMockRecorder) WaitForStatus(ctx, name, status any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForStatus", reflect.TypeOf((*MockDeployer)(nil).WaitForStatus), ctx, name, status)
}

// MockKubernetesInfoService is a mock of KubernetesInfoService interface.
type MockKubernetesInfoService struct {
	ctrl     *gomock.Controller
	recorder *MockKubernetesInfoServiceMockRecorder
}

// MockKubernetesInfoServiceMockRecorder is the mock recorder for MockKubernetesInfoService.
type MockKubernetesInfoServiceMockRecorder struct {
	mock *MockKubernetesInfoService
}

// NewMockKubernetesInfoService creates a new mock instance.
func NewMockKubernetesInfoService(ctrl *gomock.Controller) *MockKubernetesInfoService {
	mock := &MockKubernetesInfoService{ctrl: ctrl}
	mock.recorder = &MockKubernetesInfoServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKubernetesInfoService) EXPECT() *MockKubernetesInfoServiceMockRecorder {
	return m.recorder
}

// GetInformationFromKubernetesCluster mocks base method.
func (m *MockKubernetesInfoService) GetInformationFromKubernetesCluster() (*agent.RuntimeConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInformationFromKubernetesCluster")
	ret0, _ := ret[0].(*agent.RuntimeConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInformationFromKubernetesCluster indicates an expected call of GetInformationFromKubernetesCluster.
func (mr *MockKubernetesInfoServiceMockRecorder) GetInformationFromKubernetesCluster() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInformationFromKubernetesCluster", reflect.TypeOf((*MockKubernetesInfoService)(nil).GetInformationFromKubernetesCluster))
}

// MockOptionParser is a mock of OptionParser interface.
type MockOptionParser struct {
	ctrl     *gomock.Controller
	recorder *MockOptionParserMockRecorder
}

// MockOptionParserMockRecorder is the mock recorder for MockOptionParser.
type MockOptionParserMockRecorder struct {
	mock *MockOptionParser
}

// NewMockOptionParser creates a new mock instance.
func NewMockOptionParser(ctrl *gomock.Controller) *MockOptionParser {
	mock := &MockOptionParser{ctrl: ctrl}
	mock.recorder = &MockOptionParserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOptionParser) EXPECT() *MockOptionParserMockRecorder {
	return m.recorder
}

// Options mocks base method.
func (m *MockOptionParser) Options() (*agent.Options, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Options")
	ret0, _ := ret[0].(*agent.Options)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Options indicates an expected call of Options.
func (mr *MockOptionParserMockRecorder) Options() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Options", reflect.TypeOf((*MockOptionParser)(nil).Options))
}

// MockReverseTunnelClient is a mock of ReverseTunnelClient interface.
type MockReverseTunnelClient struct {
	ctrl     *gomock.Controller
	recorder *MockReverseTunnelClientMockRecorder
}

// MockReverseTunnelClientMockRecorder is the mock recorder for MockReverseTunnelClient.
type MockReverseTunnelClientMockRecorder struct {
	mock *MockReverseTunnelClient
}

// NewMockReverseTunnelClient creates a new mock instance.
func NewMockReverseTunnelClient(ctrl *gomock.Controller) *MockReverseTunnelClient {
	mock := &MockReverseTunnelClient{ctrl: ctrl}
	mock.recorder = &MockReverseTunnelClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReverseTunnelClient) EXPECT() *MockReverseTunnelClientMockRecorder {
	return m.recorder
}

// CloseTunnel mocks base method.
func (m *MockReverseTunnelClient) CloseTunnel() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseTunnel")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseTunnel indicates an expected call of CloseTunnel.
func (mr *MockReverseTunnelClientMockRecorder) CloseTunnel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseTunnel", reflect.TypeOf((*MockReverseTunnelClient)(nil).CloseTunnel))
}

// CreateTunnel mocks base method.
func (m *MockReverseTunnelClient) CreateTunnel(config agent.TunnelConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTunnel", config)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTunnel indicates an expected call of CreateTunnel.
func (mr *MockReverseTunnelClientMockRecorder) CreateTunnel(config any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTunnel", reflect.TypeOf((*MockReverseTunnelClient)(nil).CreateTunnel), config)
}

// IsTunnelOpen mocks base method.
func (m *MockReverseTunnelClient) IsTunnelOpen() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsTunnelOpen")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsTunnelOpen indicates an expected call of IsTunnelOpen.
func (mr *MockReverseTunnelClientMockRecorder) IsTunnelOpen() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsTunnelOpen", reflect.TypeOf((*MockReverseTunnelClient)(nil).IsTunnelOpen))
}

// MockScheduler is a mock of Scheduler interface.
type MockScheduler struct {
	ctrl     *gomock.Controller
	recorder *MockSchedulerMockRecorder
}

// MockSchedulerMockRecorder is the mock recorder for MockScheduler.
type MockSchedulerMockRecorder struct {
	mock *MockScheduler
}

// NewMockScheduler creates a new mock instance.
func NewMockScheduler(ctrl *gomock.Controller) *MockScheduler {
	mock := &MockScheduler{ctrl: ctrl}
	mock.recorder = &MockSchedulerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScheduler) EXPECT() *MockSchedulerMockRecorder {
	return m.recorder
}

// AddSchedule mocks base method.
func (m *MockScheduler) AddSchedule(schedule agent.Schedule) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSchedule", schedule)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddSchedule indicates an expected call of AddSchedule.
func (mr *MockSchedulerMockRecorder) AddSchedule(schedule any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSchedule", reflect.TypeOf((*MockScheduler)(nil).AddSchedule), schedule)
}

// ProcessScheduleLogsCollection mocks base method.
func (m *MockScheduler) ProcessScheduleLogsCollection() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ProcessScheduleLogsCollection")
}

// ProcessScheduleLogsCollection indicates an expected call of ProcessScheduleLogsCollection.
func (mr *MockSchedulerMockRecorder) ProcessScheduleLogsCollection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessScheduleLogsCollection", reflect.TypeOf((*MockScheduler)(nil).ProcessScheduleLogsCollection))
}

// RemoveSchedule mocks base method.
func (m *MockScheduler) RemoveSchedule(schedule agent.Schedule) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveSchedule", schedule)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveSchedule indicates an expected call of RemoveSchedule.
func (mr *MockSchedulerMockRecorder) RemoveSchedule(schedule any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSchedule", reflect.TypeOf((*MockScheduler)(nil).RemoveSchedule), schedule)
}

// Schedule mocks base method.
func (m *MockScheduler) Schedule(schedules []agent.Schedule) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Schedule", schedules)
	ret0, _ := ret[0].(error)
	return ret0
}

// Schedule indicates an expected call of Schedule.
func (mr *MockSchedulerMockRecorder) Schedule(schedules any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Schedule", reflect.TypeOf((*MockScheduler)(nil).Schedule), schedules)
}

// MockSystemService is a mock of SystemService interface.
type MockSystemService struct {
	ctrl     *gomock.Controller
	recorder *MockSystemServiceMockRecorder
}

// MockSystemServiceMockRecorder is the mock recorder for MockSystemService.
type MockSystemServiceMockRecorder struct {
	mock *MockSystemService
}

// NewMockSystemService creates a new mock instance.
func NewMockSystemService(ctrl *gomock.Controller) *MockSystemService {
	mock := &MockSystemService{ctrl: ctrl}
	mock.recorder = &MockSystemServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSystemService) EXPECT() *MockSystemServiceMockRecorder {
	return m.recorder
}

// GetDiskInfo mocks base method.
func (m *MockSystemService) GetDiskInfo() ([]agent.PhysicalDisk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDiskInfo")
	ret0, _ := ret[0].([]agent.PhysicalDisk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDiskInfo indicates an expected call of GetDiskInfo.
func (mr *MockSystemServiceMockRecorder) GetDiskInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDiskInfo", reflect.TypeOf((*MockSystemService)(nil).GetDiskInfo))
}

// GetPciDevices mocks base method.
func (m *MockSystemService) GetPciDevices() ([]agent.PciDevice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPciDevices")
	ret0, _ := ret[0].([]agent.PciDevice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPciDevices indicates an expected call of GetPciDevices.
func (mr *MockSystemServiceMockRecorder) GetPciDevices() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPciDevices", reflect.TypeOf((*MockSystemService)(nil).GetPciDevices))
}
