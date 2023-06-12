package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type DictResponse struct {
	Translation      string `json:"translation"`
	DetectedLanguage string `json:"detected_language"`
	Probability      int    `json:"probability"`
	BaseResp         struct {
		StatusCode    int    `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"base_resp"`
}

type DictRequest struct {
	SourceLanguage     string        `json:"source_language"`
	TargetLanguage     string        `json:"target_language"`
	Text               string        `json:"text"`
	HomeLanguage       string        `json:"home_language"`
	Category           string        `json:"category"`
	GlossaryList       []interface{} `json:"glossary_list"`
	EnableUserGlossary bool          `json:"enable_user_glossary"`
}

func query(text string) {
	client := &http.Client{}
	emp := make([]interface{}, 0)
	request := DictRequest{SourceLanguage: "detect", TargetLanguage: "zh", Text: text, HomeLanguage: "zh", Category: "", GlossaryList: emp, EnableUserGlossary: false}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(request)
	if err != nil {
		log.Fatal(err)
	}
	var data = &buf
	fmt.Println(data)

	req, err := http.NewRequest("POST", "https://translate.volcengine.com/web/translate/v1/?msToken=&X-Bogus=DFSzswVLQDap4Tl0trhkWTXyYJl1&_signature=_02B4Z6wo00001NiqiaQAAIDB-wunCEf-wtTYqo0AAFK00k1YioX9oREyHRFFxlEQRolBnKAZ4bI-bPGfYTOHAIyWMR57FNtRF6LhBYdFyoq6khUBntCBjB2zq7sN9FESyzR--MWlLuF.n48o65", data)

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "translate.volcengine.com")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-US;q=0.7,zh-TW;q=0.6")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("cookie", "x-jupiter-uuid=16864716238183286; i18next=zh-CN; s_v_web_id=verify_lir5ntdq_hMUdcY8F_51c5_4IIb_Adjp_539kDGfdOIBv; ttcid=df34f3c9aac04073ae66fd31e8e41ee968; tt_scid=XQeLZIxQh-DHIXKxlOrfKgXLHKlTe2wdxiCF0sNf4m-VBklu2jxT6SInvQ8OMydC064f")
	req.Header.Set("origin", "https://translate.volcengine.com")
	req.Header.Set("referer", "https://translate.volcengine.com/?category=&home_language=zh&source_language=detect&target_language=zh&text=")
	req.Header.Set("sec-ch-ua", `"Not.A/Brand";v="8", "Chromium";v="114", "Google Chrome";v="114"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode:", resp.StatusCode, "body", string(bodyText))
	}
	var dictResponse DictResponse
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}

func main() {
	query("has multiple digits")
}
