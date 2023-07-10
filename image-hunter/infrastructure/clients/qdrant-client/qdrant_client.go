package qdrantclient

import (
	"context"
	"crypto/sha256"
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
				Distance: pb.Distance_Dot,
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

func (q QdrantClient) GetSimilarVector(vector *entities.Vector) ([]*entities.Vector, error) {
	conn, err := GetConnection(q.config.QDRANT_ENDPOINT)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pb.NewPointsClient(conn)

	unfilteredSearchResult, err := client.Search(context.Background(), &pb.SearchPoints{
		CollectionName: q.config.QDRANT_COLLECTION_NAME,
		Vector:         vector.Vector,
		Limit:          20,
		WithVectors:    &pb.WithVectorsSelector{SelectorOptions: &pb.WithVectorsSelector_Enable{Enable: true}},
		WithPayload:    &pb.WithPayloadSelector{SelectorOptions: &pb.WithPayloadSelector_Enable{Enable: true}},
	})
	if err != nil {
		return nil, fmt.Errorf("could not search points: %v", err)
	}

	var vectors []*entities.Vector
	for _, result := range unfilteredSearchResult.Result {
		payloadValue := result.Payload["url"]
		vectors = append(vectors, &entities.Vector{
			ID:         result.Id.String(),
			Path:       payloadValue.GetStringValue(),
			Similarity: result.GetScore(),
			Vector:     result.Vectors.GetVector().Data,
		})

	}
	return vectors, nil

}

func GenerateId(text string) int64 {
	hash := sha256.New().Sum([]byte(text))

	var result int64
	for _, b := range hash {
		result = (result << 8) | int64(b)
	}
	return result
}

func (q QdrantClient) SaveVector(vector *entities.Vector) (*entities.Vector, error) {

	upsertPoints := []*pb.PointStruct{
		{
			Id: &pb.PointId{
				PointIdOptions: &pb.PointId_Num{Num: uint64(GenerateId(vector.Path))},
			},
			Vectors: &pb.Vectors{
				VectorsOptions: &pb.Vectors_Vector{Vector: &pb.Vector{Data: vector.Vector}},
			},
			Payload: map[string]*pb.Value{
				"url": {
					Kind: &pb.Value_StringValue{StringValue: vector.Path},
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

	return vector, nil

}

func (q QdrantClient) GetAllVectors() ([]*entities.Vector, error) {
	return nil, nil

}
