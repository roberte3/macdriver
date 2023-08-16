// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [GridView] class.
var GridViewClass = _GridViewClass{objc.GetClass("NSGridView")}

type _GridViewClass struct {
	objc.Class
}

// An interface definition for the [GridView] class.
type IGridView interface {
	IView
	IndexOfColumn(column IGridColumn) int
	RemoveColumnAtIndex(index int)
	IndexOfRow(row IGridRow) int
	InsertColumnAtIndexWithViews(index int, views []IView) GridColumn
	ColumnAtIndex(index int) GridColumn
	MergeCellsInHorizontalRangeVerticalRange(hRange foundation.Range, vRange foundation.Range)
	AddColumnWithViews(views []IView) GridColumn
	AddRowWithViews(views []IView) GridRow
	CellForView(view IView) GridCell
	InsertRowAtIndexWithViews(index int, views []IView) GridRow
	MoveRowAtIndexToIndex(fromIndex int, toIndex int)
	CellAtColumnIndexRowIndex(columnIndex int, rowIndex int) GridCell
	MoveColumnAtIndexToIndex(fromIndex int, toIndex int)
	RemoveRowAtIndex(index int)
	RowAtIndex(index int) GridRow
	RowAlignment() GridRowAlignment
	SetRowAlignment(value GridRowAlignment)
	ColumnSpacing() float64
	SetColumnSpacing(value float64)
	YPlacement() GridCellPlacement
	SetYPlacement(value GridCellPlacement)
	XPlacement() GridCellPlacement
	SetXPlacement(value GridCellPlacement)
	RowSpacing() float64
	SetRowSpacing(value float64)
	NumberOfRows() int
	NumberOfColumns() int
}

// A container that aligns views in a flexible grid of rows and columns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridview?language=objc
type GridView struct {
	View
}

func GridViewFrom(ptr unsafe.Pointer) GridView {
	return GridView{
		View: ViewFrom(ptr),
	}
}

func (gc _GridViewClass) GridViewWithNumberOfColumnsRows(columnCount int, rowCount int) GridView {
	rv := objc.Call[GridView](gc, objc.Sel("gridViewWithNumberOfColumns:rows:"), columnCount, rowCount)
	return rv
}

// Creates a newly allocated grid view object with the specified number of columns and rows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridview/1639714-gridviewwithnumberofcolumns?language=objc
func GridView_GridViewWithNumberOfColumnsRows(columnCount int, rowCount int) GridView {
	return GridViewClass.GridViewWithNumberOfColumnsRows(columnCount, rowCount)
}

func (g_ GridView) InitWithFrame(frameRect foundation.Rect) GridView {
	rv := objc.Call[GridView](g_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Creates a newly allocated grid view object with the specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridview/1639692-initwithframe?language=objc
func GridView_InitWithFrame(frameRect foundation.Rect) GridView {
	return GridViewClass.Alloc().InitWithFrame(frameRect)
}

func (gc _GridViewClass) GridViewWithViews(rows [][]IView) GridView {
	rv := objc.Call[GridView](gc, objc.Sel("gridViewWithViews:"), rows)
	return rv
}

// Creates a newly allocated grid view object with the specified array of arrays of views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridview/1639782-gridviewwithviews?language=objc
func GridView_GridViewWithViews(rows [][]IView) GridView {
	return GridViewClass.GridViewWithViews(rows)
}

func (gc _GridViewClass) Alloc() GridView {
	rv := objc.Call[GridView](gc, objc.Sel("alloc"))
	return rv
}

func GridView_Alloc() GridView {
	return GridViewClass.Alloc()
}

func (gc _GridViewClass) New() GridView {
	rv := objc.Call[GridView](gc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewGridView() GridView {
	return GridViewClass.New()
}

func (g_ GridView) Init() GridView {
	rv := objc.Call[GridView](g_, objc.Sel("init"))
	return rv
}

// Returns the index of the specified grid column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridview/1639667-indexofcolumn?language=objc
func (g_ GridView) IndexOfColumn(column IGridColumn) int {
	rv := objc.Call[int](g_, objc.Sel("indexOfColumn:"), objc.Ptr(column))
	return rv
}

// Removes the column from the grid view at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridview/1639759-removecolumnatindex?language=objc
func (g_ GridView) RemoveColumnAtIndex(index int) {
	objc.Call[objc.Void](g_, objc.Sel("removeColumnAtIndex:"), index)
}

// Returns the index of the specified grid row. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridview/1639661-indexofrow?language=objc
func (g_ GridView) IndexOfRow(row IGridRow) int {
	rv := objc.Call[int](g_, objc.Sel("indexOfRow:"), objc.Ptr(row))
	return rv
}

// Inserts the array of view objects at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridview/1639700-insertcolumnatindex?language=objc
func (g_ GridView) InsertColumnAtIndexWithViews(index int, views []IView) GridColumn {
	rv := objc.Call[GridColumn](g_, objc.Sel("insertColumnAtIndex:withViews:"), index, views)
	return rv
}

// Returns the grid column object at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridview/1639674-columnatindex?language=objc
func (g_ GridView) ColumnAtIndex(index int) GridColumn {
	rv := objc.Call[GridColumn](g_, objc.Sel("columnAtIndex:"), index)
	return rv
}

// Expands the cell at the top-leading corner of the horizontal and vertical range to cover the entire area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridview/1639749-mergecellsinhorizontalrange?language=objc
func (g_ GridView) MergeCellsInHorizontalRangeVerticalRange(hRange foundation.Range, vRange foundation.Range) {
	objc.Call[objc.Void](g_, objc.Sel("mergeCellsInHorizontalRange:verticalRange:"), hRange, vRange)
}

// Adds a new column containing the array of views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridview/1639783-addcolumnwithviews?language=objc
func (g_ GridView) AddColumnWithViews(views []IView) GridColumn {
	rv := objc.Call[GridColumn](g_, objc.Sel("addColumnWithViews:"), views)
	return rv
}

// Adds an array of views to a new row. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridview/1639690-addrowwithviews?language=objc
func (g_ GridView) AddRowWithViews(views []IView) GridRow {
	rv := objc.Call[GridRow](g_, objc.Sel("addRowWithViews:"), views)
	return rv
}

// Returns the grid cell object that contains the given view or one of its ancestors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridview/1639703-cellforview?language=objc
func (g_ GridView) CellForView(view IView) GridCell {
	rv := objc.Call[GridCell](g_, objc.Sel("cellForView:"), objc.Ptr(view))
	return rv
}

// Inserts the array of view objects into the grid view at the index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridview/1639787-insertrowatindex?language=objc
func (g_ GridView) InsertRowAtIndexWithViews(index int, views []IView) GridRow {
	rv := objc.Call[GridRow](g_, objc.Sel("insertRowAtIndex:withViews:"), index, views)
	return rv
}

// Moves the specified row to the new row location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridview/1639687-moverowatindex?language=objc
func (g_ GridView) MoveRowAtIndexToIndex(fromIndex int, toIndex int) {
	objc.Call[objc.Void](g_, objc.Sel("moveRowAtIndex:toIndex:"), fromIndex, toIndex)
}

// Returns the grid cell object at the specified column and row index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridview/1639778-cellatcolumnindex?language=objc
func (g_ GridView) CellAtColumnIndexRowIndex(columnIndex int, rowIndex int) GridCell {
	rv := objc.Call[GridCell](g_, objc.Sel("cellAtColumnIndex:rowIndex:"), columnIndex, rowIndex)
	return rv
}

// Moves the specified column to a new column location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridview/1639659-movecolumnatindex?language=objc
func (g_ GridView) MoveColumnAtIndexToIndex(fromIndex int, toIndex int) {
	objc.Call[objc.Void](g_, objc.Sel("moveColumnAtIndex:toIndex:"), fromIndex, toIndex)
}

// Removes the row from the grid view at the index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridview/1639771-removerowatindex?language=objc
func (g_ GridView) RemoveRowAtIndex(index int) {
	objc.Call[objc.Void](g_, objc.Sel("removeRowAtIndex:"), index)
}

// Returns the grid row object at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridview/1639761-rowatindex?language=objc
func (g_ GridView) RowAtIndex(index int) GridRow {
	rv := objc.Call[GridRow](g_, objc.Sel("rowAtIndex:"), index)
	return rv
}

// The row alignment for the grid view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridview/1823691-rowalignment?language=objc
func (g_ GridView) RowAlignment() GridRowAlignment {
	rv := objc.Call[GridRowAlignment](g_, objc.Sel("rowAlignment"))
	return rv
}

// The row alignment for the grid view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridview/1823691-rowalignment?language=objc
func (g_ GridView) SetRowAlignment(value GridRowAlignment) {
	objc.Call[objc.Void](g_, objc.Sel("setRowAlignment:"), value)
}

// The column spacing for the grid view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridview/1639776-columnspacing?language=objc
func (g_ GridView) ColumnSpacing() float64 {
	rv := objc.Call[float64](g_, objc.Sel("columnSpacing"))
	return rv
}

// The column spacing for the grid view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridview/1639776-columnspacing?language=objc
func (g_ GridView) SetColumnSpacing(value float64) {
	objc.Call[objc.Void](g_, objc.Sel("setColumnSpacing:"), value)
}

// The placement of the cell within the grid row. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridview/1639767-yplacement?language=objc
func (g_ GridView) YPlacement() GridCellPlacement {
	rv := objc.Call[GridCellPlacement](g_, objc.Sel("yPlacement"))
	return rv
}

// The placement of the cell within the grid row. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridview/1639767-yplacement?language=objc
func (g_ GridView) SetYPlacement(value GridCellPlacement) {
	objc.Call[objc.Void](g_, objc.Sel("setYPlacement:"), value)
}

// The placement of the cell within the grid column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridview/1639732-xplacement?language=objc
func (g_ GridView) XPlacement() GridCellPlacement {
	rv := objc.Call[GridCellPlacement](g_, objc.Sel("xPlacement"))
	return rv
}

// The placement of the cell within the grid column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridview/1639732-xplacement?language=objc
func (g_ GridView) SetXPlacement(value GridCellPlacement) {
	objc.Call[objc.Void](g_, objc.Sel("setXPlacement:"), value)
}

// The row spacing for the grid view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridview/1639730-rowspacing?language=objc
func (g_ GridView) RowSpacing() float64 {
	rv := objc.Call[float64](g_, objc.Sel("rowSpacing"))
	return rv
}

// The row spacing for the grid view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridview/1639730-rowspacing?language=objc
func (g_ GridView) SetRowSpacing(value float64) {
	objc.Call[objc.Void](g_, objc.Sel("setRowSpacing:"), value)
}

// The number of rows in the grid view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridview/1639705-numberofrows?language=objc
func (g_ GridView) NumberOfRows() int {
	rv := objc.Call[int](g_, objc.Sel("numberOfRows"))
	return rv
}

// The number of columns in the grid view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgridview/1639698-numberofcolumns?language=objc
func (g_ GridView) NumberOfColumns() int {
	rv := objc.Call[int](g_, objc.Sel("numberOfColumns"))
	return rv
}