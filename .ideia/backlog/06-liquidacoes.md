# Épico 6 — Liquidações

## Objetivo

Permitir pagar contas a pagar e receber contas a receber, gerando movimentações financeiras reais.

---

## História 6.1 — Pagar conta a pagar

Como usuário,
quero pagar uma obrigação do tipo `payable`,
para registrar a saída de dinheiro.

### Dados

```text
obligation_id
financial_account_id
amount
payment_date
idempotency_key
```

### Fluxo

1. Buscar obrigação.
2. Validar se é `payable`.
3. Validar se não está cancelada.
4. Validar saldo pendente.
5. Criar movimentação de saída.
6. Criar liquidação.
7. Atualizar status da obrigação.

### Regras

- Valor deve ser maior que zero.
- Valor não pode ser maior que saldo pendente.
- Conta financeira deve estar ativa.
- Obrigação cancelada não pode ser paga.
- Obrigação já liquidada não pode receber novo pagamento.

### Status resultante

```text
se total liquidado < valor original: partially_settled
se total liquidado = valor original: settled
```

### Critérios de aceite

- Pagamento total marca obrigação como `settled`.
- Pagamento parcial marca como `partially_settled`.
- Pagamento acima do pendente retorna erro.
- Saldo da conta diminui.
- Movimentação gerada tem `movement_type = exit`.

---

## História 6.2 — Receber conta a receber

Como usuário,
quero receber uma obrigação do tipo `receivable`,
para registrar entrada de dinheiro.

### Regras

- Igual ao fluxo de pagamento, mas gera entrada.
- Categoria deve estar associada à obrigação.
- Movimentação gerada deve ter `movement_type = entry`.

### Critérios de aceite

- Recebimento total marca como `settled`.
- Recebimento parcial marca como `partially_settled`.
- Saldo da conta aumenta.
- Valor acima do pendente retorna erro.

---

## História 6.3 — Consultar liquidações de uma obrigação

Como usuário,
quero consultar os pagamentos/recebimentos de uma obrigação,
para entender como ela foi quitada.

### Critérios de aceite

- Retorna lista de liquidações.
- Mostra data, valor, conta e movimentação.
- Ignora liquidações canceladas por padrão.

---

## História 6.4 — Recalcular status da obrigação

Como sistema,
quero recalcular o status da obrigação após cada liquidação,
para manter o estado correto.

### Regras

```text
total liquidado = 0 -> pending
total liquidado > 0 e < original_amount -> partially_settled
total liquidado = original_amount -> settled
```

### Critérios de aceite

- Status é atualizado corretamente após cada pagamento.
- Status volta corretamente após estorno de liquidação.
