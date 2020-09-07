package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
)

func main() {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	// ファイルではないコンテンツの登録
	writer.WriteField("name", "Michael Jackson")
	// 他に読み込むものとして書き込み用のwriter作成
	// Content-Typeのvoid型のapplication/octet-streamではなく任意の型を指定
	part := make(textproto.MIMEHeader)
	part.Set("Content-Type", "image/jpeg")
	part.Set("Content-Disposition", `form-data; name="thumbnail"; filename="photo.jpg"`)
	fileWriter, err := writer.CreatePart(part)
	if err != nil {
		panic(err)
	}
	// ファイルを開く
	readFile, err := os.Open("photo.jpg")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()
	// 開いたファイルをwriterにコピー
	io.Copy(fileWriter, readFile)
	writer.Close()

	// writerがbufferに書き込む様になっているのでそのまま渡す
	// FormDataContentTypeがboundary文字列の生成メソッド
	resp, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}
