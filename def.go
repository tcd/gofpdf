/*
 * Copyright (c) 2013-2014 Kurt Jung (Gmail: kurt.w.jung)
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package gofpdf

import (
	"bytes"
	"crypto/sha1"
	"encoding/gob"
	"fmt"
	"io"
	"time"

	"github.com/jung-kurt/gofpdf/internal/font"
)

// Version of FPDF from which this package is derived
const (
	cnFpdfVersion = "1.7"
)

type blendModeType struct {
	strokeStr, fillStr, modeStr string
	objNum                      int
}

type gradientType struct {
	tp                int // 2: linear, 3: radial
	clr1Str, clr2Str  string
	x1, y1, x2, y2, r float64
	objNum            int
}

const (
	// OrientationPortrait represents the portrait orientation.
	OrientationPortrait = "portrait"

	// OrientationLandscape represents the landscape orientation.
	OrientationLandscape = "landscape"
)

const (
	// UnitPoint represents the size unit point
	UnitPoint = "pt"
	// UnitMillimeter represents the size unit millimeter
	UnitMillimeter = "mm"
	// UnitCentimeter represents the size unit centimeter
	UnitCentimeter = "cm"
	// UnitInch represents the size unit inch
	UnitInch = "inch"
)

const (
	// PageSizeA3 represents DIN/ISO A3 page size
	PageSizeA3 = "A3"
	// PageSizeA4 represents DIN/ISO A4 page size
	PageSizeA4 = "A4"
	// PageSizeA5 represents DIN/ISO A5 page size
	PageSizeA5 = "A5"
	// PageSizeLetter represents US Letter page size
	PageSizeLetter = "Letter"
	// PageSizeLegal represents US Legal page size
	PageSizeLegal = "Legal"
)

const (
	// BorderNone set no border
	BorderNone = ""
	// BorderFull sets a full border
	BorderFull = "1"
	// BorderLeft sets the border on the left side
	BorderLeft = "L"
	// BorderTop sets the border at the top
	BorderTop = "T"
	// BorderRight sets the border on the right side
	BorderRight = "R"
	// BorderBottom sets the border on the bottom
	BorderBottom = "B"
)

const (
	// LineBreakNone disables linebreak
	LineBreakNone = 0
	// LineBreakNormal enables normal linebreak
	LineBreakNormal = 1
	// LineBreakBelow enables linebreak below
	LineBreakBelow = 2
)

const (
	// AlignLeft left aligns the cell
	AlignLeft = "L"
	// AlignRight right aligns the cell
	AlignRight = "R"
	// AlignCenter centers the cell
	AlignCenter = "C"
	// AlignTop aligns the cell to the top
	AlignTop = "T"
	// AlignBottom aligns the cell to the bottom
	AlignBottom = "B"
	// AlignMiddle aligns the cell to the middle
	AlignMiddle = "M"
	// AlignBaseline aligns the cell to the baseline
	AlignBaseline = "B"
)

type colorMode int

const (
	colorModeRGB colorMode = iota
	colorModeSpot
	colorModeCMYK
)

type colorType struct {
	r, g, b    float64
	ir, ig, ib int
	mode       colorMode
	spotStr    string // name of current spot color
	gray       bool
	str        string
}

// SpotColorType specifies a named spot color value
type spotColorType struct {
	id, objID int
	val       cmykColorType
}

// CMYKColorType specifies an ink-based CMYK color value
type cmykColorType struct {
	c, m, y, k byte // 0% to 100%
}

// SizeType fields Wd and Ht specify the horizontal and vertical extents of a
// document element such as a page.
type SizeType struct {
	Wd, Ht float64
}

// PointType fields X and Y specify the horizontal and vertical coordinates of
// a point, typically used in drawing.
type PointType struct {
	X, Y float64
}

// XY returns the X and Y components of the receiver point.
func (p PointType) XY() (float64, float64) {
	return p.X, p.Y
}

// ImageInfoType contains size, color and other information about an image.
// Changes to this structure should be reflected in its GobEncode and GobDecode
// methods.
type ImageInfoType struct {
	data  []byte
	smask []byte
	n     int
	w     float64
	h     float64
	cs    string
	pal   []byte
	bpc   int
	f     string
	dp    string
	trns  []int
	scale float64 // document scaling factor
	dpi   float64
	i     string
}

func generateImageID(info *ImageInfoType) (string, error) {
	b, err := info.GobEncode()
	return fmt.Sprintf("%x", sha1.Sum(b)), err
}

// GobEncode encodes the receiving image to a byte slice.
func (info *ImageInfoType) GobEncode() (buf []byte, err error) {
	fields := []interface{}{info.data, info.smask, info.n, info.w, info.h, info.cs,
		info.pal, info.bpc, info.f, info.dp, info.trns, info.scale, info.dpi}
	w := new(bytes.Buffer)
	encoder := gob.NewEncoder(w)
	for j := 0; j < len(fields) && err == nil; j++ {
		err = encoder.Encode(fields[j])
	}
	if err == nil {
		buf = w.Bytes()
	}
	return
}

// GobDecode decodes the specified byte buffer (generated by GobEncode) into
// the receiving image.
func (info *ImageInfoType) GobDecode(buf []byte) (err error) {
	fields := []interface{}{&info.data, &info.smask, &info.n, &info.w, &info.h,
		&info.cs, &info.pal, &info.bpc, &info.f, &info.dp, &info.trns, &info.scale, &info.dpi}
	r := bytes.NewBuffer(buf)
	decoder := gob.NewDecoder(r)
	for j := 0; j < len(fields) && err == nil; j++ {
		err = decoder.Decode(fields[j])
	}

	info.i, err = generateImageID(info)
	return
}

// PointConvert returns the value of pt, expressed in points (1/72 inch), as a
// value expressed in the unit of measure specified in New(). Since font
// management in Fpdf uses points, this method can help with line height
// calculations and other methods that require user units.
func (f *Fpdf) PointConvert(pt float64) (u float64) {
	return pt / f.k
}

// PointToUnitConvert is an alias for PointConvert.
func (f *Fpdf) PointToUnitConvert(pt float64) (u float64) {
	return pt / f.k
}

// UnitToPointConvert returns the value of u, expressed in the unit of measure
// specified in New(), as a value expressed in points (1/72 inch). Since font
// management in Fpdf uses points, this method can help with setting font sizes
// based on the sizes of other non-font page elements.
func (f *Fpdf) UnitToPointConvert(u float64) (pt float64) {
	return u * f.k
}

// Extent returns the width and height of the image in the units of the Fpdf
// object.
func (info *ImageInfoType) Extent() (wd, ht float64) {
	return info.Width(), info.Height()
}

// Width returns the width of the image in the units of the Fpdf object.
func (info *ImageInfoType) Width() float64 {
	return info.w / (info.scale * info.dpi / 72)
}

// Height returns the height of the image in the units of the Fpdf object.
func (info *ImageInfoType) Height() float64 {
	return info.h / (info.scale * info.dpi / 72)
}

// SetDpi sets the dots per inch for an image. PNG images MAY have their dpi
// set automatically, if the image specifies it. DPI information is not
// currently available automatically for JPG and GIF images, so if it's
// important to you, you can set it here. It defaults to 72 dpi.
func (info *ImageInfoType) SetDpi(dpi float64) {
	info.dpi = dpi
}

type linkType struct {
	x, y, wd, ht float64
	link         int    // Auto-generated internal link ID or...
	linkStr      string // ...application-provided external link string
}

type intLinkType struct {
	page int
	y    float64
}

// outlineType is used for a sidebar outline of bookmarks
type outlineType struct {
	text                                   string
	level, parent, first, last, next, prev int
	y                                      float64
	p                                      int
}

// InitType is used with NewCustom() to customize an Fpdf instance.
// OrientationStr, UnitStr, SizeStr and FontDirStr correspond to the arguments
// accepted by New(). If the Wd and Ht fields of Size are each greater than
// zero, Size will be used to set the default page size rather than SizeStr. Wd
// and Ht are specified in the units of measure indicated by UnitStr.
type InitType struct {
	OrientationStr string
	UnitStr        string
	SizeStr        string
	Size           SizeType
	FontDirStr     string
}

// FontLoader is used to read fonts (JSON font specification and zlib compressed font binaries)
// from arbitrary locations (e.g. files, zip files, embedded font resources).
//
// Open provides an io.Reader for the specified font file (.json or .z). The file name
// never includes a path. Open returns an error if the specified file cannot be opened.
type FontLoader interface {
	Open(name string) (io.Reader, error)
}

// Pdf defines the interface used for various methods. It is implemented by the
// main FPDF instance as well as templates.
type Pdf interface {
	AddFont(familyStr, styleStr, fileStr string)
	AddFontFromBytes(familyStr, styleStr string, jsonFileBytes, zFileBytes []byte)
	AddFontFromReader(familyStr, styleStr string, r io.Reader)
	AddLayer(name string, visible bool) (layerID int)
	AddLink() int
	AddPage()
	AddPageFormat(orientationStr string, size SizeType)
	AddSpotColor(nameStr string, c, m, y, k byte)
	AliasNbPages(aliasStr string)
	ArcTo(x, y, rx, ry, degRotate, degStart, degEnd float64)
	Arc(x, y, rx, ry, degRotate, degStart, degEnd float64, styleStr string)
	BeginLayer(id int)
	Beziergon(points []PointType, styleStr string)
	Bookmark(txtStr string, level int, y float64)
	CellFormat(w, h float64, txtStr, borderStr string, ln int, alignStr string, fill bool, link int, linkStr string)
	Cellf(w, h float64, fmtStr string, args ...interface{})
	Cell(w, h float64, txtStr string)
	Circle(x, y, r float64, styleStr string)
	ClearError()
	ClipCircle(x, y, r float64, outline bool)
	ClipEllipse(x, y, rx, ry float64, outline bool)
	ClipEnd()
	ClipPolygon(points []PointType, outline bool)
	ClipRect(x, y, w, h float64, outline bool)
	ClipRoundedRect(x, y, w, h, r float64, outline bool)
	ClipText(x, y float64, txtStr string, outline bool)
	Close()
	ClosePath()
	CreateTemplateCustom(corner PointType, size SizeType, fn func(*Tpl)) Template
	CreateTemplate(fn func(*Tpl)) Template
	CurveBezierCubicTo(cx0, cy0, cx1, cy1, x, y float64)
	CurveBezierCubic(x0, y0, cx0, cy0, cx1, cy1, x1, y1 float64, styleStr string)
	CurveCubic(x0, y0, cx0, cy0, x1, y1, cx1, cy1 float64, styleStr string)
	CurveTo(cx, cy, x, y float64)
	Curve(x0, y0, cx, cy, x1, y1 float64, styleStr string)
	DrawPath(styleStr string)
	Ellipse(x, y, rx, ry, degRotate float64, styleStr string)
	EndLayer()
	Err() bool
	Error() error
	GetAlpha() (alpha float64, blendModeStr string)
	GetAutoPageBreak() (auto bool, margin float64)
	GetCellMargin() float64
	GetConversionRatio() float64
	GetDrawColor() (int, int, int)
	GetDrawSpotColor() (name string, c, m, y, k byte)
	GetFillColor() (int, int, int)
	GetFillSpotColor() (name string, c, m, y, k byte)
	GetFontDesc(familyStr, styleStr string) font.FontDescType
	GetFontSize() (ptSize, unitSize float64)
	GetImageInfo(imageStr string) (info *ImageInfoType)
	GetLineWidth() float64
	GetMargins() (left, top, right, bottom float64)
	GetPageSizeStr(sizeStr string) (size SizeType)
	GetPageSize() (width, height float64)
	GetStringWidth(s string) float64
	GetTextColor() (int, int, int)
	GetTextSpotColor() (name string, c, m, y, k byte)
	GetX() float64
	GetXY() (float64, float64)
	GetY() float64
	HTMLBasicNew() (html HTMLBasicType)
	Image(imageNameStr string, x, y, w, h float64, flow bool, tp string, link int, linkStr string)
	ImageOptions(imageNameStr string, x, y, w, h float64, flow bool, options ImageOptions, link int, linkStr string)
	ImageTypeFromMime(mimeStr string) (tp string)
	LinearGradient(x, y, w, h float64, r1, g1, b1, r2, g2, b2 int, x1, y1, x2, y2 float64)
	LineTo(x, y float64)
	Line(x1, y1, x2, y2 float64)
	LinkString(x, y, w, h float64, linkStr string)
	Link(x, y, w, h float64, link int)
	Ln(h float64)
	MoveTo(x, y float64)
	MultiCell(w, h float64, txtStr, borderStr, alignStr string, fill bool)
	Ok() bool
	OpenLayerPane()
	OutputAndClose(w io.WriteCloser) error
	OutputFileAndClose(fileStr string) error
	Output(w io.Writer) error
	PageCount() int
	PageNo() int
	PageSize(pageNum int) (wd, ht float64, unitStr string)
	PointConvert(pt float64) (u float64)
	PointToUnitConvert(pt float64) (u float64)
	Polygon(points []PointType, styleStr string)
	RadialGradient(x, y, w, h float64, r1, g1, b1, r2, g2, b2 int, x1, y1, x2, y2, r float64)
	RawWriteBuf(r io.Reader)
	RawWriteStr(str string)
	Rect(x, y, w, h float64, styleStr string)
	RegisterAlias(alias, replacement string)
	RegisterImage(fileStr, tp string) (info *ImageInfoType)
	RegisterImageOptions(fileStr string, options ImageOptions) (info *ImageInfoType)
	RegisterImageOptionsReader(imgName string, options ImageOptions, r io.Reader) (info *ImageInfoType)
	RegisterImageReader(imgName, tp string, r io.Reader) (info *ImageInfoType)
	SetAcceptPageBreakFunc(fnc func() bool)
	SetAlpha(alpha float64, blendModeStr string)
	SetAuthor(authorStr string, isUTF8 bool)
	SetAutoPageBreak(auto bool, margin float64)
	SetCatalogSort(flag bool)
	SetCellMargin(margin float64)
	SetCompression(compress bool)
	SetCreationDate(tm time.Time)
	SetCreator(creatorStr string, isUTF8 bool)
	SetDashPattern(dashArray []float64, dashPhase float64)
	SetDisplayMode(zoomStr, layoutStr string)
	SetDrawColor(r, g, b int)
	SetDrawSpotColor(nameStr string, tint byte)
	SetError(err error)
	SetErrorf(fmtStr string, args ...interface{})
	SetFillColor(r, g, b int)
	SetFillSpotColor(nameStr string, tint byte)
	SetFont(familyStr, styleStr string, size float64)
	SetFontLoader(loader FontLoader)
	SetFontLocation(fontDirStr string)
	SetFontSize(size float64)
	SetFontStyle(styleStr string)
	SetFontUnitSize(size float64)
	SetFooterFunc(fnc func())
	SetFooterFuncLpi(fnc func(lastPage bool))
	SetHeaderFunc(fnc func())
	SetHeaderFuncMode(fnc func(), homeMode bool)
	SetHomeXY()
	SetJavascript(script string)
	SetKeywords(keywordsStr string, isUTF8 bool)
	SetLeftMargin(margin float64)
	SetLineCapStyle(styleStr string)
	SetLineJoinStyle(styleStr string)
	SetLineWidth(width float64)
	SetLink(link int, y float64, page int)
	SetMargins(left, top, right float64)
	SetPageBoxRec(t string, pb PageBox)
	SetPageBox(t string, x, y, wd, ht float64)
	SetPage(pageNum int)
	SetProtection(actionFlag byte, userPassStr, ownerPassStr string)
	SetRightMargin(margin float64)
	SetSubject(subjectStr string, isUTF8 bool)
	SetTextColor(r, g, b int)
	SetTextSpotColor(nameStr string, tint byte)
	SetTitle(titleStr string, isUTF8 bool)
	SetTopMargin(margin float64)
	SetUnderlineThickness(thickness float64)
	SetXmpMetadata(xmpStream []byte)
	SetX(x float64)
	SetXY(x, y float64)
	SetY(y float64)
	SplitLines(txt []byte, w float64) [][]byte
	String() string
	SVGBasicWrite(sb *SVGBasicType, scale float64)
	Text(x, y float64, txtStr string)
	TransformBegin()
	TransformEnd()
	TransformMirrorHorizontal(x float64)
	TransformMirrorLine(angle, x, y float64)
	TransformMirrorPoint(x, y float64)
	TransformMirrorVertical(y float64)
	TransformRotate(angle, x, y float64)
	TransformScale(scaleWd, scaleHt, x, y float64)
	TransformScaleX(scaleWd, x, y float64)
	TransformScaleXY(s, x, y float64)
	TransformScaleY(scaleHt, x, y float64)
	TransformSkew(angleX, angleY, x, y float64)
	TransformSkewX(angleX, x, y float64)
	TransformSkewY(angleY, x, y float64)
	Transform(tm TransformMatrix)
	TransformTranslate(tx, ty float64)
	TransformTranslateX(tx float64)
	TransformTranslateY(ty float64)
	UnicodeTranslatorFromDescriptor(cpStr string) (rep func(string) string)
	UnitToPointConvert(u float64) (pt float64)
	UseTemplateScaled(t Template, corner PointType, size SizeType)
	UseTemplate(t Template)
	WriteAligned(width, lineHeight float64, textStr, alignStr string)
	Writef(h float64, fmtStr string, args ...interface{})
	Write(h float64, txtStr string)
	WriteLinkID(h float64, displayStr string, linkID int)
	WriteLinkString(h float64, displayStr, targetStr string)
}

// PageBox defines the coordinates and extent of the various page box types
type PageBox struct {
	SizeType
	PointType
}

// Fpdf is the principal structure for creating a single PDF document
type Fpdf struct {
	isCurrentUTF8    bool                         // is current font used in utf-8 mode
	isRTL            bool                         // is is right to left mode enabled
	page             int                          // current page number
	n                int                          // current object number
	offsets          []int                        // array of object offsets
	templates        map[string]Template          // templates used in this document
	templateObjects  map[string]int               // template object IDs within this document
	importedObjs     map[string][]byte            // imported template objects (gofpdi)
	importedObjPos   map[string]map[int]string    // imported template objects hashes and their positions (gofpdi)
	importedTplObjs  map[string]string            // imported template names and IDs (hashed) (gofpdi)
	importedTplIDs   map[string]int               // imported template ids hash to object id int (gofpdi)
	buffer           fmtBuffer                    // buffer holding in-memory PDF
	pages            []*bytes.Buffer              // slice[page] of page content; 1-based
	state            int                          // current document state
	compress         bool                         // compression flag
	k                float64                      // scale factor (number of points in user unit)
	defOrientation   string                       // default orientation
	curOrientation   string                       // current orientation
	stdPageSizes     map[string]SizeType          // standard page sizes
	defPageSize      SizeType                     // default page size
	defPageBoxes     map[string]PageBox           // default page size
	curPageSize      SizeType                     // current page size
	pageSizes        map[int]SizeType             // used for pages with non default sizes or orientations
	pageBoxes        map[int]map[string]PageBox   // used to define the crop, trim, bleed and art boxes
	unitStr          string                       // unit of measure for all rendered objects except fonts
	wPt, hPt         float64                      // dimensions of current page in points
	w, h             float64                      // dimensions of current page in user unit
	lMargin          float64                      // left margin
	tMargin          float64                      // top margin
	rMargin          float64                      // right margin
	bMargin          float64                      // page break margin
	cMargin          float64                      // cell margin
	x, y             float64                      // current position in user unit
	lasth            float64                      // height of last printed cell
	lineWidth        float64                      // line width in user unit
	fontpath         string                       // path containing fonts
	fontLoader       FontLoader                   // used to load font files from arbitrary locations
	coreFonts        map[string]bool              // array of core font names
	fonts            map[string]font.FontDefType  // array of used fonts
	fontFiles        map[string]font.FontFileType // array of font files
	diffs            []string                     // array of encoding differences
	fontFamily       string                       // current font family
	fontStyle        string                       // current font style
	underline        bool                         // underlining flag
	currentFont      font.FontDefType             // current font info
	fontSizePt       float64                      // current font size in points
	fontSize         float64                      // current font size in user unit
	ws               float64                      // word spacing
	images           map[string]*ImageInfoType    // array of used images
	aliasMap         map[string]string            // map of alias->replacement
	pageLinks        [][]linkType                 // pageLinks[page][link], both 1-based
	links            []intLinkType                // array of internal links
	outlines         []outlineType                // array of outlines
	outlineRoot      int                          // root of outlines
	autoPageBreak    bool                         // automatic page breaking
	acceptPageBreak  func() bool                  // returns true to accept page break
	pageBreakTrigger float64                      // threshold used to trigger page breaks
	inHeader         bool                         // flag set when processing header
	headerFnc        func()                       // function provided by app and called to write header
	headerHomeMode   bool                         // set position to home after headerFnc is called
	inFooter         bool                         // flag set when processing footer
	footerFnc        func()                       // function provided by app and called to write footer
	footerFncLpi     func(bool)                   // function provided by app and called to write footer with last page flag
	zoomMode         string                       // zoom display mode
	layoutMode       string                       // layout display mode
	xmp              []byte                       // XMP metadata
	producer         string                       // producer
	title            string                       // title
	subject          string                       // subject
	author           string                       // author
	keywords         string                       // keywords
	creator          string                       // creator
	creationDate     time.Time                    // override for dcoument CreationDate value
	aliasNbPagesStr  string                       // alias for total number of pages
	pdfVersion       string                       // PDF version number
	fontDirStr       string                       // location of font definition files
	capStyle         int                          // line cap style: butt 0, round 1, square 2
	joinStyle        int                          // line segment join style: miter 0, round 1, bevel 2
	dashArray        []float64                    // dash array
	dashPhase        float64                      // dash phase
	blendList        []blendModeType              // slice[idx] of alpha transparency modes, 1-based
	blendMap         map[string]int               // map into blendList
	blendMode        string                       // current blend mode
	alpha            float64                      // current transpacency
	gradientList     []gradientType               // slice[idx] of gradient records
	clipNest         int                          // Number of active clipping contexts
	transformNest    int                          // Number of active transformation contexts
	err              error                        // Set if error occurs during life cycle of instance
	protect          protectType                  // document protection structure
	layer            layerRecType                 // manages optional layers in document
	catalogSort      bool                         // sort resource catalogs in document
	nJs              int                          // JavaScript object number
	javascript       *string                      // JavaScript code to include in the PDF
	colorFlag        bool                         // indicates whether fill and text colors are different
	color            struct {
		// Composite values of colors
		draw, fill, text colorType
	}
	spotColorMap           map[string]spotColorType // Map of named ink-based colors
	userUnderlineThickness float64                  // A custom user underline thickness multiplier.
}
