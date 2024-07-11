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

type ProductTag struct {
	Key  string `json:"key"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

var ProductTagList = map[string]ProductTag{
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
