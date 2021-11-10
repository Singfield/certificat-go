package pdf

import (
	"certificat-go/internal/cert"
	"fmt"
	"os"
	"path"

	"github.com/jung-kurt/gofpdf"
)

type PdfSaver struct {
	OutPutDir string
}

func New(outPutdir string) (*PdfSaver, error) {
	var p *PdfSaver
	err := os.MkdirAll(outPutdir, os.ModePerm)
	if err != nil {
		return p, err
	}

	p = &PdfSaver{
		OutPutDir: outPutdir,
	}

	return p, nil

}

func (p *PdfSaver) Save(cert cert.Cert) error {
	pdf := gofpdf.New(gofpdf.OrientationLandscape, "mm", "A4", "")

	pdf.SetTitle(cert.LabelTitle, false)
	pdf.AddPage()

	//background
	background(pdf)

	//header
	header(pdf, &cert)
	pdf.Ln(30)
	// Body
	pdf.SetFont("Helvetica","I", 20)
	pdf.WriteAligned(0, 50, cert.LabelPresented, "C")
	pdf.Ln(30)
	// Body - StudentName

	pdf.SetFont("Helvetica","B", 40)
	pdf.WriteAligned(0, 50, cert.Name, "C")
	pdf.Ln(30)
	// Body - Participation

	pdf.SetFont("Helvetica","I", 20)
	pdf.WriteAligned(0, 50, cert.LabelParticipation, "C")
	pdf.Ln(30)
	// Body - Date

	pdf.SetFont( "Helvetica", "I", 15)
	pdf.WriteAligned(0, 50, cert.LabelDate, "C")

	// footer

	footer(pdf)

	// save file
	fileName := fmt.Sprintf("%v.pdf", cert.LabelTitle)
	path := path.Join(p.OutPutDir, fileName)
	err := pdf.OutputFileAndClose(path)

	if err != nil {
		return err
	}

	fmt.Printf("Saved certificate to '%v' \n", path)

	return nil
}

func background(pdf *gofpdf.Fpdf) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}

	pageWidth, pageHeight := pdf.GetPageSize()

	pdf.ImageOptions("assets/img/background.png",
		0, 0,
		pageWidth, pageHeight,
		false, opts, 0, "",
	)
}

func header(pdf *gofpdf.Fpdf, c *cert.Cert) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}

	margin := 30.0
	x := 0.0
	imageWidth := 30.0
	fileName := "assets/img/gopher.png"
	pdf.ImageOptions(fileName,
		x+margin, 20,
		imageWidth, 0,
		false, opts, 0, "",
	)
	pageWidth,_ := pdf.GetPageSize()
	x = pageWidth - imageWidth

	pdf.ImageOptions(fileName,
		x-margin, 20,
		imageWidth, 0,
		false, opts, 0, "",
	)

	pdf.SetFont("Helvetica","", 40)
	pdf.WriteAligned(0,50, c.LabelCompletion, "C")
}

func footer(pdf *gofpdf.Fpdf) {
	opts := gofpdf.ImageOptions {
		ImageType :"png",
	}

	imageWidth :=50.0
	fileName :="assets/img/stamp.png"
	pageWidth, pageHeight := pdf.GetPageSize()
	x := pageWidth - imageWidth -60.0
	y := pageHeight - imageWidth -20.5

	pdf.ImageOptions(fileName,
		x,y,
		imageWidth, 0,
		false, opts, 0, "",
	)
}