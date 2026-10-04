BEGIN;

DROP TABLE IF EXISTS obligation_settlements CASCADE;
DROP TABLE IF EXISTS financial_transactions CASCADE;
DROP TABLE IF EXISTS financial_accounts CASCADE;
DROP TABLE IF EXISTS categories CASCADE;
DROP TABLE IF EXISTS financial_obligations CASCADE;
DROP TABLE IF EXISTS financial_transfer CASCADE;

DROP TYPE IF EXISTS financial_accounts_type;
DROP TYPE IF EXISTS categories_tpyes;
DROP TYPE IF EXISTS financial_obligations_tpyes;
DROP TYPE IF EXISTS financial_obligations_status;
DROP TYPE IF EXISTS financial_transactions_movement_type;
DROP TYPE IF EXISTS financial_transactions_operation_type;
DROP TYPE IF EXISTS financial_transactions_status;

COMMIT;