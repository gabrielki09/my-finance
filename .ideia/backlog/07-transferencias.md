# Épico 7 — Transferências

## Objetivo

Permitir mover dinheiro entre contas próprias sem classificar como receita ou despesa.

---

## História 7.1 — Criar transferência entre contas

Como usuário,
quero transferir dinheiro de uma conta para outra,
para controlar movimentações entre meus próprios saldos.

### Exemplo

```text
Nubank -> Reserva de emergência
R$ 300,00
```

### Dados

```text
source_account_id
destination_account_id
amount
description
transfer_date
idempotency_key
```

### Regras

- Conta origem deve existir e estar ativa.
- Conta destino deve existir e estar ativa.
- Origem e destino não podem ser iguais.
- Valor deve ser maior que zero.
- Transferência gera duas movimentações:
  - saída na origem;
  - entrada no destino.
- Não usa categoria obrigatória.
- Não entra como receita/despesa em resultado mensal.
- Deve usar transação de banco.

### Critérios de aceite

- Transferência válida retorna 201.
- Saldo da origem diminui.
- Saldo do destino aumenta.
- Resultado mensal não considera transferência como gasto/receita.
- Duplicidade por idempotência não gera outra transferência.

---

## História 7.2 — Listar transferências

Como usuário,
quero listar transferências,
para auditar movimentações entre contas.

### Filtros

```text
source_account_id
destination_account_id
start_date
end_date
```

### Critérios de aceite

- Lista transferências por data.
- Mostra origem, destino e valor.
- Mostra se foi cancelada.

---

## História 7.3 — Cancelar transferência

Como usuário,
quero cancelar uma transferência incorreta,
para reverter seus efeitos.

### Regras

- Gerar estorno da saída original.
- Gerar estorno da entrada original.
- Marcar transferência como cancelada.
- Impedir cancelamento duplicado.
- Usar transação de banco.

### Critérios de aceite

- Saldos voltam ao estado anterior.
- Histórico original permanece.
- Transferência fica com `canceled_at`.
