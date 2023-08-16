// AUTO-GENERATED CODE, DO NOT MODIFY

package quartzcore

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RemoteLayerClient] class.
var RemoteLayerClientClass = _RemoteLayerClientClass{objc.GetClass("CARemoteLayerClient")}

type _RemoteLayerClientClass struct {
	objc.Class
}

// An interface definition for the [RemoteLayerClient] class.
type IRemoteLayerClient interface {
	objc.IObject
	Invalidate()
	Layer() Layer
	SetLayer(value ILayer)
	ClientId() uint32
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caremotelayerclient?language=objc
type RemoteLayerClient struct {
	objc.Object
}

func RemoteLayerClientFrom(ptr unsafe.Pointer) RemoteLayerClient {
	return RemoteLayerClient{
		Object: objc.ObjectFrom(ptr),
	}
}

func (rc _RemoteLayerClientClass) Alloc() RemoteLayerClient {
	rv := objc.Call[RemoteLayerClient](rc, objc.Sel("alloc"))
	return rv
}

func RemoteLayerClient_Alloc() RemoteLayerClient {
	return RemoteLayerClientClass.Alloc()
}

func (rc _RemoteLayerClientClass) New() RemoteLayerClient {
	rv := objc.Call[RemoteLayerClient](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRemoteLayerClient() RemoteLayerClient {
	return RemoteLayerClientClass.New()
}

func (r_ RemoteLayerClient) Init() RemoteLayerClient {
	rv := objc.Call[RemoteLayerClient](r_, objc.Sel("init"))
	return rv
}

// Invalidates a remote layer client. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caremotelayerclient/1418372-invalidate?language=objc
func (r_ RemoteLayerClient) Invalidate() {
	objc.Call[objc.Void](r_, objc.Sel("invalidate"))
}

// The layer associated with the remote client. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caremotelayerclient/1418373-layer?language=objc
func (r_ RemoteLayerClient) Layer() Layer {
	rv := objc.Call[Layer](r_, objc.Sel("layer"))
	return rv
}

// The layer associated with the remote client. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caremotelayerclient/1418373-layer?language=objc
func (r_ RemoteLayerClient) SetLayer(value ILayer) {
	objc.Call[objc.Void](r_, objc.Sel("setLayer:"), objc.Ptr(value))
}

// The ID of the remote layer client. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caremotelayerclient/1418375-clientid?language=objc
func (r_ RemoteLayerClient) ClientId() uint32 {
	rv := objc.Call[uint32](r_, objc.Sel("clientId"))
	return rv
}