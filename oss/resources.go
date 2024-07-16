package oss

// Resource OSS 资源
type Resource struct {
	Name    string // 文件名
	Url     string // 文件 URL
	Content string // 文件内容
}

func NewOssResource(name, url, content string) *Resource {
	return &Resource{
		Name:    name,
		Url:     url,
		Content: content,
	}
}

// Upload2Oss 上传文件到 OSS
func (r *Resource) Upload2Oss() {

}
