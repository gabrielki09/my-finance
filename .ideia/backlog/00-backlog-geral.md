# Backlog Geral

## Épicos

| Ordem | Épico | Objetivo | Prioridade |
|---:|---|---|---|
| 1 | Infraestrutura | Base técnica do projeto | Alta |
| 2 | Contas financeiras | Controlar onde o dinheiro fica | Alta |
| 3 | Categorias | Classificar entradas e saídas | Alta |
| 4 | Movimentações | Registrar entradas e saídas reais | Alta |
| 5 | Obrigações | Controlar contas a pagar/receber | Alta |
| 6 | Liquidações | Pagar/receber obrigações | Alta |
| 7 | Transferências | Mover dinheiro entre contas | Média |
| 8 | Estornos e cancelamentos | Corrigir erros mantendo histórico | Alta |
| 9 | Dashboard | Exibir resumo financeiro | Média |
| 10 | Relatórios | Consultas históricas | Baixa/Média |

---

## MVP mínimo utilizável

O MVP mínimo deve permitir:

1. cadastrar contas financeiras;
2. cadastrar categorias;
3. registrar entrada avulsa;
4. registrar saída avulsa;
5. consultar saldo por conta;
6. cadastrar conta a pagar;
7. pagar conta a pagar;
8. cadastrar conta a receber;
9. receber conta a receber;
10. estornar uma movimentação incorreta.

---

## Definição de pronto geral

Uma história só deve ser considerada pronta quando:

- endpoint funciona;
- validações principais estão implementadas;
- erros retornam JSON padronizado;
- regra de negócio foi respeitada;
- banco mantém integridade;
- operação crítica usa transação quando necessário;
- fluxo foi testado manualmente via HTTP client;
- documentação do endpoint foi atualizada.
