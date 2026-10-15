# Ordem de Implementação

## Fase 1 — Fundação técnica

Objetivo: ter API rodando, banco funcionando e padrão mínimo de projeto.

Entregas:

1. Criar repositório Go.
2. Criar `compose.yaml` com PostgreSQL.
3. Criar estrutura de pastas.
4. Criar conexão com banco.
5. Criar servidor HTTP.
6. Criar endpoint `/health`.
7. Criar helpers de resposta JSON.
8. Criar helpers de decode de JSON.
9. Criar padrão de erros de validação.

---

## Fase 2 — Contas financeiras

Objetivo: cadastrar onde o dinheiro fica.

Entregas:

1. Criar tabela `financial_accounts`.
2. Criar CRUD de contas financeiras.
3. Permitir informar saldo inicial.
4. Listar contas ativas.
5. Desativar conta.
6. Bloquear desativação se houver regra futura necessária.

---

## Fase 3 — Categorias

Objetivo: classificar entradas e saídas.

Entregas:

1. Criar tabela `categories`.
2. Criar CRUD de categorias.
3. Permitir categoria pai.
4. Validar tipo.
5. Validar ciclos.
6. Desativar categoria pai e filhas.
7. Não permitir uso de categoria inativa em novos lançamentos.

---

## Fase 4 — Movimentações avulsas

Objetivo: registrar entradas e saídas simples.

Entregas:

1. Criar tabela `financial_transactions`.
2. Registrar entrada avulsa.
3. Registrar saída avulsa.
4. Listar extrato.
5. Consultar saldo por conta.
6. Consultar saldo total.
7. Impedir valores inválidos.
8. Aplicar idempotência.

---

## Fase 5 — Obrigações financeiras

Objetivo: controlar contas futuras a pagar/receber.

Entregas:

1. Criar tabela `financial_obligations`.
2. Cadastrar conta a pagar.
3. Cadastrar conta a receber.
4. Listar pendentes.
5. Listar vencidas.
6. Cancelar obrigação.
7. Atualizar obrigação pendente.

---

## Fase 6 — Liquidações

Objetivo: pagar ou receber obrigações.

Entregas:

1. Criar tabela `obligation_settlements`.
2. Liquidar obrigação total.
3. Liquidar obrigação parcial.
4. Atualizar status para `settled`.
5. Atualizar status para `partially_settled`.
6. Impedir pagamento maior que saldo pendente.
7. Gerar movimentação automaticamente.
8. Usar transação de banco.

---

## Fase 7 — Transferências

Objetivo: mover dinheiro entre contas próprias.

Entregas:

1. Criar tabela `financial_transfers`.
2. Criar transferência.
3. Gerar saída na origem.
4. Gerar entrada no destino.
5. Impedir conta origem igual destino.
6. Aplicar idempotência.
7. Cancelar/estornar transferência.

---

## Fase 8 — Estornos e cancelamentos

Objetivo: corrigir erros mantendo histórico.

Entregas:

1. Estornar movimentação avulsa.
2. Estornar liquidação.
3. Cancelar obrigação sem liquidação.
4. Cancelar transferência.
5. Impedir estorno duplicado.
6. Vincular transação original e transação reversa.

---

## Fase 9 — Dashboard

Objetivo: enxergar situação financeira.

Entregas:

1. Saldo total.
2. Saldo por conta.
3. Entradas do mês.
4. Saídas do mês.
5. Resultado mensal.
6. Obrigações vencidas.
7. Próximos vencimentos.
8. Gastos por categoria.
9. Reserva financeira.

---

## Fase 10 — Relatórios básicos

Objetivo: consultar histórico de forma útil.

Entregas:

1. Extrato por conta.
2. Extrato por categoria.
3. Resultado por mês.
4. Obrigações por período.
5. Gastos agrupados por categoria.
6. Exportação futura opcional.
