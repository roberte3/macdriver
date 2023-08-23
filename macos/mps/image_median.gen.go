// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageMedian] class.
var ImageMedianClass = _ImageMedianClass{objc.GetClass("MPSImageMedian")}

type _ImageMedianClass struct {
	objc.Class
}

// An interface definition for the [ImageMedian] class.
type IImageMedian interface {
	IUnaryImageKernel
	KernelDiameter() uint
}

// A filter that applies a median filter in a square region centered around each pixel in the source image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagemedian?language=objc
type ImageMedian struct {
	UnaryImageKernel
}

func ImageMedianFrom(ptr unsafe.Pointer) ImageMedian {
	return ImageMedian{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}

func (i_ ImageMedian) InitWithDeviceKernelDiameter(device metal.PDevice, kernelDiameter uint) ImageMedian {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageMedian](i_, objc.Sel("initWithDevice:kernelDiameter:"), po0, kernelDiameter)
	return rv
}

// Initializes a filter for a particular kernel size and device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagemedian/1618837-initwithdevice?language=objc
func NewImageMedianWithDeviceKernelDiameter(device metal.PDevice, kernelDiameter uint) ImageMedian {
	instance := ImageMedianClass.Alloc().InitWithDeviceKernelDiameter(device, kernelDiameter)
	instance.Autorelease()
	return instance
}

func (ic _ImageMedianClass) Alloc() ImageMedian {
	rv := objc.Call[ImageMedian](ic, objc.Sel("alloc"))
	return rv
}

func ImageMedian_Alloc() ImageMedian {
	return ImageMedianClass.Alloc()
}

func (ic _ImageMedianClass) New() ImageMedian {
	rv := objc.Call[ImageMedian](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageMedian() ImageMedian {
	return ImageMedianClass.New()
}

func (i_ ImageMedian) Init() ImageMedian {
	rv := objc.Call[ImageMedian](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageMedian) InitWithDevice(device metal.PDevice) ImageMedian {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageMedian](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/2866332-initwithdevice?language=objc
func NewImageMedianWithDevice(device metal.PDevice) ImageMedian {
	instance := ImageMedianClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (i_ ImageMedian) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageMedian {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageMedian](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageMedian_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageMedian {
	instance := ImageMedianClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

// Queries the minimum diameter, in pixels, of the filter window supported by the median filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagemedian/1618864-minkerneldiameter?language=objc
func (ic _ImageMedianClass) MinKernelDiameter() uint {
	rv := objc.Call[uint](ic, objc.Sel("minKernelDiameter"))
	return rv
}

// Queries the minimum diameter, in pixels, of the filter window supported by the median filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagemedian/1618864-minkerneldiameter?language=objc
func ImageMedian_MinKernelDiameter() uint {
	return ImageMedianClass.MinKernelDiameter()
}

// Queries the maximum diameter, in pixels, of the filter window supported by the median filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagemedian/1618830-maxkerneldiameter?language=objc
func (ic _ImageMedianClass) MaxKernelDiameter() uint {
	rv := objc.Call[uint](ic, objc.Sel("maxKernelDiameter"))
	return rv
}

// Queries the maximum diameter, in pixels, of the filter window supported by the median filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagemedian/1618830-maxkerneldiameter?language=objc
func ImageMedian_MaxKernelDiameter() uint {
	return ImageMedianClass.MaxKernelDiameter()
}

// The diameter, in pixels, of the filter window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagemedian/1618909-kerneldiameter?language=objc
func (i_ ImageMedian) KernelDiameter() uint {
	rv := objc.Call[uint](i_, objc.Sel("kernelDiameter"))
	return rv
}