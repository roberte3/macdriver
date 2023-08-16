// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TextTable] class.
var TextTableClass = _TextTableClass{objc.GetClass("NSTextTable")}

type _TextTableClass struct {
	objc.Class
}

// An interface definition for the [TextTable] class.
type ITextTable interface {
	ITextBlock
	BoundsRectForBlockContentRectInRectTextContainerCharacterRange(block ITextTableBlock, contentRect foundation.Rect, rect foundation.Rect, textContainer ITextContainer, charRange foundation.Range) foundation.Rect
	DrawBackgroundForBlockWithFrameInViewCharacterRangeLayoutManager(block ITextTableBlock, frameRect foundation.Rect, controlView IView, charRange foundation.Range, layoutManager ILayoutManager)
	RectForBlockLayoutAtPointInRectTextContainerCharacterRange(block ITextTableBlock, startingPoint foundation.Point, rect foundation.Rect, textContainer ITextContainer, charRange foundation.Range) foundation.Rect
	CollapsesBorders() bool
	SetCollapsesBorders(value bool)
	HidesEmptyCells() bool
	SetHidesEmptyCells(value bool)
	LayoutAlgorithm() TextTableLayoutAlgorithm
	SetLayoutAlgorithm(value TextTableLayoutAlgorithm)
	NumberOfColumns() uint
	SetNumberOfColumns(value uint)
}

// An object that represents a text table as a whole. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstexttable?language=objc
type TextTable struct {
	TextBlock
}

func TextTableFrom(ptr unsafe.Pointer) TextTable {
	return TextTable{
		TextBlock: TextBlockFrom(ptr),
	}
}

func (tc _TextTableClass) Alloc() TextTable {
	rv := objc.Call[TextTable](tc, objc.Sel("alloc"))
	return rv
}

func TextTable_Alloc() TextTable {
	return TextTableClass.Alloc()
}

func (tc _TextTableClass) New() TextTable {
	rv := objc.Call[TextTable](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTextTable() TextTable {
	return TextTableClass.New()
}

func (t_ TextTable) Init() TextTable {
	rv := objc.Call[TextTable](t_, objc.Sel("init"))
	return rv
}

// Returns the rectangle the text table block actually occupies, including padding, borders, and margins. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstexttable/1525956-boundsrectforblock?language=objc
func (t_ TextTable) BoundsRectForBlockContentRectInRectTextContainerCharacterRange(block ITextTableBlock, contentRect foundation.Rect, rect foundation.Rect, textContainer ITextContainer, charRange foundation.Range) foundation.Rect {
	rv := objc.Call[foundation.Rect](t_, objc.Sel("boundsRectForBlock:contentRect:inRect:textContainer:characterRange:"), objc.Ptr(block), contentRect, rect, objc.Ptr(textContainer), charRange)
	return rv
}

// Draws any colors and other decorations for a text table block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstexttable/1534234-drawbackgroundforblock?language=objc
func (t_ TextTable) DrawBackgroundForBlockWithFrameInViewCharacterRangeLayoutManager(block ITextTableBlock, frameRect foundation.Rect, controlView IView, charRange foundation.Range, layoutManager ILayoutManager) {
	objc.Call[objc.Void](t_, objc.Sel("drawBackgroundForBlock:withFrame:inView:characterRange:layoutManager:"), objc.Ptr(block), frameRect, objc.Ptr(controlView), charRange, objc.Ptr(layoutManager))
}

// Returns the rectangle within which glyphs should be laid out for a text table block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstexttable/1534161-rectforblock?language=objc
func (t_ TextTable) RectForBlockLayoutAtPointInRectTextContainerCharacterRange(block ITextTableBlock, startingPoint foundation.Point, rect foundation.Rect, textContainer ITextContainer, charRange foundation.Range) foundation.Rect {
	rv := objc.Call[foundation.Rect](t_, objc.Sel("rectForBlock:layoutAtPoint:inRect:textContainer:characterRange:"), objc.Ptr(block), startingPoint, rect, objc.Ptr(textContainer), charRange)
	return rv
}

// A Boolean value indicating whether the text table borders are collapsible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstexttable/1534170-collapsesborders?language=objc
func (t_ TextTable) CollapsesBorders() bool {
	rv := objc.Call[bool](t_, objc.Sel("collapsesBorders"))
	return rv
}

// A Boolean value indicating whether the text table borders are collapsible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstexttable/1534170-collapsesborders?language=objc
func (t_ TextTable) SetCollapsesBorders(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setCollapsesBorders:"), value)
}

// A Boolean value indicating whether the text table hides empty cells. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstexttable/1526288-hidesemptycells?language=objc
func (t_ TextTable) HidesEmptyCells() bool {
	rv := objc.Call[bool](t_, objc.Sel("hidesEmptyCells"))
	return rv
}

// A Boolean value indicating whether the text table hides empty cells. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstexttable/1526288-hidesemptycells?language=objc
func (t_ TextTable) SetHidesEmptyCells(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setHidesEmptyCells:"), value)
}

// The text table layout algorithm. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstexttable/1531734-layoutalgorithm?language=objc
func (t_ TextTable) LayoutAlgorithm() TextTableLayoutAlgorithm {
	rv := objc.Call[TextTableLayoutAlgorithm](t_, objc.Sel("layoutAlgorithm"))
	return rv
}

// The text table layout algorithm. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstexttable/1531734-layoutalgorithm?language=objc
func (t_ TextTable) SetLayoutAlgorithm(value TextTableLayoutAlgorithm) {
	objc.Call[objc.Void](t_, objc.Sel("setLayoutAlgorithm:"), value)
}

// The number of columns in the text table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstexttable/1532413-numberofcolumns?language=objc
func (t_ TextTable) NumberOfColumns() uint {
	rv := objc.Call[uint](t_, objc.Sel("numberOfColumns"))
	return rv
}

// The number of columns in the text table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstexttable/1532413-numberofcolumns?language=objc
func (t_ TextTable) SetNumberOfColumns(value uint) {
	objc.Call[objc.Void](t_, objc.Sel("setNumberOfColumns:"), value)
}