# Fonts

## Font Types

### OpenType Font (OTF)

OpenType is the evolution of TTF. 
It is the result of a joint effort between Adobe and Microsoft.

- Glyph Limit ≈ 65,535
- OpenType fonts may have the extension .OTF, .TTF, .OTC or .TTC. 
  The extensions .OTC and .TTC should only be used for font collection files.

#### Tables:

- `head`
  - [Font Header Table](https://docs.microsoft.com/en-us/typography/opentype/spec/head)
  - This table gives global information about the font.
- `hhea`
  - [Horizontal Header Table](https://docs.microsoft.com/en-us/typography/opentype/spec/hhea)
  - This table contains information for horizontal layout. 
- `maxp`
  - [Maximum Profile](https://docs.microsoft.com/en-us/typography/opentype/spec/maxp)
  - This table establishes the memory requirements for this font. 
  - Fonts with CFF data must use Version 0.5 of this table.
  - Fonts with TrueType outlines must use Version 1.0 of this table.
- `hmtx`
  - [Horizontal Metrics](https://docs.microsoft.com/en-us/typography/opentype/spec/hmtx)
  - The horizontal metrics table provides glyph advance widths and left side bearings.
- `cmap`
  - [Character to Glyph Index Mapping Table](https://docs.microsoft.com/en-us/typography/opentype/spec/cmap)
  - This table defines the mapping of character codes to the glyph index values used in the font. 
  - It may contain more than one subtable, in order to support more than one character encoding scheme.
- `name`
  - [Naming Table](https://docs.microsoft.com/en-us/typography/opentype/spec/name)
  - The naming table allows multilingual strings to be associated with the OpenType font. 
  - These strings can represent copyright notices, font names, family names, style names, and so on. 
- `OS/2`
  - [OS/2 and Windows Metrics Table](https://docs.microsoft.com/en-us/typography/opentype/spec/os2)
  - The OS/2 table consists of a set of metrics and other data that are required in OpenType fonts.
- `post`
  - [PostScript Table](https://docs.microsoft.com/en-us/typography/opentype/spec/post)
  - This table contains additional information needed to use TrueType or OpenType fonts on PostScript printers. 
  - This includes data for the FontInfo dictionary entry and the PostScript names of all the glyphs. 
  - For more information about PostScript names, see the Adobe [Glyph List Specification](https://github.com/adobe-type-tools/agl-specification).

#### Also:

- `loca`
  - [Index to Location](https://docs.microsoft.com/en-us/typography/opentype/spec/loca)
  - The indexToLoc table stores the offsets to the locations of the glyphs in the font, relative to the beginning of the glyphData table.

### TrueType Font (TTF)

The TrueType font format was developed by Apple and Microsoft as a response to the PostScript font format.

#### Tables:

Technically, a simple font is a font that only contains the standard required tables ('cmap', 'glyf', 'head', 'hhea', 'hmtx', 'loca', 'maxp', 'post') and possibly the non-layout related optional tables ('cvt ', 'fpgm', 'hdmx', 'prep'). All TrueType fonts released with Macintosh System 7 are simple fonts.

### PostScript Font (Type1)

- Limited to 256 glyphs
- File Extension: Two files required.
  - On Windows, .pfb (Printer Font Binary, the actual font outline) and .pfm (Printer Font Metrics, the metric data such as kerning).
  - On Mac, .lwfn (font outline) and .ffil (font suitcase). 
  - On Linux, .pfa (font file) and .afm (metrics).
- Names:
  - PostScript font
  - Type 1 font
  - Type 2 font
  - CID font (Character Identifier font)
  - CFF font (Compact Font Format) (This one is more of a font table)

### SVG Font

Currently not supported by PDFs.

### Web Open Font Format (WOFF)

WOFF is essentially OpenType or TrueType with compression and additional metadata.
It was created to live on the web.

## Terms

- [cmap](https://docs.microsoft.com/en-us/typography/opentype/spec/cmap#format-14-unicode-variation-sequences)
  - Character to Glyph Index Mapping Table
  - The cmap table is one of the OpenType font tables, which are required to enable correct font functioning. 
    It "defines the mapping of character codes to the glyph index values used in the font." [2]
- (OpenType) Font Collection (*previously TrueType Collection*)
- .notdef glyph
  - (□)
  - An interpretable but unrenderable character
- CFF
  - CFF stands for the Type1 font format. Strictly speaking, it refers to the Compact Font Format, which is used in the compression processes for the Type2 fonts.
- Typeface (Font Family)
  - A set of one or more fonts each composed of glyphs that share common design features.
- [Code Page](https://en.wikipedia.org/wiki/Code_page)
  - A code page is a character encoding and as such it is a specific association of a set of printable characters and control characters with unique numbers. 
  - Code pages were used before Unicode, and Unicode should be used instead (or so it seems to me; there may be an issue with PDF support).
  - [Windows-1252](https://en.wikipedia.org/wiki/Windows-1252)
- Subset Font (pdf term?)
- Glyph Index
- UTF8
- Font Table
  - The Font table contains the information for registering font files with the system.

## Font Filetypes

| Extension |               Name               |                  Notes                   |
| --------- | -------------------------------- | ---------------------------------------- |
| .TTF      | TrueType Font                    |                                          |
| .WOFF     | Web Open Font Format File        |                                          |
| .AMFM     | Adobe Multiple Font Metrics File |                                          |
| .FNT      | Windows Font File                |                                          |
| .OTF      | OpenType Font                    |                                          |
| .AFM      | Adobe Font Metrics File          |                                          |
| .EOT      | Embedded OpenType Font           | (Only supported by Interned Explorer...) |
| .PFB      | Printer Font Binary File         |                                          |
| .FOT      | Font Resource File               |                                          |
| .OTC      | and .TTC Font Collection Files   |                                          |

## Links

- [Understanding the PDF file Format – What are CID fonts](https://blog.idrsolutions.com/2011/03/understanding-the-pdf-file-format-%E2%80%93-what-are-cid-fonts/)
- [A Simple Guide to Font File Types](https://www.suttle-straus.com/blog/a-simple-guide-to-font-file-types)
- [What Are the Different Types of Font Files?](https://www.lifewire.com/types-of-fonts-1697695)
- [OpenType Font File Spec (Microsoft)](https://docs.microsoft.com/en-us/typography/opentype/spec/otff)
- [The Missing Guide to Font Formats: TTF, OTF, WOFF, EOT & SVG](https://creativemarket.com/blog/the-missing-guide-to-font-formats)
- [TrueType Spec (Apple)](https://developer.apple.com/fonts/TrueType-Reference-Manual/RM06/Chap6AATIntro.html)

[1]: https://en.wikipedia.org/wiki/PostScript_fonts#CID
[2]: https://en.wikipedia.org/wiki/Cmap_(font)
