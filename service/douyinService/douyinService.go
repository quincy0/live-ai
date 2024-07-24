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
	"strings"
)

type Douyin struct {
	cookies string
}

const (
	// GetPlatformCalendar 获取csrf Token
	GetPlatformCalendar = "https://buyin.jinritemai.com/api/buyin/marketing/anchor_redpacket/get_platform_calendar"
	CreateRedPacket     = "https://buyin.jinritemai.com/api/buyin/marketing/anchor_redpacket/create"
	DisplayRedPacket    = "https://buyin.jinritemai.com/api/buyin/marketing/anchor_redpacket/edit_display_time"
	DeleteRedPacket     = "https://buyin.jinritemai.com/api/buyin/marketing/anchor_redpacket/update_status"
	CheckDisplayTime    = "https://buyin.jinritemai.com/api/buyin/marketing/anchor_redpacket/check_display_time"
	// 选品车
	GetMyCard = "https://buyin.jinritemai.com/selection_cart_pc_api/card/m_get_card"
	AddMyCard = "https://buyin.jinritemai.com/selection_cart_pc_api/card/m_add_card"
	DelMyCard = "https://buyin.jinritemai.com/selection_cart_pc_api/card/m_del_card"
	CarCount  = "https://buyin.jinritemai.com/selectionpc api/comm/m get cart sum"
	// PromotionList 达人红包
	PromotionList      = "https://buyin.jinritemalin/marketing/anchor_coupon/promotion_list"
	CreateExportCoupon = "https://buyin.jinritemai.com/api/buyin/marketing/anchor coupon/create"
	ExportCouponList   = "https://buyin.jinritemai.com/api/buvin/marketing/anchor coupon/list"
	// 违规警告

	MarginControl = "https://buyin.jinritemai.com/aweme/v2/governance/margin_control/creator/violations"
	//获取抖音商品信息
	MaterialList = "https://buyin.jinritemai.com/pc/selection/common/material list"
	// pin
	Pin = "https://buyin.jinritemai.com/api/anchor/livepc/setcurrent"
)

// 创建红包

func (s *Douyin) CreateRedPack(name string, startTime, endTime int64, totalAmount, totalCredit int) (*CreateRedPacketResp, error) {

	//template := "{\"redpacket_data\":{\"redpacket_activity\":{\"activity_biz_type\":1,\"redpacket_activity_name\":\"0619\",\"max_apply_times\":1,\"validity_type\":1,\"live_redpack_activity_sub_type\":3,\"kol_user_tag\":0,\"redpacket_sub_type\":1,\"redpack_type\":11,\"valid_start_time\":1718801453,\"valid_end_time\":1718805053,\"total_credit\":10000},\"redpacket_meta_list\":[{\"total_amount\":20,\"total_credit\":10000,\"credit_type\":1,\"extra_info\":{\"strategy_goal\":0},\"avg_credit\":500}]}}"

	input := InitRedpacket()
	input.RedpacketData.RedpacketActivity.RedpacketActivityName = name
	input.RedpacketData.RedpacketActivity.ValidStartTime = startTime
	input.RedpacketData.RedpacketActivity.ValidEndTime = endTime
	input.RedpacketData.RedpacketMetaList[0].TotalAmount = totalAmount
	input.RedpacketData.RedpacketMetaList[0].TotalCredit = totalCredit
	input.RedpacketData.RedpacketMetaList[0].AvgCredit = totalCredit / totalAmount
	b, _ := json.Marshal(input)

	req, err := http.NewRequest("POST", CreateRedPacket, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	csrfToken, _, err := s.GetCsrfToken()
	req.Header.Add("x-secsdk-csrf-token", csrfToken)
	req.Header.Add("cookie", s.cookies)
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("origin", "https://buyin.jinritemai.com")
	req.Header.Add("priority", "u=1, i")
	req.Header.Add("referer", "https://buyin.jinritemai.com/dashboard/marketing/red-envelope?pre_universal_page_params_id=&universal_page_params_id=ccf76af9-afcf-47ac-8c0f-435ee987938c")
	req.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"125\", \"Chromium\";v=\"125\", \"Not.A/Brand\";v=\"24\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Println("create:===", string(body))

	ret := &CreateRedPacketResp{}
	_ = json.Unmarshal(body, ret)
	return ret, nil
}

//查询红包列表

//func (s *Douyin) GetRedPacketList() {
//
//}

// 橱窗列表
//func (s *Douyin) GetTransferList() {
//
//}

//红包投放 redpacket_activity_id=7382090484015300914&display_time_type=1&redpack_type=11&period_display_after_now=120&apply_period=120

func (s *Douyin) DisplayRedPacket(activityId string, periodDisplayAfterNow, applyPeriod string) (*DisplayResp, error) {

	// 将数据编码为URL形式
	data := "redpacket_activity_id=" + activityId + "&display_time_type=1&redpack_type=11&period_display_after_now=" + periodDisplayAfterNow + "&apply_period=" + applyPeriod

	fmt.Println("****", data)
	req, err := http.NewRequest("POST", DisplayRedPacket, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return nil, err
	}

	csrfToken, _, err := s.GetCsrfToken()
	req.Header.Add("x-secsdk-csrf-token", csrfToken)
	req.Header.Add("cookie", s.cookies)
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("origin", "https://buyin.jinritemai.com")
	req.Header.Add("priority", "u=1, i")
	req.Header.Add("referer", "https://buyin.jinritemai.com/dashboard/marketing/red-envelope?pre_universal_page_params_id=&universal_page_params_id=ccf76af9-afcf-47ac-8c0f-435ee987938c")
	req.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"125\", \"Chromium\";v=\"125\", \"Not.A/Brand\";v=\"24\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Println("display====", string(body))

	ret := &DisplayResp{}
	_ = json.Unmarshal(body, ret)
	return ret, nil

}

func (s *Douyin) CheckDisplayTime(activityId string, periodDisplayAfterNow, applyPeriod string) (*CheckDisplayTimeResp, error) {

	// 将数据编码为URL形式
	data := "redpacket_activity_id=" + activityId + "&display_time_type=1&redpack_type=11&period_display_after_now=" + periodDisplayAfterNow + "&apply_period=" + applyPeriod

	fmt.Println("****", data)
	req, err := http.NewRequest("POST", CheckDisplayTime, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return nil, err
	}

	csrfToken, _, err := s.GetCsrfToken()
	req.Header.Add("x-secsdk-csrf-token", csrfToken)
	req.Header.Add("cookie", s.cookies)
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("origin", "https://buyin.jinritemai.com")
	req.Header.Add("priority", "u=1, i")
	req.Header.Add("referer", "https://buyin.jinritemai.com/dashboard/marketing/red-envelope?pre_universal_page_params_id=&universal_page_params_id=ccf76af9-afcf-47ac-8c0f-435ee987938c")
	req.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"125\", \"Chromium\";v=\"125\", \"Not.A/Brand\";v=\"24\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36")

	client := http.Client{}
	resp, err := client.Do(req)
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

//作废红包

func (s *Douyin) GetCsrfToken() (string, string, error) {
	reqInner, err := http.NewRequest("HEAD", GetPlatformCalendar, nil)
	if err != nil {
		return "", "", err
	}

	// 应该加上cookie，虽然不加也能获取到csrf-token但是却不能在接下来到接口中使用
	reqInner.Header.Add("cookie", s.cookies)
	reqInner.Header.Add("accept", "*/*")
	reqInner.Header.Add("accept-language", "zh-CN,zh;q=0.9,en;q=0.8")
	reqInner.Header.Add("priority", "u=1, i")
	reqInner.Header.Add("referer", "https://buyin.jinritemai.com/dashboard/marketing/red-envelope?pre_universal_page_params_id=&universal_page_params_id=ccf76af9-afcf-47ac-8c0f-435ee987938c")

	reqInner.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"125\", \"Chromium\";v=\"125\", \"Not.A/Brand\";v=\"24\"")
	reqInner.Header.Add("sec-ch-ua-mobile", "?0")
	reqInner.Header.Add("sec-ch-ua-platform", "macOS")
	reqInner.Header.Add("sec-fetch-mode", "cors")
	reqInner.Header.Add("sec-fetch-site", "same-origin")
	reqInner.Header.Add("sec-fetch-dest", "empty")
	reqInner.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36")

	reqInner.Header.Add("x-secsdk-csrf-request", "1")
	reqInner.Header.Add("x-secsdk-csrf-version", "1.2.22")
	client := http.Client{}
	respInner, err := client.Do(reqInner)
	if err != nil {
		return "", "", err
	}

	csrfToken := respInner.Header.Get("x-ware-csrf-token")
	setCookie := respInner.Header.Get("set-cookie")
	if len(csrfToken) > 0 && len(setCookie) > 0 {
		list := strings.Split(csrfToken, ",")
		setCookieList := strings.Split(setCookie, ";")
		return list[1], setCookieList[0], nil
	} else {
		return "", "", nil
	}
}

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
