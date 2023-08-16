// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [LayoutAnchor] class.
var LayoutAnchorClass = _LayoutAnchorClass{objc.GetClass("NSLayoutAnchor")}

type _LayoutAnchorClass struct {
	objc.Class
}

// An interface definition for the [LayoutAnchor] class.
type ILayoutAnchor interface {
	objc.IObject
	ConstraintLessThanOrEqualToAnchor(anchor ILayoutAnchor) LayoutConstraint
	ConstraintEqualToAnchor(anchor ILayoutAnchor) LayoutConstraint
	ConstraintGreaterThanOrEqualToAnchor(anchor ILayoutAnchor) LayoutConstraint
	Item() objc.Object
	Name() string
	ConstraintsAffectingLayout() []LayoutConstraint
	HasAmbiguousLayout() bool
}

// A factory class for creating layout constraint objects using a fluent API. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutanchor?language=objc
type LayoutAnchor struct {
	objc.Object
}

func LayoutAnchorFrom(ptr unsafe.Pointer) LayoutAnchor {
	return LayoutAnchor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (lc _LayoutAnchorClass) Alloc() LayoutAnchor {
	rv := objc.Call[LayoutAnchor](lc, objc.Sel("alloc"))
	return rv
}

func LayoutAnchor_Alloc() LayoutAnchor {
	return LayoutAnchorClass.Alloc()
}

func (lc _LayoutAnchorClass) New() LayoutAnchor {
	rv := objc.Call[LayoutAnchor](lc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewLayoutAnchor() LayoutAnchor {
	return LayoutAnchorClass.New()
}

func (l_ LayoutAnchor) Init() LayoutAnchor {
	rv := objc.Call[LayoutAnchor](l_, objc.Sel("init"))
	return rv
}

// Returns a constraint that defines one item’s attribute as less than or equal to another. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutanchor/1500953-constraintlessthanorequaltoancho?language=objc
func (l_ LayoutAnchor) ConstraintLessThanOrEqualToAnchor(anchor ILayoutAnchor) LayoutConstraint {
	rv := objc.Call[LayoutConstraint](l_, objc.Sel("constraintLessThanOrEqualToAnchor:"), objc.Ptr(anchor))
	return rv
}

// Returns a constraint that defines one item’s attribute as equal to another. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutanchor/1500946-constraintequaltoanchor?language=objc
func (l_ LayoutAnchor) ConstraintEqualToAnchor(anchor ILayoutAnchor) LayoutConstraint {
	rv := objc.Call[LayoutConstraint](l_, objc.Sel("constraintEqualToAnchor:"), objc.Ptr(anchor))
	return rv
}

// Returns a constraint that defines one item’s attribute as greater than or equal to another. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutanchor/1500936-constraintgreaterthanorequaltoan?language=objc
func (l_ LayoutAnchor) ConstraintGreaterThanOrEqualToAnchor(anchor ILayoutAnchor) LayoutConstraint {
	rv := objc.Call[LayoutConstraint](l_, objc.Sel("constraintGreaterThanOrEqualToAnchor:"), objc.Ptr(anchor))
	return rv
}

// The layout item used to calculate the anchor’s position. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutanchor/2870025-item?language=objc
func (l_ LayoutAnchor) Item() objc.Object {
	rv := objc.Call[objc.Object](l_, objc.Sel("item"))
	return rv
}

// The name assigned to the anchor for debugging purposes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutanchor/2870022-name?language=objc
func (l_ LayoutAnchor) Name() string {
	rv := objc.Call[string](l_, objc.Sel("name"))
	return rv
}

// The constraints that impact the layout of the anchor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutanchor/2870023-constraintsaffectinglayout?language=objc
func (l_ LayoutAnchor) ConstraintsAffectingLayout() []LayoutConstraint {
	rv := objc.Call[[]LayoutConstraint](l_, objc.Sel("constraintsAffectingLayout"))
	return rv
}

// A Boolean value indicating whether the constraints impacting the anchor specify its location ambiguously. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutanchor/2870024-hasambiguouslayout?language=objc
func (l_ LayoutAnchor) HasAmbiguousLayout() bool {
	rv := objc.Call[bool](l_, objc.Sel("hasAmbiguousLayout"))
	return rv
}