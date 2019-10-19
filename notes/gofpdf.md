# gofpdf

## Files

|        file         |                                        description                                        |                notes                 |
| ------------------- | ----------------------------------------------------------------------------------------- | ------------------------------------ |
| `list/list.go`      |                                                                                           | Unused?                              |
| `compare.go`        | Functions for comparing Bytes & Pdfs                                                      | Unused?                              |
| `def.go`            | Variable and Type declarations                                                            |                                      |
| `doc.go`            | package documentation for `godoc`                                                         |                                      |
| `embedded.go`       | Data for embedded standard fonts                                                          |                                      |
| `font.go`           | Functions for parsing and embedding font files in PDFs                                    | `MakeFont()` is defined in this file |
| `fpdf.go`           | Main file                                                                                 |                                      |
| `fpdf_test.go`      | *tests*                                                                                   | test                                 |
| `fpdftrans.go`      | Functions for transforming, translating, & scaling PDF content (text, drawings, & images) |                                      |
| `grid.go`           | Utility methods for drawing graphs                                                        |                                      |
| `htmlbasic.go`      | Interface for drawing basic HTML to a PDF                                                 |                                      |
| `label.go`          | Functions used in `grid.go` for labeling graphs                                           |                                      |
| `layer.go`          | Methods adding [PDF layer][1] functionality                                               |                                      |
| `png.go`            | PNG parsing functions                                                                     |                                      |
| `protect.go`        | Functions for encrypting & password protecting PDF files                                  |                                      |
| `splittext.go`      | `fpdf.SplitText()` Definition                                                             | Single Function                      |
| `spotcolor.go`      | Methods enabling the use of [spot colors][2] (used in professional printing)              |                                      |
| `subwrite.go`       | Method for writing subscripts and superscripts                                            | Single Function                      |
| `svgbasic.go`       | Super basic SVG parsing                                                                   |                                      |
| `svgwrite.go`       | Super basic SVG rendering                                                                 | Single Function                      |
| `template.go`       |                                                                                           |                                      |
| `template_impl.go`  |                                                                                           |                                      |
| `ttfparser.go`      | Utility to parse TTF font files                                                           |                                      |
| `ttfparser_test.go` | *tests*                                                                                   | test                                 |
| `utf8fontfile.go`   |                                                                                           |                                      |
| `util.go`           |                                                                                           |                                      |


[1]: https://helpx.adobe.com/acrobat/using/pdf-layers.html
[2]: https://en.wikipedia.org/wiki/Spot_color
