package consoleService

import (
	"context"
	"encoding/json"

	"github.com/quincy0/live-ai/dto"
	"github.com/quincy0/qpro/cryption"
	"github.com/quincy0/qpro/qHttp"
	"github.com/quincy0/qpro/qLog"
	"go.uber.org/zap"
)

const (
	Domain      = "http://101.126.93.55:8080"
	LiveRoomURL = "/inner/room/detail/"
)

func GetRoomDetailData(ctx context.Context, roomId string) (*RoomDetail, error) {
	url := Domain + LiveRoomURL + roomId
	resp, err := qHttp.Get(ctx, url, nil, qHttp.WithTimeout(6000)).Result()
	if err != nil {
		qLog.TraceError(ctx, "do request failed", zap.String("roomId", roomId), zap.Error(err))
		return nil, err
	}

	var data RoomDetailResponse
	err = json.Unmarshal(resp, &data)
	if err != nil {
		qLog.TraceError(ctx, "unmarshal failed", zap.Error(err))
		return nil, err
	}
	return data.Data, nil
}

func ParseRoomData(ctx context.Context, roomId string) (*dto.RoomData, error) {
	detail, err := GetRoomDetailData(ctx, roomId)
	if err != nil {
		return nil, err
	}
	tts := &dto.TTS{
		Name: detail.TtsConfig.Name,
	}
	scriptList := make([]*dto.Script, len(detail.ScriptList))

	for scriptIndex, script := range detail.ScriptList {
		scriptInfo := &dto.Script{
			ScriptId:  script.RoomScriptId,
			SceneList: make([]*dto.Scene, len(script.ScriptConfig.Scenes)),
		}
		for sceneIndex, scene := range script.ScriptConfig.Scenes {
			sum := cryption.Md5Encode([]byte(tts.Name + scene.Text.Content))

			sceneInfo := &dto.Scene{
				SceneId: scene.SceneId,
				Name:    scene.Text.Name,
				Content: scene.Text.Content,
				Sum:     sum,
				Audio:   "",
			}
			scriptInfo.SceneList[sceneIndex] = sceneInfo
		}
		scriptList[scriptIndex] = scriptInfo
	}

	return &dto.RoomData{
		TTS:        tts,
		ScriptList: scriptList,
	}, nil
}
