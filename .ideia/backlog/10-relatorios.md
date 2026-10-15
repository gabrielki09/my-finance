# Épico 10 — Relatórios Básicos

## Objetivo

Permitir consultas úteis para análise financeira.

---

## História 10.1 — Extrato por conta

Como usuário,
quero ver o extrato de uma conta,
para conferir todas as movimentações dela.

### Filtros

```text
account_id
start_date
end_date
movement_type
```

### Critérios de aceite

- Lista movimentações ordenadas por data.
- Mostra saldo final do período.
- Mostra entradas e saídas.

---

## História 10.2 — Resultado mensal

Como usuário,
quero ver resultado por mês,
para saber se sobrou ou faltou dinheiro.

### Critérios de aceite

- Agrupa por mês.
- Mostra entradas.
- Mostra saídas.
- Mostra resultado líquido.

---

## História 10.3 — Obrigações por período

Como usuário,
quero ver contas por vencimento,
para planejar o mês.

### Critérios de aceite

- Filtra por data de vencimento.
- Separa pagar e receber.
- Mostra status.
- Mostra valor pendente.

---

## História 10.4 — Gastos por categoria e subcategoria

Como usuário,
quero analisar gastos por categoria,
para identificar onde cortar despesas.

### Critérios de aceite

- Agrupa por categoria filha.
- Permite futuramente agrupar por categoria pai.
- Mostra total e percentual.
