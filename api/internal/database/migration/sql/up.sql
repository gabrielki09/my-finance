CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE TYPE financial_accounts_type AS ENUM('cash', 'checking', 'savings', 'digital');
CREATE TABLE financial_accounts (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    type financial_accounts_type NOT NULL,
    initial_balance DECIMAL(14, 2) NOT NULL DEFAULT 0,
    opened_at DATE NOT NULL DEFAULT CURRENT_DATE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ 
);
CREATE TYPE categories_tpyes AS ENUM('income', 'expense', 'both');
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    type categories_tpyes NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ 
);
ALTER TABLE categories
ADD COLUMN parent_id INTEGER;
ALTER TABLE categories
ADD CONSTRAINT categories_parent_id_foreign
FOREIGN KEY (parent_id)
REFERENCES categories(id)
ON DELETE SET NULL
ON UPDATE CASCADE;

CREATE TYPE financial_obligations_tpyes AS ENUM('payable', 'receivable');
CREATE TYPE financial_obligations_status AS ENUM('pending', 'partially_settled', 'settled', 'canceled');
CREATE TABLE financial_obligations (
    id SERIAL PRIMARY KEY,
    category_id INTEGER REFERENCES categories(id),
    description VARCHAR(255) NOT NULL,
    type financial_obligations_tpyes NOT NULL,
    status financial_obligations_status DEFAULT 'pending',
    original_amount DECIMAL(14, 2) NOT NULL,
    due_date DATE NOT NULL,
	competence_date DATE NULL,
    notes TEXT NULL,
    canceled_at  TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ 
);
CREATE TYPE financial_transactions_movement_type AS ENUM('entry', 'exit');
CREATE TYPE financial_transactions_operation_type AS ENUM('original', 'adjustment', 'reversal');
CREATE TYPE financial_transactions_status AS ENUM('confirmed', 'canceled');
CREATE TABLE financial_transactions (
    id SERIAL PRIMARY KEY,   
    financial_account_id INTEGER NOT NULL REFERENCES financial_accounts(id),
    category_id INTEGER NOT NULL REFERENCES categories(id),
    reversed_transaction_id INTEGER NULL REFERENCES financial_transactions(id),
    description VARCHAR(255) NOT NULL,
    movement_type financial_transactions_movement_type NOT NULL,
    operation_type financial_transactions_operation_type NOT NULL,
    status financial_transactions_status NOT NULL DEFAULT 'confirmed',
    amount DECIMAL(14, 2) NOT NULL,
    movement_date TIMESTAMPTZ NOT NULL,
    reference_date DATE,
    origin_type VARCHAR(50) NULL,
    origin_id INTEGER NULL,
    idempotency_key UUID DEFAULT gen_random_uuid() UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    canceled_at  TIMESTAMPTZ
);
CREATE TABLE obligation_settlements (
    id SERIAL PRIMARY KEY,
    obligation_id INTEGER NOT NULL REFERENCES financial_obligations(id),
    transaction_id INTEGER NOT NULL REFERENCES financial_transactions(id),
    amount DECIMAL(14, 2) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    canceled_at  TIMESTAMPTZ
);
CREATE TABLE financial_transfer (
    id SERIAL PRIMARY KEY,
    transfer_date DATE,
    source_account_id INTEGER NOT NULL REFERENCES financial_accounts(id),
    destination_account_id INTEGER NOT NULL REFERENCES financial_accounts(id),
    idempotency_key UUID DEFAULT gen_random_uuid() UNIQUE,
    amount DECIMAL(14, 2) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    canceled_at  TIMESTAMPTZ
);  