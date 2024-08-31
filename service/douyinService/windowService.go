package douyinService

import (
	"encoding/json"
	"fmt"
	"github.com/PaesslerAG/jsonpath"
	"io"
	"net/http"
	"strings"
)

func GetWindow(douyinId string) ([]ProductDetail, error) {

	fileContents, err := getCookie(douyinId)
	if err != nil {
		return nil, err
	}
	// Convert file contents to string

	cookie := string(fileContents)
	fmt.Println("cookie:", cookie)

	url := "https://buyin.jinritemai.com/api/author/shop/products?product_type=0&page=1&page_size=200&filter=false&verifyFp=verify_lyh1wlcm_7vVhuv21_rChj_40I5_8Tjk_a1HLz1zdEsUf&fp=verify_lyh1wlcm_7vVhuv21_rChj_40I5_8Tjk_a1HLz1zdEsUf&msToken=v-6XPe5aQ155CnCC_0iDcQpGPml4i6xrN5_oxk8Dtw5wMzriV5ATm4j1iylVvgHYKTtYKCc04VIEP53n7ea_agV9SBhLBfLSvSz4b5_8EYEJlYw6tfxsbg%253D%253D&a_bogus=mf8ZBR0kDDfBhfLg5I9LfY3qfJM3YBno0SVkMDgbcVV8Gy39HMO99exElYGvBrDjNs%252FDIeEjy4hCTNHMiOCyA3vXHWgKW9o%252F-ggmKl3hso0j53inCy8mE0ii-7sAtePQsvHlEKiQoXpHKm8h09oHmhK4b1dzFgf3qJLziE%253D%253D"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	req.Header.Add("authority", "buyin.jinritemai.com")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Add("cookie", cookie)
	req.Header.Add("referer", "https://buyin.jinritemai.com/dashboard/shopwindow/goods-list?pre_universal_page_params_id=&universal_page_params_id=a4b29656-9968-48a4-8db4-cb277632e523")
	req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"116\", \"Not)A;Brand\";v=\"24\", \"Google Chrome\";v=\"116\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		fmt.Printf("Error: status code is %d, expected 200\n", res.StatusCode)
		return nil, err
	}

	var jsonResponse Response
	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	if jsonResponse.Code != 0 {
		return nil, fmt.Errorf("error: API response code is %d, expected 0", jsonResponse.Code)
	}

	fmt.Println("GetWindow ====", string(body))

	var products []ProductDetail

	fmt.Println("GetWindow jsonResponse.Data.List:", jsonResponse.Data.List)

	for _, item := range jsonResponse.Data.List {
		product := ProductDetail{
			//EntityType:  deepGet(item, "$.entity_type").(string),
			PromotionID: getString(deepGet(item, "$.promotion_id")),
			ProductID:   getString(deepGet(item, "$.product_id")),
			Title:       getString(deepGet(item, "$.name")),
			Cover:       getString(deepGet(item, "$.cover")),
			//IsDisabled:      getInt(deepGet(item, "$.is_disabled")),
			//DetailURL:       getString(deepGet(item, "$.detail_url")),
			Price: getInt(deepGet(item, "$.price")),
			//PromotionStatus: getString(deepGet(item, "$.promotion_info.promotion_status")),
			//ItemType:        getString(deepGet(item, "$.promotion_info.item_type")),
			CosFee:   getFloat(deepGet(item, "$.cos_fee")),
			CosRatio: getInt(deepGet(item, "$.commission.cos_ratio")),
			//MonthlySale:     getInt(deepGet(item, "$.business_operation_info.monthly_sale")),
			//GoodRatio:       getFloat(deepGet(item, "$.business_operation_info.good_ratio")),
			ShopID:   getString(deepGet(item, "$.shop_info_data.shop_base_info.shop_id")),
			ShopName: getString(deepGet(item, "$.shop_info_data.shop_base_info.shop_name")),
			//ExpScore:        getFloat(deepGet(item, "$.base_shop_info.exp_score")),
		}
		fmt.Printf("item: %+v\n", item)
		products = append(products, product)
	}

	for _, product := range products {
		fmt.Printf("Product: %+v\n", product)
	}

	return products, nil
}

func AddWindow(douyinId string, productIdStr string) (*CheckDisplayTimeResp, error) {

	fileContents, err := getCookie(douyinId)
	if err != nil {
		return nil, err
	}
	// Convert file contents to string

	cookie := string(fileContents)
	fmt.Println("cookie:", cookie)

	url := "https://buyin.jinritemai.com/api/shop/bind/?verifyFp=verify_lyh1wlcm_7vVhuv21_rChj_40I5_8Tjk_a1HLz1zdEsUf&fp=verify_lyh1wlcm_7vVhuv21_rChj_40I5_8Tjk_a1HLz1zdEsUf&msToken=wpvKosiNYs9TX-MZgBQ7BsMf3z7_6H0wzjSlVM6L7M9tSzea--IGt8AEwfBXo-RwjtL9EvFUWOQso-47JL8KfZX_r90-KgD-8zJz2ohHjkIkn9cfaA5_JushqZKpZus%253D&a_bogus=QX80%252FVz6dkfskd8D5I9LfY3qfCFgYBrm0SVkMDgblaO8tg39HMY19exElKkvMsRjNs%252FDIeEjy4hST3BMiOCyA3vXHWgKW9o%252F-ggmKl3hso0j53inCy8mE0ii-7sAtePQsvHlEKimoweHKm8h09oHmhK4b1dzFgf3qJLz4E%253D%253D"

	payload := strings.NewReader(fmt.Sprintf(`{"pmts":[{"promotion_id":"%s","bind_source":25}],"hide_status":2}`, productIdStr))

	client := &http.Client{}

	method := "POST"
	req, err := http.NewRequest(method, url, payload)

	req.Header.Add("authority", "buyin.jinritemai.com")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cookie", cookie)
	req.Header.Add("origin", "https://buyin.jinritemai.com")
	req.Header.Add("referer", "https://buyin.jinritemai.com/dashboard/merch-picking-cart?btm_ppre=a10091.b089178.c0.d0&btm_pre=a10091.b24215.c809509.d0&btm_show_id=87ab81a2-65c9-4e49-8333-31f076337390&pre_universal_page_params_id=&universal_page_params_id=e596e1f6-2d0b-4c28-b8bf-08a8d780c6c7")
	req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"116\", \"Not)A;Brand\";v=\"24\", \"Google Chrome\";v=\"116\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36")
	req.Header.Add("x-secsdk-csrf-token", "0001000000012929cdff70e09f9f1bb1fe5729eb03af508fa72869d2dbdc9209012e52aa33a417e56b60b2459416")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	fmt.Println("AddWindow ====", string(body))
	ret := &CheckDisplayTimeResp{}
	_ = json.Unmarshal(body, ret)
	return ret, nil
}

func DeleteWindow(douyinId string, productIdStr string) (*CheckDisplayTimeResp, error) {

	fileContents, err := getCookie(douyinId)
	if err != nil {
		return nil, err
	}
	// Convert file contents to string

	cookie := string(fileContents)
	fmt.Println("cookie:", cookie)

	url := "https://buyin.jinritemai.com/api/anchor/shop/batch_unbind?verifyFp=verify_lyh1wlcm_7vVhuv21_rChj_40I5_8Tjk_a1HLz1zdEsUf&fp=verify_lyh1wlcm_7vVhuv21_rChj_40I5_8Tjk_a1HLz1zdEsUf&msToken=oXJivZOPi6ovPFkzKJxWMu-k0GPEnVSEhdkpfo3MQPbRsFdQe_6y507eVEvVjrO9YnPrCgn0rXp6Qq-FohSZCrfgubKYhVOYXpPG3LtuMc75wy6RNjbKpA%3D%3D&a_bogus=m6mh%2Ffh6dDIPffLX5IKLfY3qf73JYBrz0SVkMDgbBcO82639HMTI9exEA8XvyUDjNs%2FDIejjy4hCTNHMiOCyA3vXHWgKW9o%2F-ggmKl3hsom7-3intL0grUvq-hs1Sl925kp-EKigq7lHKRj2099c5kIlO6ZCcHgjxiSmtn3FvIW%3D"

	payload := strings.NewReader(fmt.Sprintf(`{"pmts":[{"promotion_id":"%s","bind_source":25}],"hide_status":2}`, productIdStr))

	client := &http.Client{}

	method := "POST"
	req, err := http.NewRequest(method, url, payload)

	//--data-raw '{"promotion_ids":["3683420290857218177"]}' \
	//--compressed

	req.Header.Add("authority", "buyin.jinritemai.com")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cookie", cookie)
	req.Header.Add("origin", "https://buyin.jinritemai.com")
	req.Header.Add("referer", "https://buyin.jinritemai.com/dashboard/shopwindow/goods-list?pre_universal_page_params_id=&universal_page_params_id=ad232f55-99c4-4733-9cd2-0c70684b164c")
	req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"116\", \"Not)A;Brand\";v=\"24\", \"Google Chrome\";v=\"116\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36")
	req.Header.Add("x-secsdk-csrf-token", "0001000000012929cdff70e09f9f1bb1fe5729eb03af508fa72869d2dbdc9209012e52aa33a417e56b60b2459416")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	fmt.Println("DeleteWindow ====", string(body))
	ret := &CheckDisplayTimeResp{}
	_ = json.Unmarshal(body, ret)
	return ret, nil
}

// Helper function to extract nested values from JSON
func deepGet(data interface{}, path string) interface{} {
	result, err := jsonpath.Get(path, data)
	if err != nil {
		return ""
	}
	return result
}

type Response struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Extra struct {
		Now   int64  `json:"now"`
		LogID string `json:"log_id"`
	} `json:"extra"`
	Data struct {
		List []interface{} `json:"list"`
	} `json:"data"`
}

type ProductDetail struct {
	//EntityType  string
	PromotionID string
	ProductID   string
	Title       string
	Cover       string
	//IsDisabled      int
	//DetailURL       string
	Price int
	//PromotionStatus string
	//ItemType        string
	CosFee   float64
	CosRatio int
	//MonthlySale     int
	//GoodRatio       float64
	ShopID   string
	ShopName string
	//ExpScore        float64
}

func getString(value interface{}) string {
	if str, ok := value.(string); ok {
		return str
	}
	return ""
}

func getInt(value interface{}) int {
	if num, ok := value.(float64); ok {
		return int(num)
	}
	return 0
}

func getFloat(value interface{}) float64 {
	if num, ok := value.(float64); ok {
		return num
	}
	return 0.0
}
