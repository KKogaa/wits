package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/KKogaa/image-hunter/docs"
	"github.com/gin-contrib/cors"

	"github.com/KKogaa/image-hunter/infrastructure/config"

	"github.com/KKogaa/image-hunter/infrastructure/clients/clip-client"
	// "github.com/KKogaa/image-hunter/infrastructure/clients/elastic-search-client"
	"github.com/KKogaa/image-hunter/infrastructure/clients/image-download-client"
	"github.com/KKogaa/image-hunter/infrastructure/clients/minio-client"
	"github.com/KKogaa/image-hunter/infrastructure/clients/qdrant-client"
	"github.com/KKogaa/image-hunter/infrastructure/controllers/http-2"
	"github.com/KKogaa/image-hunter/usecases"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	router           *gin.Engine
	searchController *http2.SearchController
	saveController   *http2.SaveController
	vectorController *http2.VectorController
}

func NewServer(config *config.Config) *Server {
	router := gin.Default()

	//adding middleware
	router.Use(cors.Default())

	//injecting infrastructure to usecases
	hasherClient := clipclient.NewClipClient(config)

	minioClient, err := minioclient.NewMinioClient(config)
	if err != nil {
		log.Fatalf("error in minio client: %s", err.Error())
	}

	imageDownloadClient := imagedownloadclient.NewImageDownloadClient()

	// esClient, err := elasticsearchclient.NewElasticSearchClient(config)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	qdrantClient, err := qdrantclient.NewQdrantClient(config)
	if err != nil {
		log.Fatal(err)
	}

	searchImageUsecase := usecases.NewSearchImageUsecase(hasherClient, qdrantClient, minioClient)
	searchController := http2.NewSearchController(searchImageUsecase)

	saveContentUsecase := usecases.NewSaveContentUsecase(hasherClient,
		minioClient, qdrantClient, imageDownloadClient)
	saveController := http2.NewSaveController(saveContentUsecase)

	getAllVectorsUsecase := usecases.NewGetAllVectorsUsecase(qdrantClient, minioClient)
	vectorController := http2.NewVectorController(getAllVectorsUsecase)

	return &Server{
		router:           router,
		searchController: searchController,
		saveController:   saveController,
		vectorController: vectorController,
	}
}

// @title Image Hunter
// @version 0.0
// @description

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name
// @license.url

// @host localhost:8080
// @BasePath /
// @query.collection.format multi
func SetupRoutes(server *Server) {
	// docs.SwaggerInfo.BasePath = "/api/v1"
	server.searchController.SetupRoutes(server.router)
	server.saveController.SetupRoutes(server.router)
	server.vectorController.SetupRoutes(server.router)
	server.router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func Start(server *Server, config *config.Config) {

	SetupRoutes(server)

	srv := &http.Server{
		Addr:    ":" + config.PORT,
		Handler: server.router,
	}

	log.Printf("server started on port: %s\n ", config.PORT)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

func main() {
	config := config.LoadConfig()
	server := NewServer(config)
	Start(server, config)
}
