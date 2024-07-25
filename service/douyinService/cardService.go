package douyinService

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
)

func AddCard(userId string, productIdStr string) (*CheckDisplayTimeResp, error) {

	fileContents, err := getCookie(userId)

	if err != nil {
		return nil, err
	}
	// Convert file contents to string
	cookie := string(fileContents)

	fmt.Println("cookie:", cookie)

	payload := map[string]interface{}{
		"promotion_identities": []map[string]interface{}{
			{
				"entity_id":   productIdStr,
				"entity_type": 1,
			},
		},
		"need_total": true,
		"step_plan":  false,
		"extra": map[string]string{
			"path": "/dashboard/merch-picking-library",
		},
	}

	jsonData, err := json.Marshal(payload)

	fmt.Println(string(jsonData))

	req, err := http.NewRequest("POST", AddMyCard, bytes.NewBuffer([]byte(jsonData)))

	// Set headers
	setReqHeader(req, cookie)

	client := http.Client{}
	resp, err := client.Do(req)
	fmt.Println("resp:", resp)
	fmt.Println("resp Body :", resp.Body)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Println("check_display====", string(body))
	ret := &CheckDisplayTimeResp{}
	_ = json.Unmarshal(body, ret)
	return ret, nil
}

func DeleteCard(douyinId string, productIdStr string) (*CheckDisplayTimeResp, error) {

	fileContents, err := getCookie(douyinId)
	if err != nil {
		return nil, err
	}
	// Convert file contents to string
	cookie := string(fileContents)

	fmt.Println("cookie:", cookie)

	// 创建一个map结构，用于保存要转换为JSON的数据
	payload := map[string]interface{}{
		"identifiers": []map[string]interface{}{
			{
				"entity_id":   productIdStr,
				"entity_type": 1,
			},
		},
	}

	jsonData, err := json.Marshal(payload)

	fmt.Println(string(jsonData))

	req, err := http.NewRequest("POST", DelMyCard, bytes.NewBuffer([]byte(jsonData)))

	setReqHeader(req, cookie)

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Println("resp:", resp)
	fmt.Println("resp Body :", resp.Body)

	body, _ := io.ReadAll(resp.Body)

	fmt.Println("check_display====", string(body))
	ret := &CheckDisplayTimeResp{}
	_ = json.Unmarshal(body, ret)
	return ret, nil
}

func setReqHeader(req *http.Request, cookie string) {
	// Set headers
	req.Header.Set("authority", "buyin.jinritemai.com")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("cookie", cookie)
	req.Header.Set("origin", "https://buyin.jinritemai.com")
	req.Header.Set("referer", "https://buyin.jinritemai.com/dashboard/merch-picking-cart?btm_ppre=a10091.b82437.c0.d0&btm_pre=a10091.b24215.c68160.d839440_i606647&btm_show_id=9c3235cc-ea6f-405c-9af0-757f8bc5b7cd&pre_universal_page_params_id=&universal_page_params_id=64b58710-ca85-41e8-ac4f-a581c36f4b20")
	req.Header.Set("sec-ch-ua", `"Chromium";v="116", "Not)A;Brand";v="24", "Google Chrome";v="116"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36")
}

func getCookie(userId string) ([]byte, error) {
	usr, _ := user.Current()
	dir := usr.HomeDir

	douyinLoginPath := "/Work/data/broadcast/login/douyin/"

	// 组装文件路径
	filePath := filepath.Join(dir, douyinLoginPath, fmt.Sprintf("%s_cookie.txt", userId))

	fmt.Println("File path:", filePath)

	// Read the entire file content
	fileContents, err := os.ReadFile(filePath)
	return fileContents, err
}
