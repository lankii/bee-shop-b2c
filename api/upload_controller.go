package api

import (
	"cleverbamboo.com/bee-shop-b2c/common"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type UploadController struct {
	BaseController
}

// URLMapping ...
func (c *UploadController) URLMapping() {
	c.Mapping("Upload", c.Upload)
}

// @Title Upload image
// @router / [post]
func (c *UploadController) Upload() {
	// TODO 图片 source、large、medium、thumbnail
	fileData, fileHeader, err := c.GetFile("upload")
	if err != nil {
		c.ServerError(err)
		return
	}
	now := time.Now()

	fileType := "other"
	// 判断后缀为图片的文件，如果是图片我们才存入到数据库中
	fileExt := filepath.Ext(fileHeader.Filename)
	if fileExt == ".jpg" || fileExt == ".png" || fileExt == ".gif" || fileExt == ".jpeg" {
		fileType = "image"
	} else {
		c.JsonResult(common.GetHttpStatus("ok"), common.ErrOK, common.Fail, "仅支持 jpg/png/gif/jpeg 格式")
		return
	}
	// 文件夹路径
	fileDir := fmt.Sprintf("static/upload/%s/%d/%d/%d", fileType, now.Year(), now.Month(), now.Day())
	// ModePerm是0777，这样拥有该文件夹路径的执行权限
	err = os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		c.ServerError(err)
		return
	}
	// 文件路径
	timeStamp := time.Now().Unix()
	fileName := fmt.Sprintf("%d-%s", timeStamp, fileHeader.Filename)
	filePathStr := filepath.Join(fileDir, fileName)
	desFile, err := os.Create(filePathStr)
	if err != nil {
		c.ServerError(err)
		return
	}
	// 将浏览器客户端上传的文件拷贝到本地路径的文件里面
	_, err = io.Copy(desFile, fileData)
	if err != nil {
		c.ServerError(err)
		return
	}
	filePathStr = "/" + strings.Replace(filePathStr, "\\", "/", -1)

	type image struct {
		Uid      interface{} `json:"uid"`
		Name     string      `json:"name"`
		Status   string      `json:"status"`
		ThumbUrl string      `json:"thumbUrl"`
		Url      string      `json:"url"`
	}

	var imageList []image
	u1 := uuid.Must(uuid.NewV4())
	imageList = append(imageList, image{
		Uid:      u1,
		Name:     fileHeader.Filename,
		Status:   "done",
		ThumbUrl: filePathStr,
		Url:      filePathStr,
	})

	c.JsonResult(common.GetHttpStatus("created"), common.ErrOK, common.Success, imageList)
}
