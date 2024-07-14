package audioService

import (
	"context"

	"github.com/quincy0/live-ai/dto"
	"github.com/quincy0/live-ai/service/consoleService"
	"github.com/quincy0/live-ai/tts"
	"github.com/quincy0/live-ai/util"
	"github.com/quincy0/qpro/qRoutine"
)

func AudioList(ctx context.Context, roomId string) (*dto.RoomData, error) {
	roomData, err := consoleService.ParseRoomData(ctx, roomId)
	if err != nil {
		return nil, err
	}
	for _, script := range roomData.ScriptList {
		for _, scene := range script.SceneList {
			if url := tts.FindAudioURL(ctx, scene.Sum); len(url) > 0 {
				scene.Audio = url
			} else {
				scene.Audio = ""
				qRoutine.GoSafe(func() {
					ctxNew := util.InitContextWithSameTrace(ctx)
					_, _ = tts.CreateChatAudioV2(ctxNew, 0, roomData.TTS.Name, scene.Content)
				})
			}
		}
	}
	return roomData, nil
}
