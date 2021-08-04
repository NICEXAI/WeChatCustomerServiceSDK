package WeChatCustomerServiceSDK

import (
	"encoding/json"
	"fmt"
	"github.com/NICEXAI/WeChatCustomerServiceSDK/util"
	"mime/multipart"
)

const (
	// 上传临时素材
	mediaUploadAddr = "https://qyapi.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=%s"
	//获取临时素材
	mediaGetAddr = "https://qyapi.weixin.qq.com/cgi-bin/media/get?access_token=%s&media_id=%s"
)

// MediaUploadOptions 上传临时素材请求参数
type MediaUploadOptions struct {
	//上传文件类型
	Type string `json:"type"`							// 媒体文件类型，分别有图片（image）、语音（voice）、视频（video），普通文件（file）
	//文件名
	FileName string `json:"fileName"`					// 文件名
	//文件大小
	FileSize int64 `json:"fileSize"`					// 文件大小
	//文件内容
	File multipart.File									// 文件内容
}

// MediaUploadSchema 上传临时素材响应内容
type MediaUploadSchema struct {
	BaseModel
	Type string `json:"type"`							// 媒体文件类型，分别有图片（image）、语音（voice）、视频（video），普通文件(file)
	MediaID string `json:"media_id"`					// 媒体文件上传后获取的唯一标识，3天内有效
	CreatedAt string `json:"created_at"`				// 媒体文件上传时间戳
}

// MediaUpload 上传临时素材
//上传的媒体文件限制
//所有文件size必须大于5个字节
//图片（image）：2MB，支持JPG,PNG格式
//语音（voice） ：2MB，播放长度不超过60s，仅支持AMR格式
//视频（video） ：10MB，支持MP4格式
//普通文件（file）：20MB
func (r *Client) MediaUpload(options MediaUploadOptions) (info MediaUploadSchema, err error) {
	fileOptions := util.FileOptions{
		FileName: options.FileName,
		FileSize: options.FileSize,
		File:     options.File,
	}
	target := fmt.Sprintf(mediaUploadAddr, r.accessToken, options.Type)
	r.recordUpdate(target)

	data, err := util.HttpPostFile(target, fileOptions)
	if err != nil {
		return info, err
	}
	_ = json.Unmarshal(data, &info)
	fmt.Println(string(data))
	if info.ErrCode != 0 {
		return info, NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}

// MediaOriginUpload 上传临时素材
//上传的媒体文件限制
//所有文件size必须大于5个字节
//图片（image）：2MB，支持JPG,PNG格式
//语音（voice） ：2MB，播放长度不超过60s，仅支持AMR格式
//视频（video） ：10MB，支持MP4格式
//普通文件（file）：20MB
func (r *Client) MediaOriginUpload(fileName, fileType string, size int, body []byte) (info MediaUploadSchema, err error) {
	target := fmt.Sprintf(mediaUploadAddr, r.accessToken, fileType)
	r.recordUpdate(target)

	data, err := util.HttpPostOriginFile(target, fileName, size, body)
	if err != nil {
		return info, err
	}
	_ = json.Unmarshal(data, &info)
	fmt.Println(string(data))
	if info.ErrCode != 0 {
		return info, NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}

// MediaGet 获取临时素材
func (r *Client) MediaGet(mediaID string) string {
	return fmt.Sprintf(mediaGetAddr, r.accessToken, mediaID)
}