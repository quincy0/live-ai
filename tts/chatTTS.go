package tts

import (
	"context"
	"errors"
	"time"

	"github.com/quincy0/live-ai/store"
	"github.com/quincy0/qpro/cryption"
	"github.com/quincy0/qpro/qHttp"
	"github.com/quincy0/qpro/qLog"
	"github.com/quincy0/qpro/qRedis"
	"go.uber.org/zap"
)

// http://events.vicp.net/?spk=aeyes&text=你好文本
func CreateChatAudio(ctx context.Context, recreate int, spk string, text string) (string, error) {
	sum := cryption.Md5Encode([]byte(spk + text))
	fileName := sum + ".wav"
	if recreate == 0 && store.IsExist(ctx, fileName) {
		return store.FilePrefix + fileName, nil
	}
	params := map[string]any{
		"spk_id": spk,
		"text":   text,
	}
	data, err := qHttp.Post(
		ctx,
		"http://events.vicp.net",
		params,
		qHttp.WithTimeout(600000),
	).Result()
	if err != nil {
		return "", err
	}
	err = store.Upload(ctx, fileName, string(data))
	if err != nil {
		return "", err
	}
	return store.FilePrefix + fileName, nil
}

func CreateAudio(ctx context.Context, recreate int, spkId, text string) (string, error) {
	sum := cryption.Md5Encode([]byte(spkId + text))
	lockKey := "lock" + sum
	isSucc, err := qRedis.Client.SetNX(ctx, lockKey, "1", 10*time.Minute)
	if err != nil {
		return "", err
	}
	if !isSucc && recreate == 0 {
		return "", errors.New("音频生成中，请稍后尝试")
	}
	fileURL, err := CreateChatAudio(ctx, recreate, spkId, text)

	qRedis.Client.Del(ctx, lockKey)
	if err != nil {
		qLog.Error("create audio error", zap.Error(err))
		return "", err
	}
	qLog.Info("create audio success", zap.String("fileURL", fileURL))
	qRedis.Client.Set(ctx, sum, fileURL, 0)
	return fileURL, nil
}
