package dto

type RoomData struct {
	TTS        *TTS      `json:"TTS"`
	ScriptList []*Script `json:"scriptList"`
}

type TTS struct {
	Name string `json:"name"`
}

type Script struct {
	ScriptId  string   `json:"scriptId"`
	SceneList []*Scene `json:"sceneList"`
}

type Scene struct {
	SceneId string `json:"sceneId"`
	Name    string `json:"name"`
	Content string `json:"content"`
	Sum     string `json:"sum"`
	Audio   string `json:"audio"`
}
