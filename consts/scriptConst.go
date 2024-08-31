package consts

const (
	TimbreMing     = "xiaoming"
	TimbreYun      = "yunxi"
	TimbreZhaoYuan = "zhaoyuan"
)

var TimbreList = map[string]string{
	TimbreMing:     "小明",
	TimbreYun:      "云汐",
	TimbreZhaoYuan: "兆远",
}

type Tag struct {
	Key         string `json:"key"`
	Name        string `json:"name"`
	Desc        string `json:"desc"`
	Value       string `json:"value"`
	IndustryId  string `json:"industry_id"`
	Category    string `json:"category"`
	all_cate_id string `json:"all_cate_id"`
}

const (
	ProductTagClothes   = "clothes"
	ProductTagMakeups   = "makeups"
	ProductTagFood      = "food"
	ProductTagPregnancy = "pregnancy" // 婴孕
	ProductTagTravel    = ""
	ProductTagClean     = "clean"
	ProductTagFootwear  = "footwear"
	ProductTagWatch     = "watch"
	ProductTagBook      = "book"
	ProductTag3C        = "3c"
	ProductTagOutdoor   = "outdoor"
	ProductTagToy       = "toy"
	ProductTagFresh     = "fresh"
)

var ProductTagList = map[string]Tag{
	ProductTagClothes: {
		Key:         ProductTagClothes,
		Name:        "服饰内衣",
		Desc:        "男女装/内衣/饰品",
		Value:       "4",
		IndustryId:  "4",
		Category:    "4",
		all_cate_id: "all",
	},
	ProductTagMakeups: {
		Key:         ProductTagMakeups,
		Name:        "美妆",
		Desc:        "美容/护肤/彩妆/香水",
		Value:       "9",
		IndustryId:  "9",
		Category:    "9",
		all_cate_id: "all",
	},
	ProductTagFood: {
		Key:         ProductTagFood,
		Name:        "食品饮料",
		Desc:        "零食/营养滋补/粮油米面",
		Value:       "13",
		IndustryId:  "13",
		Category:    "13",
		all_cate_id: "all",
	},
	ProductTagPregnancy: {
		Key:         ProductTagPregnancy,
		Name:        "母婴宠物",
		Desc:        "婴孕用品/宠物用品",
		Value:       "10",
		IndustryId:  "10",
		Category:    "10",
		all_cate_id: "all",
	},
	ProductTagTravel: {
		Key:         ProductTagTravel,
		Name:        "智能家具",
		Desc:        "家具/家装/厨房/百货",
		Value:       "7",
		IndustryId:  "7",
		Category:    "7",
		all_cate_id: "all",
	},
	ProductTagClean: {
		Key:         ProductTagClean,
		Name:        "个护家清",
		Desc:        "洗护清洁/个人护理",
		Value:       "5",
		IndustryId:  "5",
		Category:    "5",
		all_cate_id: "all",
	},
	ProductTagFootwear: {
		Key:         ProductTagFootwear,
		Name:        "鞋靴箱包",
		Desc:        "箱包/男鞋/女鞋",
		Value:       "16",
		IndustryId:  "16",
		Category:    "16",
		all_cate_id: "all",
	},
	ProductTagWatch: {
		Key:         ProductTagWatch,
		Name:        "钟表配饰",
		Desc:        "钟表/配饰/饰品",
		Value:       "19",
		IndustryId:  "19",
		Category:    "19",
		all_cate_id: "all",
	},
	ProductTagBook: {
		Key:         ProductTagBook,
		Name:        "图书教育",
		Desc:        "学习用品/书籍杂志",
		Value:       "15",
		IndustryId:  "15",
		Category:    "15",
		all_cate_id: "all",
	},
	ProductTag3C: {
		Key:         ProductTag3C,
		Name:        "3C数码产品",
		Desc:        "家电/数码/电脑/配件",
		Value:       "14",
		IndustryId:  "14",
		Category:    "14",
		all_cate_id: "all",
	},
	ProductTagOutdoor: {
		Key:         ProductTagOutdoor,
		Name:        "户外运动",
		Desc:        "运动服/户外/健身",
		Value:       "2",
		IndustryId:  "2",
		Category:    "2",
		all_cate_id: "all",
	},
	//ProductTagOther: {
	//	Key:  ProductTagOther,
	//	Name: "其他",
	//	Desc: "非遗/书籍/文化",
	//},
	ProductTagToy: {
		Key:         ProductTagToy,
		Name:        "玩具",
		Desc:        "玩具/乐器/模玩游戏",
		Value:       "2",
		IndustryId:  "2",
		Category:    "2",
		all_cate_id: "all",
	},
	ProductTagFresh: {
		Key:         ProductTagFresh,
		Name:        "生鲜",
		Desc:        "生鲜/水果",
		Value:       "8",
		IndustryId:  "8",
		Category:    "8",
		all_cate_id: "all",
	},

	//"key": "智能家具",
	//"value": "7",
	//"desc": "",
	//"industry_id": "7",
	//"category": "7",
	//"all_cate_id": "all",
}

const (
	ScriptTagPrologue    = "ST1"
	ScriptTagLiveRoom    = "ST2"
	ScriptTagPrice       = "ST3"
	ScriptTagVoucher     = "ST4"
	ScriptTagPromotion   = "ST5"
	ScriptTagFocus       = "ST6"
	ScriptTagInviteEnter = "ST7"
	ScriptTagInviteWait  = "ST8"
	ScriptTagInviteTime  = "ST9"
	ScriptTagEnd         = "ST10"
)

var ScriptTagList = []Tag{
	{
		Key:  ScriptTagPrologue,
		Name: "开场白",
		Desc: "开场白",
	},
	{
		Key:  ScriptTagLiveRoom,
		Name: "直播间介绍",
		Desc: "直播间介绍",
	},
	{
		Key:  ScriptTagPrice,
		Name: "商品价格优势讲解",
		Desc: "商品价格优势讲解",
	},
	{
		Key:  ScriptTagVoucher,
		Name: "达人券讲解",
		Desc: "达人券讲解",
	},
	{
		Key:  ScriptTagPromotion,
		Name: "促单",
		Desc: "促单",
	},
	{
		Key:  ScriptTagFocus,
		Name: "求关注",
		Desc: "求关注",
	},
	{
		Key:  ScriptTagInviteEnter,
		Name: "拉入场（直播间外停留用户）",
		Desc: "拉入场（直播间外停留用户）",
	},
	{
		Key:  ScriptTagInviteWait,
		Name: "拉停留（直播间外停留用户）",
		Desc: "拉停留（直播间外停留用户）",
	},
	{
		Key:  ScriptTagInviteTime,
		Name: "拉时长（直播间外停留用户）",
		Desc: "拉时长（直播间外停留用户）",
	},
	{
		Key:  ScriptTagEnd,
		Name: "结束语",
		Desc: "结束语",
	},
}

type RoomTemplate struct {
	TemplateId int    `json:"templateId"`
	Name       string `json:"name"`
	List       []Tag  `json:"list"`
}

var RoomTemplateList = map[int]RoomTemplate{
	1: {
		TemplateId: 1,
		Name:       "破价法",
		List:       ScriptTagList,
	},
}
