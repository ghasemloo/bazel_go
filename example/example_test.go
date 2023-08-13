package example

import (
	"testing"

	glog "github.com/golang/glog"

	anypb "google.golang.org/protobuf/types/known/anypb"
	epb "google.golang.org/protobuf/types/known/emptypb"
)

func Test(t *testing.T) {
	resp, err := anypb.New(&epb.Empty{})
	if err != nil {
		glog.Fatalf("anypb.New() failed: %v", err)
	}

	_, err = resp.UnmarshalNew()
	if err != nil {
		t.Fatalf("UnmarshalNew() failed: %v", err)
	}
}
