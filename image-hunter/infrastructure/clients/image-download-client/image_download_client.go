package imagedownloadclient

import (
	"io/ioutil"
	"net/http"
)

type ImageDownloadClient struct {
}

func NewImageDownloadClient() *ImageDownloadClient {
	return &ImageDownloadClient{}
}

func (i ImageDownloadClient) GetImage(imageUrl string) ([]byte, error) {
	response, err := http.Get(imageUrl)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	imageData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return imageData, nil

}
