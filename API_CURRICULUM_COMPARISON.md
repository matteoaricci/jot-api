# API Curriculum vs Jot-API Implementation - Comprehensive Comparison

**Generated:** June 10, 2026  
**Last Updated:** July 16, 2026  
**Project:** jot-api (Journaling Application)  
**Curriculum:** Build an API - 78-page training curriculum (TODO List application)

---

## Executive Summary

The jot-api project has implemented **most of the critical curriculum tasks**. Since the original June 10 report, authorization enforcement, SQL indexes, a version endpoint, real password hashing, and expanded integration tests have all landed. The application now demonstrates core API development concepts end-to-end and is missing mostly production-observability and advanced-integration features.

### Key Findings:
- ✅ **Strong Foundation**: Database, migrations, CRUD, authentication, **resource ownership/authorization**, **SQL indexes**, **bcrypt password hashing**, **version endpoint**, Docker support
- ⚠️  **Missing Production Features**: Telemetry, API documentation (Swagger — in progress), admin endpoints
- ❌ **No Advanced Features**: Email, queues, scheduling, external API integration, multiple output formats

### Changes Since Original Report (July 16, 2026):
- ✅ **Authorization / resource ownership** — journal actions now scoped by `userId`, threaded handler→service→repo (was ⚠️ "enforcement unclear")
- ✅ **SQL indexes** — migration `20260715180632_add_performance_indexes.sql` (was ❌ missing)
- ✅ **Version endpoint** — `GET /api/public/version` (was ❌ missing)
- ✅ **Password hashing** — real bcrypt on sign-up/login (was plaintext, despite the schema comment "should be bcrypt hashed")
- ✅ **Unique constraint** — partial unique index on `usr.email` for live rows (was ❌ none)
- ✅ **Integration tests** — testcontainers suite covering auth/journal/entry/user, including ownership scoping and wrong-password login

### Domain Adaptation:
The curriculum teaches building a TODO list API, but jot-api is a **journaling application**:
- **TODO Lists** → **Journals** 
- **TODO Items** → **Journal Entries**
- Core concepts translate well, domain logic differs appropriately

---

## Complete Task-by-Task Comparison

| # | Curriculum Task | Priority | Curriculum Page | Implemented? | Jot-API Equivalent | Status | Notes |
|---|----------------|----------|----------------|--------------|-------------------|--------|-------|
| 1 | **Init** | 🔺 Critical | 5 | ✅ Yes | `/api/healthz` endpoint | ✅ Complete | Returns `{"status":"OK"}` |
| 2 | **Request Logger** | 🔺 Critical | 8 | ✅ Yes | `middleware/request_logger.go` | ✅ Complete | Structured JSON logging, correlation ID, all required fields |
| 3 | **TODO CRUD** | 🔺 Critical | 12 | ✅ Yes | Journal CRUD (`/api/journals`) | ✅ Complete | POST, GET, PUT, DELETE all implemented |
| 4 | **Database Storage** | 🔺 Critical | 16 | ✅ Yes | Postgres with GORM | ✅ Complete | Proper connection handling |
| 5 | **Boolean Filtering** | 🔺 Critical | 19 | ✅ Yes | `?completed=true/false` query param | ✅ Complete | Custom `IsCompleted` enum type |
| 6 | **Pagination** | 🔺 Critical | 21 | ✅ Yes | `?page=X&size=Y` query params | ✅ Complete | Returns `totalRecords`, defaults to page=1, size=10 |
| 7 | **Database Migration** | 🔺 Critical | 24 | ✅ Yes | Goose migrations in `db/migrations/` | ✅ Complete | 10+ migration files, proper up/down |
| 8 | **Dockerfile** | 🔺 Critical | 26 | ✅ Yes | `Dockerfile` + `compose.yml` | ✅ Complete | Multi-stage could be better |
| 9 | **Swagger Docs** | ℹ️ Important | 29 | 🚧 In progress | `/api/docs` (planned) | 🚧 In progress | OpenAPI spec being added |
| 10 | **Versioning** | ℹ️ Important | 31 | ✅ Yes | `GET /api/public/version` | ✅ Complete | Returns build/version info via `version.GetInfo()` |
| 11 | **Users/Auth** | 🔺 Critical | 33 | ✅ Yes | `POST /api/authenticate`, JWT | ✅ Complete | JWT with email/password |
| 12 | **TODO Lists** | ℹ️ Important | 36 | ✅ Yes | Journals (one level) | ⚠️ Adapted | Jot doesn't need nested lists, just journal→entries |
| 13 | **Audit Columns** | 🔺 Critical | 38 | ✅ Yes | `created_at`, `updated_at`, `created_by`, `updated_by` | ✅ Complete | Migration 20250420000200 |
| 14 | **List Ownership** | 🔺 Critical | 40 | ✅ Yes | `user_id` scoping in repo/service/handler | ✅ Complete | All journal actions filter by authenticated `userId`; enforced via JWT middleware |
| 15 | **Title Uniqueness** | ℹ️ Important | 42 | ⚠️ Partial | Unique index on `usr.email` | ⚠️ Partial | Email uniqueness done; `(user_id, title)` journal uniqueness still not enforced |
| 16 | **Admin Role** | ℹ️ Important | 44 | ⚠️ Partial | `role` column in `usr` table | ⚠️ Incomplete | Column exists, no admin endpoints |
| 17 | **Admin Stats** | ℹ️ Important | 46 | ❌ No | - | ❌ Missing | No `/api/admin/system-stats` or `/api/admin/user-stats` |
| 18 | **SQL Indexing** | ℹ️ Important | 48 | ✅ Yes | `20260715180632_add_performance_indexes.sql` | ✅ Complete | Indexes on `journal(user_id)`, `entry(journal_id)`, `journal(completed)`, `usr(email)`, composite `(user_id, completed)` |
| 19 | **External API** | 🔺 Critical | 50 | ❌ No | - | ❌ Missing | No external API integration (PokéAPI example) |
| 20 | **Telemetry** | ℹ️ Important | 53 | ❌ No | - | ❌ Missing | No Sentry/DataDog/AWS tracing |
| 21 | **Soft Deletes** | ✅ Useful | 55 | ✅ Yes | `deleted_at` column | ✅ Complete | In `usr` table, also in journal/entry |
| 22 | **Sharing** | ✅ Useful | 57 | ❌ No | - | ❌ Missing | No editors table or sharing endpoints |
| 23 | **Transactions** | ℹ️ Important | 59 | ❓ Unknown | - | ⚠️ Unknown | Need to check service layer code |
| 24 | **Audit Tables** | ✅ Useful | 61 | ❌ No | - | ❌ Missing | No separate audit log table |
| 25 | **Sending Email** | ℹ️ Important | 62 | ❌ No | - | ❌ Missing | No email service integration |
| 26 | **Async Queue** | ✅ Useful | 64 | ❌ No | - | ❌ Missing | No SQS/RabbitMQ |
| 27 | **Schedule Reminders** | ✅ Useful | 65 | ❌ No | - | ❌ Missing | No cron jobs or scheduled tasks |
| 28 | **User Invite** | ℹ️ Important | 67 | ❌ No | - | ❌ Missing | No `/api/user-invitations` |
| 29 | **Social Login** | ℹ️ Important | 69 | ❌ No | - | ❌ Missing | No OAuth/IDP integration |
| 30 | **CSV Output** | ℹ️ Important | 70 | ❌ No | - | ❌ Missing | No `Accept: text/csv` handling |
| 31 | **Excel Output** | ✅ Useful | 72 | ❌ No | - | ❌ Missing | No `.xlsx` generation |
| 32 | **S3 Backup** | ℹ️ Important | 74 | ❌ No | - | ❌ Missing | No cloud storage backup |
| 33 | **SFTP** | ✅ Useful | 76 | ❌ No | - | ❌ Missing | No SFTP implementation |
| 34 | **Webhooks** | ✅ Useful | 78 | ❌ No | - | ❌ Missing | No webhook configuration |
| 35 | **Bonus: PDF Generation** | ✅ Bonus | 75 | ❌ No | - | ❌ Missing | - |
| 36 | **Bonus: Kafka** | ✅ Bonus | 75 | ❌ No | - | ❌ Missing | - |

---

## Implementation Inventory

### ✅ What's Implemented (Strong Areas)

#### 1. Core API Structure
- **Health Check**: `GET /api/healthz` returns `{"status":"OK"}`
- **Framework**: Echo v4 (good choice, production-ready)
- **Architecture**: Clean layered architecture
  - `api/` - HTTP handlers
  - `service/` - Business logic
  - `repo/` - Data access
  - `models/` - DTOs and view models

#### 2. Database & Persistence
- **Database**: PostgreSQL
- **ORM**: GORM
- **Migrations**: Goose (10+ migration files)
- **Tables**:
  - `journal` (id, title, description, completed, user_id, timestamps, audit columns)
  - `entry` (id, content, journal_id, timestamps, audit columns)
  - `usr` (id, email, password, first_name, last_name, role, deleted_at, timestamps)
- **Audit Columns**: ✅ `created_at`, `updated_at`, `created_by`, `updated_by`
- **Soft Deletes**: ✅ `deleted_at` on user, journal, and entry tables
- **Triggers**: ✅ Auto-update journal's `updated_at` when entries change

#### 3. REST Endpoints

**Journals (TODO Lists equivalent):**
```
GET    /api/journals              # List journals (paginated, filterable)
POST   /api/journals              # Create journal
GET    /api/journals/:id          # Get single journal
PUT    /api/journals/:id          # Update journal
DELETE /api/journals/:id          # Delete journal (soft delete)
```

**Entries (TODO Items equivalent):**
```
GET    /api/journals/:id/entries  # List entries for a journal
POST   /api/journals/:id/entries  # Create entry in journal
```

**Users / Auth:**
```
POST   /api/public/authenticate   # Login (returns JWT cookie)
POST   /api/public/sign-up        # Register new user (bcrypt hashed)
GET    /api/users/:id/journals    # Get user's journals
DELETE /api/users/:id             # Delete user
```

**Public / Meta:**
```
GET    /api/public/healthz        # Health check → {"status":"OK"}
GET    /api/public/version        # Build/version info
```

> Note: routes under `/api/public/` bypass the auth middleware; all others require a valid JWT cookie.

#### 4. Middleware & Cross-Cutting Concerns
- **Request Logging**: ✅ Structured JSON logs with:
  - `correlationId` (via Echo's RequestID middleware)
  - `time`, `level`, `msg`
  - `app.name`, `app.version`, `app.commit`
  - `http.method`, `http.route`, `http.remoteAddr`
  - `status`, `ms` (latency), `bytes`
  - Source field (empty but present)
  - UserId field (present but not populated yet)

#### 5. Models & Validation
- **Validation**: go-playground/validator
- **Custom Types**: `IsCompleted` enum (`true`, `false`, `unknown`)
- **View Models**: Separate DTOs for requests/responses
- **Pagination Models**: `PageOfJournalVMs` with metadata

#### 6. Testing
- Test files exist: `api/journal_test.go`, `api/user_test.go`, `api/entry_test.go`
- Integration test data: `testdata/integration_test_init.sql`
- *(Need to examine test coverage in detail)*

#### 7. Deployment
- **Docker**: ✅ Dockerfile (could use multi-stage build)
- **Docker Compose**: ✅ Full stack (db + web)
  - Health checks on database
  - Volume persistence
  - Environment variable configuration
- **Lambda Support**: ✅ `lambdaAdapter.go` for AWS Lambda
- **Environment Config**: ✅ `.env` file, defaults for local dev

---

### ❌ What's Missing (Gaps from Curriculum)

#### Critical Gaps (Production Readiness)
1. **No API Documentation**
   - Missing Swagger/OpenAPI spec
   - Curriculum page 29 - Important for frontend devs and API consumers

2. **No Telemetry/Observability**
   - Curriculum page 53 - Critical for production
   - No Sentry, DataDog, AWS X-Ray integration
   - Missing error monitoring, distributed tracing

3. **No External API Integration**
   - Curriculum page 50 - Critical skill
   - No example of calling external APIs, SDK usage, error handling

4. ~~**Incomplete Authorization**~~ ✅ **RESOLVED (July 16)**
   - Journal ownership now enforced: all actions scoped by authenticated `userId`, threaded handler→service→repo
   - JWT middleware parses the `userId` claim and populates request context

5. ~~**No Versioning Endpoint**~~ ✅ **RESOLVED**
   - `GET /api/public/version` returns build/version info

#### Important Gaps (Best Practices)
6. ~~**No SQL Indexes**~~ ✅ **RESOLVED**
   - Migration `20260715180632_add_performance_indexes.sql` adds indexes on:
     - `journal.user_id` (foreign key)
     - `entry.journal_id` (foreign key)
     - `journal.completed` (frequently filtered)
     - `usr.email` (lookup by email) — now a partial *unique* index
     - composite `journal(user_id, completed)`

7. **No Admin Endpoints**
   - Curriculum pages 44-47
   - Missing:
     - `GET /api/admin/lists` (all journals across users)
     - `GET /api/admin/system-stats` (total journals, users, etc.)
     - `GET /api/admin/user-stats` (per-user stats with pagination)

8. **Partial Uniqueness Constraints**
   - Curriculum page 42
   - ✅ `usr.email` now has a partial unique index (live rows only, so soft-deleted emails can be reused)
   - ❌ Journal titles still not unique per user (`(user_id, title)`)

9. **No Audit Table**
   - Curriculum page 61
   - Audit columns exist, but no separate audit log table
   - Can't see full history of changes

10. **No Alternative Output Formats**
    - Curriculum pages 70-72
    - Missing CSV export (`Accept: text/csv`)
    - Missing Excel export (`.xlsx`)

#### Features Not Implemented (Nice-to-Have)
11. **No Email Integration** (page 62)
    - Can't send emails for sharing, invites, reminders
    - No AWS SES, SendGrid, or SMTP

12. **No Async Processing** (page 64)
    - No message queue (SQS, RabbitMQ)
    - Email/webhooks must be synchronous

13. **No Scheduled Jobs** (page 65)
    - No cron jobs or scheduled tasks
    - Can't send daily reminders

14. **No Sharing/Collaboration** (page 57)
    - No editors table
    - No `POST /api/lists/:id/editors`
    - No `DELETE /api/lists/:id/editors/:userId`

15. **No User Invite Flow** (page 67)
    - No `POST /api/user-invitations`
    - No `POST /api/user-invitations/acceptance`
    - No invitation codes

16. **No Social Login** (page 69)
    - No OAuth (Google, GitHub)
    - No identity provider integration

17. **No Backup Endpoints** (pages 74-76)
    - No S3 backup (`POST /api/todos-backup`)
    - No SFTP upload

18. **No Webhooks** (page 78)
    - No webhook configuration
    - No external notifications on events

19. **No Transactions** (page 59) - *Need to verify service layer*

---

## Database Schema Analysis

### Current Schema

**`journal` table:**
```sql
id              SERIAL PRIMARY KEY
created_at      TIMESTAMP
updated_at      TIMESTAMP
title           VARCHAR
description     VARCHAR
completed       VARCHAR          -- enum: 'true', 'false', 'unknown'
user_id         INTEGER REFERENCES usr(id)
deleted_at      TIMESTAMP
created_by      VARCHAR
updated_by      VARCHAR
```

**`entry` table:**
```sql
id              SERIAL PRIMARY KEY
created_at      TIMESTAMP
updated_at      TIMESTAMP
content         TEXT
journal_id      INTEGER REFERENCES journal(id)
deleted_at      TIMESTAMP
created_by      VARCHAR
updated_by      VARCHAR
```

**`usr` table:**
```sql
id              SERIAL PRIMARY KEY
created_at      TIMESTAMP
updated_at      TIMESTAMP
deleted_at      TIMESTAMP
email           VARCHAR          -- partial UNIQUE index (deleted_at IS NULL)
password        VARCHAR          -- bcrypt hashed ✅
first_name      VARCHAR
last_name       VARCHAR
role            VARCHAR          -- present; used in JWT claims, no admin endpoints yet
```

### Database Strengths
✅ Foreign key constraints  
✅ Soft deletes via `deleted_at`  
✅ Audit columns (`created_by`, `updated_by`)  
✅ Database trigger to update journal when entries change  
✅ Proper migration history  
✅ Indexes on foreign keys and filtered columns  
✅ Partial unique index on `usr.email`  

### Database Weaknesses
❌ No unique constraint on journal title per user (`(user_id, title)`)  
❌ No `CASCADE` behavior defined for foreign keys  
❌ No default values for timestamps at database level  
❌ No separate audit log table  
❌ No editors/sharing table  
❌ No user settings table  

---

## Architecture & Code Quality Assessment

### ✅ Strengths
1. **Clean Architecture**: Handler → Service → Repo separation
2. **Validation**: Using go-playground/validator
3. **Error Handling**: Custom HTTP errors with proper status codes
4. **Environment Config**: 12-factor app principles (env vars, defaults)
5. **Docker Support**: Compose file with health checks
6. **Migrations**: Proper up/down migrations with Goose
7. **Structured Logging**: JSON logs ready for aggregation

### ⚠️ Areas for Improvement
1. **No Tests Running in CI**: Check `.github/workflows/`
2. **Dockerfile Not Optimized**: Could use multi-stage build to reduce image size
3. **No URL API Versioning**: No `/api/v1/` prefix (a `/api/public/version` endpoint exists, but not path versioning)
4. **Hard-Coded Strings**: App name in middleware (version now sourced from `version` package)
5. ~~**UserId in Logs**~~ ✅ JWT `userId` now parsed in middleware and available in context
6. ~~**Authorization**~~ ✅ Enforced — journal actions scoped by `userId`
7. ~~**Password Storage**~~ ✅ Confirmed bcrypt (was plaintext; fixed July 16)

---

## Recommendations by Priority

### ✅ Completed Since Original Report

1. ~~**Implement Authorization Middleware**~~ ✅ Done — JWT `userId` extracted in middleware, journal actions scoped by owner
2. ~~**Add SQL Indexes**~~ ✅ Done — `20260715180632_add_performance_indexes.sql`
3. ~~**Create `/api/version` Endpoint**~~ ✅ Done — `GET /api/public/version`
4. **Add Integration Tests** — ✅ Largely done (testcontainers suite covering auth/journal/entry/user, ownership scoping, wrong-password login); CI wiring still open

### 🔴 High Priority (Do Next)

1. **API Documentation** — 🚧 in progress
   - Generate OpenAPI/Swagger spec
   - Serve at `/api/docs` or `/api/docs/ui`
   - Use annotations or code generation

2. **Wire Tests into CI**
   - Run the testcontainers suite in `.github/workflows/`

3. **External API Example**
   - One integration with proper error handling, retries, timeouts (see Medium Priority #9)

### 🟡 Medium Priority (Production Readiness)

6. **Admin Endpoints**
   - `GET /api/admin/system-stats`
   - `GET /api/admin/user-stats`
   - `GET /api/admin/journals` (all journals, paginated)
   - Require `role=admin` in JWT

7. **Telemetry Integration**
   - Add Sentry for error tracking
   - Or AWS X-Ray for distributed tracing
   - Log errors with stack traces to external service

8. **Uniqueness Constraints**
   - Add unique constraint on `(user_id, title)` for journals
   - Return 400 with helpful error message on conflict

9. **External API Example**
   - Add one external API integration (doesn't have to be PokéAPI)
   - Show proper error handling, retries, timeouts
   - Log external requests with correlation ID

10. **CSV Export**
    - Support `Accept: text/csv` header
    - Export journals and entries to CSV

### 🟢 Low Priority (Nice-to-Have)

11. **Sharing/Collaboration** - Add editors table and endpoints
12. **Email Integration** - AWS SES or SendGrid for notifications
13. **Async Queue** - SQS or RabbitMQ for background jobs
14. **Scheduled Reminders** - Cron job or CloudWatch Events
15. **User Invite Flow** - Invitation codes and acceptance
16. **Social Login** - OAuth with Google/GitHub
17. **Excel Export** - Generate `.xlsx` files
18. **Backup to S3** - Periodic backups endpoint
19. **Webhooks** - User-configurable webhooks
20. **Audit Table** - Separate table for full change history

---

## Curriculum Learning Checklist

Use this to track what concepts from the curriculum you've learned:

### Critical Concepts (Must Know)
- [x] Creating a REST API with health check
- [x] Structured logging and correlation IDs
- [x] CRUD operations (Create, Read, Update, Delete)
- [x] Database persistence (Postgres + ORM)
- [x] Query parameters (filtering, pagination)
- [x] Database migrations
- [x] Containerization (Docker, compose)
- [ ] **API documentation (OpenAPI/Swagger)** — 🚧 in progress
- [x] Authentication (JWT)
- [x] Password hashing (bcrypt)
- [x] Authorization (resource ownership)
- [x] Soft deletes
- [x] Audit columns
- [ ] **External API integration**
- [ ] **Telemetry and observability**

### Important Concepts (Should Know)
- [x] Boolean filtering
- [x] Pagination with metadata
- [x] API versioning (version endpoint)
- [x] Foreign key relationships
- [x] SQL indexing for performance
- [x] Uniqueness constraints (email; journal title still pending)
- [ ] Admin-only endpoints with role checks
- [ ] **Group by queries for stats**
- [ ] CSV output (Accept header negotiation)
- [ ] Email sending
- [ ] User invite flows

### Useful Concepts (Nice to Know)
- [x] Custom enum types
- [x] Database triggers
- [x] Lambda deployment
- [ ] Async message queues
- [ ] Scheduled jobs (cron)
- [ ] Sharing and permissions
- [ ] Transactions for complex updates
- [ ] Separate audit tables
- [ ] Social login (OAuth)
- [ ] Excel output
- [ ] S3 file storage
- [ ] SFTP integration
- [ ] Webhooks
- [ ] PDF generation
- [ ] Kafka event streaming

---

## Next Steps

### Immediate Actions
1. **Read this entire document** to understand gaps
2. **Prioritize** which features you want to add next
3. **Create issues/tickets** for high-priority items
4. **Test your API** manually against the curriculum's testing sections
5. **Check if authorization is actually enforced** (can user A access user B's journals?)

### For Skill Development
If you're using this project to learn:
- ✅ You've completed ~40% of the critical curriculum
- 🎯 Focus on the "High Priority" recommendations next
- 📚 The curriculum has excellent "Self Review" questions for each task - go back and ask yourself those questions

### For Production Deployment
If you want to deploy this to production:
- 🔴 **MUST HAVE**: Authorization, indexes, API docs, tests, telemetry
- 🟡 **SHOULD HAVE**: Admin endpoints, versioning, external API example
- 🟢 **NICE TO HAVE**: Everything else

---

## Conclusion

The jot-api project demonstrates a solid understanding of core REST API principles and has implemented the fundamental building blocks well. The domain adaptation from TODO lists to journals is appropriate and well-executed.

**Main strengths**: Database design, CRUD operations, authentication, **authorization/resource ownership**, **bcrypt password hashing**, **SQL indexing**, **version endpoint**, Docker support, structured logging, migrations, integration tests.

**Main gaps**: API documentation (in progress), telemetry, admin features, advanced integrations (email, queues, external APIs).

Since the original report, the "High Priority" list has been substantially cleared — authorization, indexes, versioning, password hashing, and integration tests are all in place. Remaining focus is API documentation, then production observability and admin/advanced features.

**Overall Assessment**: Strong foundation with core production hardening complete (A-); remaining work is observability, docs, and advanced integrations.

---

*End of Report*
