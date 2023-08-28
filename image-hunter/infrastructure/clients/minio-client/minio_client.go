package minioclient

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"time"

	"github.com/KKogaa/image-hunter/infrastructure/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioClient struct {
	config *config.Config
	client *minio.Client
}

func GenerateId(text string) string {
	hash := sha256.New().Sum([]byte(text))
	hashString := hex.EncodeToString(hash[:])
	return hashString
}

//TODO: later refactor move this to a file storage service
func NewMinioClient(config *config.Config) (*MinioClient, error) {

	useSSL := false

	minioClient, clientErr := minio.New(config.MINIO_ENDPOINT, &minio.Options{
		Creds: credentials.NewStaticV4(config.MINIO_ACCESS_KEY_ID,
			config.MINIO_SECRET_KEY, ""),
		Secure: useSSL,
	})

	if clientErr != nil {
		return &MinioClient{}, clientErr
	}
	log.Printf("image name: %s", text)

	if bucketErr := minioClient.MakeBucket(context.Background(), config.MINIO_BUCKET_NAME,
		minio.MakeBucketOptions{
			Region:        config.MINIO_LOCATION,
			ObjectLocking: false,
		}); bucketErr != nil {

		err := minioClient.SetBucketPolicy(context.Background(), config.MINIO_BUCKET_NAME, `{"Version":"2012-10-17","Statement":[{"Action":["s3:GetObject"],"Effect":"Allow","Principal":"*","Resource":["arn:aws:s3:::`+config.MINIO_BUCKET_NAME+`/*"]}],"Version":"2012-10-17"}`)
		if err != nil {
			return nil, err
		}

		exists, errBucketExists := minioClient.BucketExists(context.Background(),
			config.MINIO_BUCKET_NAME)

		if !(errBucketExists == nil && exists) {
			return &MinioClient{}, bucketErr
		}
	}

	return &MinioClient{
		config: config,
		client: minioClient,
	}, nil
}

func (c MinioClient) SaveObject(data []byte, imageName string) (string, error) {

	minioId := GenerateId(imageName)

	contentType := "images/jpeg"
	objectSize := len(data)

	reader := bytes.NewReader(data)

	info, err := c.client.PutObject(context.Background(),
		c.config.MINIO_BUCKET_NAME, minioId, reader, int64(objectSize),
		minio.PutObjectOptions{ContentType: contentType})

	if err != nil {
		return "", err
	}

	log.Printf("successfully uploaded %s of size %d\n", minioId,
		info.Size)

	return minioId, nil
}

func (c MinioClient) GetObject(contentId string) (string, error) {

	presignedUrl, err := c.client.PresignedGetObject(context.Background(),
		c.config.MINIO_BUCKET_NAME, contentId, time.Duration(7*24)*time.Hour,
		nil)

	if err != nil {
		return "", err
	}

	log.Printf("obtained presignedUrl: %s\n", presignedUrl.String())

	return presignedUrl.String(), nil

}
