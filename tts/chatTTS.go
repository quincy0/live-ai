package tts

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/quincy0/live-ai/dto"
	"github.com/quincy0/live-ai/store"
	"github.com/quincy0/qpro/cryption"
	"github.com/quincy0/qpro/qHttp"
	"github.com/quincy0/qpro/qLog"
	"github.com/quincy0/qpro/qRedis"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type ChatRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		TaskId string `json:"task_id"`
	} `json:"data"`
}

func FindAudioURL(ctx context.Context, sum string) string {
	url, err := qRedis.Client.Get(ctx, "audio"+sum)
	if err != nil {
		if err != redis.Nil {
			qLog.Error("get audio from redis failed", zap.String("sum", sum), zap.Error(err))
		}
		return ""
	}
	return url
}

func CreateChatAudioV2(ctx context.Context, recreate int, spk string, text string) (string, error) {
	sum := cryption.Md5Encode([]byte(spk + text))
	if recreate == 0 {
		url := FindAudioURL(ctx, sum)
		if len(url) > 0 {
			return url, nil
		}
	}
	if qRedis.Client.Exists(ctx, "lock"+sum) {
		return "", nil
	}
	params := map[string]any{
		"spk_id":   spk,
		"text":     text,
		"callback": "https://console.wangqiong.vip/dh/audio/notify/" + sum,
	}
	data, err := qHttp.Post(
		ctx,
		"http://events.vicp.net/generate",
		params,
		qHttp.WithTimeout(600000),
	).Result()
	if err != nil {
		return "", err
	}
	var res ChatRes
	err = json.Unmarshal(data, &res)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Msg)
	}
	qRedis.Client.SetEx(ctx, "lock"+sum, "pending", 10*time.Minute)
	return "", nil
}

func CreateChatAudioNotify(ctx context.Context, sum string, params dto.AudioNotifyParam) error {
	if err := qRedis.Client.Del(ctx, "lock"+sum); err != nil {
		return err
	}
	if params.Code != 200 {
		return nil
	}
	return qRedis.Client.Set(ctx, "audio"+sum, params.Data.VideoUrl, 0)
}

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
		qLog.Info("create is locking", zap.String("text", text))
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
