package main

import (
	"letcode/myreflect/services"
	"reflect"
	"unicode"
)

type callback struct {
	fn       reflect.Value  //函数的reflect.Value
	rcvr     reflect.Value  //服务实例的reflect.Value
	argTypes []reflect.Type //参数的reflect.Type
	outTypes []reflect.Type //函数返回值reflect.Type
}

type Mux struct {
	fnMap map[string]*callback
}

var DefaultMux = NewMux()

func NewMux() *Mux {
	return &Mux{fnMap: make(map[string]*callback)}
}

func main() {
	var user services.User
	RegisterService(&user)
	CallServerMethod("Say", []string{"hello"})
}

//访问注册的服务user.say
func CallServerMethod(call string, argc []string) {
	callback := DefaultMux.FindCallBack(call)
	args := parserArgc(argc)
	if callback != nil {
		callback.fn.Call(args)
	}
}

func parserArgc(argc []string) []reflect.Value {
	res := make([]reflect.Value, 0)
	for _, item := range argc {
		res = append(res, reflect.ValueOf(item))
	}
	return res
}

func RegisterService(server interface{}) {
	v := reflect.ValueOf(server)
	callmap := suitableCallbacks(v)
	for key := range callmap {
		DefaultMux.fnMap[key] = callmap[key]
	}
	return
}

func (m *Mux) FindCallBack(method string) *callback {
	if callback, ok := DefaultMux.fnMap[method]; ok {
		return callback
	}
	return nil
}

func suitableCallbacks(receiver reflect.Value) map[string]*callback {

	typ := receiver.Type()
	callbacks := make(map[string]*callback)
	for m := 0; m < typ.NumMethod(); m++ {
		method := typ.Method(m)
		if method.PkgPath != "" {
			continue // method not exported
		}
		cb := newCallback(receiver, method.Func)
		if cb == nil {
			continue // function invalid
		}
		methodName := formatName(method.Name)
		callbacks[methodName] = cb
	}
	return callbacks
}

func newCallback(receiver, fn reflect.Value) *callback {
	fntype := fn.Type()
	c := &callback{fn: fn, rcvr: receiver}

	//获取函数参数reflect.Type
	c.makeArgTypes()

	//获取函数的返回值reflect.Type
	c.outTypes = make([]reflect.Type, fntype.NumOut())
	for i := 0; i < fntype.NumOut(); i++ {
		c.outTypes[i] = fntype.Out(i)
	}
	return c
}

//首字母小写
func formatName(name string) string {
	ret := []rune(name)
	if len(ret) > 0 {
		ret[0] = unicode.ToLower(ret[0])
	}
	return string(ret)
}

//获取函数参数的reflect.Type
func (c *callback) makeArgTypes() {
	fntype := c.fn.Type()
	// Skip receiver and context.Context parameter (if present).
	firstArg := 0
	if c.rcvr.IsValid() {
		firstArg++
	}
	// Add all remaining parameters.
	c.argTypes = make([]reflect.Type, fntype.NumIn()-firstArg)
	for i := firstArg; i < fntype.NumIn(); i++ {
		c.argTypes[i-firstArg] = fntype.In(i)
	}
}
