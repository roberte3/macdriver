// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AccessibilityElement] class.
var AccessibilityElementClass = _AccessibilityElementClass{objc.GetClass("NSAccessibilityElement")}

type _AccessibilityElementClass struct {
	objc.Class
}

// An interface definition for the [AccessibilityElement] class.
type IAccessibilityElement interface {
	objc.IObject
	AccessibilityAddChildElement(childElement IAccessibilityElement)
	AccessibilityParent() objc.Object
	AccessibilityFrame() foundation.Rect
	IsAccessibilityFocused() bool
	AccessibilityIdentifier() string
	AccessibilityFrameInParentSpace() foundation.Rect
	SetAccessibilityFrameInParentSpace(value foundation.Rect)
}

// The basic infrastructure necessary for interacting with an assistive app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityelement?language=objc
type AccessibilityElement struct {
	objc.Object
}

func AccessibilityElementFrom(ptr unsafe.Pointer) AccessibilityElement {
	return AccessibilityElement{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AccessibilityElementClass) Alloc() AccessibilityElement {
	rv := objc.Call[AccessibilityElement](ac, objc.Sel("alloc"))
	return rv
}

func AccessibilityElement_Alloc() AccessibilityElement {
	return AccessibilityElementClass.Alloc()
}

func (ac _AccessibilityElementClass) New() AccessibilityElement {
	rv := objc.Call[AccessibilityElement](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAccessibilityElement() AccessibilityElement {
	return AccessibilityElementClass.New()
}

func (a_ AccessibilityElement) Init() AccessibilityElement {
	rv := objc.Call[AccessibilityElement](a_, objc.Sel("init"))
	return rv
}

// Adds a child to the accessibility element in the accessibility hierarchy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityelement/1533717-accessibilityaddchildelement?language=objc
func (a_ AccessibilityElement) AccessibilityAddChildElement(childElement IAccessibilityElement) {
	objc.Call[objc.Void](a_, objc.Sel("accessibilityAddChildElement:"), objc.Ptr(childElement))
}

// Returns the accessibility element’s parent in the accessibility hierarchy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/1534023-nsaccessibilityelement/1529078-accessibilityparent?language=objc
func (a_ AccessibilityElement) AccessibilityParent() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityParent"))
	return rv
}

// Returns the accessibility element’s frame in screen coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/1534023-nsaccessibilityelement/1528055-accessibilityframe?language=objc
func (a_ AccessibilityElement) AccessibilityFrame() foundation.Rect {
	rv := objc.Call[foundation.Rect](a_, objc.Sel("accessibilityFrame"))
	return rv
}

// Returns a Boolean value that indicates whether the accessibility element has the keyboard focus. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/1534023-nsaccessibilityelement/1525133-isaccessibilityfocused?language=objc
func (a_ AccessibilityElement) IsAccessibilityFocused() bool {
	rv := objc.Call[bool](a_, objc.Sel("isAccessibilityFocused"))
	return rv
}

// Instantiates and configures a new accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityelement/1531178-accessibilityelementwithrole?language=objc
func (ac _AccessibilityElementClass) AccessibilityElementWithRoleFrameLabelParent(role AccessibilityRole, frame foundation.Rect, label string, parent objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](ac, objc.Sel("accessibilityElementWithRole:frame:label:parent:"), role, frame, label, parent)
	return rv
}

// Instantiates and configures a new accessibility element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityelement/1531178-accessibilityelementwithrole?language=objc
func AccessibilityElement_AccessibilityElementWithRoleFrameLabelParent(role AccessibilityRole, frame foundation.Rect, label string, parent objc.IObject) objc.Object {
	return AccessibilityElementClass.AccessibilityElementWithRoleFrameLabelParent(role, frame, label, parent)
}

// Returns the accessibility element’s identity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/1534023-nsaccessibilityelement/1533707-accessibilityidentifier?language=objc
func (a_ AccessibilityElement) AccessibilityIdentifier() string {
	rv := objc.Call[string](a_, objc.Sel("accessibilityIdentifier"))
	return rv
}

// The accessibility element’s frame in its parent’s coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityelement/1569648-accessibilityframeinparentspace?language=objc
func (a_ AccessibilityElement) AccessibilityFrameInParentSpace() foundation.Rect {
	rv := objc.Call[foundation.Rect](a_, objc.Sel("accessibilityFrameInParentSpace"))
	return rv
}

// The accessibility element’s frame in its parent’s coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityelement/1569648-accessibilityframeinparentspace?language=objc
func (a_ AccessibilityElement) SetAccessibilityFrameInParentSpace(value foundation.Rect) {
	objc.Call[objc.Void](a_, objc.Sel("setAccessibilityFrameInParentSpace:"), value)
}