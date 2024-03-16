package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"mime/multipart"
	"goproject/global"
	"goproject/utils/rspcode"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
)

func UpLoadFile(file multipart.File, filesize int64) (string, int) {
	var (
		// Accesskey = global.Config.System.Accesskey
		// Secretkey = global.Config.System.SecretKey
		Bucket = global.Config.System.Bucket
		ImgUrl = global.Config.System.Qiniuserver
		Mac    = auth.New(global.Config.System.Accesskey, global.Config.System.SecretKey)
	)

	PutPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	uptoken := PutPolicy.UploadToken(Mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuabei,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}

	PutExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err := formUploader.PutWithoutKey(context.Background(), &ret, uptoken, file, filesize, &PutExtra)
	if err != nil {
		return "", rspcode.ERROR
	}
	url := ImgUrl + "/"+ret.Key
	return url, rspcode.SUCCESS
}

func UpLoadBase64Image(base64Data string) (string, int) {
	var (
		Bucket = global.Config.System.Bucket
		ImgUrl = global.Config.System.Qiniuserver
		Mac    = auth.New(global.Config.System.Accesskey, global.Config.System.SecretKey)
	)

	// 解码 base64 编码的图片数据
	imageData, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return "", rspcode.ERROR
	}

	PutPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	uptoken := PutPolicy.UploadToken(Mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuabei,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}

	PutExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err = formUploader.PutWithoutKey(context.Background(), &ret, uptoken, bytes.NewReader(imageData), int64(len(imageData)), &PutExtra)
	if err != nil {
		return "", rspcode.ERROR
	}

	url := ImgUrl + ret.Key
	return url, rspcode.SUCCESS
}
