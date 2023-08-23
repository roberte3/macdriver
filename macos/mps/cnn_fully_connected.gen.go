// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNFullyConnected] class.
var CNNFullyConnectedClass = _CNNFullyConnectedClass{objc.GetClass("MPSCNNFullyConnected")}

type _CNNFullyConnectedClass struct {
	objc.Class
}

// An interface definition for the [CNNFullyConnected] class.
type ICNNFullyConnected interface {
	ICNNConvolution
}

// A fully connected convolution layer, also known as an inner product layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnfullyconnected?language=objc
type CNNFullyConnected struct {
	CNNConvolution
}

func CNNFullyConnectedFrom(ptr unsafe.Pointer) CNNFullyConnected {
	return CNNFullyConnected{
		CNNConvolution: CNNConvolutionFrom(ptr),
	}
}

func (c_ CNNFullyConnected) InitWithDeviceWeights(device metal.PDevice, weights PCNNConvolutionDataSource) CNNFullyConnected {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	po1 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", weights)
	rv := objc.Call[CNNFullyConnected](c_, objc.Sel("initWithDevice:weights:"), po0, po1)
	return rv
}

// Initializes a fully connected convolution layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnfullyconnected/2867198-initwithdevice?language=objc
func NewCNNFullyConnectedWithDeviceWeights(device metal.PDevice, weights PCNNConvolutionDataSource) CNNFullyConnected {
	instance := CNNFullyConnectedClass.Alloc().InitWithDeviceWeights(device, weights)
	instance.Autorelease()
	return instance
}

func (cc _CNNFullyConnectedClass) Alloc() CNNFullyConnected {
	rv := objc.Call[CNNFullyConnected](cc, objc.Sel("alloc"))
	return rv
}

func CNNFullyConnected_Alloc() CNNFullyConnected {
	return CNNFullyConnectedClass.Alloc()
}

func (cc _CNNFullyConnectedClass) New() CNNFullyConnected {
	rv := objc.Call[CNNFullyConnected](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNFullyConnected() CNNFullyConnected {
	return CNNFullyConnectedClass.New()
}

func (c_ CNNFullyConnected) Init() CNNFullyConnected {
	rv := objc.Call[CNNFullyConnected](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNFullyConnected) InitWithDevice(device metal.PDevice) CNNFullyConnected {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNFullyConnected](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewCNNFullyConnectedWithDevice(device metal.PDevice) CNNFullyConnected {
	instance := CNNFullyConnectedClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNFullyConnected) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNFullyConnected {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNFullyConnected](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNFullyConnected_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNFullyConnected {
	instance := CNNFullyConnectedClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}