-- =========================================================
-- UP MIGRATION
-- =========================================================

CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- =========================
-- users
-- =========================
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email TEXT NOT NULL,
    password_hash TEXT NOT NULL,
    full_name TEXT NOT NULL,
    role TEXT NOT NULL DEFAULT 'user',
    is_deleted BOOLEAN NOT NULL DEFAULT false,
    deleted_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    CONSTRAINT chk_users_role
        CHECK (role IN ('user', 'admin'))
);

CREATE UNIQUE INDEX uq_users_email_active
ON users(email)
WHERE is_deleted = false;

CREATE INDEX idx_users_email
ON users(email);

CREATE INDEX idx_users_is_deleted
ON users(is_deleted);


-- =========================
-- plans
-- =========================
CREATE TABLE plans (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    code TEXT NOT NULL,
    name TEXT NOT NULL,
    description TEXT,
    price_amount BIGINT NOT NULL DEFAULT 0,
    currency TEXT NOT NULL DEFAULT 'IDR',
    billing_interval TEXT NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT true,
    is_deleted BOOLEAN NOT NULL DEFAULT false,
    deleted_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    CONSTRAINT chk_plans_price_amount
        CHECK (price_amount >= 0),

    CONSTRAINT chk_plans_billing_interval
        CHECK (billing_interval IN ('monthly', 'yearly')),

    CONSTRAINT chk_plans_currency
        CHECK (currency <> '')
);

CREATE UNIQUE INDEX uq_plans_code_active
ON plans(code)
WHERE is_deleted = false;

CREATE INDEX idx_plans_code
ON plans(code);

CREATE INDEX idx_plans_is_active
ON plans(is_active);

CREATE INDEX idx_plans_is_deleted
ON plans(is_deleted);


-- =========================
-- features
-- =========================
CREATE TABLE features (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    code TEXT NOT NULL,
    name TEXT NOT NULL,
    description TEXT,
    is_deleted BOOLEAN NOT NULL DEFAULT false,
    deleted_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX uq_features_code_active
ON features(code)
WHERE is_deleted = false;

CREATE INDEX idx_features_code
ON features(code);

CREATE INDEX idx_features_is_deleted
ON features(is_deleted);


-- =========================
-- plan_features
-- =========================
CREATE TABLE plan_features (
    plan_id UUID NOT NULL REFERENCES plans(id),
    feature_id UUID NOT NULL REFERENCES features(id),
    limit_value INTEGER,
    is_deleted BOOLEAN NOT NULL DEFAULT false,
    deleted_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    PRIMARY KEY (plan_id, feature_id),

    CONSTRAINT chk_plan_features_limit_value
        CHECK (limit_value IS NULL OR limit_value >= 0)
);

CREATE INDEX idx_plan_features_plan_id
ON plan_features(plan_id);

CREATE INDEX idx_plan_features_feature_id
ON plan_features(feature_id);

CREATE INDEX idx_plan_features_is_deleted
ON plan_features(is_deleted);


-- =========================
-- subscriptions
-- =========================
CREATE TABLE subscriptions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id),
    plan_id UUID NOT NULL REFERENCES plans(id),
    status TEXT NOT NULL DEFAULT 'incomplete',
    started_at TIMESTAMPTZ,
    current_period_start TIMESTAMPTZ,
    current_period_end TIMESTAMPTZ,
    cancel_at_period_end BOOLEAN NOT NULL DEFAULT false,
    canceled_at TIMESTAMPTZ,
    ended_at TIMESTAMPTZ,
    is_deleted BOOLEAN NOT NULL DEFAULT false,
    deleted_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    CONSTRAINT chk_subscriptions_status
        CHECK (
            status IN (
                'incomplete',
                'trialing',
                'active',
                'past_due',
                'canceled',
                'expired',
                'paused'
            )
        ),

    CONSTRAINT chk_subscription_period
        CHECK (
            current_period_start IS NULL
            OR current_period_end IS NULL
            OR current_period_end > current_period_start
        )
);

CREATE UNIQUE INDEX uq_user_current_subscription
ON subscriptions(user_id)
WHERE status IN ('incomplete', 'trialing', 'active', 'past_due')
AND is_deleted = false;

CREATE INDEX idx_subscriptions_user_id
ON subscriptions(user_id);

CREATE INDEX idx_subscriptions_plan_id
ON subscriptions(plan_id);

CREATE INDEX idx_subscriptions_status
ON subscriptions(status);

CREATE INDEX idx_subscriptions_current_period_end
ON subscriptions(current_period_end);

CREATE INDEX idx_subscriptions_is_deleted
ON subscriptions(is_deleted);


-- =========================
-- invoices
-- =========================
CREATE TABLE invoices (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id),
    subscription_id UUID NOT NULL REFERENCES subscriptions(id),
    invoice_number TEXT NOT NULL,
    amount BIGINT NOT NULL,
    currency TEXT NOT NULL DEFAULT 'IDR',
    status TEXT NOT NULL DEFAULT 'open',
    due_at TIMESTAMPTZ NOT NULL,
    paid_at TIMESTAMPTZ,
    is_deleted BOOLEAN NOT NULL DEFAULT false,
    deleted_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    CONSTRAINT chk_invoices_amount
        CHECK (amount >= 0),

    CONSTRAINT chk_invoices_status
        CHECK (
            status IN (
                'draft',
                'open',
                'paid',
                'failed',
                'expired',
                'void'
            )
        )
);

CREATE UNIQUE INDEX uq_invoices_invoice_number_active
ON invoices(invoice_number)
WHERE is_deleted = false;

CREATE INDEX idx_invoices_user_id
ON invoices(user_id);

CREATE INDEX idx_invoices_subscription_id
ON invoices(subscription_id);

CREATE INDEX idx_invoices_status
ON invoices(status);

CREATE INDEX idx_invoices_due_at
ON invoices(due_at);

CREATE INDEX idx_invoices_is_deleted
ON invoices(is_deleted);


-- =========================
-- payments
-- =========================
CREATE TABLE payments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    invoice_id UUID NOT NULL REFERENCES invoices(id),
    provider TEXT NOT NULL DEFAULT 'mock',
    provider_payment_id TEXT NOT NULL,
    amount BIGINT NOT NULL,
    currency TEXT NOT NULL DEFAULT 'IDR',
    status TEXT NOT NULL DEFAULT 'pending',
    paid_at TIMESTAMPTZ,
    failed_at TIMESTAMPTZ,
    is_deleted BOOLEAN NOT NULL DEFAULT false,
    deleted_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    CONSTRAINT chk_payments_amount
        CHECK (amount >= 0),

    CONSTRAINT chk_payments_status
        CHECK (
            status IN (
                'pending',
                'succeeded',
                'failed',
                'expired',
                'canceled'
            )
        )
);

CREATE UNIQUE INDEX uq_payments_provider_payment_id_active
ON payments(provider, provider_payment_id)
WHERE is_deleted = false;

CREATE INDEX idx_payments_invoice_id
ON payments(invoice_id);

CREATE INDEX idx_payments_status
ON payments(status);

CREATE INDEX idx_payments_provider_payment_id
ON payments(provider, provider_payment_id);

CREATE INDEX idx_payments_is_deleted
ON payments(is_deleted);


-- =========================
-- provider_events
-- =========================
CREATE TABLE provider_events (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    provider TEXT NOT NULL,
    provider_event_id TEXT NOT NULL,
    event_type TEXT NOT NULL,
    payload JSONB NOT NULL,
    signature TEXT,
    processed_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    CONSTRAINT uq_provider_events_provider_event_id
        UNIQUE (provider, provider_event_id),

    CONSTRAINT chk_provider_events_provider
        CHECK (provider <> ''),

    CONSTRAINT chk_provider_events_event_type
        CHECK (event_type <> '')
);

CREATE INDEX idx_provider_events_provider_event_id
ON provider_events(provider, provider_event_id);

CREATE INDEX idx_provider_events_event_type
ON provider_events(event_type);

CREATE INDEX idx_provider_events_processed_at
ON provider_events(processed_at);

CREATE INDEX idx_provider_events_created_at
ON provider_events(created_at);


-- =========================
-- audit_logs
-- =========================
CREATE TABLE audit_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    actor_type TEXT NOT NULL,
    actor_id UUID,
    action TEXT NOT NULL,
    entity_type TEXT NOT NULL,
    entity_id UUID,
    metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    CONSTRAINT chk_audit_logs_actor_type
        CHECK (actor_type IN ('user', 'admin', 'system', 'provider')),

    CONSTRAINT chk_audit_logs_action
        CHECK (action <> ''),

    CONSTRAINT chk_audit_logs_entity_type
        CHECK (entity_type <> '')
);

CREATE INDEX idx_audit_logs_actor
ON audit_logs(actor_type, actor_id);

CREATE INDEX idx_audit_logs_entity
ON audit_logs(entity_type, entity_id);

CREATE INDEX idx_audit_logs_action
ON audit_logs(action);

CREATE INDEX idx_audit_logs_created_at
ON audit_logs(created_at);