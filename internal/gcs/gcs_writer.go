package gcs

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
)

type GCSService struct {
	bucketName    string
	storageClient *storage.Client
}

func NewGCSService(bucketName string) GCSService {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Errorf("unable to initialize GCS client %v", err)
		panic(err)
	}
	return GCSService{bucketName, client}
}

func (service *GCSService) WriteFileToBucket(file []byte, fileName string) {
	ctx := context.Background()
	fmt.Printf("Writing file '%s' to GCS bucket\n", fileName)
	wc := service.storageClient.Bucket(service.bucketName).Object(fileName).NewWriter(ctx)
	defer wc.Close()
	if _, err := wc.Write(file); err != nil {
		panic(fmt.Errorf("unable to create GCS object %w", err))
	}
}
