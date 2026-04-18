# Lovers Backend

## Overview

Couples' expense-tracking app. Go/Gin backend running on AWS Lambda, backed by Supabase PostgreSQL and AWS Cognito.

## Tech Stack

| Layer | Technology |
|---|---|
| Language | Go |
| Framework | Gin |
| Deploy | AWS Lambda |
| Database | Supabase PostgreSQL |
| Auth | AWS Cognito (`ADMIN_USER_PASSWORD_AUTH`) |
| DI | Wire |
| Task runner | justfile |

## Architecture

Clean Architecture + DDD. Layer boundaries must be respected.

```
cmd/
  di/          # Wire DI configuration per domain
  initialize/  # DI initializer functions
internal/
  domain/
    models/
      aggregates/          # Aggregate roots
      valueobjects/        # Shared value objects (createdAt, updatedAt)
      <domain>/
        <valueoject>/      # Domain-specific value objects
    repositories/          # Repository interfaces
    events/                # Domain events (past tense: ExpenseAdded)
  usecases/
    dto/
      <domain>/
        request/           # Request DTOs
        response/          # Response DTOs
  infrastructure/
    repositories/          # Repository implementations
    db/
  shared/
    infrastructure/
      logger/
      security/
        userid/            # JWT user ID extraction
```

## Naming Conventions

### Value Objects

- **Package name**: lowercase, single word matching the concept (e.g., `groupid`, `groupname`, `userid`, `paymentdate`, `createdat`)
- **Type name**: PascalCase (e.g., `GroupId`, `GroupName`, `UserId`, `PaymentDate`, `CreatedAt`)
- **Constructor**: `New<TypeName>()` for generated values, `New<TypeName>FromString()` for parsing from external input
- **Directory**: `internal/domain/models/<domain>/<valueoject>/` for domain-specific, `internal/domain/models/valueobjects/<valueoject>/` for shared

```go
// Example: lovers/internal/domain/models/group/groupid/groupid.go
package groupid

type GroupId struct { ... }

func NewGroupId() (*GroupId, error) { ... }
func NewGroupIdFromString(s string) (*GroupId, error) { ... }
```

### Other Conventions

- Domain events: past tense (`ExpenseAdded`, not `AddExpense`)
- Repository scan variables: `raw` prefix (e.g., `rawExpenseId`, `rawCreatedAt`)

## Key Principles

- Aggregates must have complete identity from creation; IDs belong in the domain layer, not infrastructure
- Repository serves the write side and returns the aggregate root
- `expense_log` stays within the main app (not a microservice) to preserve transactional atomicity
- Never trust user-supplied `user_id` for self-identification; always extract from JWT
- For cross-user lookups (group context), accept explicit `target_user_id` with server-side group membership validation
- DTOs must not use domain models; accept primitives and convert in the use case layer

## Authentication

- Flow: `ADMIN_USER_PASSWORD_AUTH` via Cognito
- Tokens: AccessToken + RefreshToken stored in HttpOnly cookies
- JWT verification: JWKS public keys cached locally in Lambda
- Token refresh: triggered by 401 response from frontend
- Logout: `RevokeToken` API (already-issued AccessTokens remain valid until expiry)

## DI Pattern

Each domain has its own Wire provider set under `cmd/di/<domain>/`. Initializer functions live in `cmd/initialize/`.

```go
// Pattern: same structure as group for all new domains
func InitTemplate(ctx context.Context, d *db.DbClient) *template.TemplateSet {
    templateSet := template.Initialize(d)
    return templateSet
}
```

## justfile Commands

- `just di_all` — runs all wire generation commands
- `just di_<domain>` — runs wire for a specific domain

## Code Generation Rules

- No inline comments in generated code
- Use `raw` prefix for variables declared before `Scan` in repository implementations
- Use `sql.NullString` / `sql.NullTime` for nullable columns
