package file_api

import (
	"fast_gin/global"
	"fast_gin/utils/find"
	"fast_gin/utils/res"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var whiteList = []string{
	".pdf",
}

func (FileApi) UploadView(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		res.FailWithMsg("请选择文件", c)
		return
	}

	// 大小限制
	if fileHeader.Size > 20*1024*1024 {
		res.FailWithMsg("上传文件过大", c)
		return
	}

	// 后缀判断
	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	if !find.InList(whiteList, ext) {
		res.FailWithMsg("不支持该文件上传", c)
		return
	}
	basename := strings.TrimSuffix(fileHeader.Filename, ext)
	newFilename := basename + "_" + time.Now().Format("20060102150405") + "_" + uuid.NewString() + ext

	filePath := path.Join("uploads", global.Config.Upload.Dir, newFilename)

	c.SaveUploadedFile(fileHeader, filePath)
	res.Ok("/"+filePath, "上传成功", c)

}
