package image_svc

type ImageService struct {
}

// 图片上传白名单
var WhiteImageList = []string{
	".jpg",
	".png",
	".jpeg",
	".ico",
	".tiff",
	".gif",
	".svg",
	".webp",
}
