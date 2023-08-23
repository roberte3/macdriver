// AUTO-GENERATED CODE, DO NOT MODIFY

package mpsgraph

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/mps"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RandomOpDescriptor] class.
var RandomOpDescriptorClass = _RandomOpDescriptorClass{objc.GetClass("MPSGraphRandomOpDescriptor")}

type _RandomOpDescriptorClass struct {
	objc.Class
}

// An interface definition for the [RandomOpDescriptor] class.
type IRandomOpDescriptor interface {
	objc.IObject
	Distribution() RandomDistribution
	SetDistribution(value RandomDistribution)
	MinInteger() int
	SetMinInteger(value int)
	Min() float64
	SetMin(value float64)
	Max() float64
	SetMax(value float64)
	StandardDeviation() float64
	SetStandardDeviation(value float64)
	DataType() mps.DataType
	SetDataType(value mps.DataType)
	MaxInteger() int
	SetMaxInteger(value int)
	SamplingMethod() RandomNormalSamplingMethod
	SetSamplingMethod(value RandomNormalSamplingMethod)
	Mean() float64
	SetMean(value float64)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphrandomopdescriptor?language=objc
type RandomOpDescriptor struct {
	objc.Object
}

func RandomOpDescriptorFrom(ptr unsafe.Pointer) RandomOpDescriptor {
	return RandomOpDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (rc _RandomOpDescriptorClass) DescriptorWithDistributionDataType(distribution RandomDistribution, dataType mps.DataType) RandomOpDescriptor {
	rv := objc.Call[RandomOpDescriptor](rc, objc.Sel("descriptorWithDistribution:dataType:"), distribution, dataType)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphrandomopdescriptor/3901499-descriptorwithdistribution?language=objc
func RandomOpDescriptor_DescriptorWithDistributionDataType(distribution RandomDistribution, dataType mps.DataType) RandomOpDescriptor {
	return RandomOpDescriptorClass.DescriptorWithDistributionDataType(distribution, dataType)
}

func (rc _RandomOpDescriptorClass) Alloc() RandomOpDescriptor {
	rv := objc.Call[RandomOpDescriptor](rc, objc.Sel("alloc"))
	return rv
}

func RandomOpDescriptor_Alloc() RandomOpDescriptor {
	return RandomOpDescriptorClass.Alloc()
}

func (rc _RandomOpDescriptorClass) New() RandomOpDescriptor {
	rv := objc.Call[RandomOpDescriptor](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRandomOpDescriptor() RandomOpDescriptor {
	return RandomOpDescriptorClass.New()
}

func (r_ RandomOpDescriptor) Init() RandomOpDescriptor {
	rv := objc.Call[RandomOpDescriptor](r_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphrandomopdescriptor/3901500-distribution?language=objc
func (r_ RandomOpDescriptor) Distribution() RandomDistribution {
	rv := objc.Call[RandomDistribution](r_, objc.Sel("distribution"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphrandomopdescriptor/3901500-distribution?language=objc
func (r_ RandomOpDescriptor) SetDistribution(value RandomDistribution) {
	objc.Call[objc.Void](r_, objc.Sel("setDistribution:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphrandomopdescriptor/3901505-mininteger?language=objc
func (r_ RandomOpDescriptor) MinInteger() int {
	rv := objc.Call[int](r_, objc.Sel("minInteger"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphrandomopdescriptor/3901505-mininteger?language=objc
func (r_ RandomOpDescriptor) SetMinInteger(value int) {
	objc.Call[objc.Void](r_, objc.Sel("setMinInteger:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphrandomopdescriptor/3901504-min?language=objc
func (r_ RandomOpDescriptor) Min() float64 {
	rv := objc.Call[float64](r_, objc.Sel("min"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphrandomopdescriptor/3901504-min?language=objc
func (r_ RandomOpDescriptor) SetMin(value float64) {
	objc.Call[objc.Void](r_, objc.Sel("setMin:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphrandomopdescriptor/3901501-max?language=objc
func (r_ RandomOpDescriptor) Max() float64 {
	rv := objc.Call[float64](r_, objc.Sel("max"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphrandomopdescriptor/3901501-max?language=objc
func (r_ RandomOpDescriptor) SetMax(value float64) {
	objc.Call[objc.Void](r_, objc.Sel("setMax:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphrandomopdescriptor/3901507-standarddeviation?language=objc
func (r_ RandomOpDescriptor) StandardDeviation() float64 {
	rv := objc.Call[float64](r_, objc.Sel("standardDeviation"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphrandomopdescriptor/3901507-standarddeviation?language=objc
func (r_ RandomOpDescriptor) SetStandardDeviation(value float64) {
	objc.Call[objc.Void](r_, objc.Sel("setStandardDeviation:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphrandomopdescriptor/3901498-datatype?language=objc
func (r_ RandomOpDescriptor) DataType() mps.DataType {
	rv := objc.Call[mps.DataType](r_, objc.Sel("dataType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphrandomopdescriptor/3901498-datatype?language=objc
func (r_ RandomOpDescriptor) SetDataType(value mps.DataType) {
	objc.Call[objc.Void](r_, objc.Sel("setDataType:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphrandomopdescriptor/3901502-maxinteger?language=objc
func (r_ RandomOpDescriptor) MaxInteger() int {
	rv := objc.Call[int](r_, objc.Sel("maxInteger"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphrandomopdescriptor/3901502-maxinteger?language=objc
func (r_ RandomOpDescriptor) SetMaxInteger(value int) {
	objc.Call[objc.Void](r_, objc.Sel("setMaxInteger:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphrandomopdescriptor/3901506-samplingmethod?language=objc
func (r_ RandomOpDescriptor) SamplingMethod() RandomNormalSamplingMethod {
	rv := objc.Call[RandomNormalSamplingMethod](r_, objc.Sel("samplingMethod"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphrandomopdescriptor/3901506-samplingmethod?language=objc
func (r_ RandomOpDescriptor) SetSamplingMethod(value RandomNormalSamplingMethod) {
	objc.Call[objc.Void](r_, objc.Sel("setSamplingMethod:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphrandomopdescriptor/3901503-mean?language=objc
func (r_ RandomOpDescriptor) Mean() float64 {
	rv := objc.Call[float64](r_, objc.Sel("mean"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphrandomopdescriptor/3901503-mean?language=objc
func (r_ RandomOpDescriptor) SetMean(value float64) {
	objc.Call[objc.Void](r_, objc.Sel("setMean:"), value)
}