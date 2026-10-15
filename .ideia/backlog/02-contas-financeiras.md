# Épico 2 — Contas Financeiras

## Objetivo

Permitir cadastrar e controlar os locais onde o dinheiro está armazenado.

Exemplos:

- Nubank
- Carteira
- Sicoob
- Reserva de emergência
- Poupança

---

## História 2.1 — Cadastrar conta financeira

Como usuário,
quero cadastrar uma conta financeira,
para informar onde meu dinheiro está.

### Dados necessários

```text
name
type
initial_balance
opened_at
active
```

### Tipos aceitos

```text
cash
checking
savings
digital
investment
```

### Regras

- Nome é obrigatório.
- Nome não pode ser vazio.
- Tipo é obrigatório.
- Tipo deve ser válido.
- Saldo inicial pode ser zero.
- Saldo inicial pode ser positivo.
- Permitir saldo inicial negativo somente se quiser representar cheque especial/dívida em conta. Para MVP, recomenda-se bloquear negativo.
- `opened_at` pode assumir data atual quando não informado.
- Conta começa ativa por padrão.

### Critérios de aceite

- Ao cadastrar conta válida, retorna 201.
- Ao cadastrar sem nome, retorna 422.
- Ao cadastrar tipo inválido, retorna 422.
- Ao cadastrar saldo inválido, retorna 422.
- Conta aparece na listagem após cadastro.

---

## História 2.2 — Listar contas financeiras

Como usuário,
quero listar minhas contas financeiras,
para visualizar onde tenho dinheiro.

### Filtros sugeridos

- `active=true`
- `active=false`
- sem filtro: listar todas não deletadas

### Critérios de aceite

- Retorna lista ordenada por nome.
- Retorna contas ativas.
- Não retorna contas com `deleted_at` preenchido, salvo se criado filtro futuro.

---

## História 2.3 — Consultar conta financeira por ID

Como usuário,
quero consultar uma conta específica,
para ver seus dados.

### Critérios de aceite

- ID existente retorna 200.
- ID inexistente retorna 404.
- ID inválido retorna 400.

---

## História 2.4 — Atualizar conta financeira

Como usuário,
quero alterar dados da conta financeira,
para corrigir nome, tipo ou saldo inicial.

### Regras

- Pode alterar nome.
- Pode alterar tipo.
- Pode alterar saldo inicial enquanto não houver movimentações, ou permitir alteração com cuidado.
- Se já houver movimentações, alterar saldo inicial muda o saldo atual histórico.
- Recomendação: no MVP permitir alteração, pois é uso pessoal, mas saber que isso recalcula o saldo.

### Critérios de aceite

- Atualização válida retorna 200.
- Nome vazio retorna 422.
- Tipo inválido retorna 422.

---

## História 2.5 — Desativar conta financeira

Como usuário,
quero desativar uma conta financeira,
para não usá-la mais em novos lançamentos.

### Regras

- Desativar conta não apaga transações antigas.
- Conta inativa não pode ser usada em novas movimentações.
- Conta inativa pode aparecer em relatórios históricos.
- Se conta possuir saldo diferente de zero, exibir alerta/regra no retorno, mas para API pode permitir.

### Critérios de aceite

- Conta desativada não aparece em seleção de novas operações.
- Transações antigas continuam consultáveis.
- Tentativa de usar conta inativa em nova movimentação retorna erro.
