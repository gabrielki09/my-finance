# Épico 4 — Movimentações Financeiras

## Objetivo

Registrar entradas e saídas reais de dinheiro.

---

## História 4.1 — Registrar entrada avulsa

Como usuário,
quero registrar uma entrada,
para atualizar o saldo de uma conta.

### Exemplos

- salário;
- freelance;
- reembolso;
- venda eventual.

### Dados

```text
financial_account_id
category_id
description
amount
movement_date
reference_date
idempotency_key
```

### Regras

- Conta financeira deve existir e estar ativa.
- Categoria deve existir, estar ativa e aceitar `income` ou `both`.
- Valor deve ser maior que zero.
- `movement_type` será `entry`.
- `operation_type` será `original`.
- `idempotency_key` deve ser único.

### Critérios de aceite

- Entrada válida retorna 201.
- Saldo da conta aumenta.
- Categoria de despesa não deve ser usada em entrada.
- Duplicidade de `idempotency_key` não cria outra movimentação.

---

## História 4.2 — Registrar saída avulsa

Como usuário,
quero registrar uma saída,
para controlar gastos.

### Exemplos

- mercado;
- lanche;
- combustível;

### Regras

- Conta financeira deve existir e estar ativa.
- Categoria deve existir, estar ativa e aceitar `expense` ou `both`.
- Valor deve ser maior que zero.
- `movement_type` será `exit`.
- `operation_type` será `original`.
- Não bloquear saldo negativo no MVP, a menos que você queira controle rígido.

### Critérios de aceite

- Saída válida retorna 201.
- Saldo da conta diminui.
- Categoria de receita não deve ser usada em saída.

---

## História 4.3 — Listar extrato

Como usuário,
quero listar movimentações por período,
para conferir meu histórico financeiro.

### Filtros

```text
financial_account_id
category_id
movement_type
start_date
end_date
operation_type
```

### Critérios de aceite

- Lista ordenada por data decrescente.
- Permite filtrar por conta.
- Permite filtrar por categoria.
- Permite filtrar por período.
- Não mistura movimentações canceladas, salvo se filtro futuro for criado.

---

## História 4.4 — Consultar saldo por conta

Como usuário,
quero consultar saldo de uma conta,
para saber quanto tenho nela.

### Fórmula

```text
saldo atual =
initial_balance
+ total de entradas não canceladas
- total de saídas não canceladas
```

### Critérios de aceite

- Retorna saldo correto.
- Considera saldo inicial.
- Ignora movimentações canceladas.
- Considera estornos como movimentações normais inversas.

---

## História 4.5 — Consultar saldo total

Como usuário,
quero consultar saldo total,
para saber quanto possuo somando minhas contas.

### Regra

Somar o saldo de todas as contas ativas ou todas as contas, conforme filtro escolhido.

### Critérios de aceite

- Retorna soma correta.
- Permite detalhar por conta no dashboard.
