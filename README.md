# goAuthApi

Uma API de autenticação simples escrita em Go (Golang) utilizando o framework Gin e JWT (JSON Web Tokens) para autenticação.

## 📋 Visão Geral

Este projeto fornece uma API de autenticação básica com as seguintes funcionalidades:
- Autenticação de usuários via login.
- Geração de tokens JWT para acesso seguro a endpoints protegidos.

### Tecnologias utilizadas:
- **Gin**: Framework HTTP rápido e eficiente para Go.
- **JWT**: Para geração e validação de tokens de autenticação.
- **InMemoryUserRepository**: Repositório em memória para gerenciamento de usuários (somente para demonstração).

## 🚀 Como Executar

### Pré-requisitos

- Go 1.20 ou superior instalado.
- Git instalado (para clonar o repositório).

### Passos para Configuração

1. Clone o repositório:
   ```bash
   git clone https://github.com/pedroaraldidev/goAuthApi.git
   cd goAuthApi
   ```

2. Crie um arquivo `.env` na raiz do projeto e adicione a chave secreta JWT:
   ```env
   JWT_SECRET=suaChaveSecretaMuitoSegura
   ```

3. Instale as dependências:
   ```bash
   go mod tidy
   ```

4. Execute a API:
   ```bash
   go run main.go
   ```

A API estará disponível em `http://localhost:8080`.

## 📌 Endpoints

### **Autenticação**
#### `POST /login`
Autentica um usuário e retorna um token JWT.

**Request Body:**
```json
{
  "username": "user1",
  "password": "password1"
}
```

**Response:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

## 🛠 Estrutura do Projeto

```
goAuthApi/
├── controllers/      # Controladores para lidar com as requisições HTTP
├── repositories/     # Repositórios para gerenciar dados (ex: usuários)
├── routes/           # Definição das rotas da API
├── services/         # Lógica de negócio e serviços
├── main.go           # Ponto de entrada da aplicação
├── go.mod            # Arquivo de dependências do Go
├── go.sum            # Arquivo de checksum das dependências
└── README.md         # Este arquivo
```


