package main

import (
	"log"
	"os"
	"vend/docs"
	"vend/internal/delivery/http"
	"vend/internal/infrastructure/chatgpt"
	"vend/internal/infrastructure/mongodb"
	"vend/internal/repository"
	"vend/internal/usecase"

	_ "vend/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Vend API
// @version         1.0
// @description     API para gerenciamento de vendas com integração ChatGPT
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Aviso: Arquivo .env não encontrado")
	}

	// Inicializa o cliente MongoDB
	mongoClient, err := mongodb.NewMongoClient()
	if err != nil {
		log.Fatalf("Erro ao conectar ao MongoDB: %v", err)
	}
	defer mongoClient.Disconnect(nil)

	// Inicializa o repositório
	pessoaRepo := repository.NewPessoaRepository(mongoClient)

	// Inicializa os casos de uso
	pessoaUseCase := usecase.NewPessoaUseCase(pessoaRepo)
	telefoneUseCase := usecase.NewTelefoneUseCase(pessoaRepo)
	contextoUseCase := usecase.NewContextoUseCase(pessoaRepo)
	promptUseCase := usecase.NewPromptUseCase(pessoaRepo)

	// Inicializa o handler
	handler := http.NewHandler(pessoaUseCase, telefoneUseCase, contextoUseCase, promptUseCase)

	// Inicializa o serviço do ChatGPT (será usado posteriormente)
	_ = chatgpt.NewChatGPTService()

	// Configurar router
	r := gin.Default()

	// Configurar CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Grupo de rotas da API
	v1 := r.Group("/api/v1")
	{
		// Rotas de Pessoas
		pessoas := v1.Group("/pessoas")
		{
			pessoas.GET("", handler.ListPessoas)
			pessoas.POST("", handler.CreatePessoa)
			pessoas.GET("/:id", handler.GetPessoa)
			pessoas.PUT("/:id", handler.UpdatePessoa)
			pessoas.DELETE("/:id", handler.DeletePessoa)
		}

		// Rotas de Telefones
		telefones := v1.Group("/telefones")
		{
			telefones.GET("", handler.ListTelefones)
			telefones.POST("", handler.CreateTelefone)
			telefones.GET("/:id", handler.GetTelefone)
			telefones.PUT("/:id", handler.UpdateTelefone)
			telefones.DELETE("/:id", handler.DeleteTelefone)
		}

		// Rotas de Contextos
		contextos := v1.Group("/contextos")
		{
			contextos.GET("", handler.ListContextos)
			contextos.POST("", handler.CreateContexto)
			contextos.GET("/:id", handler.GetContexto)
			contextos.PUT("/:id", handler.UpdateContexto)
			contextos.DELETE("/:id", handler.DeleteContexto)
		}

		// Rotas de Prompts
		prompts := v1.Group("/prompts")
		{
			prompts.GET("", handler.ListPrompts)
			prompts.POST("", handler.CreatePrompt)
			prompts.GET("/:id", handler.GetPrompt)
			prompts.PUT("/:id", handler.UpdatePrompt)
			prompts.DELETE("/:id", handler.DeletePrompt)
		}
	}

	// Configurar Swagger
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Iniciar servidor
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor iniciado na porta %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
