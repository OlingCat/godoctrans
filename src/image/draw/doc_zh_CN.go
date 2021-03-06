// Copyright The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

// Package draw provides image composition functions.
//
// See "The Go image/draw package" for an introduction to this package:
// http://golang.org/doc/articles/image_draw.html

// draw 包提供组装图片的方法.
//
// 参考 "The Go image/draw package" 获取这个包的简介：
// http://golang.org/doc/articles/image_draw.html
package draw

// Draw calls DrawMask with a nil mask.
func Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point, op Op)

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the
// rectangle r in dst with the result of a Porter-Duff composition. A nil mask is
// treated as opaque.

// DrawMask将dst上的r.Min，src上的sp，mask上的mp对齐，然后对dst上的r型矩阵区域执行Porter-Duff合并操作。
// mask设置为nil就代表完全不透明。
func DrawMask(dst Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, op Op)

// Drawer contains the Draw method.

// Drawer 包含 Draw 方法。
type Drawer interface {
	// Draw aligns r.Min in dst with sp in src and then replaces the
	// rectangle r in dst with the result of drawing src on dst.
	Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point)
}

// FloydSteinberg is a Drawer that is the Src Op with Floyd-Steinberg error
// diffusion.

// FloydSteinberg 是一个 Drawer，它对 Src 进行 Floyd-Steinberg 误差扩散操作。
var FloydSteinberg Drawer = floydSteinberg{}

// Image is an image.Image with a Set method to change a single pixel.
type Image interface {
	image.Image
	Set(x, y int, c color.Color)
}

// Op is a Porter-Duff compositing operator.

// Op是Porter-Duff合成操作。
type Op int

const (
	// Over specifies ``(src in mask) over dst''.
	Over Op = iota
	// Src specifies ``src in mask''.
	Src
)

// Draw implements the Drawer interface by calling the Draw function with this Op.

// Draw 通过用此 Op 调用 Draw 函数实现了 Drawer 接口。
func (op Op) Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point)

// Quantizer produces a palette for an image.
type Quantizer interface {
	// Quantize appends up to cap(p) - len(p) colors to p and returns the
	// updated palette suitable for converting m to a paletted image.
	Quantize(p color.Palette, m image.Image) color.Palette
}
