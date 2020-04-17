package chineseController

import (
	"github.com/kataras/iris/v12"
	"github.com/signintech/gopdf"
	"log"
	"study-service/app/controllers"
)
var pdfController PdfController

type PdfController struct {
	controllers.BaseController
}

func Abc(ctr iris.Context)  {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{ PageSize: *gopdf.PageSizeA4 })
	pdf.AddPage()
	err := pdf.AddTTFFont("wts11", "./data/ttf/wts11.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.SetFont("wts11", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.Cell(nil, "您好")
	pdf.WritePdf("hello.pdf")
}