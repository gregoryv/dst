package dst

import (
	"net/url"
	"os"
	"strings"
	"testing"
	"time"
)

func TestSetContent(t *testing.T) {
	var dst []byte
	if err := SetContent(&dst, "oups"); err == nil {
		t.Error("expected error")
	}
	if err := SetContent(&dst, "README"); err != nil {
		t.Error("unexpected", err)
	}
	if !strings.Contains(string(dst), "Package") {
		os.Stderr.Write(dst)
		t.Error("unexpected file content")
	}
}

func TestSetFloat(t *testing.T) {
	var dst float64
	if err := SetFloat(&dst, "oups"); err == nil {
		t.Error("expected error")
	}
	if err := SetFloat(&dst, "7.6"); err != nil {
		t.Error("unexpected", err)
	}
	if dst != 7.6 {
		t.Error(dst)
	}
}

func TestSetInt(t *testing.T) {
	var dst int
	if err := SetInt(&dst, "oups"); err == nil {
		t.Error("expected error")
	}
	if err := SetInt(&dst, "1"); err != nil {
		t.Error("unexpected", err)
	}
	if dst != 1 {
		t.Error(dst)
	}
}

func TestSetURL(t *testing.T) {
	var dst url.URL
	if err := SetURL(&dst, "htt p://"); err == nil {
		t.Error("expected error")
	}
	if err := SetURL(&dst, "localhost:80"); err != nil {
		t.Error("unexpected", err)
	}
	if dst.String() != "localhost:80" {
		t.Fail()
	}
}

func TestSetDuration(t *testing.T) {
	var dst time.Duration
	if err := SetDuration(&dst, "oups"); err == nil {
		t.Error("expected error")
	}
	if err := SetDuration(&dst, ""); err == nil {
		t.Error("expected error")
	}
	if err := SetDuration(&dst, "4h2m3s"); err != nil {
		t.Error("unexpected", err)
	}
	if v := dst.String(); v != "4h2m3s" {
		t.Error(v)
	}
}

func TestSetUint32(t *testing.T) {
	var dst uint32
	if err := SetUint32(&dst, "oups"); err == nil {
		t.Error("expected error")
	}
	if err := SetUint32(&dst, "-1"); err == nil {
		t.Error("expected error")
	}
	if err := SetUint32(&dst, "1"); err != nil {
		t.Error("unexpected", err)
	}
	if dst != 1 {
		t.Error(dst)
	}
}

func TestSetBool(t *testing.T) {
	var dst bool
	if err := SetBool(&dst, "oups"); err == nil {
		t.Error("expected error")
	}
	if err := SetBool(&dst, "true"); err != nil {
		t.Error("unexpected", err)
	}
	if err := SetBool(&dst, ""); err != nil {
		t.Error("unexpected", err)
	}
	if dst != true {
		t.Error(dst)
	}
}
