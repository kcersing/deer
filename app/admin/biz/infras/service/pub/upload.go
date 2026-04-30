package pub

import (
	"context"
	"fmt"

	"github.com/astaxie/beego/httplib"
	"github.com/cloudwego/hertz/pkg/app"
)

// https://sjqzhang.github.io/go-fastdfs/api.html

// 上传
func Upload(ctx context.Context, c *app.RequestContext) {
	var obj interface{}
	req := httplib.Post("http://127.0.0.1:9109/group/upload")
	//file:上传的文件
	req.PostFile("file", "filename")
	//output:输出
	req.Param("output", "json")
	//scene:场景
	req.Param("scene", "img")
	//path:自定义路径
	req.Param("path", "")
	req.ToJSON(&obj)

	fmt.Print(obj)
}

// 删除
func Delete(ctx context.Context, c *app.RequestContext) {
	var obj interface{}
	req := httplib.Post("http://127.0.0.1:9109/group/delete")
	//path:路径
	req.Param("path", "")
	req.ToJSON(&obj)

	fmt.Print(obj)
}

// 文件信息
func getFileInfo(ctx context.Context, c *app.RequestContext) {
	var obj interface{}
	req := httplib.Get("http://127.0.0.1:9109/group/get_file_info")
	//path:路径
	req.Param("path", "")
	req.ToJSON(&obj)

	fmt.Print(obj)
}

// 列表
func listDir(ctx context.Context, c *app.RequestContext) {

	ctx.Deadline()
	var obj interface{}
	req := httplib.Get("http://127.0.0.1:9109/group/get_file_info")
	//dir:路径
	req.Param("dir", "")
	req.ToJSON(&obj)

	fmt.Print(obj)
}
