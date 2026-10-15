# Épico 3 — Categorias

## Objetivo

Permitir classificar receitas e despesas.

Exemplos:

```text
Alimentação
  Mercado
  Delivery
Moradia
  Luz
  Água
Salário
Freelance
```

---

## História 3.1 — Cadastrar categoria

Como usuário,
quero cadastrar categorias,
para classificar meus lançamentos financeiros.

### Dados

```text
parent_id
name
type
active
```

### Tipos

```text
income
expense
both
```

### Regras

- Nome é obrigatório.
- Tipo é obrigatório.
- Tipo deve ser válido.
- `parent_id` é opcional.
- Se `parent_id` for informado, a categoria pai deve existir.
- Categoria pai deve estar ativa.
- Categoria filha deve ter tipo compatível com o pai.
- Categoria começa ativa.

### Compatibilidade de tipo sugerida

| Filha | Pai permitido |
|---|---|
| income | income, both |
| expense | expense, both |
| both | both |

### Critérios de aceite

- Categoria sem pai é cadastrada.
- Categoria com pai válido é cadastrada.
- Categoria com pai inexistente retorna 422 ou 404.
- Categoria com tipo incompatível retorna 422.
- Categoria sem nome retorna 422.

---

## História 3.2 — Listar categorias

Como usuário,
quero listar categorias,
para selecionar nos lançamentos e visualizar a árvore.

### Regras

- Permitir listar em formato plano.
- Futuramente permitir listar em formato árvore.
- Permitir filtro por tipo.
- Permitir filtro por ativas.

### Critérios de aceite

- Lista categorias cadastradas.
- Filtra por `type=income`.
- Filtra por `type=expense`.
- Filtra por `active=true`.

---

## História 3.3 — Atualizar categoria

Como usuário,
quero atualizar uma categoria,
para corrigir nome, tipo ou categoria pai.

### Regras

- Pode alterar nome.
- Pode alterar tipo se não quebrar compatibilidade das filhas.
- Pode alterar categoria pai.
- Não pode definir ela mesma como pai.
- Não pode criar ciclo.
- Não pode vincular em pai inativo.
- Não deve alterar transações antigas.

### Critérios de aceite

- Alteração válida retorna 200.
- Definir a própria categoria como pai retorna 422.
- Criar ciclo retorna 422.
- Pai inexistente retorna erro.
- Pai inativo retorna erro.

---

## História 3.4 — Desativar categoria

Como usuário,
quero desativar uma categoria,
para não usá-la em novos lançamentos.

### Regras

- Desativar categoria pai desativa filhas.
- Desativar categoria não altera transações antigas.
- Categoria inativa não pode ser usada em novas movimentações.
- Categoria inativa não pode ser usada em novas obrigações.
- Histórico deve continuar exibindo a categoria.

### Critérios de aceite

- Categoria pai e filhas ficam inativas.
- Transações antigas continuam com `category_id`.
- Nova movimentação com categoria inativa retorna erro.
