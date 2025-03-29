# Vend - Sistema de Gerenciamento de Pessoas e Contextos

Sistema desenvolvido em Go com React para gerenciamento de pessoas, telefones e contextos, com integração ao ChatGPT.

## Estrutura do Projeto

```
.
├── cmd/
│   └── api/             # Ponto de entrada da aplicação
├── internal/
│   ├── domain/          # Entidades e interfaces do domínio
│   ├── usecase/         # Casos de uso da aplicação
│   ├── repository/      # Interfaces dos repositórios
│   ├── delivery/        # Camada de entrega (HTTP)
│   └── infrastructure/  # Implementações concretas (PostgreSQL, ChatGPT)
├── pkg/                 # Pacotes compartilhados
├── test/               # Testes unitários e de integração
└── deployments/        # Configurações de deploy (Kubernetes, GitHub Actions)
```

## Requisitos

- Go 1.21 ou superior
- PostgreSQL 13 ou superior
- Docker
- Kubernetes
- Azure CLI (para deploy)

## Configuração

1. Clone o repositório:
```bash
git clone https://github.com/seu-usuario/vend.git
cd vend
```

2. Instale as dependências:
```bash
go mod download
```

3. Configure as variáveis de ambiente:
```bash
cp .env.example .env
# Edite o arquivo .env com suas configurações
```

4. Execute as migrações do banco de dados:
```bash
go run cmd/api/main.go
```

## Executando Localmente

1. Inicie o PostgreSQL:
```bash
docker-compose up -d postgres
```

2. Execute a aplicação:
```bash
go run cmd/api/main.go
```

A API estará disponível em `http://localhost:8080`

## Testes

### Testes Unitários
```bash
go test ./test/unit/...
```

### Testes de Integração
```bash
go test ./test/integration/...
```

## Deploy

O deploy é automatizado através do GitHub Actions para o Azure Kubernetes Service (AKS).

### Configuração do Azure

1. Crie um cluster AKS:
```bash
az group create --name vend-rg --location eastus
az aks create --resource-group vend-rg --name vend-cluster --node-count 3 --enable-addons monitoring
```

2. Configure as secrets no GitHub:
- ACR_LOGIN_SERVER
- ACR_USERNAME
- ACR_PASSWORD
- AZURE_CREDENTIALS
- AZURE_SUBSCRIPTION_ID
- AZURE_RESOURCE_GROUP
- AKS_CLUSTER_NAME

### Deploy Manual

```bash
# Construir a imagem
docker build -t vend-api .

# Fazer push para o Azure Container Registry
docker tag vend-api seu-registro.azurecr.io/vend-api:latest
docker push seu-registro.azurecr.io/vend-api:latest

# Aplicar as configurações do Kubernetes
kubectl apply -f deployments/kubernetes/
```

## API Endpoints

### Pessoas
- GET /pessoas - Lista todas as pessoas
- POST /pessoas - Cria uma nova pessoa
- GET /pessoas/:id - Obtém uma pessoa específica
- PUT /pessoas/:id - Atualiza uma pessoa
- DELETE /pessoas/:id - Remove uma pessoa

### Telefones
- GET /telefones - Lista todos os telefones
- POST /telefones - Cria um novo telefone
- GET /telefones/:id - Obtém um telefone específico
- PUT /telefones/:id - Atualiza um telefone
- DELETE /telefones/:id - Remove um telefone

### Contextos
- GET /contextos - Lista todos os contextos
- POST /contextos - Cria um novo contexto
- GET /contextos/:id - Obtém um contexto específico
- PUT /contextos/:id - Atualiza um contexto
- DELETE /contextos/:id - Remove um contexto

### Prompts
- GET /prompts - Lista todos os prompts
- POST /prompts - Cria um novo prompt
- GET /prompts/:id - Obtém um prompt específico
- PUT /prompts/:id - Atualiza um prompt
- DELETE /prompts/:id - Remove um prompt

## Contribuindo

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## Licença

Este projeto está licenciado sob a licença MIT - veja o arquivo [LICENSE](LICENSE) para detalhes. 




 



docker stop portainer && docker rm portainer
docker volume rm portainer_data
docker-compose down -v

echo -n "admin" | sha256sum
8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918

docker run -d -p 8000:8000 -p 9000:9000 --name=portainer --restart=always -v /var/run/docker.sock:/var/run/docker.sock -v portainer_data:/data portainer/portainer-ce

admin
adminadmin123

http://localhost:8080/swagger/index.html


go clean
go mod tidy
go build 
