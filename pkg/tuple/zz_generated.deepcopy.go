//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by deepcopy-gen. DO NOT EDIT.

package tuple

import (
	bpf "github.com/go-faster/cilium/pkg/bpf"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TupleKey4) DeepCopyInto(out *TupleKey4) {
	*out = *in
	in.DestAddr.DeepCopyInto(&out.DestAddr)
	in.SourceAddr.DeepCopyInto(&out.SourceAddr)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TupleKey4.
func (in *TupleKey4) DeepCopy() *TupleKey4 {
	if in == nil {
		return nil
	}
	out := new(TupleKey4)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyMapKey is an autogenerated deepcopy function, copying the receiver, creating a new bpf.MapKey.
func (in *TupleKey4) DeepCopyMapKey() bpf.MapKey {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TupleKey4Global) DeepCopyInto(out *TupleKey4Global) {
	*out = *in
	in.TupleKey4.DeepCopyInto(&out.TupleKey4)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TupleKey4Global.
func (in *TupleKey4Global) DeepCopy() *TupleKey4Global {
	if in == nil {
		return nil
	}
	out := new(TupleKey4Global)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyMapKey is an autogenerated deepcopy function, copying the receiver, creating a new bpf.MapKey.
func (in *TupleKey4Global) DeepCopyMapKey() bpf.MapKey {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TupleKey6) DeepCopyInto(out *TupleKey6) {
	*out = *in
	in.DestAddr.DeepCopyInto(&out.DestAddr)
	in.SourceAddr.DeepCopyInto(&out.SourceAddr)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TupleKey6.
func (in *TupleKey6) DeepCopy() *TupleKey6 {
	if in == nil {
		return nil
	}
	out := new(TupleKey6)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyMapKey is an autogenerated deepcopy function, copying the receiver, creating a new bpf.MapKey.
func (in *TupleKey6) DeepCopyMapKey() bpf.MapKey {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TupleKey6Global) DeepCopyInto(out *TupleKey6Global) {
	*out = *in
	in.TupleKey6.DeepCopyInto(&out.TupleKey6)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TupleKey6Global.
func (in *TupleKey6Global) DeepCopy() *TupleKey6Global {
	if in == nil {
		return nil
	}
	out := new(TupleKey6Global)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyMapKey is an autogenerated deepcopy function, copying the receiver, creating a new bpf.MapKey.
func (in *TupleKey6Global) DeepCopyMapKey() bpf.MapKey {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TupleValStub) DeepCopyInto(out *TupleValStub) {
	*out = *in
	in.buff.DeepCopyInto(&out.buff)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TupleValStub.
func (in *TupleValStub) DeepCopy() *TupleValStub {
	if in == nil {
		return nil
	}
	out := new(TupleValStub)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyMapValue is an autogenerated deepcopy function, copying the receiver, creating a new bpf.MapValue.
func (in *TupleValStub) DeepCopyMapValue() bpf.MapValue {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
