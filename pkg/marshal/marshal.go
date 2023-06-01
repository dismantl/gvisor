// Copyright 2019 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package marshal defines the Marshallable interface for
// serialize/deserializing go data structures to/from memory, according to the
// Linux ABI.
//
// Implementations of this interface are typically automatically generated by
// tools/go_marshal. See the go_marshal README for details.
package marshal

import (
	"io"

	"github.com/dismantl/gvisor/pkg/hostarch"
)

// CopyContext defines the memory operations required to marshal to and from
// user memory. Typically, kernel.Task is used to provide implementations for
// these operations.
type CopyContext interface {
	// CopyScratchBuffer provides a task goroutine-local scratch buffer. See
	// kernel.CopyScratchBuffer.
	CopyScratchBuffer(size int) []byte

	// CopyOutBytes writes the contents of b to the task's memory. See
	// kernel.CopyOutBytes.
	CopyOutBytes(addr hostarch.Addr, b []byte) (int, error)

	// CopyInBytes reads the contents of the task's memory to b. See
	// kernel.CopyInBytes.
	CopyInBytes(addr hostarch.Addr, b []byte) (int, error)
}

// Marshallable represents operations on a type that can be marshalled to and
// from memory.
//
// go-marshal automatically generates implementations for this interface for
// types marked as '+marshal'.
type Marshallable interface {
	io.WriterTo

	// SizeBytes is the size of the memory representation of a type in
	// marshalled form.
	//
	// SizeBytes must handle a nil receiver. Practically, this means SizeBytes
	// cannot deference any fields on the object implementing it (but will
	// likely make use of the type of these fields).
	SizeBytes() int

	// MarshalBytes serializes a copy of a type to dst and returns the remaining
	// buffer.
	// Precondition: dst must be at least SizeBytes() in length.
	MarshalBytes(dst []byte) []byte

	// UnmarshalBytes deserializes a type from src and returns the remaining
	// buffer.
	// Precondition: src must be at least SizeBytes() in length.
	UnmarshalBytes(src []byte) []byte

	// Packed returns true if the marshalled size of the type is the same as the
	// size it occupies in memory. This happens when the type has no fields
	// starting at unaligned addresses (should always be true by default for ABI
	// structs, verified by automatically generated tests when using
	// go_marshal), and has no fields marked `marshal:"unaligned"`.
	//
	// Packed must return the same result for all possible values of the type
	// implementing it. Violating this constraint implies the type doesn't have
	// a static memory layout, and will lead to memory corruption.
	// Go-marshal-generated code reuses the result of Packed for multiple values
	// of the same type.
	Packed() bool

	// MarshalUnsafe serializes a type by bulk copying its in-memory
	// representation to the dst buffer. This is only safe to do when the type
	// has no implicit padding, see Marshallable.Packed. When Packed would
	// return false, MarshalUnsafe should fall back to the safer but slower
	// MarshalBytes.
	// Precondition: dst must be at least SizeBytes() in length.
	MarshalUnsafe(dst []byte) []byte

	// UnmarshalUnsafe deserializes a type by directly copying to the underlying
	// memory allocated for the object by the runtime.
	//
	// This allows much faster unmarshalling of types which have no implicit
	// padding, see Marshallable.Packed. When Packed would return false,
	// UnmarshalUnsafe should fall back to the safer but slower unmarshal
	// mechanism implemented in UnmarshalBytes.
	// Precondition: src must be at least SizeBytes() in length.
	UnmarshalUnsafe(src []byte) []byte

	// CopyIn deserializes a Marshallable type from a task's memory. This may
	// only be called from a task goroutine. This is more efficient than calling
	// UnmarshalUnsafe on Marshallable.Packed types, as the type being
	// marshalled does not escape. The implementation should avoid creating
	// extra copies in memory by directly deserializing to the object's
	// underlying memory.
	//
	// If the copy-in from the task memory is only partially successful, CopyIn
	// should still attempt to deserialize as much data as possible. See comment
	// for UnmarshalBytes.
	CopyIn(cc CopyContext, addr hostarch.Addr) (int, error)

	// CopyOut serializes a Marshallable type to a task's memory. This may only
	// be called from a task goroutine. This is more efficient than calling
	// MarshalUnsafe on Marshallable.Packed types, as the type being serialized
	// does not escape. The implementation should avoid creating extra copies in
	// memory by directly serializing from the object's underlying memory.
	//
	// The copy-out to the task memory may be partially successful, in which
	// case CopyOut returns how much data was serialized. See comment for
	// MarshalBytes for implications.
	CopyOut(cc CopyContext, addr hostarch.Addr) (int, error)

	// CopyOutN is like CopyOut, but explicitly requests a partial
	// copy-out. Note that this may yield unexpected results for non-packed
	// types and the caller may only want to allow this for packed types. See
	// comment on MarshalBytes.
	//
	// The limit must be less than or equal to SizeBytes().
	CopyOutN(cc CopyContext, addr hostarch.Addr, limit int) (int, error)
}

// CheckedMarshallable represents operations on a type that can be marshalled
// to and from memory and additionally does bound checking.
type CheckedMarshallable interface {
	// CheckedMarshal is the same as Marshallable.MarshalUnsafe but without the
	// precondition that dst must at least have some appropriate length. Similar
	// to Marshallable.MarshalBytes, it returns a shifted slice according to how
	// much data is consumed. Additionally it returns a bool indicating whether
	// marshalling was successful. Unsuccessful marshalling doesn't consume any
	// data.
	CheckedMarshal(dst []byte) ([]byte, bool)

	// CheckedUnmarshal is the same as Marshallable.UmarshalUnsafe but without
	// the precondition that src must at least have some appropriate length.
	// Similar to Marshallable.UnmarshalBytes, it returns a shifted slice
	// according to how much data is consumed. Additionally it returns a bool
	// indicating whether marshalling was successful. Unsuccessful marshalling
	// doesn't consume any data.
	CheckedUnmarshal(src []byte) ([]byte, bool)
}

// go-marshal generates additional functions for a type based on additional
// clauses to the +marshal directive. They are documented below.
//
// Slice API
// =========
//
// Adding a "slice" clause to the +marshal directive for structs or newtypes on
// primitives like this:
//
// // +marshal slice:FooSlice
// type Foo struct { ... }
//
// Generates four additional functions for marshalling slices of Foos like this:
//
// // MarshalUnsafeFooSlice is like Foo.MarshalUnsafe, buf for a []Foo. It
// // might be more efficient that repeatedly calling Foo.MarshalUnsafe
// // over a []Foo in a loop if the type is Packed.
// // Preconditions: dst must be at least len(src)*Foo.SizeBytes() in length.
// func MarshalUnsafeFooSlice(src []Foo, dst []byte) []byte { ... }
//
// // UnmarshalUnsafeFooSlice is like Foo.UnmarshalUnsafe, buf for a []Foo. It
// // might be more efficient that repeatedly calling Foo.UnmarshalUnsafe
// // over a []Foo in a loop if the type is Packed.
// // Preconditions: src must be at least len(dst)*Foo.SizeBytes() in length.
// func UnmarshalUnsafeFooSlice(dst []Foo, src []byte) []byte { ... }
//
// // CopyFooSliceIn copies in a slice of Foo objects from the task's memory.
// func CopyFooSliceIn(cc marshal.CopyContext, addr hostarch.Addr, dst []Foo) (int, error) { ... }
//
// // CopyFooSliceIn copies out a slice of Foo objects to the task's memory.
// func CopyFooSliceOut(cc marshal.CopyContext, addr hostarch.Addr, src []Foo) (int, error) { ... }
//
// The name of the functions are of the format "Copy%sIn" and "Copy%sOut", where
// %s is the first argument to the slice clause. This directive is not supported
// for newtypes on arrays.
//
// Note: Partial copies are not supported for Slice API UnmarshalUnsafe and
// MarshalUnsafe.
//
// The slice clause also takes an optional second argument, which must be the
// value "inner":
//
// // +marshal slice:Int32Slice:inner
// type Int32 int32
//
// This is only valid on newtypes on primitives, and causes the generated
// functions to accept slices of the inner type instead:
//
// func CopyInt32SliceIn(cc marshal.CopyContext, addr hostarch.Addr, dst []int32) (int, error) { ... }
//
// Without "inner", they would instead be:
//
// func CopyInt32SliceIn(cc marshal.CopyContext, addr hostarch.Addr, dst []Int32) (int, error) { ... }
//
// This may help avoid a cast depending on how the generated functions are used.
//
// Bound Checking
// ==============
//
// Some users might want to do bound checking on marshal and unmarshal. This is
// is useful when the user does not control the buffer size. To prevent
// repeated bound checking code around Marshallable, users can add a
// "boundCheck" clause to the +marshal directive. go_marshal will generate the
// CheckedMarshallable interface methods on the type.
