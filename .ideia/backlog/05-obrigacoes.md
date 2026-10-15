# Épico 5 — Obrigações Financeiras

## Objetivo

Cadastrar contas a pagar e a receber antes de o dinheiro realmente movimentar.

---

## História 5.1 — Cadastrar conta a pagar

Como usuário,
quero cadastrar uma conta a pagar,
para controlar compromissos futuros.

### Exemplos

- aluguel;
- energia;
- internet;
- cartão;
- parcela.

### Dados

```text
category_id
description
type = payable
original_amount
due_date
competence_date
notes
```

### Regras

- Categoria deve existir e estar ativa.
- Categoria deve aceitar `expense` ou `both`.
- Valor deve ser maior que zero.
- Data de vencimento é obrigatória.
- Status inicial será `pending`.

### Critérios de aceite

- Conta a pagar válida retorna 201.
- Conta cadastrada não altera saldo.
- Conta aparece na listagem de pendentes.

---

## História 5.2 — Cadastrar conta a receber

Como usuário,
quero cadastrar uma conta a receber,
para controlar valores futuros de entrada.

### Exemplos

- salário previsto;
- freelance a receber;
- reembolso;
- dinheiro emprestado.

### Regras

- Categoria deve existir e estar ativa.
- Categoria deve aceitar `income` ou `both`.
- Valor deve ser maior que zero.
- Status inicial será `pending`.

### Critérios de aceite

- Conta a receber válida retorna 201.
- Conta cadastrada não altera saldo.
- Conta aparece na listagem de pendentes.

---

## História 5.3 — Listar obrigações

Como usuário,
quero listar obrigações,
para saber o que tenho a pagar ou receber.

### Filtros

```text
type
status
start_due_date
end_due_date
category_id
```

### Critérios de aceite

- Lista contas a pagar.
- Lista contas a receber.
- Filtra por vencimento.
- Filtra por status.
- Identifica vencidas via cálculo.

---

## História 5.4 — Atualizar obrigação pendente

Como usuário,
quero alterar uma obrigação pendente,
para corrigir dados antes da liquidação.

### Regras

- Permitir alteração se status for `pending`.
- Se parcialmente liquidada, limitar alterações.
- Se liquidada, não permitir alterar valor original.
- Se cancelada, não permitir alterar.

### Critérios de aceite

- Obrigação pendente pode ser alterada.
- Obrigação liquidada não permite alteração sensível.
- Obrigação cancelada não permite alteração.

---

## História 5.5 — Cancelar obrigação

Como usuário,
quero cancelar uma obrigação,
para remover uma conta futura que não será mais paga/recebida.

### Regras

- Pode cancelar se não houver liquidação ativa.
- Se houver liquidação, deve estornar antes ou criar fluxo de cancelamento com estorno.
- Status vira `canceled`.
- Preencher `canceled_at`.

### Critérios de aceite

- Obrigação sem liquidação é cancelada.
- Obrigação liquidada não é cancelada diretamente.
- Obrigação cancelada não pode ser liquidada.
