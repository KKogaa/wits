package elasticsearchclient

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/KKogaa/image-hunter/domain/entities"
	"github.com/KKogaa/image-hunter/infrastructure/config"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type ElasticSearchClient struct {
	client *elasticsearch.Client
	config *config.Config
}

func InitESClient(esEndpoint string) (*elasticsearch.Client, error) {

	var addresses []string
	addresses = append(addresses, esEndpoint)
	log.Printf("try connecting to elastic search: %s\n", esEndpoint)

	cfg := elasticsearch.Config{
		Addresses: addresses,
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("error creating elastic search client: %s", err)
	}

	res, err := client.Info()
	if err != nil {
		return nil, fmt.Errorf("error connecting to elastic search client: %s",
			err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("elastic search ping failed: %s", res.Status())
	}

	log.Printf("success connecting to elastic search\n")

	return client, nil

}

func InitIndex(esClient *elasticsearch.Client, indexName string) error {

	//TODO: convert to struct
	mapping := `
    {
      "settings": {
        "number_of_shards": 1
      },
      "mappings": {
        "properties": {
          "text": {
            "type": "text"
          },
          "embedding": {
              "type": "dense_vector",
              "dims": 512,
              "index": true,
              "similarity": "cosine"
          }
        }
      }
    }`



	indexExistsRes, err := esClient.Indices.Exists([]string{indexName})
	if err != nil {
		return err
	}

	if !indexExistsRes.IsError() {
		log.Printf("index already exists: %s", indexExistsRes.String())
		return nil
	}

	req := esapi.IndicesCreateRequest{
		Index: indexName,
		Body:  strings.NewReader(string(mapping)),
	}

	res, err := req.Do(context.Background(), esClient)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	log.Printf(res.String())

	return nil

}

func NewElasticSearchClient(config *config.Config) (*ElasticSearchClient, error) {

	client, err := InitESClient(config.ELASTIC_SEARCH_ENDPOINT)
	if err != nil {
		return nil, err
	}

	err = InitIndex(client, config.ELASTIC_SEARCH_INDEX_NAME)
	if err != nil {
		return nil, err
	}

	return &ElasticSearchClient{
		client: client,
		config: config,
	}, nil
}

type ImageSearchRequest struct {
	KNN    KNNParams `json:"knn"`
	Fields []string  `json:"fields"`
}

type KNNParams struct {
	Field         string    `json:"field"`
	QueryVector   []float32 `json:"query_vector"`
	K             int       `json:"k"`
	NumCandidates int       `json:"num_candidates"`
}

type Hit struct {
	ID     string  `json:"_id"`
	Score  float32 `json:"_score"`
	Source struct {
		Text      string    `json:"text"`
		Embedding []float32 `json:"embedding"`
	} `json:"_source"`
}

type SearchResponse struct {
	Hits struct {
		Hits []Hit `json:"hits"`
	} `json:"hits"`
}

func (e ElasticSearchClient) GetSimilarVector(vector *entities.Vector) ([]*entities.Vector, error) {

	request := ImageSearchRequest{
		KNN: KNNParams{
			Field:         "embedding",
			QueryVector:   vector.Vector,
			K:             10,
			NumCandidates: 100,
		},
		Fields: []string{"text"},
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON: %s", err)
	}

	req := esapi.SearchRequest{
		Index: []string{e.config.ELASTIC_SEARCH_INDEX_NAME},
		Body:  strings.NewReader(string(jsonData)),
	}

	res, err := req.Do(context.Background(), e.client)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("error occured while searching for vectors: %s",
			res.String())
	}

	var searchResponse SearchResponse
	err = json.NewDecoder(res.Body).Decode(&searchResponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing search response: %s", err)
	}

	var vectors []*entities.Vector
	for _, hit := range searchResponse.Hits.Hits {
		vectors = append(vectors, &entities.Vector{
			ID:         hit.ID,
			Vector:     hit.Source.Embedding,
			Path:       hit.Source.Text,
			Similarity: hit.Score,
		})
	}

	return vectors, nil
}

type Document struct {
	Text      string    `json:"text"`
	Embedding []float32 `json:"embedding"`
}

func (e ElasticSearchClient) GenerateDocumentId(sentence string) string {
	hasher := sha256.New()
	hasher.Write([]byte(sentence))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (e ElasticSearchClient) SaveVector(vector *entities.Vector) (*entities.Vector, error) {

	document := Document{
		Text:      vector.Path,
		Embedding: vector.Vector,
	}

	jsonDoc, err := json.Marshal(document)
	if err != nil {
		return nil, err

	}

	req := esapi.CreateRequest{
		Index:      e.config.ELASTIC_SEARCH_INDEX_NAME,
		Refresh:    "true",
		Body:       strings.NewReader(string(jsonDoc)),
		DocumentID: e.GenerateDocumentId(vector.Path),
	}

	res, err := req.Do(context.Background(), e.client)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("error indexing document %s", res.String())
	}

	log.Printf("sucess uploading to elasticsearch")

	return vector, nil
}
