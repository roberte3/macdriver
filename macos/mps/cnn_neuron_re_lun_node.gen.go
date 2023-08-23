// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNNeuronReLUNNode] class.
var CNNNeuronReLUNNodeClass = _CNNNeuronReLUNNodeClass{objc.GetClass("MPSCNNNeuronReLUNNode")}

type _CNNNeuronReLUNNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNNeuronReLUNNode] class.
type ICNNNeuronReLUNNode interface {
	ICNNNeuronNode
}

// A representation a ReLUN neuron filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronrelunnode?language=objc
type CNNNeuronReLUNNode struct {
	CNNNeuronNode
}

func CNNNeuronReLUNNodeFrom(ptr unsafe.Pointer) CNNNeuronReLUNNode {
	return CNNNeuronReLUNNode{
		CNNNeuronNode: CNNNeuronNodeFrom(ptr),
	}
}

func (cc _CNNNeuronReLUNNodeClass) NodeWithSource(sourceNode INNImageNode) CNNNeuronReLUNNode {
	rv := objc.Call[CNNNeuronReLUNNode](cc, objc.Sel("nodeWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronrelunnode/2921594-nodewithsource?language=objc
func CNNNeuronReLUNNode_NodeWithSource(sourceNode INNImageNode) CNNNeuronReLUNNode {
	return CNNNeuronReLUNNodeClass.NodeWithSource(sourceNode)
}

func (c_ CNNNeuronReLUNNode) InitWithSource(sourceNode INNImageNode) CNNNeuronReLUNNode {
	rv := objc.Call[CNNNeuronReLUNNode](c_, objc.Sel("initWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronrelunnode/2921593-initwithsource?language=objc
func NewCNNNeuronReLUNNodeWithSource(sourceNode INNImageNode) CNNNeuronReLUNNode {
	instance := CNNNeuronReLUNNodeClass.Alloc().InitWithSource(sourceNode)
	instance.Autorelease()
	return instance
}

func (cc _CNNNeuronReLUNNodeClass) Alloc() CNNNeuronReLUNNode {
	rv := objc.Call[CNNNeuronReLUNNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNNeuronReLUNNode_Alloc() CNNNeuronReLUNNode {
	return CNNNeuronReLUNNodeClass.Alloc()
}

func (cc _CNNNeuronReLUNNodeClass) New() CNNNeuronReLUNNode {
	rv := objc.Call[CNNNeuronReLUNNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNNeuronReLUNNode() CNNNeuronReLUNNode {
	return CNNNeuronReLUNNodeClass.New()
}

func (c_ CNNNeuronReLUNNode) Init() CNNNeuronReLUNNode {
	rv := objc.Call[CNNNeuronReLUNNode](c_, objc.Sel("init"))
	return rv
}

func (cc _CNNNeuronReLUNNodeClass) NodeWithSourceDescriptor(sourceNode INNImageNode, descriptor INNNeuronDescriptor) CNNNeuronReLUNNode {
	rv := objc.Call[CNNNeuronReLUNNode](cc, objc.Sel("nodeWithSource:descriptor:"), objc.Ptr(sourceNode), objc.Ptr(descriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronnode/3019333-nodewithsource?language=objc
func CNNNeuronReLUNNode_NodeWithSourceDescriptor(sourceNode INNImageNode, descriptor INNNeuronDescriptor) CNNNeuronReLUNNode {
	return CNNNeuronReLUNNodeClass.NodeWithSourceDescriptor(sourceNode, descriptor)
}