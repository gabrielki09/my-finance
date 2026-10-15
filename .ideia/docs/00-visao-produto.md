# Visão do Produto

## Produto

Aplicação pessoal para gerenciamento financeiro, com foco em controle de contas, entradas, saídas, obrigações, saldos e planejamento básico.

## Objetivo principal

Permitir que o usuário saiba:

- quanto dinheiro possui;
- onde o dinheiro está;
- quanto entrou;
- quanto saiu;
- quais contas ainda precisa pagar;
- quais valores ainda espera receber;
- quais contas estão vencidas;
- quanto consegue guardar;
- como os saldos evoluem ao longo do tempo.

## Escopo inicial

O produto será usado por uma única pessoa, portanto o MVP não terá:

- autenticação;
- tela de login;
- múltiplos usuários;
- permissões;
- compartilhamento;
- multiempresa;
- plano de assinatura;
- frontend.

## Stack definida

- Go
- `net/http`
- PostgreSQL
- Docker
- SQL manual
- API REST
- JSON

## Módulos do MVP

1. Infraestrutura
2. Contas financeiras
3. Categorias
4. Movimentações avulsas
5. Obrigações financeiras
6. Liquidações
7. Transferências
8. Estornos e cancelamentos
9. Dashboard
10. Relatórios básicos

## Conceitos centrais

### Conta financeira

Representa onde o dinheiro fica.

Exemplos:

- Nubank
- Carteira física
- Sicoob
- Reserva de emergência
- Poupança
- Conta investimento

### Categoria

Representa o motivo/classificação da entrada ou saída.

Exemplos:

- Alimentação
- Moradia
- Transporte
- Salário
- Freelance
- Saúde

### Movimentação financeira

Representa dinheiro que realmente entrou ou saiu de uma conta.

Exemplos:

- pagamento do mercado;
- recebimento de salário;
- pagamento de conta de luz;
- recebimento de freelance.

### Obrigação financeira

Representa algo a pagar ou receber, mas que ainda não necessariamente movimentou dinheiro.

Exemplos:

- conta de internet vencendo no dia 10;
- salário previsto para receber no dia 5;
- parcela do notebook;
- aluguel do mês.

### Liquidação

Representa o pagamento ou recebimento de uma obrigação.

Uma obrigação de R$ 300,00 pode ser liquidada parcialmente:

- pagamento 1: R$ 100,00;
- pagamento 2: R$ 200,00.

### Transferência

Representa movimentação entre contas próprias.

Exemplo:

- Nubank → Reserva de emergência.

Transferência não é receita nem despesa.

### Estorno

Representa a reversão de uma movimentação efetivada.
Não deve apagar histórico.
