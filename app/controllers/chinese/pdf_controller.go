package chineseController

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/kataras/iris/v12"
	"study-service/app/controllers"
)
var pdfController PdfController

type PdfController struct {
	controllers.BaseController
}

func Abc(ctr iris.Context)  {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world")
	_ = pdf.OutputFileAndClose("hello.pdf")
}