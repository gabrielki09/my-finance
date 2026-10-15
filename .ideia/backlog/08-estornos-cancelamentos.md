# Épico 8 — Estornos e Cancelamentos

## Objetivo

Corrigir erros financeiros sem apagar histórico.

---

## História 8.1 — Estornar movimentação avulsa

Como usuário,
quero estornar uma movimentação incorreta,
para corrigir o saldo mantendo rastreabilidade.

### Regras

- Movimentação original não pode estar cancelada.
- Movimentação original não pode já ter estorno.
- Criar movimentação inversa:
  - entrada vira saída;
  - saída vira entrada.
- Operação reversa deve ter `operation_type = reversal`.
- Vincular estorno à movimentação original.
- Marcar original como cancelada, se essa for a decisão do modelo.
- Usar transação de banco.

### Critérios de aceite

- Estorno de saída aumenta saldo.
- Estorno de entrada diminui saldo.
- Não permite estornar duas vezes.
- Histórico mostra original e reversa.

---

## História 8.2 — Estornar liquidação

Como usuário,
quero estornar o pagamento/recebimento de uma obrigação,
para corrigir uma baixa feita errado.

### Regras

- Encontrar liquidação.
- Estornar movimentação vinculada.
- Marcar liquidação como cancelada.
- Recalcular status da obrigação.
- Usar transação de banco.

### Critérios de aceite

- Obrigação volta para `pending` se não restar liquidação.
- Obrigação volta para `partially_settled` se ainda houver liquidação parcial.
- Saldo da conta é corrigido.
- Não permite estorno duplicado.

---

## História 8.3 — Cancelar obrigação

Como usuário,
quero cancelar uma obrigação sem pagamento,
para remover uma previsão que não existe mais.

### Regras

- Só permitir cancelamento direto se não houver liquidação ativa.
- Se houver liquidação, exigir estorno primeiro.
- Status vira `canceled`.
- Preencher `canceled_at`.

### Critérios de aceite

- Obrigação pendente é cancelada.
- Obrigação cancelada não aparece em pendentes.
- Obrigação cancelada não pode ser paga/recebida.

---

## História 8.4 — Cancelar conta/categoria

Como usuário,
quero desativar cadastros,
para impedir uso futuro sem apagar histórico.

### Regras

- Conta financeira desativada não pode ser usada em novos movimentos.
- Categoria desativada não pode ser usada em novos movimentos.
- Histórico permanece.
- Desativar categoria pai desativa filhas.

### Critérios de aceite

- Cadastro fica inativo.
- Histórico permanece consultável.
- Novas operações bloqueiam cadastros inativos.
