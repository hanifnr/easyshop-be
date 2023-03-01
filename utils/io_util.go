package utils

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type StatusReturn struct {
	ErrCode int
	Message string
}

func StatusReturnOK() StatusReturn {
	return StatusReturn{ErrCode: 0, Message: "OK"}
}

func MessageListData(status bool, data, page interface{}) map[string]interface{} {
	resp := Message(status)
	resp["data"] = data
	resp["page"] = page
	return resp
}

func MessageData(status bool, data interface{}) map[string]interface{} {
	resp := Message(status)
	resp["data"] = data
	return resp
}

func Message(status bool) map[string]interface{} {
	return map[string]interface{}{"status": status}
}

func MessageErr(status bool, code int, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "code": code, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func RespondError(w http.ResponseWriter, data map[string]interface{}, statusCode int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func ConvertToXls(response []map[string]interface{}, header []string) *excelize.File {
	xlsx := excelize.NewFile()

	sheet1Name := "Sheet 1"
	xlsx.SetSheetName(xlsx.GetSheetName(1), sheet1Name)

	for i, value := range header {
		axis := string('A' - 1 + i + 1)
		xlsx.SetCellValue(sheet1Name, axis+"1", value)
	}

	for i, each := range response {
		for a, cell := range header {
			axis := string('A' - 1 + a + 1)
			xlsx.SetCellValue(sheet1Name, axis+strconv.Itoa(i+2), each[cell])
		}
	}

	return xlsx
}

func SetXlsHeader(w http.ResponseWriter, filename string) {
	w.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
	w.Header().Set("Content-Type", "application/xlsx")
	w.Header().Set("Content-Disposition", "attachment; filename="+filename+".xlsx")
	w.Header().Set("Content-Transfer-Encoding", "binary")
}

func SetJsonHeader(w http.ResponseWriter, filename string) {
	w.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Disposition", "attachment; filename="+filename+".json")
	w.Header().Set("Content-Transfer-Encoding", "binary")
}

const ErrValidate = 1001
const ErrSQLCreate = 1002
const ErrSQLLoad = 1003
const ErrSQLList = 1004
const ErrSQLSave = 1005
const ErrSQLDelete = 1006
const ErrSQLUpdate = 1007
const ErrExist = 1008
const ErrRequest = 1009
const ErrIO = 1010

const RESPONSE_NOT_FOUND = "data not found"
