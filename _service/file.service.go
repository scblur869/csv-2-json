package _service

import (
	"encoding/csv"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type CSVData struct {
	Header []string                 `json:"header"`
	Rows   []map[string]interface{} `json:"data"`
}

func HandleUpload(ctx *gin.Context) {
	var csvData CSVData
	file, header, err := ctx.Request.FormFile("upload")
	defer file.Close()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	csvData.Header = records[0]
	for x := 1; x < len(records); x++ {
		row := make(map[string]interface{})
		for y := 0; y < len(records[0]); y++ {
			key := records[0][y]
			value := records[x][y]
			row[key] = value
		}
		csvData.Rows = append(csvData.Rows, row)
	}

	fmt.Println(header.Filename)
	ctx.JSON(http.StatusOK, csvData)

}
