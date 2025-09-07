package insideservice

import "fmt"

type DataService interface {
	SendPDFDocument() 
} 

type PDFDocument struct {
}

func (doc PDFDocument) SendPDFDocument() {
	fmt.Println("sending pdf document...")
}
