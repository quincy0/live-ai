package consoleService

import (
	"time"
)

type RoomDetailResponse struct {
	Code int         `json:"code"`
	Data *RoomDetail `json:"data"`
}

type RoomDetail struct {
	RoomId         string `json:"roomId"`   //直播间ID
	RoomType       int    `json:"roomType"` // 直播间类型 1：文本 2：音频
	RoomName       string `json:"roomName"` //主播名称
	UserId         int64  `json:"userId"`
	PlayOrder      int    `json:"playOrder"` // 1 随机
	LivingPlatform string `json:"livingPlatform"`
	PlatformLink   string `json:"platformLink"`                       // 平台链接
	PlatformUserId string `json:"platformUserId"`                     //平台用户标识
	AnchorId       int    `json:"anchorId,string" binding:"required"` //主播ID
	AnchorName     string `json:"anchorName" binding:"required"`      //主播名称
	AnchorModel    string `json:"anchorModel" binding:"required"`     //主播模型名称
	AnchorSource   int    `json:"anchorSource"`
	TtsSource      int    `json:"ttsSource"`
	TtsId          string `json:"tts_id"`
	Platform       int    `json:"platform" binding:"required"`   //平台：1、tiktok
	Tone           string `json:"tone" binding:"required"`       //音色
	Intonation     string `json:"intonation" binding:"required"` //语调
	Speed          string `json:"speed" binding:"required"`      //语速
	Volume         string `json:"volume" binding:"required"`     //音量
	Lang           string `json:"lang" binding:"required"`       //语言
	TtsConfig      struct {
		Volume      string `json:"vol"`
		Intonation  string `json:"pit"`
		Speed       string `json:"spd"`
		Tone        string `json:"tone"`
		Lang        string `json:"lang"`
		Gesture     string `json:"gesture"`
		ShowName    string `json:"showName"`
		TtsSource   int    `json:"source"`
		TtsId       string `json:"ttsId"`
		Name        string `json:"name"`
		Style       string `json:"style"`
		Styledegree string `json:"styledegree"`
		Role        string `json:"role"`
	} `json:"ttsConfig"`
	Width                  int                      `json:"width" binding:"required"`  //直播间宽
	Height                 int                      `json:"height" binding:"required"` //直播间高
	Status                 int                      `json:"status"`                    //直播间状态：1、待直播 2、直播中
	CoverUrl               string                   `json:"coverUrl"`                  //封面链接
	ItemCount              int                      `json:"itemCount"`                 //商品数量
	LiveDuration           string                   `json:"liveDuration"`              //直播时长
	ScriptDuration         string                   `json:"scriptDuration"`            //剧本时长
	VoiceOnly              bool                     `json:"voiceOnly"`                 //只用声音
	AutoRemoveOpen         bool                     `json:"autoRemoveOpen"`            //只用声音
	CreatedAt              time.Time                `json:"createdAt"`                 //创建时间
	ScriptList             []*GetRoomScriptResponse `json:"scriptList"`                //剧本列表
	Interaction            []InteractionItem        `json:"interaction"`               //互动配置
	IsOpen                 int                      `json:"isOpen"`
	OnlyAnswerCurrentGoods int                      `json:"onlyAnswerCurrentGoods"`
	GlobalAnswerInterval   int                      `json:"globalAnswerInterval"` //所有意向时间间隔
	SameAnswerInterval     int                      `json:"sameAnswerInterval"`   //同意向时间价格
	AutoMountGoods         bool                     `json:"autoMountGoods"`       //是否自动挂载商品
}

type InteractionItem struct {
	ActionType int `json:"actionType"` // 类型
	Frequency  int `json:"frequency"`  // 频率
	Enable     int `json:"enable"`     //  是否开启
}

type GetRoomScriptResponse struct {
	RoomScriptId  string        `json:"roomScriptId"`  //直播间剧本ID
	ScriptId      string        `json:"scriptId"`      //剧本ID
	GoodsId       string        `json:"goodsId"`       //三方平台商品ID
	ScriptName    string        `json:"scriptName"`    //剧本ID
	ScriptType    int           `json:"scriptType"`    //剧本类型：1、商品剧本 2、通用剧本
	CoverImageUrl string        `json:"coverImageUrl"` //封面图链接
	ScriptConfig  *ScriptConfig `json:"ScriptConfig"`  //剧本配置
}

type ScriptConfig struct {
	AnswerInterval int      `json:"answerInterval"`            //问答间隔
	Faqs           []*Faq   `json:"faqs" binding:"required"`   //问答库
	Scenes         []*Scene `json:"scenes" binding:"required"` //分镜
}

type Faq struct {
	Intention string   `json:"intention" binding:"required"` //意图
	Questions []string `json:"questions" binding:"required"` //问题列表
	Answers   []string `json:"answers" binding:"required"`   //答案列表
}

type Scene struct {
	SceneId       string   `json:"sceneId"`
	Text          *Text    `json:"text" binding:"required"` //文案
	Audios        *Audios  `json:"audios"`
	Images        []*Image `json:"images"`        //图片
	CoverImageUrl string   `json:"coverImageUrl"` //封面图
}

type Text struct {
	Name    string `json:"name" binding:"required"`    //文案名称
	Content string `json:"content" binding:"required"` //文案内容
	Active  bool   `json:"active"`                     //是否选中
}

type Audios struct {
	AudioId string `json:"audioId"`
	Name    string `json:"name"`
	URL     string `json:"url"`
}

const (
	UnknownImageType = iota
	BackgroundImageType
	AnchorImageType
	ChartletImageType
)

type Image struct {
	MaterialId   int64   `json:"materialId,string" binding:"required"` //素材ID
	ImageType    int     `json:"ImageType" binding:"required"`         //图层类型：1、背景 2、主播 3、贴图
	ImageSubType int     `json:"imageSubType"`                         // 贴图类型：1、静图 2、动图 3、视频 0、历史遗留，理论为静图
	Url          string  `json:"url" binding:"required"`               //图片链接
	Width        int     `json:"width" binding:"required"`             //宽
	Height       int     `json:"height" binding:"required"`            //高
	Top          int     `json:"top" binding:"required"`               //上边距
	Left         int     `json:"left" binding:"required"`              //左边距
	Size         float32 `json:"size"`
}
