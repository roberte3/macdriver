// AUTO-GENERATED CODE, DO NOT MODIFY

package quartzcore

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [BasicAnimation] class.
var BasicAnimationClass = _BasicAnimationClass{objc.GetClass("CABasicAnimation")}

type _BasicAnimationClass struct {
	objc.Class
}

// An interface definition for the [BasicAnimation] class.
type IBasicAnimation interface {
	IPropertyAnimation
	ByValue() objc.Object
	SetByValue(value objc.IObject)
	FromValue() objc.Object
	SetFromValue(value objc.IObject)
	ToValue() objc.Object
	SetToValue(value objc.IObject)
}

// An object that provides basic, single-keyframe animation capabilities for a layer property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cabasicanimation?language=objc
type BasicAnimation struct {
	PropertyAnimation
}

func BasicAnimationFrom(ptr unsafe.Pointer) BasicAnimation {
	return BasicAnimation{
		PropertyAnimation: PropertyAnimationFrom(ptr),
	}
}

func (bc _BasicAnimationClass) Alloc() BasicAnimation {
	rv := objc.Call[BasicAnimation](bc, objc.Sel("alloc"))
	return rv
}

func BasicAnimation_Alloc() BasicAnimation {
	return BasicAnimationClass.Alloc()
}

func (bc _BasicAnimationClass) New() BasicAnimation {
	rv := objc.Call[BasicAnimation](bc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewBasicAnimation() BasicAnimation {
	return BasicAnimationClass.New()
}

func (b_ BasicAnimation) Init() BasicAnimation {
	rv := objc.Call[BasicAnimation](b_, objc.Sel("init"))
	return rv
}

func (bc _BasicAnimationClass) AnimationWithKeyPath(path string) BasicAnimation {
	rv := objc.Call[BasicAnimation](bc, objc.Sel("animationWithKeyPath:"), path)
	return rv
}

// Creates and returns an CAPropertyAnimation instance for the specified key path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/capropertyanimation/1412534-animationwithkeypath?language=objc
func BasicAnimation_AnimationWithKeyPath(path string) BasicAnimation {
	return BasicAnimationClass.AnimationWithKeyPath(path)
}

func (bc _BasicAnimationClass) Animation() BasicAnimation {
	rv := objc.Call[BasicAnimation](bc, objc.Sel("animation"))
	return rv
}

// Creates and returns a new CAAnimation instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimation/1412479-animation?language=objc
func BasicAnimation_Animation() BasicAnimation {
	return BasicAnimationClass.Animation()
}

// Defines the value the receiver uses to perform relative interpolation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cabasicanimation/1412445-byvalue?language=objc
func (b_ BasicAnimation) ByValue() objc.Object {
	rv := objc.Call[objc.Object](b_, objc.Sel("byValue"))
	return rv
}

// Defines the value the receiver uses to perform relative interpolation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cabasicanimation/1412445-byvalue?language=objc
func (b_ BasicAnimation) SetByValue(value objc.IObject) {
	objc.Call[objc.Void](b_, objc.Sel("setByValue:"), value)
}

// Defines the value the receiver uses to start interpolation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cabasicanimation/1412519-fromvalue?language=objc
func (b_ BasicAnimation) FromValue() objc.Object {
	rv := objc.Call[objc.Object](b_, objc.Sel("fromValue"))
	return rv
}

// Defines the value the receiver uses to start interpolation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cabasicanimation/1412519-fromvalue?language=objc
func (b_ BasicAnimation) SetFromValue(value objc.IObject) {
	objc.Call[objc.Void](b_, objc.Sel("setFromValue:"), value)
}

// Defines the value the receiver uses to end interpolation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cabasicanimation/1412523-tovalue?language=objc
func (b_ BasicAnimation) ToValue() objc.Object {
	rv := objc.Call[objc.Object](b_, objc.Sel("toValue"))
	return rv
}

// Defines the value the receiver uses to end interpolation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cabasicanimation/1412523-tovalue?language=objc
func (b_ BasicAnimation) SetToValue(value objc.IObject) {
	objc.Call[objc.Void](b_, objc.Sel("setToValue:"), value)
}