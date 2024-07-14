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
	Key  string `json:"key"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

const (
	ProductTagClothes   = "clothes"
	ProductTagMakeups   = "makeups"
	ProductTagFood      = "food"
	ProductTagPregnancy = "pregnancy" // 婴孕
	ProductTagHealth    = "health"
	ProductTagTravel    = "travel"
	ProductTag3C        = "3c"
	ProductTagOutdoor   = "outdoor"
	ProductTagOther     = "other"
)

var ProductTagList = map[string]Tag{
	ProductTagClothes: {
		Key:  ProductTagClothes,
		Name: "服装",
		Desc: "男女装/内衣/饰品",
	},
	ProductTagMakeups: {
		Key:  ProductTagMakeups,
		Name: "美妆",
		Desc: "男女装/内衣/饰品",
	},
	ProductTagFood: {
		Key:  ProductTagFood,
		Name: "美食",
		Desc: "餐饮/零食/特产",
	},
	ProductTagPregnancy: {
		Key:  ProductTagPregnancy,
		Name: "婴孕",
		Desc: "婴孕用品/玩具",
	},
	ProductTagHealth: {
		Key:  ProductTagHealth,
		Name: "保健品",
		Desc: "养生/护理/保健",
	},
	ProductTagTravel: {
		Key:  ProductTagTravel,
		Name: "旅游",
		Desc: "5A/长途/海外",
	},
	ProductTag3C: {
		Key:  ProductTag3C,
		Name: "3C产品",
		Desc: "家电/数码/电脑/配件",
	},
	ProductTagOutdoor: {
		Key:  ProductTagOutdoor,
		Name: "户外运动",
		Desc: "帐篷/渔具",
	},
	ProductTagOther: {
		Key:  ProductTagOther,
		Name: "其他",
		Desc: "非遗/书籍/文化",
	},
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
