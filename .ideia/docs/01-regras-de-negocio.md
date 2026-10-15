# Regras de Negócio

## RN001 — Dinheiro deve ser tratado em centavos

Todos os valores monetários devem ser armazenados como inteiros em centavos.

Exemplos:

| Valor real | Valor salvo |
|---|---:|
| R$ 10,00 | 1000 |
| R$ 150,50 | 15050 |
| R$ 0,99 | 99 |

Não usar `float32` ou `float64` para valores monetários.

---

## RN002 — Saldo inicial não é receita

Ao cadastrar uma conta financeira, o campo `initial_balance` representa o valor existente antes do sistema começar a controlar aquela conta.

Esse valor:

- não é receita;
- não é movimentação comum;
- não entra em relatórios de entrada;
- serve apenas como base para cálculo de saldo.

Saldo atual:

```text
saldo atual = saldo inicial + entradas efetivadas - saídas efetivadas
```

---

## RN003 — Toda movimentação precisa afetar uma conta financeira

Toda entrada ou saída deve informar uma conta financeira.

Exemplo:

```text
Recebi salário no Nubank.
Paguei mercado com dinheiro da Carteira.
```

---

## RN004 — Categoria classifica motivo, conta financeira classifica origem/destino

Exemplo:

```text
Conta financeira: Nubank
Categoria: Alimentação
Movimentação: Saída de R$ 80,00
Descrição: Mercado
```

---

## RN005 — Categorias inativas não podem ser usadas em novos lançamentos

Categorias desativadas devem continuar aparecendo em histórico, mas não devem ser selecionáveis em novas movimentações ou obrigações.

---

## RN006 — Desativar categoria pai deve desativar filhas

Quando uma categoria pai for desativada, suas categorias filhas também devem ser desativadas.

Transações antigas não devem ser alteradas.

---

## RN007 — Alterar categoria pai é permitido com validações

Ao alterar `parent_id` de uma categoria:

1. o pai deve existir;
2. o pai deve estar ativo;
3. a categoria não pode ser pai dela mesma;
4. a alteração não pode criar ciclo;
5. o tipo do pai deve ser compatível com o tipo da filha.

---

## RN008 — Obrigação financeira não movimenta saldo sozinha

Cadastrar uma conta a pagar ou receber não altera saldo.

O saldo só muda quando houver liquidação, ou seja, pagamento/recebimento.

---

## RN009 — Obrigação pode ser paga parcialmente

Uma obrigação pode ter uma ou várias liquidações.

Status possíveis:

- `pending`
- `partially_settled`
- `settled`
- `canceled`

---

## RN010 — Obrigação vencida deve ser condição calculada

Não salvar `overdue` como status permanente.

Uma obrigação está vencida quando:

```text
status em pending ou partially_settled
e due_date < data atual
```

---

## RN011 — Movimentação efetivada não deve ser excluída fisicamente

Movimentações efetivadas devem ser corrigidas por estorno ou cancelamento.

Não apagar histórico financeiro.

---

## RN012 — Estorno deve gerar movimentação inversa

Se a movimentação original foi saída, o estorno será entrada.

Se a movimentação original foi entrada, o estorno será saída.

Exemplo:

```text
Original: Saída R$ 100,00
Estorno: Entrada R$ 100,00
```

---

## RN013 — Transferência gera duas movimentações

Uma transferência entre contas próprias gera:

1. saída na conta de origem;
2. entrada na conta de destino.

Não altera resultado mensal de receitas/despesas.

---

## RN014 — Idempotência deve evitar duplicidade em operações sensíveis

Pagamentos, recebimentos, transferências e estornos devem ter uma chave de idempotência para evitar duplicidade por duplo clique ou retry.

---

## RN015 — Operações críticas devem usar transação de banco

Devem usar transação no PostgreSQL:

- liquidar obrigação;
- transferir entre contas;
- estornar movimentação;
- cancelar liquidação;
- atualizar status financeiro derivado.
