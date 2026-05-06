# Subscription System`

## Features

- User registration & login
- JWT-based authentication
- Role-based access control: `user` dan `admin`
- Plan management untuk admin
- Public plan listing untuk user/guest
- Feature/entitlement management per plan
- Create subscription ke plan tertentu
- Support subscription status lifecycle:
  - `incomplete`
  - `trialing`
  - `active`
  - `past_due`
  - `canceled`
  - `expired`
  - `paused`
- Get current user subscription
- Cancel subscription secara langsung
- Cancel subscription at period end
- Upgrade subscription plan
- Downgrade subscription plan
- Automatic invoice generation saat subscribe ke paid plan
- User invoice history
- Invoice detail
- Mock payment provider untuk simulasi payment
- Simulate payment success
- Simulate payment failure
- Webhook endpoint untuk payment provider
- Webhook signature verification
- Idempotent webhook processing untuk mencegah double processing
- Duplicate webhook detection
- Transaction-safe payment processing
- Payment status tracking:
  - `pending`
  - `succeeded`
  - `failed`
  - `expired`
  - `canceled`
- Invoice status tracking:
  - `draft`
  - `open`
  - `paid`
  - `failed`
  - `expired`
  - `void`
- Premium feature access control berdasarkan subscription aktif
- Example protected premium endpoint
- Admin dashboard API untuk melihat semua subscriptions
- Admin dapat melihat detail subscription user
- Admin dapat extend subscription user
- Admin dapat cancel subscription user
- Audit log untuk aktivitas penting
- Soft delete untuk data utama
- PostgreSQL relational schema
- Redis-ready untuk cache, lock, atau background job
- Docker Compose-ready untuk local development
- OpenAPI/Swagger-ready API documentation
- Clean architecture / layered architecture friendly

---

## Database Structure

### 1. users

| Field | Type |
|---|---|
| id | UUID |
| email | TEXT |
| password_hash | TEXT |
| full_name | TEXT |
| role | TEXT |
| is_deleted | BOOLEAN |
| deleted_at | TIMESTAMPTZ |
| created_at | TIMESTAMPTZ |
| updated_at | TIMESTAMPTZ |

### 2. plans

| Field | Type |
|---|---|
| id | UUID |
| code | TEXT |
| name | TEXT |
| description | TEXT |
| price_amount | BIGINT |
| currency | TEXT |
| billing_interval | TEXT |
| is_active | BOOLEAN |
| is_deleted | BOOLEAN |
| deleted_at | TIMESTAMPTZ |
| created_at | TIMESTAMPTZ |
| updated_at | TIMESTAMPTZ |

### 3. features

| Field | Type |
|---|---|
| id | UUID |
| code | TEXT |
| name | TEXT |
| description | TEXT |
| is_deleted | BOOLEAN |
| deleted_at | TIMESTAMPTZ |
| created_at | TIMESTAMPTZ |
| updated_at | TIMESTAMPTZ |

### 4. plan_features

| Field | Type |
|---|---|
| plan_id | UUID |
| feature_id | UUID |
| is_deleted | BOOLEAN |
| deleted_at | TIMESTAMPTZ |
| created_at | TIMESTAMPTZ |
| updated_at | TIMESTAMPTZ |

### 5. subscriptions

| Field | Type |
|---|---|
| id | UUID |
| user_id | UUID |
| plan_id | UUID |
| status | TEXT |
| started_at | TIMESTAMPTZ |
| current_period_start | TIMESTAMPTZ |
| current_period_end | TIMESTAMPTZ |
| cancel_at_period_end | BOOLEAN |
| canceled_at | TIMESTAMPTZ |
| ended_at | TIMESTAMPTZ |
| is_deleted | BOOLEAN |
| deleted_at | TIMESTAMPTZ |
| created_at | TIMESTAMPTZ |
| updated_at | TIMESTAMPTZ |

### 6. invoices

| Field | Type |
|---|---|
| id | UUID |
| user_id | UUID |
| subscription_id | UUID |
| invoice_number | TEXT |
| amount | BIGINT |
| currency | TEXT |
| status | TEXT |
| due_at | TIMESTAMPTZ |
| paid_at | TIMESTAMPTZ |
| is_deleted | BOOLEAN |
| deleted_at | TIMESTAMPTZ |
| created_at | TIMESTAMPTZ |
| updated_at | TIMESTAMPTZ |

### 7. payments

| Field | Type |
|---|---|
| id | UUID |
| invoice_id | UUID |
| provider | TEXT |
| provider_payment_id | TEXT |
| amount | BIGINT |
| currency | TEXT |
| status | TEXT |
| paid_at | TIMESTAMPTZ |
| failed_at | TIMESTAMPTZ |
| is_deleted | BOOLEAN |
| deleted_at | TIMESTAMPTZ |
| created_at | TIMESTAMPTZ |
| updated_at | TIMESTAMPTZ |


### 8. provider_events

| Field | Type |
|---|---|
| id | UUID |
| provider | TEXT |
| provider_event_id | TEXT |
| event_type | TEXT |
| payload | JSONB |
| signature | TEXT |
| processed_at | TIMESTAMPTZ |
| created_at | TIMESTAMPTZ |

### 9. audit_logs

| Field | Type |
|---|---|
| id | UUID |
| actor_type | TEXT |
| actor_id | UUID |
| action | TEXT |
| entity_type | TEXT |
| entity_id | UUID |
| metadata | JSONB |
| created_at | TIMESTAMPTZ |

---

## API Specification

### Auth API

| Method | Endpoint | Access | Description |
|---|---|---|---|
| POST | `/api/v1/auth/register` | Public | Register user |
| POST | `/api/v1/auth/login` | Public | Login user |
| GET | `/api/v1/auth/me` | User | Get current user |


### Plan API

| Method | Endpoint | Access | Description |
|---|---|---|---|
| GET | `/api/v1/plans` | Public | List active plans |
| GET | `/api/v1/plans/:id` | Public | Get plan detail |
| POST | `/api/v1/admin/plans` | Admin | Create plan |
| PATCH | `/api/v1/admin/plans/:id` | Admin | Update plan |
| DELETE | `/api/v1/admin/plans/:id` | Admin | Archive plan |

### Subscription API

| Method | Endpoint | Access | Description |
|---|---|---|---|
| POST | `/api/v1/subscriptions` | User | Create subscription |
| GET | `/api/v1/subscriptions/current` | User | Get current subscription |
| POST | `/api/v1/subscriptions/current/cancel` | User | Cancel subscription |
| POST | `/api/v1/subscriptions/current/upgrade` | User | Upgrade plan |
| POST | `/api/v1/subscriptions/current/downgrade` | User | Downgrade plan |

### Invoice API

| Method | Endpoint | Access | Description |
|---|---|---|---|
| GET | `/api/v1/invoices` | User | List own invoices |
| GET | `/api/v1/invoices/:id` | User | Get invoice detail |

### Payment API

| Method | Endpoint | Access | Description |
|---|---|---|---|
| POST | `/api/v1/payments/mock/:invoice_id` | User | Create mock payment |
| POST | `/api/v1/mock-provider/payments/:payment_id/success` | Dev Only | Simulate success |
| POST | `/api/v1/mock-provider/payments/:payment_id/fail` | Dev Only | Simulate failure |


### Webhook API

| Method | Endpoint | Access | Description |
|---|---|---|---|
| POST | `/api/v1/webhooks/mock-payment` | Provider | Receive mock payment webhook |


### Entitlement API

| Method | Endpoint | Access | Description |
|---|---|---|---|
| GET | `/api/v1/entitlements/me` | User | Get current user features |
| GET | `/api/v1/premium/reports` | User | Example protected premium feature |

### Admin API

| Method | Endpoint | Access | Description |
|---|---|---|---|
| GET | `/api/v1/admin/subscriptions` | Admin | List all subscriptions |
| GET | `/api/v1/admin/subscriptions/:id` | Admin | Get subscription detail |
| POST | `/api/v1/admin/subscriptions/:id/extend` | Admin | Extend subscription |
| POST | `/api/v1/admin/subscriptions/:id/cancel` | Admin | Cancel subscription |
| GET | `/api/v1/admin/audit-logs` | Admin | List audit logs |