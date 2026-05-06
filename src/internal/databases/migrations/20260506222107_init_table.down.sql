-- =========================================================
-- DOWN MIGRATION
-- =========================================================

DROP TABLE IF EXISTS audit_logs;
DROP TABLE IF EXISTS provider_events;
DROP TABLE IF EXISTS payments;
DROP TABLE IF EXISTS invoices;
DROP TABLE IF EXISTS subscriptions;
DROP TABLE IF EXISTS plan_features;
DROP TABLE IF EXISTS features;
DROP TABLE IF EXISTS plans;
DROP TABLE IF EXISTS users;

DROP EXTENSION IF EXISTS pgcrypto;