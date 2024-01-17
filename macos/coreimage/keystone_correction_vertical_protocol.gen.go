// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a keystone correction vertical filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikeystonecorrectionvertical?language=objc
type PKeystoneCorrectionVertical interface {
	// optional
	SetFocalLength(value float64)
	HasSetFocalLength() bool

	// optional
	FocalLength() float64
	HasFocalLength() bool
}

// ensure impl type implements protocol interface
var _ PKeystoneCorrectionVertical = (*KeystoneCorrectionVerticalObject)(nil)

// A concrete type for the [PKeystoneCorrectionVertical] protocol.
type KeystoneCorrectionVerticalObject struct {
	objc.Object
}

func (k_ KeystoneCorrectionVerticalObject) HasSetFocalLength() bool {
	return k_.RespondsToSelector(objc.Sel("setFocalLength:"))
}

// The 35mm equivalent focal length of the input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikeystonecorrectionvertical/3325532-focallength?language=objc
func (k_ KeystoneCorrectionVerticalObject) SetFocalLength(value float64) {
	objc.Call[objc.Void](k_, objc.Sel("setFocalLength:"), value)
}

func (k_ KeystoneCorrectionVerticalObject) HasFocalLength() bool {
	return k_.RespondsToSelector(objc.Sel("focalLength"))
}

// The 35mm equivalent focal length of the input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cikeystonecorrectionvertical/3325532-focallength?language=objc
func (k_ KeystoneCorrectionVerticalObject) FocalLength() float64 {
	rv := objc.Call[float64](k_, objc.Sel("focalLength"))
	return rv
}