package main

import (
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"log"
	"os"
)

func main(){
	
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	
	if err != nil {
		log.Fatal(err)
	}
	
	f, err := os.Open("file/template_email/transaction_history.html")
	if f != nil {
		defer f.Close()
	}
	if err != nil {
		log.Fatal(err)
	}
	
	pdfg.AddPage(wkhtmltopdf.NewPageReader(f))
	
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.Dpi.Set(80)
	pdfg.MarginLeft.Set(0)
	pdfg.MarginRight.Set(0)
	
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}
	
	err = pdfg.WriteFile("./transaction_history.pdf")
	if err != nil {
		log.Fatal(err)
	}
	
	log.Println("Done")
}