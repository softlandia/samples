package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	_ "image/png"
	"log"
	"net/http"
	"text/template"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

// Page -
type Page struct {
	Title string
}

/*
подразумеваем,
что обращаются по адресу: /getQR
в теле шлют:
{
	userID: 	uint64,
	officeID: 	uint64,
	date:		2020-11-28
}
в ответ получают png файл с картинкой
*/
func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/getQR", getQRHandler)
	http.HandleFunc("/generator/", viewCodeHandler)
	http.ListenAndServe(":8081", nil)
}

func getQRHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Post("http://localhost:8080/qr", "application/json", bytes.NewReader([]byte(`{"employeeID":1,"officeID":2,"date":"2000-02-15"}`)))
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
	/*
		b, _ := ioutil.ReadAll(resp.Body)
		ioutil.WriteFile("data-resp.png", b, 0644)
		buf := bytes.NewBuffer(b)
		img, format, _ := image.Decode(buf)
	*/
	img, format, _ := image.Decode(resp.Body)
	log.Printf("img format: '%s'", format)
	bmp, _ := gozxing.NewBinaryBitmapFromImage(img)
	qrReader := qrcode.NewQRCodeReader()
	result, _ := qrReader.Decode(bmp, nil)
	log.Println(result)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	p := Page{Title: "QR Code Generator"}

	t, _ := template.ParseFiles("generator.html")
	t.Execute(w, p)
}

func viewCodeHandler(w http.ResponseWriter, r *http.Request) {
	dataString := r.FormValue("dataString")
	qrCode, _ := qr.Encode(dataString, qr.L, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 512, 512)
	png.Encode(w, qrCode)
}
