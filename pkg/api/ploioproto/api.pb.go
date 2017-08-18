// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

/*
Package ploioproto is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	ApplicationCreate
	ApplicationGet
	ApplicationList
	Application
	Pipeline
	PipelineGet
	PipelineList
	PipelineCreate
	Stage
	StageGet
	StageList
	Env
	Cluster
	ClusterGet
	ClusterList
	ClusterCreate
	Service
	Port
*/
package ploioproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StageType int32

const (
	StageType_Deploy   StageType = 0
	StageType_Check    StageType = 1
	StageType_Teardown StageType = 2
	StageType_Wait     StageType = 3
)

var StageType_name = map[int32]string{
	0: "Deploy",
	1: "Check",
	2: "Teardown",
	3: "Wait",
}
var StageType_value = map[string]int32{
	"Deploy":   0,
	"Check":    1,
	"Teardown": 2,
	"Wait":     3,
}

func (x StageType) String() string {
	return proto.EnumName(StageType_name, int32(x))
}
func (StageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ServiceType int32

const (
	ServiceType_ClusterIP    ServiceType = 0
	ServiceType_NodePort     ServiceType = 1
	ServiceType_LoadBalancer ServiceType = 2
	ServiceType_ExternalName ServiceType = 3
)

var ServiceType_name = map[int32]string{
	0: "ClusterIP",
	1: "NodePort",
	2: "LoadBalancer",
	3: "ExternalName",
}
var ServiceType_value = map[string]int32{
	"ClusterIP":    0,
	"NodePort":     1,
	"LoadBalancer": 2,
	"ExternalName": 3,
}

func (x ServiceType) String() string {
	return proto.EnumName(ServiceType_name, int32(x))
}
func (ServiceType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ApplicationCreate struct {
	Name  string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Owner string `protobuf:"bytes,2,opt,name=Owner" json:"Owner,omitempty"`
	Repo  string `protobuf:"bytes,3,opt,name=Repo" json:"Repo,omitempty"`
	// Optional values below
	Env      []*Env     `protobuf:"bytes,6,rep,name=Env" json:"Env,omitempty"`
	Services []*Service `protobuf:"bytes,7,rep,name=Services" json:"Services,omitempty"`
}

func (m *ApplicationCreate) Reset()                    { *m = ApplicationCreate{} }
func (m *ApplicationCreate) String() string            { return proto.CompactTextString(m) }
func (*ApplicationCreate) ProtoMessage()               {}
func (*ApplicationCreate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ApplicationCreate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ApplicationCreate) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ApplicationCreate) GetRepo() string {
	if m != nil {
		return m.Repo
	}
	return ""
}

func (m *ApplicationCreate) GetEnv() []*Env {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *ApplicationCreate) GetServices() []*Service {
	if m != nil {
		return m.Services
	}
	return nil
}

type ApplicationGet struct {
	ID   int32  `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
}

func (m *ApplicationGet) Reset()                    { *m = ApplicationGet{} }
func (m *ApplicationGet) String() string            { return proto.CompactTextString(m) }
func (*ApplicationGet) ProtoMessage()               {}
func (*ApplicationGet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ApplicationGet) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *ApplicationGet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ApplicationList struct {
	Applications []*Application `protobuf:"bytes,1,rep,name=Applications" json:"Applications,omitempty"`
}

func (m *ApplicationList) Reset()                    { *m = ApplicationList{} }
func (m *ApplicationList) String() string            { return proto.CompactTextString(m) }
func (*ApplicationList) ProtoMessage()               {}
func (*ApplicationList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ApplicationList) GetApplications() []*Application {
	if m != nil {
		return m.Applications
	}
	return nil
}

type Application struct {
	ID        int32       `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Name      string      `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Owner     string      `protobuf:"bytes,3,opt,name=Owner" json:"Owner,omitempty"`
	Repo      string      `protobuf:"bytes,4,opt,name=Repo" json:"Repo,omitempty"`
	Env       []*Env      `protobuf:"bytes,5,rep,name=Env" json:"Env,omitempty"`
	Services  []*Service  `protobuf:"bytes,6,rep,name=Services" json:"Services,omitempty"`
	Pipelines []*Pipeline `protobuf:"bytes,7,rep,name=Pipelines" json:"Pipelines,omitempty"`
}

func (m *Application) Reset()                    { *m = Application{} }
func (m *Application) String() string            { return proto.CompactTextString(m) }
func (*Application) ProtoMessage()               {}
func (*Application) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Application) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Application) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Application) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Application) GetRepo() string {
	if m != nil {
		return m.Repo
	}
	return ""
}

func (m *Application) GetEnv() []*Env {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *Application) GetServices() []*Service {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *Application) GetPipelines() []*Pipeline {
	if m != nil {
		return m.Pipelines
	}
	return nil
}

type Pipeline struct {
	ID     int32    `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Name   string   `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Order  int32    `protobuf:"varint,3,opt,name=Order" json:"Order,omitempty"`
	Stages []*Stage `protobuf:"bytes,4,rep,name=Stages" json:"Stages,omitempty"`
}

func (m *Pipeline) Reset()                    { *m = Pipeline{} }
func (m *Pipeline) String() string            { return proto.CompactTextString(m) }
func (*Pipeline) ProtoMessage()               {}
func (*Pipeline) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Pipeline) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Pipeline) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Pipeline) GetOrder() int32 {
	if m != nil {
		return m.Order
	}
	return 0
}

func (m *Pipeline) GetStages() []*Stage {
	if m != nil {
		return m.Stages
	}
	return nil
}

type PipelineGet struct {
	Application string `protobuf:"bytes,1,opt,name=Application" json:"Application,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
}

func (m *PipelineGet) Reset()                    { *m = PipelineGet{} }
func (m *PipelineGet) String() string            { return proto.CompactTextString(m) }
func (*PipelineGet) ProtoMessage()               {}
func (*PipelineGet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PipelineGet) GetApplication() string {
	if m != nil {
		return m.Application
	}
	return ""
}

func (m *PipelineGet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PipelineList struct {
	Pipelines []*Pipeline `protobuf:"bytes,1,rep,name=Pipelines" json:"Pipelines,omitempty"`
}

func (m *PipelineList) Reset()                    { *m = PipelineList{} }
func (m *PipelineList) String() string            { return proto.CompactTextString(m) }
func (*PipelineList) ProtoMessage()               {}
func (*PipelineList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *PipelineList) GetPipelines() []*Pipeline {
	if m != nil {
		return m.Pipelines
	}
	return nil
}

type PipelineCreate struct {
	Application string `protobuf:"bytes,1,opt,name=Application" json:"Application,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	// optional
	Stages []*Stage `protobuf:"bytes,3,rep,name=Stages" json:"Stages,omitempty"`
}

func (m *PipelineCreate) Reset()                    { *m = PipelineCreate{} }
func (m *PipelineCreate) String() string            { return proto.CompactTextString(m) }
func (*PipelineCreate) ProtoMessage()               {}
func (*PipelineCreate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *PipelineCreate) GetApplication() string {
	if m != nil {
		return m.Application
	}
	return ""
}

func (m *PipelineCreate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PipelineCreate) GetStages() []*Stage {
	if m != nil {
		return m.Stages
	}
	return nil
}

type Stage struct {
	ID               int32  `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Name             string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Type             string `protobuf:"bytes,3,opt,name=Type" json:"Type,omitempty"`
	Order            int32  `protobuf:"varint,4,opt,name=Order" json:"Order,omitempty"`
	Strategy         string `protobuf:"bytes,5,opt,name=Strategy" json:"Strategy,omitempty"`
	Cluster          string `protobuf:"bytes,6,opt,name=Cluster" json:"Cluster,omitempty"`
	Namespace        string `protobuf:"bytes,7,opt,name=Namespace" json:"Namespace,omitempty"`
	Memory           string `protobuf:"bytes,8,opt,name=Memory" json:"Memory,omitempty"`
	Proc             string `protobuf:"bytes,9,opt,name=Proc" json:"Proc,omitempty"`
	Replicas         int32  `protobuf:"varint,10,opt,name=Replicas" json:"Replicas,omitempty"`
	MaxUnavailable   int32  `protobuf:"varint,11,opt,name=MaxUnavailable" json:"MaxUnavailable,omitempty"`
	MaxSurge         int32  `protobuf:"varint,12,opt,name=MaxSurge" json:"MaxSurge,omitempty"`
	NotifyEmail      bool   `protobuf:"varint,13,opt,name=NotifyEmail" json:"NotifyEmail,omitempty"`
	NotifySlack      bool   `protobuf:"varint,14,opt,name=NotifySlack" json:"NotifySlack,omitempty"`
	SmokeHeaderKey   string `protobuf:"bytes,15,opt,name=SmokeHeaderKey" json:"SmokeHeaderKey,omitempty"`
	SmokeHeaderVal   string `protobuf:"bytes,16,opt,name=SmokeHeaderVal" json:"SmokeHeaderVal,omitempty"`
	CanaryPercentage int32  `protobuf:"varint,17,opt,name=CanaryPercentage" json:"CanaryPercentage,omitempty"`
	WaitDuration     int64  `protobuf:"varint,18,opt,name=WaitDuration" json:"WaitDuration,omitempty"`
}

func (m *Stage) Reset()                    { *m = Stage{} }
func (m *Stage) String() string            { return proto.CompactTextString(m) }
func (*Stage) ProtoMessage()               {}
func (*Stage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Stage) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Stage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Stage) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Stage) GetOrder() int32 {
	if m != nil {
		return m.Order
	}
	return 0
}

func (m *Stage) GetStrategy() string {
	if m != nil {
		return m.Strategy
	}
	return ""
}

func (m *Stage) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *Stage) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Stage) GetMemory() string {
	if m != nil {
		return m.Memory
	}
	return ""
}

func (m *Stage) GetProc() string {
	if m != nil {
		return m.Proc
	}
	return ""
}

func (m *Stage) GetReplicas() int32 {
	if m != nil {
		return m.Replicas
	}
	return 0
}

func (m *Stage) GetMaxUnavailable() int32 {
	if m != nil {
		return m.MaxUnavailable
	}
	return 0
}

func (m *Stage) GetMaxSurge() int32 {
	if m != nil {
		return m.MaxSurge
	}
	return 0
}

func (m *Stage) GetNotifyEmail() bool {
	if m != nil {
		return m.NotifyEmail
	}
	return false
}

func (m *Stage) GetNotifySlack() bool {
	if m != nil {
		return m.NotifySlack
	}
	return false
}

func (m *Stage) GetSmokeHeaderKey() string {
	if m != nil {
		return m.SmokeHeaderKey
	}
	return ""
}

func (m *Stage) GetSmokeHeaderVal() string {
	if m != nil {
		return m.SmokeHeaderVal
	}
	return ""
}

func (m *Stage) GetCanaryPercentage() int32 {
	if m != nil {
		return m.CanaryPercentage
	}
	return 0
}

func (m *Stage) GetWaitDuration() int64 {
	if m != nil {
		return m.WaitDuration
	}
	return 0
}

type StageGet struct {
	Pipeline string `protobuf:"bytes,1,opt,name=Pipeline" json:"Pipeline,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
}

func (m *StageGet) Reset()                    { *m = StageGet{} }
func (m *StageGet) String() string            { return proto.CompactTextString(m) }
func (*StageGet) ProtoMessage()               {}
func (*StageGet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *StageGet) GetPipeline() string {
	if m != nil {
		return m.Pipeline
	}
	return ""
}

func (m *StageGet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type StageList struct {
	Stages []*Stage `protobuf:"bytes,1,rep,name=Stages" json:"Stages,omitempty"`
}

func (m *StageList) Reset()                    { *m = StageList{} }
func (m *StageList) String() string            { return proto.CompactTextString(m) }
func (*StageList) ProtoMessage()               {}
func (*StageList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *StageList) GetStages() []*Stage {
	if m != nil {
		return m.Stages
	}
	return nil
}

type Env struct {
	ID        int32  `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Key       string `protobuf:"bytes,3,opt,name=Key" json:"Key,omitempty"`
	Value     string `protobuf:"bytes,4,opt,name=Value" json:"Value,omitempty"`
	ConfigMap string `protobuf:"bytes,5,opt,name=ConfigMap" json:"ConfigMap,omitempty"`
	Secret    string `protobuf:"bytes,6,opt,name=Secret" json:"Secret,omitempty"`
}

func (m *Env) Reset()                    { *m = Env{} }
func (m *Env) String() string            { return proto.CompactTextString(m) }
func (*Env) ProtoMessage()               {}
func (*Env) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *Env) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Env) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Env) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Env) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Env) GetConfigMap() string {
	if m != nil {
		return m.ConfigMap
	}
	return ""
}

func (m *Env) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

type Cluster struct {
	ID   int32  `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	URL  string `protobuf:"bytes,3,opt,name=URL" json:"URL,omitempty"`
}

func (m *Cluster) Reset()                    { *m = Cluster{} }
func (m *Cluster) String() string            { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()               {}
func (*Cluster) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *Cluster) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Cluster) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Cluster) GetURL() string {
	if m != nil {
		return m.URL
	}
	return ""
}

type ClusterGet struct {
	Name string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
}

func (m *ClusterGet) Reset()                    { *m = ClusterGet{} }
func (m *ClusterGet) String() string            { return proto.CompactTextString(m) }
func (*ClusterGet) ProtoMessage()               {}
func (*ClusterGet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *ClusterGet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ClusterList struct {
	Clusters []*Cluster `protobuf:"bytes,1,rep,name=Clusters" json:"Clusters,omitempty"`
}

func (m *ClusterList) Reset()                    { *m = ClusterList{} }
func (m *ClusterList) String() string            { return proto.CompactTextString(m) }
func (*ClusterList) ProtoMessage()               {}
func (*ClusterList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *ClusterList) GetClusters() []*Cluster {
	if m != nil {
		return m.Clusters
	}
	return nil
}

type ClusterCreate struct {
	Name string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	URL  string `protobuf:"bytes,2,opt,name=URL" json:"URL,omitempty"`
}

func (m *ClusterCreate) Reset()                    { *m = ClusterCreate{} }
func (m *ClusterCreate) String() string            { return proto.CompactTextString(m) }
func (*ClusterCreate) ProtoMessage()               {}
func (*ClusterCreate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *ClusterCreate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ClusterCreate) GetURL() string {
	if m != nil {
		return m.URL
	}
	return ""
}

type Service struct {
	ID          int32       `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Name        string      `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Application string      `protobuf:"bytes,3,opt,name=Application" json:"Application,omitempty"`
	Type        ServiceType `protobuf:"varint,4,opt,name=Type,enum=ploioproto.ServiceType" json:"Type,omitempty"`
	Ports       []*Port     `protobuf:"bytes,5,rep,name=Ports" json:"Ports,omitempty"`
}

func (m *Service) Reset()                    { *m = Service{} }
func (m *Service) String() string            { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()               {}
func (*Service) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *Service) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetApplication() string {
	if m != nil {
		return m.Application
	}
	return ""
}

func (m *Service) GetType() ServiceType {
	if m != nil {
		return m.Type
	}
	return ServiceType_ClusterIP
}

func (m *Service) GetPorts() []*Port {
	if m != nil {
		return m.Ports
	}
	return nil
}

type Port struct {
	Name       string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Protocol   string `protobuf:"bytes,2,opt,name=Protocol" json:"Protocol,omitempty"`
	Port       string `protobuf:"bytes,3,opt,name=Port" json:"Port,omitempty"`
	TargetPort string `protobuf:"bytes,4,opt,name=TargetPort" json:"TargetPort,omitempty"`
	NodePort   string `protobuf:"bytes,5,opt,name=NodePort" json:"NodePort,omitempty"`
}

func (m *Port) Reset()                    { *m = Port{} }
func (m *Port) String() string            { return proto.CompactTextString(m) }
func (*Port) ProtoMessage()               {}
func (*Port) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *Port) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Port) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *Port) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *Port) GetTargetPort() string {
	if m != nil {
		return m.TargetPort
	}
	return ""
}

func (m *Port) GetNodePort() string {
	if m != nil {
		return m.NodePort
	}
	return ""
}

func init() {
	proto.RegisterType((*ApplicationCreate)(nil), "ploioproto.ApplicationCreate")
	proto.RegisterType((*ApplicationGet)(nil), "ploioproto.ApplicationGet")
	proto.RegisterType((*ApplicationList)(nil), "ploioproto.ApplicationList")
	proto.RegisterType((*Application)(nil), "ploioproto.Application")
	proto.RegisterType((*Pipeline)(nil), "ploioproto.Pipeline")
	proto.RegisterType((*PipelineGet)(nil), "ploioproto.PipelineGet")
	proto.RegisterType((*PipelineList)(nil), "ploioproto.PipelineList")
	proto.RegisterType((*PipelineCreate)(nil), "ploioproto.PipelineCreate")
	proto.RegisterType((*Stage)(nil), "ploioproto.Stage")
	proto.RegisterType((*StageGet)(nil), "ploioproto.StageGet")
	proto.RegisterType((*StageList)(nil), "ploioproto.StageList")
	proto.RegisterType((*Env)(nil), "ploioproto.Env")
	proto.RegisterType((*Cluster)(nil), "ploioproto.Cluster")
	proto.RegisterType((*ClusterGet)(nil), "ploioproto.ClusterGet")
	proto.RegisterType((*ClusterList)(nil), "ploioproto.ClusterList")
	proto.RegisterType((*ClusterCreate)(nil), "ploioproto.ClusterCreate")
	proto.RegisterType((*Service)(nil), "ploioproto.Service")
	proto.RegisterType((*Port)(nil), "ploioproto.Port")
	proto.RegisterEnum("ploioproto.StageType", StageType_name, StageType_value)
	proto.RegisterEnum("ploioproto.ServiceType", ServiceType_name, ServiceType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PloioAPI service

type PloioAPIClient interface {
	ListApplications(ctx context.Context, in *ApplicationGet, opts ...grpc.CallOption) (*ApplicationList, error)
	GetApplication(ctx context.Context, in *ApplicationGet, opts ...grpc.CallOption) (*Application, error)
	CreateApplication(ctx context.Context, in *ApplicationCreate, opts ...grpc.CallOption) (*Application, error)
	ListClusters(ctx context.Context, in *ClusterGet, opts ...grpc.CallOption) (*ClusterList, error)
	ListPipelines(ctx context.Context, in *PipelineGet, opts ...grpc.CallOption) (*PipelineList, error)
	CreatePipeline(ctx context.Context, in *PipelineCreate, opts ...grpc.CallOption) (*Pipeline, error)
	CreateStage(ctx context.Context, in *Stage, opts ...grpc.CallOption) (*Stage, error)
	ListStage(ctx context.Context, in *StageGet, opts ...grpc.CallOption) (*StageList, error)
}

type ploioAPIClient struct {
	cc *grpc.ClientConn
}

func NewPloioAPIClient(cc *grpc.ClientConn) PloioAPIClient {
	return &ploioAPIClient{cc}
}

func (c *ploioAPIClient) ListApplications(ctx context.Context, in *ApplicationGet, opts ...grpc.CallOption) (*ApplicationList, error) {
	out := new(ApplicationList)
	err := grpc.Invoke(ctx, "/ploioproto.PloioAPI/ListApplications", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ploioAPIClient) GetApplication(ctx context.Context, in *ApplicationGet, opts ...grpc.CallOption) (*Application, error) {
	out := new(Application)
	err := grpc.Invoke(ctx, "/ploioproto.PloioAPI/GetApplication", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ploioAPIClient) CreateApplication(ctx context.Context, in *ApplicationCreate, opts ...grpc.CallOption) (*Application, error) {
	out := new(Application)
	err := grpc.Invoke(ctx, "/ploioproto.PloioAPI/CreateApplication", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ploioAPIClient) ListClusters(ctx context.Context, in *ClusterGet, opts ...grpc.CallOption) (*ClusterList, error) {
	out := new(ClusterList)
	err := grpc.Invoke(ctx, "/ploioproto.PloioAPI/ListClusters", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ploioAPIClient) ListPipelines(ctx context.Context, in *PipelineGet, opts ...grpc.CallOption) (*PipelineList, error) {
	out := new(PipelineList)
	err := grpc.Invoke(ctx, "/ploioproto.PloioAPI/ListPipelines", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ploioAPIClient) CreatePipeline(ctx context.Context, in *PipelineCreate, opts ...grpc.CallOption) (*Pipeline, error) {
	out := new(Pipeline)
	err := grpc.Invoke(ctx, "/ploioproto.PloioAPI/CreatePipeline", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ploioAPIClient) CreateStage(ctx context.Context, in *Stage, opts ...grpc.CallOption) (*Stage, error) {
	out := new(Stage)
	err := grpc.Invoke(ctx, "/ploioproto.PloioAPI/CreateStage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ploioAPIClient) ListStage(ctx context.Context, in *StageGet, opts ...grpc.CallOption) (*StageList, error) {
	out := new(StageList)
	err := grpc.Invoke(ctx, "/ploioproto.PloioAPI/ListStage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PloioAPI service

type PloioAPIServer interface {
	ListApplications(context.Context, *ApplicationGet) (*ApplicationList, error)
	GetApplication(context.Context, *ApplicationGet) (*Application, error)
	CreateApplication(context.Context, *ApplicationCreate) (*Application, error)
	ListClusters(context.Context, *ClusterGet) (*ClusterList, error)
	ListPipelines(context.Context, *PipelineGet) (*PipelineList, error)
	CreatePipeline(context.Context, *PipelineCreate) (*Pipeline, error)
	CreateStage(context.Context, *Stage) (*Stage, error)
	ListStage(context.Context, *StageGet) (*StageList, error)
}

func RegisterPloioAPIServer(s *grpc.Server, srv PloioAPIServer) {
	s.RegisterService(&_PloioAPI_serviceDesc, srv)
}

func _PloioAPI_ListApplications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationGet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PloioAPIServer).ListApplications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ploioproto.PloioAPI/ListApplications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PloioAPIServer).ListApplications(ctx, req.(*ApplicationGet))
	}
	return interceptor(ctx, in, info, handler)
}

func _PloioAPI_GetApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationGet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PloioAPIServer).GetApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ploioproto.PloioAPI/GetApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PloioAPIServer).GetApplication(ctx, req.(*ApplicationGet))
	}
	return interceptor(ctx, in, info, handler)
}

func _PloioAPI_CreateApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PloioAPIServer).CreateApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ploioproto.PloioAPI/CreateApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PloioAPIServer).CreateApplication(ctx, req.(*ApplicationCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _PloioAPI_ListClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterGet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PloioAPIServer).ListClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ploioproto.PloioAPI/ListClusters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PloioAPIServer).ListClusters(ctx, req.(*ClusterGet))
	}
	return interceptor(ctx, in, info, handler)
}

func _PloioAPI_ListPipelines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PipelineGet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PloioAPIServer).ListPipelines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ploioproto.PloioAPI/ListPipelines",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PloioAPIServer).ListPipelines(ctx, req.(*PipelineGet))
	}
	return interceptor(ctx, in, info, handler)
}

func _PloioAPI_CreatePipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PipelineCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PloioAPIServer).CreatePipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ploioproto.PloioAPI/CreatePipeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PloioAPIServer).CreatePipeline(ctx, req.(*PipelineCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _PloioAPI_CreateStage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Stage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PloioAPIServer).CreateStage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ploioproto.PloioAPI/CreateStage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PloioAPIServer).CreateStage(ctx, req.(*Stage))
	}
	return interceptor(ctx, in, info, handler)
}

func _PloioAPI_ListStage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StageGet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PloioAPIServer).ListStage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ploioproto.PloioAPI/ListStage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PloioAPIServer).ListStage(ctx, req.(*StageGet))
	}
	return interceptor(ctx, in, info, handler)
}

var _PloioAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ploioproto.PloioAPI",
	HandlerType: (*PloioAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListApplications",
			Handler:    _PloioAPI_ListApplications_Handler,
		},
		{
			MethodName: "GetApplication",
			Handler:    _PloioAPI_GetApplication_Handler,
		},
		{
			MethodName: "CreateApplication",
			Handler:    _PloioAPI_CreateApplication_Handler,
		},
		{
			MethodName: "ListClusters",
			Handler:    _PloioAPI_ListClusters_Handler,
		},
		{
			MethodName: "ListPipelines",
			Handler:    _PloioAPI_ListPipelines_Handler,
		},
		{
			MethodName: "CreatePipeline",
			Handler:    _PloioAPI_CreatePipeline_Handler,
		},
		{
			MethodName: "CreateStage",
			Handler:    _PloioAPI_CreateStage_Handler,
		},
		{
			MethodName: "ListStage",
			Handler:    _PloioAPI_ListStage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1040 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xef, 0x6e, 0xe3, 0x44,
	0x10, 0x8f, 0xe3, 0x24, 0x4d, 0x26, 0x69, 0xce, 0x5d, 0xca, 0xdd, 0x2a, 0xfc, 0x51, 0xf0, 0x87,
	0x53, 0x29, 0x52, 0x4f, 0x2a, 0x1c, 0x1f, 0x8a, 0x04, 0xea, 0x25, 0x55, 0xa9, 0x68, 0x42, 0xe4,
	0xb4, 0xe5, 0xf3, 0x9e, 0xb3, 0x17, 0xac, 0xba, 0x5e, 0xb3, 0x71, 0x7b, 0xcd, 0x0b, 0xc0, 0x33,
	0xf0, 0x1d, 0x9e, 0x80, 0x47, 0xe2, 0x45, 0xd0, 0xac, 0xd7, 0xf6, 0xa6, 0x71, 0x7b, 0x11, 0xdf,
	0xe6, 0xcf, 0xcf, 0xb3, 0xbf, 0x99, 0x9d, 0x99, 0x35, 0xb4, 0x58, 0x1c, 0x1c, 0xc4, 0x52, 0x24,
	0x82, 0x40, 0x1c, 0x8a, 0x40, 0x28, 0xd9, 0xfd, 0xdb, 0x82, 0x9d, 0xe3, 0x38, 0x0e, 0x03, 0x9f,
	0x25, 0x81, 0x88, 0x06, 0x92, 0xb3, 0x84, 0x13, 0x02, 0xb5, 0x31, 0xbb, 0xe1, 0xd4, 0xea, 0x5b,
	0x7b, 0x2d, 0x4f, 0xc9, 0x64, 0x17, 0xea, 0x3f, 0xbf, 0x8f, 0xb8, 0xa4, 0x55, 0x65, 0x4c, 0x15,
	0x44, 0x7a, 0x3c, 0x16, 0xd4, 0x4e, 0x91, 0x28, 0x93, 0x2f, 0xc0, 0x3e, 0x89, 0xee, 0x68, 0xa3,
	0x6f, 0xef, 0xb5, 0x0f, 0x9f, 0x1d, 0x14, 0xa7, 0x1d, 0x9c, 0x44, 0x77, 0x1e, 0xfa, 0xc8, 0x2b,
	0x68, 0x4e, 0xb9, 0xbc, 0x0b, 0x7c, 0xbe, 0xa0, 0x5b, 0x0a, 0xf7, 0x91, 0x89, 0xd3, 0x3e, 0x2f,
	0x07, 0xb9, 0xdf, 0x40, 0xd7, 0xa0, 0x79, 0xca, 0x13, 0xd2, 0x85, 0xea, 0xd9, 0x50, 0x31, 0xac,
	0x7b, 0xd5, 0xb3, 0x61, 0xce, 0xb9, 0x5a, 0x70, 0x76, 0xc7, 0xf0, 0xcc, 0xf8, 0xea, 0x3c, 0x58,
	0x24, 0xe4, 0x3b, 0xe8, 0x18, 0xa6, 0x05, 0xb5, 0xd4, 0xe9, 0x2f, 0xcc, 0xd3, 0x0d, 0xbf, 0xb7,
	0x02, 0x76, 0xff, 0xb5, 0xa0, 0x6d, 0x18, 0x36, 0xe1, 0x50, 0xd4, 0xcd, 0x2e, 0xab, 0x5b, 0x6d,
	0xbd, 0x6e, 0xf5, 0x0d, 0xeb, 0xd6, 0xd8, 0xa0, 0x6e, 0xe4, 0x10, 0x5a, 0x93, 0x20, 0xe6, 0x61,
	0x10, 0xe5, 0x95, 0xde, 0x35, 0xbf, 0xc8, 0x9c, 0x5e, 0x01, 0x73, 0x05, 0x34, 0x33, 0x65, 0xe3,
	0x0c, 0xe5, 0x4c, 0x67, 0x58, 0xf7, 0x52, 0x85, 0x7c, 0x09, 0x8d, 0x69, 0xc2, 0xe6, 0x7c, 0x41,
	0x6b, 0xea, 0xd8, 0x9d, 0x15, 0xa2, 0xe8, 0xf1, 0x34, 0xc0, 0x1d, 0x40, 0x3b, 0x3b, 0x10, 0x6f,
	0xb6, 0xbf, 0x52, 0x64, 0xdd, 0x84, 0x2b, 0x75, 0x2f, 0xbb, 0xeb, 0x37, 0xd0, 0xc9, 0x82, 0xa8,
	0x8b, 0x5e, 0xc9, 0xdc, 0xda, 0x2c, 0xf3, 0xdf, 0xa0, 0x9b, 0x29, 0x7a, 0x12, 0xfe, 0x17, 0x17,
	0x23, 0x77, 0xfb, 0x43, 0xb9, 0xff, 0x59, 0x83, 0xba, 0x12, 0x37, 0x2a, 0x35, 0x81, 0xda, 0xc5,
	0x32, 0xe6, 0xd9, 0xb8, 0xa1, 0x5c, 0x94, 0xbf, 0x66, 0x96, 0xbf, 0x07, 0xcd, 0x69, 0x22, 0x59,
	0xc2, 0xe7, 0x4b, 0x5a, 0x57, 0xe8, 0x5c, 0x27, 0x14, 0xb6, 0x06, 0xe1, 0xed, 0x22, 0xe1, 0x92,
	0x36, 0x94, 0x2b, 0x53, 0xc9, 0xa7, 0xd0, 0xc2, 0x73, 0x16, 0x31, 0xf3, 0x39, 0xdd, 0x52, 0xbe,
	0xc2, 0x40, 0x9e, 0x43, 0x63, 0xc4, 0x6f, 0x84, 0x5c, 0xd2, 0xa6, 0x72, 0x69, 0x0d, 0x59, 0x4d,
	0xa4, 0xf0, 0x69, 0x2b, 0x65, 0x85, 0x32, 0x9e, 0xef, 0x71, 0x55, 0xa5, 0x05, 0x05, 0x45, 0x2c,
	0xd7, 0xc9, 0x4b, 0xe8, 0x8e, 0xd8, 0xfd, 0x65, 0xc4, 0xee, 0x58, 0x10, 0xb2, 0xb7, 0x21, 0xa7,
	0x6d, 0x85, 0x78, 0x60, 0xc5, 0x18, 0x23, 0x76, 0x3f, 0xbd, 0x95, 0x73, 0x4e, 0x3b, 0x69, 0x8c,
	0x4c, 0xc7, 0x8b, 0x19, 0x8b, 0x24, 0x78, 0xb7, 0x3c, 0xb9, 0x61, 0x41, 0x48, 0xb7, 0xfb, 0xd6,
	0x5e, 0xd3, 0x33, 0x4d, 0x05, 0x62, 0x1a, 0x32, 0xff, 0x9a, 0x76, 0x4d, 0x84, 0x32, 0x21, 0x8f,
	0xe9, 0x8d, 0xb8, 0xe6, 0x3f, 0x72, 0x36, 0xe3, 0xf2, 0x27, 0xbe, 0xa4, 0xcf, 0x54, 0x06, 0x0f,
	0xac, 0x0f, 0x70, 0x57, 0x2c, 0xa4, 0xce, 0x1a, 0xee, 0x8a, 0x85, 0x64, 0x1f, 0x9c, 0x01, 0x8b,
	0x98, 0x5c, 0x4e, 0xb8, 0xf4, 0x79, 0x84, 0xb7, 0x4a, 0x77, 0x14, 0xef, 0x35, 0x3b, 0x71, 0xa1,
	0xf3, 0x0b, 0x0b, 0x92, 0xe1, 0xad, 0x4c, 0x3b, 0x8b, 0xf4, 0xad, 0x3d, 0xdb, 0x5b, 0xb1, 0xb9,
	0x47, 0x78, 0x87, 0x6c, 0xae, 0x86, 0xa2, 0x57, 0x0c, 0xa5, 0xee, 0xc2, 0x62, 0x48, 0xcb, 0xc6,
	0xe1, 0x5b, 0x68, 0xa9, 0x6f, 0xd5, 0x2c, 0x14, 0xfd, 0x68, 0x7d, 0xa8, 0x1f, 0xff, 0xb0, 0xd4,
	0x16, 0xda, 0xa8, 0x1b, 0x1d, 0xb0, 0xb1, 0x68, 0x69, 0x33, 0xa2, 0x88, 0xbd, 0x78, 0xc5, 0xc2,
	0x5b, 0xae, 0xf7, 0x5a, 0xaa, 0x60, 0x57, 0x0d, 0x44, 0xf4, 0x2e, 0x98, 0x8f, 0x58, 0xac, 0x9b,
	0xb1, 0x30, 0x60, 0x57, 0x4d, 0xb9, 0x2f, 0x79, 0xa2, 0x9b, 0x51, 0x6b, 0xee, 0x0f, 0x79, 0x97,
	0x6e, 0x4a, 0xe6, 0xd2, 0x3b, 0xcf, 0xc8, 0x5c, 0x7a, 0xe7, 0x6e, 0x1f, 0x40, 0x07, 0xc0, 0x02,
	0x96, 0xbc, 0x69, 0xee, 0xf7, 0xd0, 0xd6, 0x08, 0x55, 0xa6, 0x57, 0xd0, 0xd4, 0x6a, 0x56, 0xa8,
	0x95, 0xed, 0xaa, 0x7d, 0x5e, 0x0e, 0x72, 0x5f, 0xc3, 0xb6, 0x96, 0x9f, 0x78, 0x38, 0x35, 0xb1,
	0x6a, 0x41, 0xec, 0x2f, 0x0b, 0xb6, 0xf4, 0x86, 0xde, 0x28, 0xb5, 0x07, 0x4b, 0xc8, 0x5e, 0x5f,
	0x42, 0x5f, 0xe9, 0xbd, 0x80, 0x65, 0xef, 0xae, 0xbe, 0x66, 0xfa, 0x20, 0x74, 0xeb, 0x85, 0xf1,
	0x12, 0xea, 0x13, 0x21, 0x93, 0x85, 0x7e, 0x69, 0x9c, 0x95, 0xad, 0x28, 0x64, 0xe2, 0xa5, 0x6e,
	0xf7, 0x77, 0x0b, 0x6a, 0x28, 0x95, 0x66, 0x85, 0xfd, 0x88, 0x5f, 0xf8, 0x22, 0xd4, 0x5c, 0x73,
	0x5d, 0xed, 0x03, 0x21, 0x93, 0x6c, 0x4b, 0xa9, 0x18, 0x9f, 0x03, 0x5c, 0x30, 0x39, 0xe7, 0x89,
	0xf2, 0xa4, 0xed, 0x61, 0x58, 0x30, 0xde, 0x58, 0xcc, 0xb8, 0xf2, 0xea, 0x7d, 0x95, 0xe9, 0xfb,
	0x47, 0xba, 0x97, 0x15, 0x7b, 0x80, 0xc6, 0x90, 0xc7, 0xa1, 0x58, 0x3a, 0x15, 0xd2, 0x82, 0xfa,
	0xe0, 0x57, 0xee, 0x5f, 0x3b, 0x16, 0xe9, 0x40, 0xf3, 0x82, 0x33, 0x39, 0x13, 0xef, 0x23, 0xa7,
	0x4a, 0x9a, 0x50, 0xc3, 0x49, 0x72, 0xec, 0xfd, 0x31, 0xb4, 0x8d, 0x0a, 0x90, 0x6d, 0x68, 0xe9,
	0x1b, 0x3b, 0x9b, 0x38, 0x15, 0xfc, 0x2a, 0x3b, 0xc5, 0xb1, 0x88, 0x03, 0x9d, 0x73, 0xc1, 0x66,
	0x6f, 0x58, 0xc8, 0x22, 0x9f, 0x4b, 0xa7, 0x8a, 0x96, 0x93, 0xfb, 0x84, 0xcb, 0x88, 0x85, 0x98,
	0xb5, 0x63, 0x1f, 0xfe, 0x53, 0x83, 0xe6, 0x04, 0xeb, 0x75, 0x3c, 0x39, 0x23, 0x23, 0x70, 0xb0,
	0x71, 0xcc, 0x7f, 0x04, 0xd2, 0x7b, 0xe4, 0x57, 0xe2, 0x94, 0x27, 0xbd, 0x4f, 0x1e, 0xf1, 0x61,
	0x10, 0xb7, 0x42, 0x4e, 0xa1, 0x7b, 0xca, 0xcd, 0x68, 0x4f, 0x06, 0x7b, 0xec, 0x9f, 0xc5, 0xad,
	0x90, 0x11, 0xec, 0xa4, 0x0d, 0x69, 0xc6, 0xfa, 0xec, 0x11, 0x7c, 0x8a, 0x7c, 0x2a, 0xdc, 0x31,
	0x74, 0x90, 0x61, 0xd6, 0xf6, 0xe4, 0x79, 0xc9, 0x54, 0xac, 0x31, 0x32, 0x06, 0xcb, 0xad, 0x90,
	0x21, 0x6c, 0xa3, 0x94, 0x3f, 0xb5, 0xe4, 0x45, 0xd9, 0x5b, 0x8c, 0x41, 0x68, 0x99, 0x23, 0x8f,
	0xd2, 0x4d, 0xd9, 0xe6, 0xab, 0xaf, 0x57, 0x86, 0xd6, 0x19, 0x95, 0x3e, 0xf7, 0x6e, 0x85, 0xbc,
	0x86, 0x76, 0x8a, 0x48, 0xdf, 0xdd, 0xf5, 0x65, 0xd8, 0x5b, 0x37, 0xb9, 0x15, 0x72, 0x04, 0x2d,
	0xa4, 0x91, 0x7e, 0xb4, 0xbb, 0x86, 0x40, 0xee, 0x1f, 0xaf, 0x59, 0x53, 0xe2, 0x6f, 0x1b, 0xca,
	0xf4, 0xf5, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x27, 0xd9, 0x07, 0x37, 0x86, 0x0b, 0x00, 0x00,
}
