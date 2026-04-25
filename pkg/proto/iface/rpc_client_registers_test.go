package iface

import "testing"

func TestFindRPCContextTupleByClazzID(t *testing.T) {
	const clazzID = uint32(0xfeed0001)
	RegisterClazzName("test_raw_method", 0, clazzID)
	RegisterClazzIDName("test_raw_method", clazzID)
	RegisterRPCContextTuple("TLTestRawMethod", "/tg.RPCTest/test.rawMethod", func() interface{} {
		return nil
	})

	tuple := FindRPCContextTupleByClazzID(clazzID)
	if tuple == nil {
		t.Fatal("FindRPCContextTupleByClazzID() = nil, want tuple")
	}
	if tuple.Method != "/tg.RPCTest/test.rawMethod" {
		t.Fatalf("tuple.Method = %q, want %q", tuple.Method, "/tg.RPCTest/test.rawMethod")
	}
	if got := tuple.ServiceName(); got != "RPCTest" {
		t.Fatalf("tuple.ServiceName() = %q, want %q", got, "RPCTest")
	}
}

func TestFindRPCContextTupleByClazzIDReturnsNilForUnknownClazzID(t *testing.T) {
	if tuple := FindRPCContextTupleByClazzID(0xfeed0002); tuple != nil {
		t.Fatalf("FindRPCContextTupleByClazzID() = %#v, want nil", tuple)
	}
}
