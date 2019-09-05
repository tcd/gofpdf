# Notes

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
| `ttfparser.go`      |                                                                                           |                                      |
| `ttfparser_test.go` | *tests*                                                                                   | test                                 |
| `utf8fontfile.go`   |                                                                                           |                                      |
| `util.go`           |                                                                                           |                                      |


## Terms

### Fonts

- CID (character identifier font)
  - The CID-keyed font (also known as CID font, CID-based font, short for Character Identifier font) is a font structure, 
    originally developed for PostScript font formats, designed to address a large number of glyphs. [3]
- cmap
  - The cmap table is one of the OpenType font tables, which are required to enable correct font functioning. 
    It "defines the mapping of character codes to the glyph index values used in the font." [4]

#### Links

- [Understanding the PDF file Format – What are CID fonts](https://blog.idrsolutions.com/2011/03/understanding-the-pdf-file-format-%E2%80%93-what-are-cid-fonts/)

[1]: https://helpx.adobe.com/acrobat/using/pdf-layers.html
[2]: https://en.wikipedia.org/wiki/Spot_color
[3]: https://en.wikipedia.org/wiki/PostScript_fonts#CID
[4]: https://en.wikipedia.org/wiki/Cmap_(font)
