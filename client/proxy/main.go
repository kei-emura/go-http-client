package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	proxyURL, err := url.Parse("http://localhost:18888")
	if err != nil {
		panic(err)
	}
	// Proxyを設定するとGetなどリクエストを投げる場合に先にProxyURLに指定したリソースへアクセスする
	// http.Clientはデフォルトで環境変数のHTTP_PROXY、HTTPS_PROXYを読み込む
	// ProxyURLにBASIC認証を設定するとBASIC認証でのユーザー名とパスワードを与えられる
	// http://ユーザー名:パスワード@github.com
	client := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
	}
	resp, err := client.Get("http://github.com")
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	log.Println(string(dump))
}
