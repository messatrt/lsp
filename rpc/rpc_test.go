package rpc_test

import (
	"lsp/rpc"
	"testing"
)

type EncodeExample struct{
	Testing bool
}

func TestEncode(t *testing.T){
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := rpc.EncodeMessage(EncodeExample{Testing: true})
	if expected != actual{
		t.Fatalf("Expected: %s, Actual: %s",expected,actual)
	}
}
