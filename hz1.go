package main

import (
  "fmt"
  "bytes"
  "mime/multipart"
  "net/http"
  "io/ioutil"
  "math/rand"
  "time"
)

type MyResponse struct {
  Result int      `json:"result"`
  Error  []string `json:"error"`
  Data   string   `json:"data"`
}

func generateRandomString(length int) string {
  rand.Seed(time.Now().UnixNano())

  const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
  result := make([]byte, length)
  for i := range result {
    result[i] = charset[rand.Intn(len(charset))]
  }
  return string(result)
}

func main() {

  url := "https://1261aaacdty111.com/req.sys.php"
  method := "POST"

  payload := &bytes.Buffer{}
  writer := multipart.NewWriter(payload)
  randomString1 := generateRandomString(8)
  randomString2 := generateRandomString(9)
  _ = writer.WriteField("JDATA", "{\"data\":{\"mod\":\"cw_r00706_html\",\"app\":\"frontaction\",\"func\":\"insertorder\",\"account\":\""+randomString1+"\",\"mActcode\":\""+randomString2+"\"}}")
  err := writer.Close()
  if err != nil {
    fmt.Println(err)
    return
  }


  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Accept", "application/json, text/plain, */*")
  req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
  req.Header.Add("Cache-Control", "no-cache")
  req.Header.Add("Connection", "keep-alive")
  req.Header.Add("Origin", "https://0306693.com")
  req.Header.Add("Pragma", "no-cache")
  req.Header.Add("Referer", "https://0306693.com/")
  req.Header.Add("Sec-Fetch-Dest", "empty")
  req.Header.Add("Sec-Fetch-Mode", "cors")
  req.Header.Add("Sec-Fetch-Site", "cross-site")
  req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
  req.Header.Add("sec-ch-ua", "\"Not A(Brand\";v=\"99\", \"Google Chrome\";v=\"121\", \"Chromium\";v=\"121\"")
  req.Header.Add("sec-ch-ua-mobile", "?0")
  req.Header.Add("sec-ch-ua-platform", "\"Windows\"")

  req.Header.Set("Content-Type", writer.FormDataContentType())
  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  


  fmt.Println(string(body))
}