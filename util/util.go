package util

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"
)

// HttpGet GET请求
func HttpGet(path string) ([]byte, error) {
	resp, err := http.Get(path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

// HttpPost POST请求
func HttpPost(path string, body interface{}) ([]byte, error) {
	params, _ := json.Marshal(body)
	resp, err := http.Post(path, "application/json;charset=utf-8", bytes.NewBuffer(params))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

// FileOptions 文件上传参数
type FileOptions struct {
	//文件名
	FileName string `json:"fileName"`					// 文件名
	//文件大小
	FileSize int64 `json:"fileSize"`					// 文件大小
	//文件内容
	File multipart.File									// 文件内容
}

// HttpPostFile POST上传文件
func HttpPostFile(path string, options FileOptions) ([]byte, error) {
	bodyBuf := bytes.Buffer{}
	bodyWriter := multipart.NewWriter(&bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile("media", options.FileName)
	if err != nil {
		return nil, err
	}

	if _, err = io.Copy(fileWriter, options.File); err != nil {
		return nil, err
	}

	contentType := bodyWriter.FormDataContentType()
	_ = bodyWriter.Close()

	_ = bodyWriter.WriteField("filelength", strconv.Itoa(int(options.FileSize)))

	resp, err := http.Post(path, contentType, &bodyBuf)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

// HttpPostOriginFile POST上传文件
func HttpPostOriginFile(path, fileName string, size int, body []byte) ([]byte, error) {
	bodyBuf := bytes.Buffer{}
	bodyWriter := multipart.NewWriter(&bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile("media", fileName)
	if err != nil {
		return nil, err
	}

	if _, err = io.Copy(fileWriter, bytes.NewReader(body)); err != nil {
		return nil, err
	}

	contentType := bodyWriter.FormDataContentType()
	_ = bodyWriter.Close()

	_ = bodyWriter.WriteField("filelength", strconv.Itoa(size))

	resp, err := http.Post(path, contentType, &bodyBuf)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}
