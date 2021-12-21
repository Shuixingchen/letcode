package services

import (
	"reflect"
	"testing"
)

type Foo int

type Args struct {
	Num1, Num2 int
}

func (f Foo) Sum(args Args, reply *int) error {
	*reply = args.Num1 + args.Num2
	return nil
}

// it's not a exported Method
func (f Foo) sum(args Args, reply *int) error {
	*reply = args.Num1 + args.Num2
	return nil
}

func TestNewService(t *testing.T) {
	var foo Foo
	s := newService(foo)
	mType := s.method["Sum"]
	if mType == nil {
		t.Error("newservice wrong")
	}
}
func TestMethodType_Call(t *testing.T) {
	var foo Foo
	s := newService(foo)
	mType := s.method["Sum"]
	if mType == nil {
		t.Error("newservice wrong")
	}
	argv := mType.newArgv()
	replyv := mType.newReplyv()
	argv.Set(reflect.ValueOf(Args{Num1: 1, Num2: 3}))
	err := s.call(mType, argv, replyv)
}
