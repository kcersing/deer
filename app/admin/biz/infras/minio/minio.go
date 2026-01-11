package minio

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"
	"net/url"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	Client *minio.Client
	err    error
)

func initMinio(cfg *MinioConfig) {
	Client, err = minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		klog.Fatalf("minio.New failed, err: %v", err)
	}
	initMakeBucket(context.Background(), cfg.ImgBucket)
	initMakeBucket(context.Background(), cfg.VideoBucket)
}

// initMakeBucket(context.Background(), cfg.Bucket)
func initMakeBucket(ctx context.Context, bucketName string) {
	hlog.Info(bucketName)
	exists, err := Client.BucketExists(ctx, bucketName)

	if err != nil {
		klog.Fatalf("minio.BucketExists failed, err: %v", err)
		return
	}

	if !exists {
		err = Client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	policy := `{"Version": "2012-10-17","Statement": [{"Action": ["s3:GetObject"],"Effect": "Allow","Principal": {"AWS": ["*"]},"Resource": ["arn:aws:s3:::` + bucketName + `/*"],"Sid": ""}]}`
	err = Client.SetBucketPolicy(context.Background(), bucketName, policy)
	if err != nil {
		klog.Fatal("set bucket policy err:%s", err)
		return
	}
}

type Manager struct {
	bucketName string
	client     *minio.Client
}

func NewManager(bucketName string) *Manager {
	return &Manager{bucketName: bucketName, client: Client}
}

func (s *Manager) GetObjectURL(ctx context.Context, objectName string, timeOut time.Duration) (string, error) {
	url, err := s.client.PresignedGetObject(ctx, s.bucketName, objectName, timeOut, nil)
	if err != nil {
		return "", err
	}
	return url.String(), err
}

func (s *Manager) PutObjectURL(ctx context.Context, objectName string, timeOut time.Duration) (string, error) {
	url, err := s.client.PresignedPutObject(ctx, s.bucketName, objectName, timeOut)
	if err != nil {
		return "", err
	}
	return url.String(), nil
}

// PutToBucket put the file into the bucket by *multipart.FileHeader
func (s *Manager) PutToBucket(ctx context.Context, bucketName string, file *multipart.FileHeader) (info minio.UploadInfo, err error) {
	fileObj, _ := file.Open()
	info, err = s.client.PutObject(ctx, bucketName, file.Filename, fileObj, file.Size, minio.PutObjectOptions{})
	fileObj.Close()
	return info, err
}

// GetObjURL get the original link of the file in minio
func (s *Manager) GetObjURL(ctx context.Context, bucketName, filename string) (u *url.URL, err error) {
	exp := time.Hour * 24
	reqParams := make(url.Values)
	u, err = s.client.PresignedGetObject(ctx, bucketName, filename, exp, reqParams)
	return u, err
}

// PutToBucketByBuf put the file into the bucket by *bytes.Buffer
func (s *Manager) PutToBucketByBuf(ctx context.Context, bucketName, filename string, buf *bytes.Buffer) (info minio.UploadInfo, err error) {
	info, err = s.client.PutObject(ctx, bucketName, filename, buf, int64(buf.Len()), minio.PutObjectOptions{})
	return info, err
}

// PutToBucketByFilePath put the file into the bucket by filepath
func (s *Manager) PutToBucketByFilePath(ctx context.Context, bucketName, filename, filepath string) (info minio.UploadInfo, err error) {
	info, err = s.client.FPutObject(ctx, bucketName, filename, filepath, minio.PutObjectOptions{})
	return info, err
}
