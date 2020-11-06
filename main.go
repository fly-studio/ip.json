package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	_path, err := os.Executable()
	if err != nil {
		panic("read current path error")
	}
	currentDir := filepath.Dir(_path)

	datFile := flag.String("qqwry", filepath.Join(currentDir, "qqwry.dat"), "纯真 IP 库的地址")
	port := flag.String("port", "2060", "HTTP 请求监听端口号")
	flag.Parse()

	IPData.FilePath = *datFile

	res := IPData.InitIPData()

	if v, ok := res.(error); ok {
		log.Panic(v)
	}
	log.Printf("IP 库加载完成 共加载:%d 条 IP 记录，启动 0.0.0.0:%s\n", IPData.IPNum, *port)

	http.HandleFunc("/", ServeHTTP)
	if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", *port), nil); err != nil {
		panic(err)
	}

}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var msg struct {
		Ip      string `json:"ip"`
		Url     string `json:"url"`
		Country string `json:"country"`
		Area    string `json:"area"`
	}
	ipStr := r.Header.Get("X-Forwarded-For")
	if ipStr == "" {
		ipStr = r.RemoteAddr
		ipStr, _, _ = net.SplitHostPort(ipStr)
	}

	log.Printf("IP: %s\n", ipStr)

	if ipStr != "" {
		qqWry := NewQQwry()
		res := qqWry.Find(ipStr)
		msg.Country = res.Country
		msg.Area = res.Area
	}

	msg.Ip = ipStr
	msg.Url = r.URL.String()

	j, err := json.Marshal(msg)
	if err != nil {
		log.Print(err.Error())
		return
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(j)
}
