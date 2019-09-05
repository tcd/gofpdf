package font

import (
	"compress/zlib"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// MakeFont generates a font definition file in JSON format. A definition file
// of this type is required to use non-core fonts in the PDF documents that
// gofpdf generates. See the makefont utility in the gofpdf package for a
// command line interface to this function.
//
// fontFileStr is the name of the TrueType file (extension .ttf), OpenType file
// (extension .otf) or binary Type1 file (extension .pfb) from which to
// generate a definition file. If an OpenType file is specified, it must be one
// that is based on TrueType outlines, not PostScript outlines; this cannot be
// determined from the file extension alone. If a Type1 file is specified, a
// metric file with the same pathname except with the extension .afm must be
// present.
//
// encodingFileStr is the name of the encoding file that corresponds to the
// font.
//
// dstDirStr is the name of the directory in which to save the definition file
// and, if embed is true, the compressed font file.
//
// msgWriter is the writer that is called to display messages throughout the
// process. Use nil to turn off messages.
//
// embed is true if the font is to be embedded in the PDF files.
func MakeFont(fontFileStr, encodingFileStr, dstDirStr string, msgWriter io.Writer, embed bool) error {
	if msgWriter == nil {
		msgWriter = ioutil.Discard
	}
	if !fileExist(fontFileStr) {
		return fmt.Errorf("font file not found: %s", fontFileStr)
	}
	extStr := strings.ToLower(fontFileStr[len(fontFileStr)-3:])
	// printf("Font file extension [%s]\n", extStr)
	var tpStr string
	switch extStr {
	case "ttf":
		fallthrough
	case "otf":
		tpStr = "TrueType"
	case "pfb":
		tpStr = "Type1"
	default:
		return fmt.Errorf("unrecognized font file extension: %s", extStr)
	}

	var info FontInfoType
	encList, err := loadMap(encodingFileStr)
	if err != nil {
		return err
	}
	// printf("Encoding table\n")
	// dump(encList)
	if tpStr == "TrueType" {
		info, err = getInfoFromTrueType(fontFileStr, msgWriter, embed, encList)
		if err != nil {
			return err
		}
	} else {
		info, err = getInfoFromType1(fontFileStr, msgWriter, embed, encList)
		if err != nil {
			return err
		}
	}
	baseStr := baseNoExt(fontFileStr)
	// fmt.Printf("Base [%s]\n", baseStr)
	if embed {
		var f *os.File
		info.File = baseStr + ".z"
		zFileStr := filepath.Join(dstDirStr, info.File)
		f, err = os.Create(zFileStr)
		if err != nil {
			return err
		}
		defer f.Close()
		cmp := zlib.NewWriter(f)
		_, err = cmp.Write(info.Data)
		if err != nil {
			return err
		}
		err = cmp.Close()
		if err != nil {
			return err
		}
		fmt.Fprintf(msgWriter, "Font file compressed: %s\n", zFileStr)
	}
	defFileStr := filepath.Join(dstDirStr, baseStr+".json")
	err = makeDefinitionFile(defFileStr, tpStr, encodingFileStr, embed, encList, info)
	if err != nil {
		return err
	}
	fmt.Fprintf(msgWriter, "Font definition file successfully generated: %s\n", defFileStr)
	return nil
}
