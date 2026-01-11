package minio

import (
	"common/pkg/errno"
	"common/pkg/utils"
	"context"
	"path"
	"strings"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/minio/minio-go/v7"
)

var imgSuffixSet = map[string]struct{}{
	"jpg": {}, "png": {}, "gif": {},
	"raw": {}, "heic": {}, "jpeg": {},
	"heif": {}, "cr3": {}, "nef": {},
	"bmp": {}, "tiff": {}, "exif": {},
	"wbmp": {}, "mbm": {},
}
var videoSuffixSet = map[string]struct{}{
	"mp4": {}, "avi": {}, "wmv": {},
	"w4v": {}, "asf": {}, "flv": {},
	"rmvb": {}, "rm": {}, "3gp": {},
	"vob": {}, "wma": {}, "mpeg": {},
	"mpg": {}, "mov": {},
}

func getFileSuffix(fileName string) (suffix string, err error) {
	lastDotIndex := strings.LastIndex(fileName, ".")
	if lastDotIndex < 0 {
		return "", errno.NewErrNo(errno.Unauthorized, "missing suffix")
	}
	suffix = fileName[lastDotIndex+1:]
	suffix = strings.ToLower(suffix)
	return suffix, nil
}

func UpLoadImg(ctx context.Context, c *app.RequestContext, bucketName string) (uploadinfo minio.UploadInfo, err error) {
	file, err := c.FormFile("files")
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	suffix, err := getFileSuffix(file.Filename)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
	}
	if _, ok := imgSuffixSet[suffix]; !ok {
		utils.SendResponse(c, errno.NewErrNo(errno.Unauthorized, "invalid image suffix"), nil, 0, "")
		return
	}

	sf, err := snowflake.NewNode(6)
	if err != nil {

		utils.SendResponse(c, errno.NewErrNo(errno.Unauthorized, err.Error()), nil, 0, "")
		return
	}
	nowTime := time.Now()
	filename := sf.Generate().String()
	dateName := nowTime.Format("2006/01/02")
	file.Filename = dateName + "/" + filename + path.Ext(file.Filename)

	uploadinfo, err = NewManager(bucketName).PutObjectURL(ctx, bucketName, file)

	return uploadinfo, err

}
