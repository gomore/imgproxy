package imagemeta

import (
	"bytes"
	"io"

	"github.com/imgproxy/imgproxy/v3/imagetype"
)

var pdfMagick = []byte("%PDF-")

type PdfFormatError string

func (e PdfFormatError) Error() string { return "invalid PDF format: " + string(e) }

func DecodePdfMeta(r io.Reader) (Meta, error) {
	var tmp [16]byte

	if _, err := io.ReadFull(r, tmp[:5]); err != nil {
		return nil, err
	}

	if !bytes.Equal(pdfMagick, tmp[:5]) {
		return nil, PdfFormatError("not a PDF image")
	}

	if _, err := io.ReadFull(r, tmp[:]); err != nil {
		return nil, err
	}

	return &meta{
		format: imagetype.PDF,
		width:  1,
		height: 1,
	}, nil
}

func init() {
	RegisterFormat(string(pdfMagick), DecodePdfMeta)
}
