package main

import (
	"fmt"
	"log"
	"os"

	_ "vend/docs"
	"vend/internal/delivery/http"
	"vend/internal/domain"
	"vend/internal/infrastructure/chatgpt"
	postgresRepo "vend/internal/infrastructure/postgres"
	"vend/internal/service"
	"vend/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title           Vend API
// @version         1.0
// @description     API para gerenciamento de pessoas, telefones e contextos com integração ChatGPT
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
	// Configuração do Viper
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Erro ao ler arquivo de configuração: %v", err)
	}

	// Configuração do logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	// Configuração do banco de dados
	dbHost := viper.GetString("DB_HOST")
	dbPort := viper.GetString("DB_PORT")
	dbUser := viper.GetString("DB_USER")
	dbPassword := viper.GetString("DB_PASSWORD")
	dbName := viper.GetString("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// Configurar banco de dados
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	// Migrar modelos
	db.AutoMigrate(&domain.Pessoa{}, &domain.Telefone{}, &domain.Contexto{}, &domain.Prompt{})

	// Inicializar repositórios e casos de uso
	repo := postgresRepo.NewRepository(db)
	pessoaUseCase := usecase.NewPessoaUseCase(repo)

	// Inicializar serviço ChatGPT (será usado posteriormente)
	_ = chatgpt.NewService(viper.GetString("OPENAI_API_KEY"))

	// Inicialização do serviço
	pessoaService := service.NewPessoaService(repo)

	// Inicialização do handler
	handler := http.NewHandler(pessoaUseCase)

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
	port := viper.GetString("PORT")
	if port == "" {
		port = "8080"
	}

	logger.Infof("Servidor iniciado na porta %s", port)
	if err := r.Run(":" + port); err != nil {
		logger.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
