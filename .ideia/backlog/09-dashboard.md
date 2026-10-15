# Épico 9 — Dashboard

## Objetivo

Exibir visão rápida da situação financeira.

---

## História 9.1 — Saldo total

Como usuário,
quero ver meu saldo total,
para saber quanto dinheiro tenho no momento.

### Fórmula

```text
saldo total = soma dos saldos das contas
```

### Critérios de aceite

- Retorna saldo em centavos.
- Pode retornar saldo formatado futuramente.
- Considera saldos iniciais e movimentações.

---

## História 9.2 — Saldo por conta

Como usuário,
quero ver saldo por conta,
para saber onde meu dinheiro está.

### Critérios de aceite

- Lista cada conta e seu saldo.
- Permite filtrar somente contas ativas.
- Ordena por nome ou maior saldo.

---

## História 9.3 — Entradas e saídas do mês

Como usuário,
quero ver entradas e saídas do mês,
para entender meu resultado financeiro.

### Regras

- Considerar apenas movimentações originais e ajustes.
- Não considerar transferências como receita/despesa.
- Estornos devem impactar resultado conforme operação inversa ou serem tratados separadamente no relatório.

### Critérios de aceite

- Retorna total de entradas do mês.
- Retorna total de saídas do mês.
- Retorna resultado: entradas - saídas.

---

## História 9.4 — Obrigações vencidas

Como usuário,
quero ver contas vencidas,
para saber o que precisa de atenção.

### Regra

Obrigação vencida:

```text
status em pending ou partially_settled
e due_date < data atual
```

### Critérios de aceite

- Lista vencidas a pagar.
- Lista vencidas a receber.
- Mostra dias de atraso.
- Mostra valor pendente.

---

## História 9.5 — Próximos vencimentos

Como usuário,
quero ver próximos vencimentos,
para me planejar.

### Critérios de aceite

- Lista obrigações com vencimento nos próximos X dias.
- X pode começar fixo em 7 ou 15 dias.
- Mostra tipo, valor e data.

---

## História 9.6 — Gastos por categoria

Como usuário,
quero ver gastos por categoria,
para entender onde estou gastando mais.

### Regras

- Considerar saídas.
- Agrupar por categoria.
- Permitir período.

### Critérios de aceite

- Retorna categoria e total gasto.
- Ordena maior gasto primeiro.
- Permite filtrar mês atual.
