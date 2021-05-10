// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AsistenciaServClient is the client API for AsistenciaServ service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AsistenciaServClient interface {
	Crear(ctx context.Context, in *Asistencia, opts ...grpc.CallOption) (*IndividualAsistencia, error)
	ObtenerPor_ID(ctx context.Context, in *Filtro, opts ...grpc.CallOption) (*AsistenciasCollection, error)
	ObtenerPor_CARNET(ctx context.Context, in *Filtro, opts ...grpc.CallOption) (*AsistenciasCollection, error)
}

type asistenciaServClient struct {
	cc grpc.ClientConnInterface
}

func NewAsistenciaServClient(cc grpc.ClientConnInterface) AsistenciaServClient {
	return &asistenciaServClient{cc}
}

func (c *asistenciaServClient) Crear(ctx context.Context, in *Asistencia, opts ...grpc.CallOption) (*IndividualAsistencia, error) {
	out := new(IndividualAsistencia)
	err := c.cc.Invoke(ctx, "/pb.AsistenciaServ/Crear", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asistenciaServClient) ObtenerPor_ID(ctx context.Context, in *Filtro, opts ...grpc.CallOption) (*AsistenciasCollection, error) {
	out := new(AsistenciasCollection)
	err := c.cc.Invoke(ctx, "/pb.AsistenciaServ/Obtener_por_ID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asistenciaServClient) ObtenerPor_CARNET(ctx context.Context, in *Filtro, opts ...grpc.CallOption) (*AsistenciasCollection, error) {
	out := new(AsistenciasCollection)
	err := c.cc.Invoke(ctx, "/pb.AsistenciaServ/Obtener_por_CARNET", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AsistenciaServServer is the server API for AsistenciaServ service.
// All implementations must embed UnimplementedAsistenciaServServer
// for forward compatibility
type AsistenciaServServer interface {
	Crear(context.Context, *Asistencia) (*IndividualAsistencia, error)
	ObtenerPor_ID(context.Context, *Filtro) (*AsistenciasCollection, error)
	ObtenerPor_CARNET(context.Context, *Filtro) (*AsistenciasCollection, error)
	mustEmbedUnimplementedAsistenciaServServer()
}

// UnimplementedAsistenciaServServer must be embedded to have forward compatible implementations.
type UnimplementedAsistenciaServServer struct {
}

func (UnimplementedAsistenciaServServer) Crear(context.Context, *Asistencia) (*IndividualAsistencia, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Crear not implemented")
}
func (UnimplementedAsistenciaServServer) ObtenerPor_ID(context.Context, *Filtro) (*AsistenciasCollection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ObtenerPor_ID not implemented")
}
func (UnimplementedAsistenciaServServer) ObtenerPor_CARNET(context.Context, *Filtro) (*AsistenciasCollection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ObtenerPor_CARNET not implemented")
}
func (UnimplementedAsistenciaServServer) mustEmbedUnimplementedAsistenciaServServer() {}

// UnsafeAsistenciaServServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AsistenciaServServer will
// result in compilation errors.
type UnsafeAsistenciaServServer interface {
	mustEmbedUnimplementedAsistenciaServServer()
}

func RegisterAsistenciaServServer(s grpc.ServiceRegistrar, srv AsistenciaServServer) {
	s.RegisterService(&AsistenciaServ_ServiceDesc, srv)
}

func _AsistenciaServ_Crear_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Asistencia)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsistenciaServServer).Crear(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AsistenciaServ/Crear",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsistenciaServServer).Crear(ctx, req.(*Asistencia))
	}
	return interceptor(ctx, in, info, handler)
}

func _AsistenciaServ_ObtenerPor_ID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filtro)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsistenciaServServer).ObtenerPor_ID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AsistenciaServ/Obtener_por_ID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsistenciaServServer).ObtenerPor_ID(ctx, req.(*Filtro))
	}
	return interceptor(ctx, in, info, handler)
}

func _AsistenciaServ_ObtenerPor_CARNET_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filtro)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsistenciaServServer).ObtenerPor_CARNET(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AsistenciaServ/Obtener_por_CARNET",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsistenciaServServer).ObtenerPor_CARNET(ctx, req.(*Filtro))
	}
	return interceptor(ctx, in, info, handler)
}

// AsistenciaServ_ServiceDesc is the grpc.ServiceDesc for AsistenciaServ service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AsistenciaServ_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AsistenciaServ",
	HandlerType: (*AsistenciaServServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Crear",
			Handler:    _AsistenciaServ_Crear_Handler,
		},
		{
			MethodName: "Obtener_por_ID",
			Handler:    _AsistenciaServ_ObtenerPor_ID_Handler,
		},
		{
			MethodName: "Obtener_por_CARNET",
			Handler:    _AsistenciaServ_ObtenerPor_CARNET_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "esquema.proto",
}

// ReporteServClient is the client API for ReporteServ service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReporteServClient interface {
	Crear(ctx context.Context, in *Report, opts ...grpc.CallOption) (*IndividualReport, error)
	ObtenerPor_ID(ctx context.Context, in *Filtro, opts ...grpc.CallOption) (*ReportCollection, error)
	ObtenerPor_CARNET(ctx context.Context, in *Filtro, opts ...grpc.CallOption) (*ReportCollection, error)
}

type reporteServClient struct {
	cc grpc.ClientConnInterface
}

func NewReporteServClient(cc grpc.ClientConnInterface) ReporteServClient {
	return &reporteServClient{cc}
}

func (c *reporteServClient) Crear(ctx context.Context, in *Report, opts ...grpc.CallOption) (*IndividualReport, error) {
	out := new(IndividualReport)
	err := c.cc.Invoke(ctx, "/pb.ReporteServ/Crear", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reporteServClient) ObtenerPor_ID(ctx context.Context, in *Filtro, opts ...grpc.CallOption) (*ReportCollection, error) {
	out := new(ReportCollection)
	err := c.cc.Invoke(ctx, "/pb.ReporteServ/Obtener_por_ID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reporteServClient) ObtenerPor_CARNET(ctx context.Context, in *Filtro, opts ...grpc.CallOption) (*ReportCollection, error) {
	out := new(ReportCollection)
	err := c.cc.Invoke(ctx, "/pb.ReporteServ/Obtener_por_CARNET", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReporteServServer is the server API for ReporteServ service.
// All implementations must embed UnimplementedReporteServServer
// for forward compatibility
type ReporteServServer interface {
	Crear(context.Context, *Report) (*IndividualReport, error)
	ObtenerPor_ID(context.Context, *Filtro) (*ReportCollection, error)
	ObtenerPor_CARNET(context.Context, *Filtro) (*ReportCollection, error)
	mustEmbedUnimplementedReporteServServer()
}

// UnimplementedReporteServServer must be embedded to have forward compatible implementations.
type UnimplementedReporteServServer struct {
}

func (UnimplementedReporteServServer) Crear(context.Context, *Report) (*IndividualReport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Crear not implemented")
}
func (UnimplementedReporteServServer) ObtenerPor_ID(context.Context, *Filtro) (*ReportCollection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ObtenerPor_ID not implemented")
}
func (UnimplementedReporteServServer) ObtenerPor_CARNET(context.Context, *Filtro) (*ReportCollection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ObtenerPor_CARNET not implemented")
}
func (UnimplementedReporteServServer) mustEmbedUnimplementedReporteServServer() {}

// UnsafeReporteServServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReporteServServer will
// result in compilation errors.
type UnsafeReporteServServer interface {
	mustEmbedUnimplementedReporteServServer()
}

func RegisterReporteServServer(s grpc.ServiceRegistrar, srv ReporteServServer) {
	s.RegisterService(&ReporteServ_ServiceDesc, srv)
}

func _ReporteServ_Crear_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Report)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReporteServServer).Crear(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ReporteServ/Crear",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReporteServServer).Crear(ctx, req.(*Report))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReporteServ_ObtenerPor_ID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filtro)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReporteServServer).ObtenerPor_ID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ReporteServ/Obtener_por_ID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReporteServServer).ObtenerPor_ID(ctx, req.(*Filtro))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReporteServ_ObtenerPor_CARNET_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filtro)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReporteServServer).ObtenerPor_CARNET(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ReporteServ/Obtener_por_CARNET",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReporteServServer).ObtenerPor_CARNET(ctx, req.(*Filtro))
	}
	return interceptor(ctx, in, info, handler)
}

// ReporteServ_ServiceDesc is the grpc.ServiceDesc for ReporteServ service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReporteServ_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ReporteServ",
	HandlerType: (*ReporteServServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Crear",
			Handler:    _ReporteServ_Crear_Handler,
		},
		{
			MethodName: "Obtener_por_ID",
			Handler:    _ReporteServ_ObtenerPor_ID_Handler,
		},
		{
			MethodName: "Obtener_por_CARNET",
			Handler:    _ReporteServ_ObtenerPor_CARNET_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "esquema.proto",
}
