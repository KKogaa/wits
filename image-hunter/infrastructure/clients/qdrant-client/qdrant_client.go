package qdrantclient

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"fmt"

	"github.com/KKogaa/image-hunter/domain/entities"
	"github.com/KKogaa/image-hunter/infrastructure/config"
	pb "github.com/qdrant/go-client/qdrant"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type QdrantClient struct {
	config *config.Config
}

func GetConnection(endpoint string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("did not connect: %v", err)
	}
	return conn, nil

}

func NewQdrantClient(config *config.Config) (*QdrantClient, error) {

	err := InitCollection(config)
	if err != nil {
		return nil, err
	}

	return &QdrantClient{
		config: config,
	}, nil

}

func InitCollection(config *config.Config) error {

	conn, err := GetConnection(config.QDRANT_ENDPOINT)
	if err != nil {
		return err
	}
	defer conn.Close()

	clientCollection := pb.NewCollectionsClient(conn)

	var defaultSegmentNumber uint64 = 2
	var embeddingSize = 512

	_, err = clientCollection.Create(context.Background(), &pb.CreateCollection{
		CollectionName: config.QDRANT_COLLECTION_NAME,
		VectorsConfig: &pb.VectorsConfig{Config: &pb.VectorsConfig_Params{
			Params: &pb.VectorParams{
				Size:     uint64(embeddingSize),
				Distance: pb.Distance_Cosine,
			},
		}},
		OptimizersConfig: &pb.OptimizersConfigDiff{
			DefaultSegmentNumber: &defaultSegmentNumber,
		},
	})

	if err != nil && !(status.Code(err) == codes.InvalidArgument) {
		return err
	}

	return nil

}

func (q QdrantClient) GetSimilarVector(embedding []float32) ([]*entities.Vector, error) {

	conn, err := GetConnection(q.config.QDRANT_ENDPOINT)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pb.NewPointsClient(conn)

	unfilteredSearchResult, err := client.Search(context.Background(), &pb.SearchPoints{
		CollectionName: q.config.QDRANT_COLLECTION_NAME,
		Vector:         embedding,
		Limit:          20,
		WithVectors:    &pb.WithVectorsSelector{SelectorOptions: &pb.WithVectorsSelector_Enable{Enable: true}},
		WithPayload:    &pb.WithPayloadSelector{SelectorOptions: &pb.WithPayloadSelector_Enable{Enable: true}},
	})
	if err != nil {
		return nil, fmt.Errorf("could not search points: %v", err)
	}

	var vectors []*entities.Vector
	for _, result := range unfilteredSearchResult.Result {
		vectors = append(vectors, &entities.Vector{
			ID:         result.Id.GetNum(),
			Similarity: result.GetScore(),
			Embedding:  result.Vectors.GetVector().Data,
			ImageID:    result.Payload["imageId"].GetStringValue(),
		})

	}
	return vectors, nil

}

func GenerateId(text string) uint64 {
	hash := sha256.Sum256([]byte(text))
	return binary.BigEndian.Uint64(hash[:8])
}

func (q QdrantClient) SaveVector(imageId string, embedding []float32) (*entities.Vector, error) {

	vector := entities.Vector{
		ID:        GenerateId(imageId),
		Embedding: embedding,
		ImageID:   imageId,
	}

	upsertPoints := []*pb.PointStruct{
		{
			Id: &pb.PointId{
				PointIdOptions: &pb.PointId_Num{Num: vector.ID},
			},
			Vectors: &pb.Vectors{
				VectorsOptions: &pb.Vectors_Vector{Vector: &pb.Vector{Data: vector.Embedding}},
			},
			Payload: map[string]*pb.Value{
				"imageId": {
					Kind: &pb.Value_StringValue{StringValue: vector.ImageID},
				},
			},
		},
	}

	conn, err := GetConnection(q.config.QDRANT_ENDPOINT)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pb.NewPointsClient(conn)

	waitUpsert := true
	_, err = client.Upsert(context.Background(), &pb.UpsertPoints{
		CollectionName: q.config.QDRANT_COLLECTION_NAME,
		Wait:           &waitUpsert,
		Points:         upsertPoints,
	})

	if err != nil {
		return nil, err
	}

	return &vector, nil

}

func (q QdrantClient) GetAllVectors() ([]*entities.Vector, error) {
	vector := make([]float32, 512)

	for i := 0; i < 512; i++ {
		vector[i] = 0.0
	}

	conn, err := GetConnection(q.config.QDRANT_ENDPOINT)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pb.NewPointsClient(conn)

	unfilteredSearchResult, err := client.Search(context.Background(), &pb.SearchPoints{
		CollectionName: q.config.QDRANT_COLLECTION_NAME,
		Vector:         vector,
		Limit:          100,
		WithVectors:    &pb.WithVectorsSelector{SelectorOptions: &pb.WithVectorsSelector_Enable{Enable: true}},
		WithPayload:    &pb.WithPayloadSelector{SelectorOptions: &pb.WithPayloadSelector_Enable{Enable: true}},
	})
	if err != nil {
		return nil, fmt.Errorf("could not search points: %v", err)
	}

	var vectors []*entities.Vector
	for _, result := range unfilteredSearchResult.Result {
		vectors = append(vectors, &entities.Vector{
			ID:         result.Id.GetNum(),
			Similarity: result.GetScore(),
			Embedding:  result.Vectors.GetVector().Data,
			ImageID:    result.Payload["imageId"].GetStringValue(),
		})

	}
	return vectors, nil
}
