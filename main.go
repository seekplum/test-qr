package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"bytes"
	"image/png"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func handleQr(w http.ResponseWriter, r *http.Request) {
	// 获取参数
	params := r.URL.Query()
	data := params.Get("data")
	if data == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("data is required"))
		return
	}
	// 创建二维码
	qrCode, _ := qr.Encode(data, qr.M, qr.Auto)
	// 将二维码缩放到200x200像素
	qrCode, _ = barcode.Scale(qrCode, 200, 200)
	buf := new(bytes.Buffer)
	// 将二维码内容写入 buf
	err := png.Encode(buf, qrCode)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	// w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Type", "image/jpeg")
	w.Write(buf.Bytes())
	return
}

func RequestLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(rw, r) // serve the original request
		duration := time.Since(start)
		log.Printf("method: %s, uri: %s, duration: %s", r.Method, r.RequestURI, duration)
	})
}
func getEnvDefault(key, defaultVal string) string {
	val, ex := os.LookupEnv(key)
	if !ex {
		return defaultVal
	}
	return val
}

func main() {
	// 设置日志格式
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetOutput(os.Stderr)

	defaultHost := getEnvDefault("HOST", ":8089")
	var host string
	flag.StringVar(&host, "host", defaultHost, "绑定的端口号")
	flag.Parse()

	log.Printf("Listening and serving HTTP on %s", host)

	mux := http.NewServeMux()
	mux.HandleFunc("/qr", handleQr)
	err := http.ListenAndServe(host, RequestLogger(mux))
	if err != nil {
		log.Println(err)
	}
}
