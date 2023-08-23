// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNConvolutionGradient] class.
var CNNConvolutionGradientClass = _CNNConvolutionGradientClass{objc.GetClass("MPSCNNConvolutionGradient")}

type _CNNConvolutionGradientClass struct {
	objc.Class
}

// An interface definition for the [CNNConvolutionGradient] class.
type ICNNConvolutionGradient interface {
	ICNNGradientKernel
	ReloadWeightsAndBiasesFromDataSource()
	ReloadWeightsAndBiasesWithCommandBufferState(commandBuffer metal.PCommandBuffer, state ICNNConvolutionWeightsAndBiasesState)
	ReloadWeightsAndBiasesWithCommandBufferObjectState(commandBufferObject objc.IObject, state ICNNConvolutionWeightsAndBiasesState)
	DataSource() CNNConvolutionDataSourceWrapper
	SourceImageFeatureChannels() uint
	ChannelMultiplier() uint
	GradientOption() CNNConvolutionGradientOption
	SetGradientOption(value CNNConvolutionGradientOption)
	SourceGradientFeatureChannels() uint
	Groups() uint
}

// A gradient convolution kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradient?language=objc
type CNNConvolutionGradient struct {
	CNNGradientKernel
}

func CNNConvolutionGradientFrom(ptr unsafe.Pointer) CNNConvolutionGradient {
	return CNNConvolutionGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}

func (c_ CNNConvolutionGradient) InitWithDeviceWeights(device metal.PDevice, weights PCNNConvolutionDataSource) CNNConvolutionGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	po1 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", weights)
	rv := objc.Call[CNNConvolutionGradient](c_, objc.Sel("initWithDevice:weights:"), po0, po1)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradient/2942414-initwithdevice?language=objc
func NewCNNConvolutionGradientWithDeviceWeights(device metal.PDevice, weights PCNNConvolutionDataSource) CNNConvolutionGradient {
	instance := CNNConvolutionGradientClass.Alloc().InitWithDeviceWeights(device, weights)
	instance.Autorelease()
	return instance
}

func (cc _CNNConvolutionGradientClass) Alloc() CNNConvolutionGradient {
	rv := objc.Call[CNNConvolutionGradient](cc, objc.Sel("alloc"))
	return rv
}

func CNNConvolutionGradient_Alloc() CNNConvolutionGradient {
	return CNNConvolutionGradientClass.Alloc()
}

func (cc _CNNConvolutionGradientClass) New() CNNConvolutionGradient {
	rv := objc.Call[CNNConvolutionGradient](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNConvolutionGradient() CNNConvolutionGradient {
	return CNNConvolutionGradientClass.New()
}

func (c_ CNNConvolutionGradient) Init() CNNConvolutionGradient {
	rv := objc.Call[CNNConvolutionGradient](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNConvolutionGradient) InitWithDevice(device metal.PDevice) CNNConvolutionGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNConvolutionGradient](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942657-initwithdevice?language=objc
func NewCNNConvolutionGradientWithDevice(device metal.PDevice) CNNConvolutionGradient {
	instance := CNNConvolutionGradientClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNConvolutionGradient) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNConvolutionGradient {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNConvolutionGradient](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNConvolutionGradient_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNConvolutionGradient {
	instance := CNNConvolutionGradientClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradient/2966659-reloadweightsandbiasesfromdataso?language=objc
func (c_ CNNConvolutionGradient) ReloadWeightsAndBiasesFromDataSource() {
	objc.Call[objc.Void](c_, objc.Sel("reloadWeightsAndBiasesFromDataSource"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradient/2953960-reloadweightsandbiaseswithcomman?language=objc
func (c_ CNNConvolutionGradient) ReloadWeightsAndBiasesWithCommandBufferState(commandBuffer metal.PCommandBuffer, state ICNNConvolutionWeightsAndBiasesState) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](c_, objc.Sel("reloadWeightsAndBiasesWithCommandBuffer:state:"), po0, objc.Ptr(state))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradient/2953960-reloadweightsandbiaseswithcomman?language=objc
func (c_ CNNConvolutionGradient) ReloadWeightsAndBiasesWithCommandBufferObjectState(commandBufferObject objc.IObject, state ICNNConvolutionWeightsAndBiasesState) {
	objc.Call[objc.Void](c_, objc.Sel("reloadWeightsAndBiasesWithCommandBuffer:state:"), objc.Ptr(commandBufferObject), objc.Ptr(state))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradient/2953959-datasource?language=objc
func (c_ CNNConvolutionGradient) DataSource() CNNConvolutionDataSourceWrapper {
	rv := objc.Call[CNNConvolutionDataSourceWrapper](c_, objc.Sel("dataSource"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradient/2947882-sourceimagefeaturechannels?language=objc
func (c_ CNNConvolutionGradient) SourceImageFeatureChannels() uint {
	rv := objc.Call[uint](c_, objc.Sel("sourceImageFeatureChannels"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradient/2966658-channelmultiplier?language=objc
func (c_ CNNConvolutionGradient) ChannelMultiplier() uint {
	rv := objc.Call[uint](c_, objc.Sel("channelMultiplier"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradient/2942432-gradientoption?language=objc
func (c_ CNNConvolutionGradient) GradientOption() CNNConvolutionGradientOption {
	rv := objc.Call[CNNConvolutionGradientOption](c_, objc.Sel("gradientOption"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradient/2942432-gradientoption?language=objc
func (c_ CNNConvolutionGradient) SetGradientOption(value CNNConvolutionGradientOption) {
	objc.Call[objc.Void](c_, objc.Sel("setGradientOption:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradient/2947880-sourcegradientfeaturechannels?language=objc
func (c_ CNNConvolutionGradient) SourceGradientFeatureChannels() uint {
	rv := objc.Call[uint](c_, objc.Sel("sourceGradientFeatureChannels"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradient/2942430-groups?language=objc
func (c_ CNNConvolutionGradient) Groups() uint {
	rv := objc.Call[uint](c_, objc.Sel("groups"))
	return rv
}