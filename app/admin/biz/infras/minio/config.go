package minio

type MinioConfig struct {
	Endpoint        string `json:"endpoint"`
	AccessKeyID     string `json:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key"`
	ImgBucket       string `json:"img_bucket"`
	VideoBucket     string `json:"video_bucket"`
	UrlPrefix       string `json:"url_prefix"`
}
