package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main() {
	client := &http.Client{}
	values := url.Values{"test": {"value"}}
	reader := strings.NewReader(values.Encode())
	request, err := http.NewRequest("DELETE", "http://localhost:18888", reader)

	// Headerの追加法
	request.Header.Add("Content-Type", "image/jpg")

	// BASIC認証
	request.SetBasicAuth("ユーザ名", "パスワード")

	// クッキーを手動で追加（クライアント側）
	// cookiejar.Jarのインスタンスに設定ではなく、手動で設定できる
	request.AddCookie(&http.Cookie{Name: "test", Value: "value"})
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	log.Println(string(dump))
}
