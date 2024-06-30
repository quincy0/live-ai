package audioService

import (
	"context"

	"github.com/quincy0/live-ai/dto"
	"github.com/quincy0/live-ai/service/consoleService"
	"github.com/quincy0/live-ai/store"
	"github.com/quincy0/live-ai/tts"
	"github.com/quincy0/qpro/qRoutine"
)

func AudioList(ctx context.Context, roomId string) (*dto.RoomData, error) {
	roomData, err := consoleService.ParseRoomData(ctx, roomId)
	if err != nil {
		return nil, err
	}
	for _, script := range roomData.ScriptList {
		for _, scene := range script.SceneList {
			fileName := scene.Sum + ".wav"
			if store.IsExist(ctx, fileName) {
				scene.Audio = store.FilePrefix + fileName
			} else {
				scene.Audio = ""
				qRoutine.GoSafe(func() {
					tts.CreateAudio(ctx, 0, roomData.TTS.Name, scene.Content)
				})
			}
		}
	}
	return roomData, nil
}
