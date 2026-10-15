# Backlog — App Pessoal de Finanças

Este pacote contém o planejamento completo do app pessoal de gerenciamento financeiro.

## Premissas técnicas

- Uso pessoal.
- Sem login.
- Sem usuário.
- Sem frontend inicialmente.
- Backend em Go.
- API usando `net/http`.
- Banco PostgreSQL.
- Docker apenas para infraestrutura.
- SQL feito manualmente por enquanto.
- Valores monetários armazenados em centavos usando inteiro.
- Histórico financeiro não deve ser apagado; correções devem ocorrer via cancelamento/estorno.

## Ordem sugerida de leitura

1. `docs/00-visao-produto.md`
2. `docs/01-regras-de-negocio.md`
3. `docs/02-modelo-conceitual.md`
4. `docs/03-ordem-de-implementacao.md`
5. `backlog/00-backlog-geral.md`
6. `backlog/01-infraestrutura.md`
7. `backlog/02-contas-financeiras.md`
8. `backlog/03-categorias.md`
9. `backlog/04-movimentacoes.md`
10. `backlog/05-obrigacoes.md`
11. `backlog/06-liquidacoes.md`
12. `backlog/07-transferencias.md`
13. `backlog/08-estornos-cancelamentos.md`
14. `backlog/09-dashboard.md`
15. `fluxogramas/`

## Como usar

Use os arquivos da pasta `backlog/` como cards de implementação.
Cada item possui objetivo, regras, fluxo, critérios de aceite e sugestões de endpoints.
Os arquivos `.mmd` da pasta `fluxogramas/` podem ser abertos em ferramentas compatíveis com Mermaid.
