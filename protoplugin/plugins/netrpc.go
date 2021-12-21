package plugins

import (
	"github.com/golang/protobuf/protoc-gen-go/generator"
	"google.golang.org/protobuf/types/descriptorpb"
)

//自定义protoc-gen-go插件，通过--go_out=plugins=netrpc 来生成go代码

type NetrpcPlugin struct {
	*generator.Generator
}

func init() {
	//注册到generator
	generator.RegisterPlugin(new(NetrpcPlugin))
}

func (p *NetrpcPlugin) Name() string {
	return "netrpc"
}
func (p *NetrpcPlugin) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *NetrpcPlugin) GenerateImports(file *generator.FileDescriptor) {
	if len(file.Service) > 0 {
		p.genImportCode(file)
	}
}

func (p *NetrpcPlugin) Generate(file *generator.FileDescriptor) {
	for _, svc := range file.Service {
		p.genServiceCode(svc)
	}
}

func (p *NetrpcPlugin) genImportCode(file *generator.FileDescriptor) {
	p.P("// TODO: import code")
}

func (p *NetrpcPlugin) genServiceCode(svc *descriptorpb.ServiceDescriptorProto) {
	p.P("// TODO: service code, Name = " + svc.GetName())
}
