package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"io"
	"net/http"
	"time"
)

func main() {
	// 参数
	ts := time.Now().Unix()
	agentHostname := "zy-fat"

	// 拼接待签名字符串
	signkey := global.GVA_CONFIG.BasicAuth.User + global.GVA_CONFIG.BasicAuth.Password
	toSign := fmt.Sprintf("%s-sign-key=%s@&target=%s@&sign-time=%s", ts, signkey, agentHostname, ts)

	// 生成 HMAC-SHA256 签名
	mac := hmac.New(sha256.New, []byte(signkey))
	mac.Write([]byte(toSign))
	signature := hex.EncodeToString(mac.Sum(nil))

	// 打印签名
	fmt.Println("Generated Signature:", signature)

	// 构建最终请求 URL
	url := fmt.Sprintf("http://localhost:8888/config/categraf?ts=%d&agent_hostname=%s&version=sss", ts, agentHostname)
	fmt.Println(url)
	// 创建 HTTP 请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// 将签名添加到 X-Sign 请求头
	req.Header.Add("X-Sign", signature)

	// 发送 HTTP 请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println("Response status:", resp.Status)
	fmt.Println("Response body:", string(body))
}
