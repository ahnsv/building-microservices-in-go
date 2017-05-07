// Code generated by protoc-gen-go.
// source: google/api/servicemanagement/v1/resources.proto
// DO NOT EDIT!

/*
Package servicemanagement is a generated protocol buffer package.

It is generated from these files:
	google/api/servicemanagement/v1/resources.proto
	google/api/servicemanagement/v1/servicemanager.proto

It has these top-level messages:
	ManagedService
	OperationMetadata
	Diagnostic
	ConfigSource
	ConfigFile
	ConfigRef
	ChangeReport
	Rollout
	ListServicesRequest
	ListServicesResponse
	GetServiceRequest
	CreateServiceRequest
	DeleteServiceRequest
	UndeleteServiceRequest
	UndeleteServiceResponse
	GetServiceConfigRequest
	ListServiceConfigsRequest
	ListServiceConfigsResponse
	CreateServiceConfigRequest
	SubmitConfigSourceRequest
	SubmitConfigSourceResponse
	CreateServiceRolloutRequest
	ListServiceRolloutsRequest
	ListServiceRolloutsResponse
	GetServiceRolloutRequest
	EnableServiceRequest
	DisableServiceRequest
	GenerateConfigReportRequest
	GenerateConfigReportResponse
*/
package servicemanagement

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_api2 "google.golang.org/genproto/googleapis/api/configchange"
import _ "google.golang.org/genproto/googleapis/api/serviceconfig"
import _ "google.golang.org/genproto/googleapis/longrunning"
import _ "github.com/golang/protobuf/ptypes/any"
import _ "google.golang.org/genproto/protobuf/field_mask"
import _ "github.com/golang/protobuf/ptypes/struct"
import google_protobuf9 "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/rpc/status"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Code describes the status of one operation step.
type OperationMetadata_Status int32

const (
	// Unspecifed code.
	OperationMetadata_STATUS_UNSPECIFIED OperationMetadata_Status = 0
	// The step has completed without errors.
	OperationMetadata_DONE OperationMetadata_Status = 1
	// The step has not started yet.
	OperationMetadata_NOT_STARTED OperationMetadata_Status = 2
	// The step is in progress.
	OperationMetadata_IN_PROGRESS OperationMetadata_Status = 3
	// The step has completed with errors.
	OperationMetadata_FAILED OperationMetadata_Status = 4
	// The step has completed with cancellation.
	OperationMetadata_CANCELLED OperationMetadata_Status = 5
)

var OperationMetadata_Status_name = map[int32]string{
	0: "STATUS_UNSPECIFIED",
	1: "DONE",
	2: "NOT_STARTED",
	3: "IN_PROGRESS",
	4: "FAILED",
	5: "CANCELLED",
}
var OperationMetadata_Status_value = map[string]int32{
	"STATUS_UNSPECIFIED": 0,
	"DONE":               1,
	"NOT_STARTED":        2,
	"IN_PROGRESS":        3,
	"FAILED":             4,
	"CANCELLED":          5,
}

func (x OperationMetadata_Status) String() string {
	return proto.EnumName(OperationMetadata_Status_name, int32(x))
}
func (OperationMetadata_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

// The kind of diagnostic information possible.
type Diagnostic_Kind int32

const (
	// Warnings and errors
	Diagnostic_WARNING Diagnostic_Kind = 0
	// Only errors
	Diagnostic_ERROR Diagnostic_Kind = 1
)

var Diagnostic_Kind_name = map[int32]string{
	0: "WARNING",
	1: "ERROR",
}
var Diagnostic_Kind_value = map[string]int32{
	"WARNING": 0,
	"ERROR":   1,
}

func (x Diagnostic_Kind) String() string {
	return proto.EnumName(Diagnostic_Kind_name, int32(x))
}
func (Diagnostic_Kind) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type ConfigFile_FileType int32

const (
	// Unknown file type.
	ConfigFile_FILE_TYPE_UNSPECIFIED ConfigFile_FileType = 0
	// YAML-specification of service.
	ConfigFile_SERVICE_CONFIG_YAML ConfigFile_FileType = 1
	// OpenAPI specification, serialized in JSON.
	ConfigFile_OPEN_API_JSON ConfigFile_FileType = 2
	// OpenAPI specification, serialized in YAML.
	ConfigFile_OPEN_API_YAML ConfigFile_FileType = 3
	// FileDescriptorSet, generated by protoc.
	//
	// To generate, use protoc with imports and source info included.
	// For an example test.proto file, the following command would put the value
	// in a new file named out.pb.
	//
	// $protoc --include_imports --include_source_info test.proto -o out.pb
	ConfigFile_FILE_DESCRIPTOR_SET_PROTO ConfigFile_FileType = 4
)

var ConfigFile_FileType_name = map[int32]string{
	0: "FILE_TYPE_UNSPECIFIED",
	1: "SERVICE_CONFIG_YAML",
	2: "OPEN_API_JSON",
	3: "OPEN_API_YAML",
	4: "FILE_DESCRIPTOR_SET_PROTO",
}
var ConfigFile_FileType_value = map[string]int32{
	"FILE_TYPE_UNSPECIFIED":     0,
	"SERVICE_CONFIG_YAML":       1,
	"OPEN_API_JSON":             2,
	"OPEN_API_YAML":             3,
	"FILE_DESCRIPTOR_SET_PROTO": 4,
}

func (x ConfigFile_FileType) String() string {
	return proto.EnumName(ConfigFile_FileType_name, int32(x))
}
func (ConfigFile_FileType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

// Status of a Rollout.
type Rollout_RolloutStatus int32

const (
	// No status specified.
	Rollout_ROLLOUT_STATUS_UNSPECIFIED Rollout_RolloutStatus = 0
	// The Rollout is in progress.
	Rollout_IN_PROGRESS Rollout_RolloutStatus = 1
	// The Rollout has completed successfully.
	Rollout_SUCCESS Rollout_RolloutStatus = 2
	// The Rollout has been cancelled. This can happen if you have overlapping
	// Rollout pushes, and the previous ones will be cancelled.
	Rollout_CANCELLED Rollout_RolloutStatus = 3
	// The Rollout has failed. It is typically caused by configuration errors.
	Rollout_FAILED Rollout_RolloutStatus = 4
	// The Rollout has not started yet and is pending for execution.
	Rollout_PENDING Rollout_RolloutStatus = 5
)

var Rollout_RolloutStatus_name = map[int32]string{
	0: "ROLLOUT_STATUS_UNSPECIFIED",
	1: "IN_PROGRESS",
	2: "SUCCESS",
	3: "CANCELLED",
	4: "FAILED",
	5: "PENDING",
}
var Rollout_RolloutStatus_value = map[string]int32{
	"ROLLOUT_STATUS_UNSPECIFIED": 0,
	"IN_PROGRESS":                1,
	"SUCCESS":                    2,
	"CANCELLED":                  3,
	"FAILED":                     4,
	"PENDING":                    5,
}

func (x Rollout_RolloutStatus) String() string {
	return proto.EnumName(Rollout_RolloutStatus_name, int32(x))
}
func (Rollout_RolloutStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{7, 0} }

// The full representation of a Service that is managed by
// Google Service Management.
type ManagedService struct {
	// The name of the service. See the [overview](/service-management/overview)
	// for naming requirements.
	ServiceName string `protobuf:"bytes,2,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
	// ID of the project that produces and owns this service.
	ProducerProjectId string `protobuf:"bytes,3,opt,name=producer_project_id,json=producerProjectId" json:"producer_project_id,omitempty"`
}

func (m *ManagedService) Reset()                    { *m = ManagedService{} }
func (m *ManagedService) String() string            { return proto.CompactTextString(m) }
func (*ManagedService) ProtoMessage()               {}
func (*ManagedService) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ManagedService) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *ManagedService) GetProducerProjectId() string {
	if m != nil {
		return m.ProducerProjectId
	}
	return ""
}

// The metadata associated with a long running operation resource.
type OperationMetadata struct {
	// The full name of the resources that this operation is directly
	// associated with.
	ResourceNames []string `protobuf:"bytes,1,rep,name=resource_names,json=resourceNames" json:"resource_names,omitempty"`
	// Detailed status information for each step. The order is undetermined.
	Steps []*OperationMetadata_Step `protobuf:"bytes,2,rep,name=steps" json:"steps,omitempty"`
	// Percentage of completion of this operation, ranging from 0 to 100.
	ProgressPercentage int32 `protobuf:"varint,3,opt,name=progress_percentage,json=progressPercentage" json:"progress_percentage,omitempty"`
	// The start time of the operation.
	StartTime *google_protobuf9.Timestamp `protobuf:"bytes,4,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
}

func (m *OperationMetadata) Reset()                    { *m = OperationMetadata{} }
func (m *OperationMetadata) String() string            { return proto.CompactTextString(m) }
func (*OperationMetadata) ProtoMessage()               {}
func (*OperationMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *OperationMetadata) GetResourceNames() []string {
	if m != nil {
		return m.ResourceNames
	}
	return nil
}

func (m *OperationMetadata) GetSteps() []*OperationMetadata_Step {
	if m != nil {
		return m.Steps
	}
	return nil
}

func (m *OperationMetadata) GetProgressPercentage() int32 {
	if m != nil {
		return m.ProgressPercentage
	}
	return 0
}

func (m *OperationMetadata) GetStartTime() *google_protobuf9.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

// Represents the status of one operation step.
type OperationMetadata_Step struct {
	// The short description of the step.
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	// The status code.
	Status OperationMetadata_Status `protobuf:"varint,4,opt,name=status,enum=google.api.servicemanagement.v1.OperationMetadata_Status" json:"status,omitempty"`
}

func (m *OperationMetadata_Step) Reset()                    { *m = OperationMetadata_Step{} }
func (m *OperationMetadata_Step) String() string            { return proto.CompactTextString(m) }
func (*OperationMetadata_Step) ProtoMessage()               {}
func (*OperationMetadata_Step) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *OperationMetadata_Step) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *OperationMetadata_Step) GetStatus() OperationMetadata_Status {
	if m != nil {
		return m.Status
	}
	return OperationMetadata_STATUS_UNSPECIFIED
}

// Represents a diagnostic message (error or warning)
type Diagnostic struct {
	// File name and line number of the error or warning.
	Location string `protobuf:"bytes,1,opt,name=location" json:"location,omitempty"`
	// The kind of diagnostic information provided.
	Kind Diagnostic_Kind `protobuf:"varint,2,opt,name=kind,enum=google.api.servicemanagement.v1.Diagnostic_Kind" json:"kind,omitempty"`
	// Message describing the error or warning.
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *Diagnostic) Reset()                    { *m = Diagnostic{} }
func (m *Diagnostic) String() string            { return proto.CompactTextString(m) }
func (*Diagnostic) ProtoMessage()               {}
func (*Diagnostic) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Diagnostic) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *Diagnostic) GetKind() Diagnostic_Kind {
	if m != nil {
		return m.Kind
	}
	return Diagnostic_WARNING
}

func (m *Diagnostic) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// Represents a source file which is used to generate the service configuration
// defined by `google.api.Service`.
type ConfigSource struct {
	// A unique ID for a specific instance of this message, typically assigned
	// by the client for tracking purpose. If empty, the server may choose to
	// generate one instead.
	Id string `protobuf:"bytes,5,opt,name=id" json:"id,omitempty"`
	// Set of source configuration files that are used to generate a service
	// configuration (`google.api.Service`).
	Files []*ConfigFile `protobuf:"bytes,2,rep,name=files" json:"files,omitempty"`
}

func (m *ConfigSource) Reset()                    { *m = ConfigSource{} }
func (m *ConfigSource) String() string            { return proto.CompactTextString(m) }
func (*ConfigSource) ProtoMessage()               {}
func (*ConfigSource) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ConfigSource) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ConfigSource) GetFiles() []*ConfigFile {
	if m != nil {
		return m.Files
	}
	return nil
}

// Generic specification of a source configuration file
type ConfigFile struct {
	// The file name of the configuration file (full or relative path).
	FilePath string `protobuf:"bytes,1,opt,name=file_path,json=filePath" json:"file_path,omitempty"`
	// The bytes that constitute the file.
	FileContents []byte `protobuf:"bytes,3,opt,name=file_contents,json=fileContents,proto3" json:"file_contents,omitempty"`
	// The type of configuration file this represents.
	FileType ConfigFile_FileType `protobuf:"varint,4,opt,name=file_type,json=fileType,enum=google.api.servicemanagement.v1.ConfigFile_FileType" json:"file_type,omitempty"`
}

func (m *ConfigFile) Reset()                    { *m = ConfigFile{} }
func (m *ConfigFile) String() string            { return proto.CompactTextString(m) }
func (*ConfigFile) ProtoMessage()               {}
func (*ConfigFile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ConfigFile) GetFilePath() string {
	if m != nil {
		return m.FilePath
	}
	return ""
}

func (m *ConfigFile) GetFileContents() []byte {
	if m != nil {
		return m.FileContents
	}
	return nil
}

func (m *ConfigFile) GetFileType() ConfigFile_FileType {
	if m != nil {
		return m.FileType
	}
	return ConfigFile_FILE_TYPE_UNSPECIFIED
}

// Represents a service configuration with its name and id.
type ConfigRef struct {
	// Resource name of a service config. It must have the following
	// format: "services/{service name}/configs/{config id}".
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *ConfigRef) Reset()                    { *m = ConfigRef{} }
func (m *ConfigRef) String() string            { return proto.CompactTextString(m) }
func (*ConfigRef) ProtoMessage()               {}
func (*ConfigRef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ConfigRef) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Change report associated with a particular service configuration.
//
// It contains a list of ConfigChanges based on the comparison between
// two service configurations.
type ChangeReport struct {
	// List of changes between two service configurations.
	// The changes will be alphabetically sorted based on the identifier
	// of each change.
	// A ConfigChange identifier is a dot separated path to the configuration.
	// Example: visibility.rules[selector='LibraryService.CreateBook'].restriction
	ConfigChanges []*google_api2.ConfigChange `protobuf:"bytes,1,rep,name=config_changes,json=configChanges" json:"config_changes,omitempty"`
}

func (m *ChangeReport) Reset()                    { *m = ChangeReport{} }
func (m *ChangeReport) String() string            { return proto.CompactTextString(m) }
func (*ChangeReport) ProtoMessage()               {}
func (*ChangeReport) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ChangeReport) GetConfigChanges() []*google_api2.ConfigChange {
	if m != nil {
		return m.ConfigChanges
	}
	return nil
}

// A rollout resource that defines how service configuration versions are pushed
// to control plane systems. Typically, you create a new version of the
// service config, and then create a Rollout to push the service config.
type Rollout struct {
	// Optional unique identifier of this Rollout. Only lower case letters, digits
	//  and '-' are allowed.
	//
	// If not specified by client, the server will generate one. The generated id
	// will have the form of <date><revision number>, where "date" is the create
	// date in ISO 8601 format.  "revision number" is a monotonically increasing
	// positive number that is reset every day for each service.
	// An example of the generated rollout_id is '2016-02-16r1'
	RolloutId string `protobuf:"bytes,1,opt,name=rollout_id,json=rolloutId" json:"rollout_id,omitempty"`
	// Creation time of the rollout. Readonly.
	CreateTime *google_protobuf9.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	// The user who created the Rollout. Readonly.
	CreatedBy string `protobuf:"bytes,3,opt,name=created_by,json=createdBy" json:"created_by,omitempty"`
	// The status of this rollout. Readonly. In case of a failed rollout,
	// the system will automatically rollback to the current Rollout
	// version. Readonly.
	Status Rollout_RolloutStatus `protobuf:"varint,4,opt,name=status,enum=google.api.servicemanagement.v1.Rollout_RolloutStatus" json:"status,omitempty"`
	// Strategy that defines which versions of service configurations should be
	// pushed
	// and how they should be used at runtime.
	//
	// Types that are valid to be assigned to Strategy:
	//	*Rollout_TrafficPercentStrategy_
	//	*Rollout_DeleteServiceStrategy_
	Strategy isRollout_Strategy `protobuf_oneof:"strategy"`
	// The name of the service associated with this Rollout.
	ServiceName string `protobuf:"bytes,8,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
}

func (m *Rollout) Reset()                    { *m = Rollout{} }
func (m *Rollout) String() string            { return proto.CompactTextString(m) }
func (*Rollout) ProtoMessage()               {}
func (*Rollout) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type isRollout_Strategy interface {
	isRollout_Strategy()
}

type Rollout_TrafficPercentStrategy_ struct {
	TrafficPercentStrategy *Rollout_TrafficPercentStrategy `protobuf:"bytes,5,opt,name=traffic_percent_strategy,json=trafficPercentStrategy,oneof"`
}
type Rollout_DeleteServiceStrategy_ struct {
	DeleteServiceStrategy *Rollout_DeleteServiceStrategy `protobuf:"bytes,200,opt,name=delete_service_strategy,json=deleteServiceStrategy,oneof"`
}

func (*Rollout_TrafficPercentStrategy_) isRollout_Strategy() {}
func (*Rollout_DeleteServiceStrategy_) isRollout_Strategy()  {}

func (m *Rollout) GetStrategy() isRollout_Strategy {
	if m != nil {
		return m.Strategy
	}
	return nil
}

func (m *Rollout) GetRolloutId() string {
	if m != nil {
		return m.RolloutId
	}
	return ""
}

func (m *Rollout) GetCreateTime() *google_protobuf9.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Rollout) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *Rollout) GetStatus() Rollout_RolloutStatus {
	if m != nil {
		return m.Status
	}
	return Rollout_ROLLOUT_STATUS_UNSPECIFIED
}

func (m *Rollout) GetTrafficPercentStrategy() *Rollout_TrafficPercentStrategy {
	if x, ok := m.GetStrategy().(*Rollout_TrafficPercentStrategy_); ok {
		return x.TrafficPercentStrategy
	}
	return nil
}

func (m *Rollout) GetDeleteServiceStrategy() *Rollout_DeleteServiceStrategy {
	if x, ok := m.GetStrategy().(*Rollout_DeleteServiceStrategy_); ok {
		return x.DeleteServiceStrategy
	}
	return nil
}

func (m *Rollout) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Rollout) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Rollout_OneofMarshaler, _Rollout_OneofUnmarshaler, _Rollout_OneofSizer, []interface{}{
		(*Rollout_TrafficPercentStrategy_)(nil),
		(*Rollout_DeleteServiceStrategy_)(nil),
	}
}

func _Rollout_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Rollout)
	// strategy
	switch x := m.Strategy.(type) {
	case *Rollout_TrafficPercentStrategy_:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TrafficPercentStrategy); err != nil {
			return err
		}
	case *Rollout_DeleteServiceStrategy_:
		b.EncodeVarint(200<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DeleteServiceStrategy); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Rollout.Strategy has unexpected type %T", x)
	}
	return nil
}

func _Rollout_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Rollout)
	switch tag {
	case 5: // strategy.traffic_percent_strategy
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Rollout_TrafficPercentStrategy)
		err := b.DecodeMessage(msg)
		m.Strategy = &Rollout_TrafficPercentStrategy_{msg}
		return true, err
	case 200: // strategy.delete_service_strategy
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Rollout_DeleteServiceStrategy)
		err := b.DecodeMessage(msg)
		m.Strategy = &Rollout_DeleteServiceStrategy_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Rollout_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Rollout)
	// strategy
	switch x := m.Strategy.(type) {
	case *Rollout_TrafficPercentStrategy_:
		s := proto.Size(x.TrafficPercentStrategy)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Rollout_DeleteServiceStrategy_:
		s := proto.Size(x.DeleteServiceStrategy)
		n += proto.SizeVarint(200<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Strategy that specifies how Google Service Control should select
// different
// versions of service configurations based on traffic percentage.
//
// One example of how to gradually rollout a new service configuration using
// this
// strategy:
// Day 1
//
//     Rollout {
//       id: "example.googleapis.com/rollout_20160206"
//       traffic_percent_strategy {
//         percentages: {
//           "example.googleapis.com/20160201": 70.00
//           "example.googleapis.com/20160206": 30.00
//         }
//       }
//     }
//
// Day 2
//
//     Rollout {
//       id: "example.googleapis.com/rollout_20160207"
//       traffic_percent_strategy: {
//         percentages: {
//           "example.googleapis.com/20160206": 100.00
//         }
//       }
//     }
type Rollout_TrafficPercentStrategy struct {
	// Maps service configuration IDs to their corresponding traffic percentage.
	// Key is the service configuration ID, Value is the traffic percentage
	// which must be greater than 0.0 and the sum must equal to 100.0.
	Percentages map[string]float64 `protobuf:"bytes,1,rep,name=percentages" json:"percentages,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"fixed64,2,opt,name=value"`
}

func (m *Rollout_TrafficPercentStrategy) Reset()         { *m = Rollout_TrafficPercentStrategy{} }
func (m *Rollout_TrafficPercentStrategy) String() string { return proto.CompactTextString(m) }
func (*Rollout_TrafficPercentStrategy) ProtoMessage()    {}
func (*Rollout_TrafficPercentStrategy) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{7, 0}
}

func (m *Rollout_TrafficPercentStrategy) GetPercentages() map[string]float64 {
	if m != nil {
		return m.Percentages
	}
	return nil
}

// Strategy used to delete a service. This strategy is a placeholder only
// used by the system generated rollout to delete a service.
type Rollout_DeleteServiceStrategy struct {
}

func (m *Rollout_DeleteServiceStrategy) Reset()         { *m = Rollout_DeleteServiceStrategy{} }
func (m *Rollout_DeleteServiceStrategy) String() string { return proto.CompactTextString(m) }
func (*Rollout_DeleteServiceStrategy) ProtoMessage()    {}
func (*Rollout_DeleteServiceStrategy) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{7, 1}
}

func init() {
	proto.RegisterType((*ManagedService)(nil), "google.api.servicemanagement.v1.ManagedService")
	proto.RegisterType((*OperationMetadata)(nil), "google.api.servicemanagement.v1.OperationMetadata")
	proto.RegisterType((*OperationMetadata_Step)(nil), "google.api.servicemanagement.v1.OperationMetadata.Step")
	proto.RegisterType((*Diagnostic)(nil), "google.api.servicemanagement.v1.Diagnostic")
	proto.RegisterType((*ConfigSource)(nil), "google.api.servicemanagement.v1.ConfigSource")
	proto.RegisterType((*ConfigFile)(nil), "google.api.servicemanagement.v1.ConfigFile")
	proto.RegisterType((*ConfigRef)(nil), "google.api.servicemanagement.v1.ConfigRef")
	proto.RegisterType((*ChangeReport)(nil), "google.api.servicemanagement.v1.ChangeReport")
	proto.RegisterType((*Rollout)(nil), "google.api.servicemanagement.v1.Rollout")
	proto.RegisterType((*Rollout_TrafficPercentStrategy)(nil), "google.api.servicemanagement.v1.Rollout.TrafficPercentStrategy")
	proto.RegisterType((*Rollout_DeleteServiceStrategy)(nil), "google.api.servicemanagement.v1.Rollout.DeleteServiceStrategy")
	proto.RegisterEnum("google.api.servicemanagement.v1.OperationMetadata_Status", OperationMetadata_Status_name, OperationMetadata_Status_value)
	proto.RegisterEnum("google.api.servicemanagement.v1.Diagnostic_Kind", Diagnostic_Kind_name, Diagnostic_Kind_value)
	proto.RegisterEnum("google.api.servicemanagement.v1.ConfigFile_FileType", ConfigFile_FileType_name, ConfigFile_FileType_value)
	proto.RegisterEnum("google.api.servicemanagement.v1.Rollout_RolloutStatus", Rollout_RolloutStatus_name, Rollout_RolloutStatus_value)
}

func init() { proto.RegisterFile("google/api/servicemanagement/v1/resources.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xef, 0x8e, 0xdb, 0x44,
	0x10, 0xaf, 0xf3, 0xe7, 0xee, 0x32, 0xb9, 0x0b, 0xee, 0x96, 0xf6, 0xd2, 0xd0, 0x3f, 0xc1, 0x15,
	0xd2, 0x49, 0x48, 0x0e, 0x0d, 0x08, 0x28, 0x95, 0x5a, 0xe5, 0x12, 0xdf, 0x61, 0xc8, 0xd9, 0xee,
	0xda, 0x07, 0x2a, 0x5f, 0xac, 0xad, 0xbd, 0x71, 0x4d, 0x13, 0xdb, 0xb2, 0x37, 0x27, 0x45, 0xfd,
	0xc8, 0x0b, 0xf0, 0x0c, 0x7c, 0x81, 0x47, 0xe1, 0x03, 0x4f, 0x00, 0x2f, 0x83, 0xbc, 0x5e, 0xdf,
	0xe5, 0xcf, 0xa1, 0x14, 0xf8, 0x92, 0xec, 0xfc, 0x66, 0xf6, 0x37, 0xb3, 0xb3, 0x33, 0xb3, 0x86,
	0x5e, 0x10, 0xc7, 0xc1, 0x94, 0xf6, 0x48, 0x12, 0xf6, 0x32, 0x9a, 0x5e, 0x84, 0x1e, 0x9d, 0x91,
	0x88, 0x04, 0x74, 0x46, 0x23, 0xd6, 0xbb, 0x78, 0xdc, 0x4b, 0x69, 0x16, 0xcf, 0x53, 0x8f, 0x66,
	0x6a, 0x92, 0xc6, 0x2c, 0x46, 0x0f, 0x8b, 0x0d, 0x2a, 0x49, 0x42, 0x75, 0x63, 0x83, 0x7a, 0xf1,
	0xb8, 0x73, 0x6f, 0x89, 0x91, 0x44, 0x51, 0xcc, 0x08, 0x0b, 0xe3, 0x48, 0x6c, 0xef, 0x3c, 0x58,
	0xd2, 0x7a, 0x71, 0x34, 0x09, 0x03, 0xd7, 0x7b, 0x4d, 0xa2, 0x80, 0x0a, 0x7d, 0x7b, 0x33, 0x1e,
	0xa1, 0x79, 0x24, 0x34, 0xd3, 0x38, 0x0a, 0xd2, 0x79, 0x14, 0x85, 0x51, 0xd0, 0x8b, 0x13, 0x9a,
	0xae, 0xd0, 0xdf, 0x15, 0x46, 0x5c, 0x7a, 0x35, 0x9f, 0xf4, 0x48, 0xb4, 0x10, 0xaa, 0xee, 0xba,
	0x6a, 0x12, 0xd2, 0xa9, 0xef, 0xce, 0x48, 0xf6, 0x46, 0x58, 0xdc, 0x5b, 0xb7, 0xc8, 0x58, 0x3a,
	0xf7, 0x98, 0xd0, 0x3e, 0x5c, 0xd7, 0xb2, 0x70, 0x46, 0x33, 0x46, 0x66, 0x89, 0x30, 0x38, 0x14,
	0x06, 0x69, 0xe2, 0xf5, 0x32, 0x46, 0xd8, 0x5c, 0x04, 0xa5, 0x78, 0xd0, 0x3a, 0xe3, 0x29, 0xf2,
	0xed, 0xe2, 0x44, 0xe8, 0x43, 0xd8, 0x17, 0x87, 0x73, 0x23, 0x32, 0xa3, 0xed, 0x4a, 0x57, 0x3a,
	0x6a, 0xe0, 0xa6, 0xc0, 0x0c, 0x32, 0xa3, 0x48, 0x85, 0x5b, 0x49, 0x1a, 0xfb, 0x73, 0x8f, 0xa6,
	0x6e, 0x92, 0xc6, 0x3f, 0x52, 0x8f, 0xb9, 0xa1, 0xdf, 0xae, 0x72, 0xcb, 0x9b, 0xa5, 0xca, 0x2a,
	0x34, 0xba, 0xaf, 0xfc, 0x55, 0x85, 0x9b, 0x66, 0x99, 0x8e, 0x33, 0xca, 0x88, 0x4f, 0x18, 0x41,
	0x1f, 0x41, 0xab, 0xbc, 0x40, 0xee, 0x29, 0x6b, 0x4b, 0xdd, 0xea, 0x51, 0x03, 0x1f, 0x94, 0x68,
	0xee, 0x2b, 0x43, 0x67, 0x50, 0xcf, 0x18, 0x4d, 0xb2, 0x76, 0xa5, 0x5b, 0x3d, 0x6a, 0xf6, 0xbf,
	0x50, 0xb7, 0x5c, 0xb2, 0xba, 0xe1, 0x49, 0xb5, 0x19, 0x4d, 0x70, 0xc1, 0x82, 0x7a, 0x3c, 0xf6,
	0x20, 0xa5, 0x59, 0xe6, 0x26, 0x34, 0xf5, 0x68, 0xc4, 0x48, 0x40, 0x79, 0xec, 0x75, 0x8c, 0x4a,
	0x95, 0x75, 0xa9, 0x41, 0x4f, 0x00, 0x32, 0x46, 0x52, 0xe6, 0xe6, 0x39, 0x6d, 0xd7, 0xba, 0xd2,
	0x51, 0xb3, 0xdf, 0x29, 0x83, 0x28, 0x13, 0xae, 0x3a, 0x65, 0xc2, 0x71, 0x83, 0x5b, 0xe7, 0x72,
	0xe7, 0x2d, 0xd4, 0x72, 0xd7, 0xa8, 0x0b, 0x4d, 0x9f, 0x66, 0x5e, 0x1a, 0x26, 0x79, 0x58, 0x65,
	0x46, 0x97, 0x20, 0xf4, 0x02, 0x76, 0x8a, 0x6b, 0xe1, 0x0e, 0x5a, 0xfd, 0x27, 0xff, 0xe9, 0x94,
	0x39, 0x01, 0x16, 0x44, 0x4a, 0x00, 0x3b, 0x05, 0x82, 0xee, 0x00, 0xb2, 0x9d, 0x81, 0x73, 0x6e,
	0xbb, 0xe7, 0x86, 0x6d, 0x69, 0x43, 0xfd, 0x44, 0xd7, 0x46, 0xf2, 0x0d, 0xb4, 0x07, 0xb5, 0x91,
	0x69, 0x68, 0xb2, 0x84, 0xde, 0x83, 0xa6, 0x61, 0x3a, 0xae, 0xed, 0x0c, 0xb0, 0xa3, 0x8d, 0xe4,
	0x4a, 0x0e, 0xe8, 0x86, 0x6b, 0x61, 0xf3, 0x14, 0x6b, 0xb6, 0x2d, 0x57, 0x11, 0xc0, 0xce, 0xc9,
	0x40, 0x1f, 0x6b, 0x23, 0xb9, 0x86, 0x0e, 0xa0, 0x31, 0x1c, 0x18, 0x43, 0x6d, 0x9c, 0x8b, 0x75,
	0xe5, 0x37, 0x09, 0x60, 0x14, 0x92, 0x20, 0x8a, 0x33, 0x16, 0x7a, 0xa8, 0x03, 0x7b, 0xd3, 0xd8,
	0xe3, 0xa1, 0xb5, 0x25, 0x7e, 0xd2, 0x4b, 0x19, 0x8d, 0xa0, 0xf6, 0x26, 0x8c, 0x7c, 0x9e, 0x81,
	0x56, 0xff, 0x93, 0xad, 0x87, 0xbc, 0xa2, 0x55, 0xbf, 0x0d, 0x23, 0x1f, 0xf3, 0xdd, 0xa8, 0x0d,
	0xbb, 0x33, 0x9a, 0x65, 0xe5, 0xb5, 0x35, 0x70, 0x29, 0x2a, 0x0f, 0xa0, 0x96, 0xdb, 0xa1, 0x26,
	0xec, 0x7e, 0x3f, 0xc0, 0x86, 0x6e, 0x9c, 0xca, 0x37, 0x50, 0x03, 0xea, 0x1a, 0xc6, 0x26, 0x96,
	0x25, 0x85, 0xc0, 0xfe, 0x90, 0x37, 0xb6, 0xcd, 0x0b, 0x0c, 0xb5, 0xa0, 0x12, 0xfa, 0xed, 0x3a,
	0x27, 0xa9, 0x84, 0x3e, 0x1a, 0x40, 0x7d, 0x12, 0x4e, 0x69, 0x59, 0x6b, 0x1f, 0x6f, 0x0d, 0xb0,
	0x60, 0x3b, 0x09, 0xa7, 0x14, 0x17, 0x3b, 0x95, 0x5f, 0x2b, 0x00, 0x57, 0x28, 0xfa, 0x00, 0x1a,
	0x39, 0xee, 0x26, 0x84, 0xbd, 0x2e, 0xd3, 0x91, 0x03, 0x16, 0x61, 0xaf, 0xd1, 0x23, 0x38, 0xe0,
	0x4a, 0x2f, 0x8e, 0x18, 0x8d, 0x58, 0xc6, 0x8f, 0xb3, 0x8f, 0xf7, 0x73, 0x70, 0x28, 0x30, 0xf4,
	0x42, 0x30, 0xb0, 0x45, 0x42, 0x45, 0x75, 0x7c, 0xf6, 0x2f, 0xe2, 0x52, 0xf3, 0x1f, 0x67, 0x91,
	0xd0, 0xc2, 0x6f, 0xbe, 0x52, 0x7e, 0x92, 0x60, 0xaf, 0x84, 0xd1, 0x5d, 0xb8, 0x7d, 0xa2, 0x8f,
	0x35, 0xd7, 0x79, 0x69, 0x69, 0x6b, 0x05, 0x72, 0x08, 0xb7, 0x6c, 0x0d, 0x7f, 0xa7, 0x0f, 0x35,
	0x77, 0x68, 0x1a, 0x27, 0xfa, 0xa9, 0xfb, 0x72, 0x70, 0x36, 0x96, 0x25, 0x74, 0x13, 0x0e, 0x4c,
	0x4b, 0x33, 0xdc, 0x81, 0xa5, 0xbb, 0xdf, 0xd8, 0xa6, 0x21, 0x57, 0x56, 0x20, 0x6e, 0x55, 0x45,
	0xf7, 0xe1, 0x2e, 0x67, 0x1e, 0x69, 0xf6, 0x10, 0xeb, 0x96, 0x63, 0x62, 0xd7, 0xd6, 0x9c, 0xbc,
	0xaa, 0x1c, 0x53, 0xae, 0x29, 0x0f, 0xa1, 0x51, 0x84, 0x89, 0xe9, 0x04, 0x21, 0xa8, 0xf1, 0x69,
	0x53, 0xa4, 0x88, 0xaf, 0x15, 0x13, 0xf6, 0x87, 0x7c, 0xfe, 0x62, 0x9a, 0xc4, 0x29, 0x43, 0xcf,
	0xa1, 0xb5, 0x32, 0x96, 0x8b, 0x81, 0xd1, 0xec, 0xb7, 0x97, 0xd3, 0x51, 0x50, 0x8a, 0x7d, 0x07,
	0xde, 0x92, 0x94, 0x29, 0x7f, 0xee, 0xc0, 0x2e, 0x8e, 0xa7, 0xd3, 0x78, 0xce, 0xd0, 0x7d, 0x80,
	0xb4, 0x58, 0xe6, 0xa3, 0xab, 0x70, 0xdb, 0x10, 0x88, 0xee, 0xa3, 0xa7, 0xd0, 0xf4, 0x52, 0x4a,
	0x18, 0x2d, 0xda, 0xbe, 0xb2, 0xb5, 0xed, 0xa1, 0x30, 0xcf, 0x81, 0x9c, 0xbb, 0x90, 0x7c, 0xf7,
	0xd5, 0x42, 0xd4, 0x68, 0x43, 0x20, 0xc7, 0x0b, 0x64, 0xac, 0x35, 0xfb, 0xe7, 0x5b, 0xaf, 0x53,
	0x04, 0x5d, 0xfe, 0xaf, 0x76, 0x3a, 0x7a, 0x0b, 0x6d, 0x96, 0x92, 0xc9, 0x24, 0xf4, 0xca, 0x89,
	0xe6, 0x66, 0x2c, 0x25, 0x8c, 0x06, 0x0b, 0x5e, 0xdb, 0xcd, 0xfe, 0xf3, 0x77, 0xf6, 0xe0, 0x14,
	0x44, 0x62, 0xfe, 0xd9, 0x82, 0xe6, 0xeb, 0x1b, 0xf8, 0x0e, 0xbb, 0x56, 0x83, 0x16, 0x70, 0xe8,
	0xd3, 0x29, 0x65, 0xd4, 0x2d, 0x5f, 0x8d, 0x4b, 0xdf, 0xbf, 0x4b, 0xdc, 0xf9, 0xb3, 0x77, 0x76,
	0x3e, 0xe2, 0x44, 0xe2, 0x21, 0x5a, 0xf2, 0x7d, 0xdb, 0xbf, 0x4e, 0xb1, 0xf1, 0x52, 0xed, 0x6d,
	0xbc, 0x54, 0x9d, 0x3f, 0x24, 0xb8, 0x73, 0xfd, 0x91, 0x50, 0x0a, 0xcd, 0xab, 0xf9, 0x5f, 0x96,
	0x92, 0xf5, 0x3f, 0x13, 0xa5, 0x5e, 0x3d, 0x1c, 0x99, 0x16, 0xb1, 0x74, 0x81, 0x97, 0x9d, 0x74,
	0x9e, 0x81, 0xbc, 0x6e, 0x80, 0x64, 0xa8, 0xbe, 0xa1, 0x0b, 0x51, 0x81, 0xf9, 0x12, 0xbd, 0x0f,
	0xf5, 0x0b, 0x32, 0x9d, 0x17, 0x55, 0x27, 0xe1, 0x42, 0xf8, 0xaa, 0xf2, 0xa5, 0xd4, 0x39, 0x84,
	0xdb, 0xd7, 0xe6, 0x48, 0x99, 0xc3, 0xc1, 0x4a, 0x6d, 0xa0, 0x07, 0xd0, 0xc1, 0xe6, 0x78, 0x6c,
	0x9e, 0xf3, 0xa9, 0xbe, 0x39, 0xfb, 0xd7, 0x06, 0xbc, 0x94, 0x8f, 0x4c, 0xfb, 0x7c, 0x38, 0xcc,
	0x85, 0xca, 0xea, 0x84, 0x5f, 0x1d, 0xfe, 0x4d, 0xd8, 0xb5, 0x34, 0x63, 0x94, 0x8f, 0xd6, 0xfa,
	0x31, 0xc0, 0x5e, 0x79, 0xdb, 0xc7, 0x3f, 0x4b, 0xf0, 0xc8, 0x8b, 0x67, 0xdb, 0x12, 0x78, 0xdc,
	0xc2, 0xe5, 0x57, 0x9b, 0x95, 0x77, 0x91, 0x25, 0xfd, 0x60, 0x89, 0x2d, 0x41, 0x3c, 0x25, 0x51,
	0xa0, 0xc6, 0x69, 0xd0, 0x0b, 0x68, 0xc4, 0x7b, 0x4c, 0x7c, 0x02, 0x92, 0x24, 0xcc, 0xfe, 0xf1,
	0x33, 0xf0, 0xe9, 0x06, 0xf8, 0x4b, 0xa5, 0x76, 0x3a, 0xb0, 0xcf, 0x5e, 0xed, 0x70, 0x8e, 0x4f,
	0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x2d, 0x64, 0x4d, 0x1e, 0x49, 0x0a, 0x00, 0x00,
}