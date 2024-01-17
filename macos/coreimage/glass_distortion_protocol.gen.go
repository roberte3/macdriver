// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciglassdistortion?language=objc
type PGlassDistortion interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() Image
	HasInputImage() bool

	// optional
	SetScale(value float64)
	HasSetScale() bool

	// optional
	Scale() float64
	HasScale() bool

	// optional
	SetCenter(value coregraphics.Point)
	HasSetCenter() bool

	// optional
	Center() coregraphics.Point
	HasCenter() bool

	// optional
	SetTextureImage(value Image)
	HasSetTextureImage() bool

	// optional
	TextureImage() Image
	HasTextureImage() bool
}

// ensure impl type implements protocol interface
var _ PGlassDistortion = (*GlassDistortionObject)(nil)

// A concrete type for the [PGlassDistortion] protocol.
type GlassDistortionObject struct {
	objc.Object
}

func (g_ GlassDistortionObject) HasSetInputImage() bool {
	return g_.RespondsToSelector(objc.Sel("setInputImage:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciglassdistortion/3600157-inputimage?language=objc
func (g_ GlassDistortionObject) SetInputImage(value Image) {
	objc.Call[objc.Void](g_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (g_ GlassDistortionObject) HasInputImage() bool {
	return g_.RespondsToSelector(objc.Sel("inputImage"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciglassdistortion/3600157-inputimage?language=objc
func (g_ GlassDistortionObject) InputImage() Image {
	rv := objc.Call[Image](g_, objc.Sel("inputImage"))
	return rv
}

func (g_ GlassDistortionObject) HasSetScale() bool {
	return g_.RespondsToSelector(objc.Sel("setScale:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciglassdistortion/3600158-scale?language=objc
func (g_ GlassDistortionObject) SetScale(value float64) {
	objc.Call[objc.Void](g_, objc.Sel("setScale:"), value)
}

func (g_ GlassDistortionObject) HasScale() bool {
	return g_.RespondsToSelector(objc.Sel("scale"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciglassdistortion/3600158-scale?language=objc
func (g_ GlassDistortionObject) Scale() float64 {
	rv := objc.Call[float64](g_, objc.Sel("scale"))
	return rv
}

func (g_ GlassDistortionObject) HasSetCenter() bool {
	return g_.RespondsToSelector(objc.Sel("setCenter:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciglassdistortion/3600156-center?language=objc
func (g_ GlassDistortionObject) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](g_, objc.Sel("setCenter:"), value)
}

func (g_ GlassDistortionObject) HasCenter() bool {
	return g_.RespondsToSelector(objc.Sel("center"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciglassdistortion/3600156-center?language=objc
func (g_ GlassDistortionObject) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](g_, objc.Sel("center"))
	return rv
}

func (g_ GlassDistortionObject) HasSetTextureImage() bool {
	return g_.RespondsToSelector(objc.Sel("setTextureImage:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciglassdistortion/3600159-textureimage?language=objc
func (g_ GlassDistortionObject) SetTextureImage(value Image) {
	objc.Call[objc.Void](g_, objc.Sel("setTextureImage:"), objc.Ptr(value))
}

func (g_ GlassDistortionObject) HasTextureImage() bool {
	return g_.RespondsToSelector(objc.Sel("textureImage"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciglassdistortion/3600159-textureimage?language=objc
func (g_ GlassDistortionObject) TextureImage() Image {
	rv := objc.Call[Image](g_, objc.Sel("textureImage"))
	return rv
}