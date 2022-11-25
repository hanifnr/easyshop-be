package utils

import (
	"io/ioutil"
	"net/http"
)

func UploadFile(w http.ResponseWriter, r *http.Request) ([]byte, StatusReturn) {
	r.ParseMultipartForm(10 << 20)

	file, _, err := r.FormFile("image")
	if err != nil {
		return nil, StatusReturn{ErrCode: ErrRequest, Message: err.Error()}
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, StatusReturn{ErrCode: ErrIO, Message: err.Error()}
	}

	return fileBytes, StatusReturnOK()
}
