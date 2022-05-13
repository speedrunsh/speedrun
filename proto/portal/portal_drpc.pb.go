// Code generated by protoc-gen-go-drpc. DO NOT EDIT.
// protoc-gen-go-drpc version: v0.0.30
// source: portal/portal.proto

package portal

import (
	context "context"
	errors "errors"
	protojson "google.golang.org/protobuf/encoding/protojson"
	proto "google.golang.org/protobuf/proto"
	drpc "storj.io/drpc"
	drpcerr "storj.io/drpc/drpcerr"
)

type drpcEncoding_File_portal_portal_proto struct{}

func (drpcEncoding_File_portal_portal_proto) Marshal(msg drpc.Message) ([]byte, error) {
	return proto.Marshal(msg.(proto.Message))
}

func (drpcEncoding_File_portal_portal_proto) MarshalAppend(buf []byte, msg drpc.Message) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(buf, msg.(proto.Message))
}

func (drpcEncoding_File_portal_portal_proto) Unmarshal(buf []byte, msg drpc.Message) error {
	return proto.Unmarshal(buf, msg.(proto.Message))
}

func (drpcEncoding_File_portal_portal_proto) JSONMarshal(msg drpc.Message) ([]byte, error) {
	return protojson.Marshal(msg.(proto.Message))
}

func (drpcEncoding_File_portal_portal_proto) JSONUnmarshal(buf []byte, msg drpc.Message) error {
	return protojson.Unmarshal(buf, msg.(proto.Message))
}

type DRPCPortalClient interface {
	DRPCConn() drpc.Conn

	ServiceRestart(ctx context.Context, in *ServiceRequest) (*ServiceResponse, error)
	ServiceStart(ctx context.Context, in *ServiceRequest) (*ServiceResponse, error)
	ServiceStop(ctx context.Context, in *ServiceRequest) (*ServiceResponse, error)
	ServiceStatus(ctx context.Context, in *ServiceRequest) (*ServiceStatusResponse, error)
	RunCommand(ctx context.Context, in *CommandRequest) (*CommandResponse, error)
	CPUusage(ctx context.Context, in *CPUusageRequest) (*CPUusageResponse, error)
	FileRead(ctx context.Context, in *FileReadRequest) (*FileReadResponse, error)
	SystemReboot(ctx context.Context, in *SystemRebootRequest) (*SystemRebootResponse, error)
}

type drpcPortalClient struct {
	cc drpc.Conn
}

func NewDRPCPortalClient(cc drpc.Conn) DRPCPortalClient {
	return &drpcPortalClient{cc}
}

func (c *drpcPortalClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcPortalClient) ServiceRestart(ctx context.Context, in *ServiceRequest) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/portal.Portal/ServiceRestart", drpcEncoding_File_portal_portal_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcPortalClient) ServiceStart(ctx context.Context, in *ServiceRequest) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/portal.Portal/ServiceStart", drpcEncoding_File_portal_portal_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcPortalClient) ServiceStop(ctx context.Context, in *ServiceRequest) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/portal.Portal/ServiceStop", drpcEncoding_File_portal_portal_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcPortalClient) ServiceStatus(ctx context.Context, in *ServiceRequest) (*ServiceStatusResponse, error) {
	out := new(ServiceStatusResponse)
	err := c.cc.Invoke(ctx, "/portal.Portal/ServiceStatus", drpcEncoding_File_portal_portal_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcPortalClient) RunCommand(ctx context.Context, in *CommandRequest) (*CommandResponse, error) {
	out := new(CommandResponse)
	err := c.cc.Invoke(ctx, "/portal.Portal/RunCommand", drpcEncoding_File_portal_portal_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcPortalClient) CPUusage(ctx context.Context, in *CPUusageRequest) (*CPUusageResponse, error) {
	out := new(CPUusageResponse)
	err := c.cc.Invoke(ctx, "/portal.Portal/CPUusage", drpcEncoding_File_portal_portal_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcPortalClient) FileRead(ctx context.Context, in *FileReadRequest) (*FileReadResponse, error) {
	out := new(FileReadResponse)
	err := c.cc.Invoke(ctx, "/portal.Portal/FileRead", drpcEncoding_File_portal_portal_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcPortalClient) SystemReboot(ctx context.Context, in *SystemRebootRequest) (*SystemRebootResponse, error) {
	out := new(SystemRebootResponse)
	err := c.cc.Invoke(ctx, "/portal.Portal/SystemReboot", drpcEncoding_File_portal_portal_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type DRPCPortalServer interface {
	ServiceRestart(context.Context, *ServiceRequest) (*ServiceResponse, error)
	ServiceStart(context.Context, *ServiceRequest) (*ServiceResponse, error)
	ServiceStop(context.Context, *ServiceRequest) (*ServiceResponse, error)
	ServiceStatus(context.Context, *ServiceRequest) (*ServiceStatusResponse, error)
	RunCommand(context.Context, *CommandRequest) (*CommandResponse, error)
	CPUusage(context.Context, *CPUusageRequest) (*CPUusageResponse, error)
	FileRead(context.Context, *FileReadRequest) (*FileReadResponse, error)
	SystemReboot(context.Context, *SystemRebootRequest) (*SystemRebootResponse, error)
}

type DRPCPortalUnimplementedServer struct{}

func (s *DRPCPortalUnimplementedServer) ServiceRestart(context.Context, *ServiceRequest) (*ServiceResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCPortalUnimplementedServer) ServiceStart(context.Context, *ServiceRequest) (*ServiceResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCPortalUnimplementedServer) ServiceStop(context.Context, *ServiceRequest) (*ServiceResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCPortalUnimplementedServer) ServiceStatus(context.Context, *ServiceRequest) (*ServiceStatusResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCPortalUnimplementedServer) RunCommand(context.Context, *CommandRequest) (*CommandResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCPortalUnimplementedServer) CPUusage(context.Context, *CPUusageRequest) (*CPUusageResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCPortalUnimplementedServer) FileRead(context.Context, *FileReadRequest) (*FileReadResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCPortalUnimplementedServer) SystemReboot(context.Context, *SystemRebootRequest) (*SystemRebootResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

type DRPCPortalDescription struct{}

func (DRPCPortalDescription) NumMethods() int { return 8 }

func (DRPCPortalDescription) Method(n int) (string, drpc.Encoding, drpc.Receiver, interface{}, bool) {
	switch n {
	case 0:
		return "/portal.Portal/ServiceRestart", drpcEncoding_File_portal_portal_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCPortalServer).
					ServiceRestart(
						ctx,
						in1.(*ServiceRequest),
					)
			}, DRPCPortalServer.ServiceRestart, true
	case 1:
		return "/portal.Portal/ServiceStart", drpcEncoding_File_portal_portal_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCPortalServer).
					ServiceStart(
						ctx,
						in1.(*ServiceRequest),
					)
			}, DRPCPortalServer.ServiceStart, true
	case 2:
		return "/portal.Portal/ServiceStop", drpcEncoding_File_portal_portal_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCPortalServer).
					ServiceStop(
						ctx,
						in1.(*ServiceRequest),
					)
			}, DRPCPortalServer.ServiceStop, true
	case 3:
		return "/portal.Portal/ServiceStatus", drpcEncoding_File_portal_portal_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCPortalServer).
					ServiceStatus(
						ctx,
						in1.(*ServiceRequest),
					)
			}, DRPCPortalServer.ServiceStatus, true
	case 4:
		return "/portal.Portal/RunCommand", drpcEncoding_File_portal_portal_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCPortalServer).
					RunCommand(
						ctx,
						in1.(*CommandRequest),
					)
			}, DRPCPortalServer.RunCommand, true
	case 5:
		return "/portal.Portal/CPUusage", drpcEncoding_File_portal_portal_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCPortalServer).
					CPUusage(
						ctx,
						in1.(*CPUusageRequest),
					)
			}, DRPCPortalServer.CPUusage, true
	case 6:
		return "/portal.Portal/FileRead", drpcEncoding_File_portal_portal_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCPortalServer).
					FileRead(
						ctx,
						in1.(*FileReadRequest),
					)
			}, DRPCPortalServer.FileRead, true
	case 7:
		return "/portal.Portal/SystemReboot", drpcEncoding_File_portal_portal_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCPortalServer).
					SystemReboot(
						ctx,
						in1.(*SystemRebootRequest),
					)
			}, DRPCPortalServer.SystemReboot, true
	default:
		return "", nil, nil, nil, false
	}
}

func DRPCRegisterPortal(mux drpc.Mux, impl DRPCPortalServer) error {
	return mux.Register(impl, DRPCPortalDescription{})
}

type DRPCPortal_ServiceRestartStream interface {
	drpc.Stream
	SendAndClose(*ServiceResponse) error
}

type drpcPortal_ServiceRestartStream struct {
	drpc.Stream
}

func (x *drpcPortal_ServiceRestartStream) SendAndClose(m *ServiceResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_portal_portal_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCPortal_ServiceStartStream interface {
	drpc.Stream
	SendAndClose(*ServiceResponse) error
}

type drpcPortal_ServiceStartStream struct {
	drpc.Stream
}

func (x *drpcPortal_ServiceStartStream) SendAndClose(m *ServiceResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_portal_portal_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCPortal_ServiceStopStream interface {
	drpc.Stream
	SendAndClose(*ServiceResponse) error
}

type drpcPortal_ServiceStopStream struct {
	drpc.Stream
}

func (x *drpcPortal_ServiceStopStream) SendAndClose(m *ServiceResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_portal_portal_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCPortal_ServiceStatusStream interface {
	drpc.Stream
	SendAndClose(*ServiceStatusResponse) error
}

type drpcPortal_ServiceStatusStream struct {
	drpc.Stream
}

func (x *drpcPortal_ServiceStatusStream) SendAndClose(m *ServiceStatusResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_portal_portal_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCPortal_RunCommandStream interface {
	drpc.Stream
	SendAndClose(*CommandResponse) error
}

type drpcPortal_RunCommandStream struct {
	drpc.Stream
}

func (x *drpcPortal_RunCommandStream) SendAndClose(m *CommandResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_portal_portal_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCPortal_CPUusageStream interface {
	drpc.Stream
	SendAndClose(*CPUusageResponse) error
}

type drpcPortal_CPUusageStream struct {
	drpc.Stream
}

func (x *drpcPortal_CPUusageStream) SendAndClose(m *CPUusageResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_portal_portal_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCPortal_FileReadStream interface {
	drpc.Stream
	SendAndClose(*FileReadResponse) error
}

type drpcPortal_FileReadStream struct {
	drpc.Stream
}

func (x *drpcPortal_FileReadStream) SendAndClose(m *FileReadResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_portal_portal_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCPortal_SystemRebootStream interface {
	drpc.Stream
	SendAndClose(*SystemRebootResponse) error
}

type drpcPortal_SystemRebootStream struct {
	drpc.Stream
}

func (x *drpcPortal_SystemRebootStream) SendAndClose(m *SystemRebootResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_portal_portal_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}
