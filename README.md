# Ecoffe 

## Pré-requisitos

Antes de começar, certifique-se de ter instalado:
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Como Executar o Projeto com Docker Compose

Siga os passos abaixo para configurar e executar o projeto usando Docker Compose:

1. **Clone o Repositório** (se ainda não o fez):
   ```bash
   git clone <URL-do-repositório>
   cd ecoffe
   ```

2. **Configurar Variáveis de Ambiente**:
   - Crie um arquivo `.env` na raiz do projeto com base no exemplo fornecido ou configure as variáveis de ambiente diretamente no seu sistema.
   - As variáveis essenciais incluem:
     - `POSTGRES_USER`: Usuário do banco de dados PostgreSQL.
     - `POSTGRES_PASSWORD`: Senha do banco de dados PostgreSQL.
     - `POSTGRES_DB_NAME`: Nome do banco de dados.

3. **Construir e Iniciar os Contêineres**:
   Execute o comando abaixo para construir a imagem Docker e iniciar os serviços:
   ```bash
   docker-compose up --build
   ```
   Isso iniciará tanto o serviço da aplicação `ecoffe` quanto o banco de dados PostgreSQL.

4. **Acessar a Aplicação**:
   Após a inicialização, a aplicação estará disponível em `http://localhost:8080`.

5. **Parar os Contêineres**:
   Quando terminar, você pode parar os serviços com:
   ```bash
   docker-compose down
   ```

## Arquitetura
-	Entidades: Modelos centrais (`order.go`).
-	Casos de Uso: Lógica de negócio em `usecases/`.
-	Adaptadores: Integrações externas (`api/`).

# Endpoints da API

Abaixo estão exemplos de como interagir com a API usando comandos `curl`:

## Criar um Novo Pedido
```bash
curl --request POST \
  --url http://localhost:8080/orders \
  --header 'Content-Type: application/json' \
  --data '{
    "client": "Maria",
    "item": "Cold Brew",
    "quantity": 1,
    "observation": "Com leite zero lactose"
}'
```

## Obter Todos os Pedidos
```bash
curl --request GET \
  --url http://localhost:8080/orders
```

## Obter Pedidos por Status
```bash
curl --request GET \
  --url 'http://localhost:8080/orders?s=Entregue'
```

**Status Disponíveis:**
- Pendente
- Preparando
- Pronto
- A Caminho
- Entregue
- Cancelado

## Atualizar Status do Pedido
```bash
curl --request PATCH \
  --url http://localhost:8080/orders/bcacfc97-7bb2-4791-acf3-95780fa40955 \
  --header 'Content-Type: application/json' \
  --data '{
    "status": "Pronto"
}'
```
