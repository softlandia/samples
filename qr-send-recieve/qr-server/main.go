package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"time"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/gin-gonic/gin"
)

func writeQR(s string, w io.Writer) {
	qrCode, _ := qr.Encode(s, qr.L, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 512, 512)
	png.Encode(w, qrCode)
}

func main() {
	router := gin.Default()
	router.POST("/qr", qrGenerate)
	router.Run(":8080")
}

func qrGenerate(c *gin.Context) {
	type qrRequest struct {
		EmployeeID uint64    `json:"employeeID"` // идентификатор сотрудника
		OfficeID   uint64    `json:"officeID"`   // идентификатор склада
		Date       time.Time `json:"date"`       // дата выхода
	}
	body, _ := ioutil.ReadAll(c.Request.Body)
	q := qrRequest{}
	if err := json.Unmarshal(body, &q); err != nil {
		log.Println(err)
		c.String(400, "%s", err.Error())
		return
	}
	log.Println(q)
	dataString := fmt.Sprintf(`{"employeeID":%d,"officeID":%d,"date":%s}`, q.EmployeeID, q.OfficeID, q.Date.Format("2006-01-02"))
	qrCode, _ := qr.Encode(dataString, qr.L, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 512, 512)
	b := bytes.NewBuffer([]byte{})
	png.Encode(b, qrCode)
	c.Data(200, "png", b.Bytes())
}
