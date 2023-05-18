package ctype

import "encoding/json"

type ImageType int

const (
	Local   ImageType = 1 // 本地
	Tencent ImageType = 2 // 七牛云
)

func (s ImageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s ImageType) String() string {
	switch s {
	case Local:
		return "本地"
	case Tencent:
		return "腾讯云"
	default:
		return "其他"
	}
}
