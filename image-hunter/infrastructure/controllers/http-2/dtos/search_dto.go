package dtos

type SearchTextDTO struct {
	Text string `form:"text"`
}

type SearchImageLinkDTO struct {
	Url string `form:"url"`
}
