# Modelo Conceitual

## Entidades principais

```text
financial_accounts
categories
financial_transactions
financial_obligations
obligation_settlements
financial_transfers
```

---

## financial_accounts

Representa onde o dinheiro está.

Campos conceituais:

```text
id
name
type
initial_balance
opened_at
active
created_at
updated_at
deleted_at
```

Tipos sugeridos:

```text
cash
checking
savings
digital
investment
```

---

## categories

Representa classificação de entradas e saídas.

Campos conceituais:

```text
id
parent_id
name
type
active
created_at
updated_at
deleted_at
```

Tipos:

```text
income
expense
both
```

---

## financial_transactions

Representa dinheiro efetivamente movimentado.

Campos conceituais:

```text
id
financial_account_id
category_id
reversed_transaction_id
description
movement_type
operation_type
amount
movement_date
reference_date
origin_type
origin_id
idempotency_key
created_at
canceled_at
```

Tipos de movimento:

```text
entry
exit
```

Tipos de operação:

```text
original
adjustment
reversal
```

---

## financial_obligations

Representa conta a pagar ou receber.

Campos conceituais:

```text
id
category_id
description
type
status
original_amount
due_date
competence_date
notes
created_at
updated_at
canceled_at
```

Tipos:

```text
payable
receivable
```

Status:

```text
pending
partially_settled
settled
canceled
```

---

## obligation_settlements

Relaciona obrigação com movimentação financeira.

Campos conceituais:

```text
id
obligation_id
transaction_id
amount
created_at
canceled_at
```

---

## financial_transfers

Representa transferência entre contas próprias.

Campos conceituais:

```text
id
source_account_id
destination_account_id
source_transaction_id
destination_transaction_id
amount
description
transfer_date
idempotency_key
created_at
canceled_at
```

---

## Relações

```text
financial_accounts 1:N financial_transactions

categories 1:N financial_transactions

categories 1:N financial_obligations

financial_obligations 1:N obligation_settlements

financial_transactions 1:1 obligation_settlements

financial_transfers 1:1 source_transaction

financial_transfers 1:1 destination_transaction
```
