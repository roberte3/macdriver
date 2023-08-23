// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/corevideo"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [VideoCompositionRenderContext] class.
var VideoCompositionRenderContextClass = _VideoCompositionRenderContextClass{objc.GetClass("AVVideoCompositionRenderContext")}

type _VideoCompositionRenderContextClass struct {
	objc.Class
}

// An interface definition for the [VideoCompositionRenderContext] class.
type IVideoCompositionRenderContext interface {
	objc.IObject
	NewPixelBuffer() corevideo.PixelBufferRef
	HighQualityRendering() bool
	VideoComposition() VideoComposition
	RenderTransform() coregraphics.AffineTransform
	PixelAspectRatio() PixelAspectRatio
	EdgeWidths() EdgeWidths
	RenderScale() float64
	Size() coregraphics.Size
}

// An object that defines the context in which custom compositors render pixel buffers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocompositionrendercontext?language=objc
type VideoCompositionRenderContext struct {
	objc.Object
}

func VideoCompositionRenderContextFrom(ptr unsafe.Pointer) VideoCompositionRenderContext {
	return VideoCompositionRenderContext{
		Object: objc.ObjectFrom(ptr),
	}
}

func (vc _VideoCompositionRenderContextClass) Alloc() VideoCompositionRenderContext {
	rv := objc.Call[VideoCompositionRenderContext](vc, objc.Sel("alloc"))
	return rv
}

func VideoCompositionRenderContext_Alloc() VideoCompositionRenderContext {
	return VideoCompositionRenderContextClass.Alloc()
}

func (vc _VideoCompositionRenderContextClass) New() VideoCompositionRenderContext {
	rv := objc.Call[VideoCompositionRenderContext](vc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewVideoCompositionRenderContext() VideoCompositionRenderContext {
	return VideoCompositionRenderContextClass.New()
}

func (v_ VideoCompositionRenderContext) Init() VideoCompositionRenderContext {
	rv := objc.Call[VideoCompositionRenderContext](v_, objc.Sel("init"))
	return rv
}

// Returns a pixel buffer to use for rendering. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocompositionrendercontext/1386802-newpixelbuffer?language=objc
func (v_ VideoCompositionRenderContext) NewPixelBuffer() corevideo.PixelBufferRef {
	rv := objc.Call[corevideo.PixelBufferRef](v_, objc.Sel("newPixelBuffer"))
	return rv
}

// The rendering quality to use. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocompositionrendercontext/1388758-highqualityrendering?language=objc
func (v_ VideoCompositionRenderContext) HighQualityRendering() bool {
	rv := objc.Call[bool](v_, objc.Sel("highQualityRendering"))
	return rv
}

// The video composition being rendered. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocompositionrendercontext/1390647-videocomposition?language=objc
func (v_ VideoCompositionRenderContext) VideoComposition() VideoComposition {
	rv := objc.Call[VideoComposition](v_, objc.Sel("videoComposition"))
	return rv
}

// A transform to apply to the source image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocompositionrendercontext/1389831-rendertransform?language=objc
func (v_ VideoCompositionRenderContext) RenderTransform() coregraphics.AffineTransform {
	rv := objc.Call[coregraphics.AffineTransform](v_, objc.Sel("renderTransform"))
	return rv
}

// The pixel aspect ratio for rendered frames. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocompositionrendercontext/1389800-pixelaspectratio?language=objc
func (v_ VideoCompositionRenderContext) PixelAspectRatio() PixelAspectRatio {
	rv := objc.Call[PixelAspectRatio](v_, objc.Sel("pixelAspectRatio"))
	return rv
}

// The width of the edge processing region on the left, top, right, and bottom edges, in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocompositionrendercontext/1387026-edgewidths?language=objc
func (v_ VideoCompositionRenderContext) EdgeWidths() EdgeWidths {
	rv := objc.Call[EdgeWidths](v_, objc.Sel("edgeWidths"))
	return rv
}

// A scaling ratio that is applied when rendering frames. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocompositionrendercontext/1387408-renderscale?language=objc
func (v_ VideoCompositionRenderContext) RenderScale() float64 {
	rv := objc.Call[float64](v_, objc.Sel("renderScale"))
	return rv
}

// The width and height for the rendering frames. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocompositionrendercontext/1389718-size?language=objc
func (v_ VideoCompositionRenderContext) Size() coregraphics.Size {
	rv := objc.Call[coregraphics.Size](v_, objc.Sel("size"))
	return rv
}