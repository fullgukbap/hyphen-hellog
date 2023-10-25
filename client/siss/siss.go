package siss

import (
	"bytes"
	"encoding/json"
	"hyphen-hellog/cerrors"
	"hyphen-hellog/model/response"
	"hyphen-hellog/verifier"

	"io"
	"mime/multipart"
	"net/http"
)

var serverURL = "http://101.101.217.155:8083/api/siss/images/"

func CreateImage(image *multipart.FileHeader) string {
	// request body 설정하는 방법
	var requestBody bytes.Buffer
	multipartWriter := multipart.NewWriter(&requestBody)

	// file field 설정
	part, err := multipartWriter.CreateFormFile("image", "image.txt")
	if err != nil {
		panic(cerrors.ErrUnknown)
	}

	// 업로드된 이미지 파일을 열기
	file, err := image.Open()
	if err != nil {
		panic(cerrors.ErrUnknown)
	}
	defer file.Close()

	// 파일 데이터를 MultiPart Form 데이터에 복사
	_, err = io.Copy(part, file)
	if err != nil {
		panic(cerrors.ErrUnknown)
	}

	// MultiPart Form 마무리
	err = multipartWriter.Close()
	if err != nil {
		panic(cerrors.ErrUnknown)
	}

	// HTTP POST 요청 만들기
	targetURL := "http://101.101.217.155:8083/api/siss/images/image"
	req, err := http.NewRequest("POST", targetURL, &requestBody)

	// Content-Type 설정
	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	// HTTP 클라이언트 생성
	client := &http.Client{}

	// 요청 보내기
	resp, err := client.Do(req)
	if err != nil {
		panic(cerrors.ErrUnknown)
	}
	defer resp.Body.Close()

	// body parsing
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(cerrors.ErrUnknown)
	}

	// json Unmarshal
	respJSON := new(response.GetSISS)
	err = json.Unmarshal(respBody, respJSON)
	if err != nil {
		panic(cerrors.ErrUnknown)
	}

	// 유효성 검사
	verifier.Validate(respJSON)

	// 응답 처리

	// 응답에 실패했으면
	if respJSON.Code != 201 {
		panic(cerrors.ErrRequestFailed)
	}

	return serverURL + respJSON.Data.ID
}

func DeleteImage(image string) {

	var requestBody bytes.Buffer

	// HTTP POST 요청 만들기
	targetURL := image
	req, err := http.NewRequest("DELETE", targetURL, &requestBody)

	// HTTP 클라이언트 생성
	client := &http.Client{}

	// 요청 보내기
	resp, err := client.Do(req)
	if err != nil {
		panic(cerrors.ErrUnknown)
	}
	defer resp.Body.Close()

	// body parsing
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(cerrors.ErrUnknown)
	}

	// json Unmarshal
	respJSON := new(response.DeleteSISS)
	err = json.Unmarshal(respBody, respJSON)
	if err != nil {
		panic(cerrors.ErrUnknown)
	}

	// 유효성 검사
	verifier.Validate(respJSON)

	// 응답에 실패했으면
	if respJSON.Code != 200 {
		panic(cerrors.ErrRequestFailed)
	}
}