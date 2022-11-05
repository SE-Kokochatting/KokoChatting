package controller

import (
	"KokoChatting/global"
	"KokoChatting/model/res"
	_ "KokoChatting/service"
	"context"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

type UploadController struct {
	baseController
}

const (
	bucket = "kokochatting"
	accessKey = "Y-K-DeCCev-2utgn7gLnOtwjS4G_j1j7bIm8SubU"
	secretKey = "DVytnQd7FYXpjaJ6-00fL4gJQi7FsR6o_DJrv84Y"
	imgUrl = "http://rkno5r6k7.hn-bkt.clouddn.com/"
)
// Login [POST]
// PATH: api/v1/upload
// Function: 上传一张图片，因为使用表单上传，需要前端将文件名指定为file，且大小需小于4M
// Feat: 可以采用分片上传的方式。但是这里暂时不考虑
func (uploadCtl *UploadController) Upload(c *gin.Context) {
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone : &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS: false,
	}

	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		global.Logger.Error("get file error", zap.Error(err))
		uploadCtl.WithErr(global.GetFileError, c)
		return
	}
	fileSize := fileHeader.Size

	putExtra := storage.PutExtra{}
	formUploder := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err = formUploder.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)

	if err != nil {
		global.Logger.Error("upload picture error", zap.Error(err))
		uploadCtl.WithErr(global.UploadPictureError, c)
		return
	}
	url := imgUrl + ret.Key

	uploadRes := &res.UploadPictureRes{
		Data: struct {
			Url       string `json:"url"`
		}{
			Url:       url,
		},
	}
	uploadCtl.WithData(uploadRes, c)
}

func NewUploadController() *UploadController {
	return &UploadController{}
}