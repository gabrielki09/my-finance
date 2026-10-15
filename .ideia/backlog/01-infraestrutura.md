# Épico 1 — Infraestrutura

## Objetivo

Criar a base técnica da API em Go com PostgreSQL no Docker, sem frontend, sem login e sem usuário.

## História 1.1 — Criar estrutura inicial do projeto

Como desenvolvedor,
quero criar a estrutura inicial do projeto,
para manter o código organizado desde o começo.

### Tarefas

- Criar diretório do projeto.
- Inicializar `go.mod`.
- Criar pasta `cmd/api`.
- Criar pasta `internal`.
- Criar pasta `sql`.
- Criar pasta `docs`, opcionalmente.
- Criar `.env.example`.
- Criar `README.md`.

### Critérios de aceite

- Projeto executa com `go run ./cmd/api`.
- Estrutura de pastas está clara.
- O projeto compila sem erro.

---

## História 1.2 — Subir PostgreSQL com Docker

Como desenvolvedor,
quero subir PostgreSQL via Docker,
para ter banco local reproduzível.

### Tarefas

- Criar `compose.yaml`.
- Definir database `finance`.
- Definir user `finance`.
- Definir password local.
- Expor porta somente em `127.0.0.1`.
- Criar volume persistente.
- Criar healthcheck.

### Critérios de aceite

- `docker compose up -d` sobe o banco.
- É possível conectar usando `psql` ou client gráfico.
- Banco persiste dados após reiniciar container.

---

## História 1.3 — Criar conexão com PostgreSQL

Como desenvolvedor,
quero conectar a API ao PostgreSQL,
para permitir operações de persistência.

### Tarefas

- Ler `DATABASE_URL`.
- Criar pool de conexões.
- Testar conexão com `Ping`.
- Encerrar pool no shutdown.
- Logar erro de conexão.

### Critérios de aceite

- API não sobe se `DATABASE_URL` não existir.
- API não sobe se banco estiver indisponível.
- API sobe corretamente com banco disponível.

---

## História 1.4 — Criar servidor HTTP

Como desenvolvedor,
quero criar servidor HTTP usando `net/http`,
para expor a API sem dependência de roteador externo.

### Tarefas

- Criar `http.NewServeMux`.
- Criar endpoint `GET /health`.
- Configurar timeouts.
- Configurar graceful shutdown.
- Retornar JSON no `/health`.

### Critérios de aceite

- `GET /health` retorna 200.
- Resposta esperada:

```json
{
  "status": "ok"
}
```

---

## História 1.5 — Padronizar respostas JSON

Como desenvolvedor,
quero padronizar respostas JSON,
para manter consistência entre endpoints.

### Modelo de sucesso

```json
{
  "status": true,
  "message": "Operação realizada com sucesso.",
  "data": {}
}
```

### Modelo de erro

```json
{
  "status": false,
  "message": "Erro ao processar operação.",
  "error": {}
}
```

### Critérios de aceite

- Todo endpoint usa o mesmo padrão.
- Erros de validação retornam `fields`.
- Erros inesperados não expõem stack trace.

---

## História 1.6 — Padronizar decode de JSON

Como desenvolvedor,
quero uma função central para ler JSON,
para tratar erros como tipo inválido, JSON malformado e campos desconhecidos.

### Tarefas

- Criar helper para `json.Decoder`.
- Usar `DisallowUnknownFields`.
- Tratar `json.SyntaxError`.
- Tratar `json.UnmarshalTypeError`.
- Tratar payload vazio.
- Tratar campo desconhecido.

### Critérios de aceite

- Enviar string onde espera número retorna erro amigável.
- Enviar campo inexistente retorna erro amigável.
- Enviar JSON inválido retorna erro amigável.
