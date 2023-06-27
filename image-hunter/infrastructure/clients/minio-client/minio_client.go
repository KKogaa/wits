package minioclient

import (
	"bytes"
	"context"
	"log"
	"time"

	"github.com/KKogaa/image-hunter/domain/entities"
	"github.com/KKogaa/image-hunter/infrastructure/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioClient struct {
	config *config.Config
	client *minio.Client
}

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

	if bucketErr := minioClient.MakeBucket(context.Background(), config.MINIO_BUCKET_NAME,
		minio.MakeBucketOptions{Region: config.MINIO_LOCATION}); bucketErr != nil {

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

func (c MinioClient) SaveObject(content *entities.Content) (*entities.Content, error) {

	contentType := "images/jpeg"
	objectSize := len(content.Data)

	reader := bytes.NewReader(content.Data)

	info, err := c.client.PutObject(context.Background(),
		c.config.MINIO_BUCKET_NAME, content.Name, reader, int64(objectSize),
		minio.PutObjectOptions{ContentType: contentType})

	if err != nil {
		return nil, err
	}

	log.Printf("successfully uploaded %s of size %d\n", content.Name,
		info.Size)

	presignedUrl, err := c.client.PresignedGetObject(context.Background(),
		c.config.MINIO_BUCKET_NAME, content.Name, time.Duration(24)*time.Hour,
		nil)

	if err != nil {
		return nil, err
	}

	content.Path = presignedUrl.String()

	log.Printf("obtained presignedUrl: %s\n", presignedUrl.String())

	return content, nil
}
