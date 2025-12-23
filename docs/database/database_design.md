```mermaid
erDiagram
    user ||--o{ group : "creates"
    user ||--o{ group_member : "belongs to"
    group ||--o{ group_member : "has"
    user ||--o{ expense : "pays"
    group ||--o{ expense : "has"
    category ||--o{ expense : "categorizes"
    expense ||--o{ expense_log : "has history"
    group ||--o{ expense_log : "tracks"
    user ||--o{ expense_log : "operates"
    group ||--o{ category : "has"
    user ||--o{ category : "creates"
    group ||--o{ template : "has"
    user ||--o{ template : "creates"
    template ||--o{ template_expense : "has"
    category ||--o{ template_expense : "categorizes"

    user {
        uuid user_id PK
        varchar email
        varchar user_name
        timestamptz created_at
        timestamptz updated_at
    }

    group {
        uuid group_id PK
        uuid created_by FK
        varchar group_name
        timestamptz created_at
        timestamptz updated_at
    }

    group_member {
        uuid group_member_id PK
        uuid group_id FK
        uuid user_id FK
        timestamptz created_at
    }

    expense {
        uuid expense_id PK
        uuid group_id FK
        uuid payment_by FK
        uuid category_id FK
        bigint amount
        varchar nominal
        timestamptz payment_date
        text description
        timestamptz deleted_at
        timestamptz created_at
        timestamptz updated_at
    }

    expense_log {
        uuid expense_log_id PK
        uuid expense_id FK
        uuid group_id FK
        uuid user_id FK
        varchar operation
        jsonb before_data
        jsonb after_data
        timestamptz created_at
    }

    category {
        uuid category_id PK
        uuid group_id FK
        uuid created_by FK
        varchar category_name
        timestamptz created_at
        timestamptz updated_at
    }

    template {
        uuid template_id PK
        uuid group_id FK
        uuid created_by FK
        varchar template_name
        timestamptz created_at
        timestamptz updated_at
    }

    template_expense {
        uuid template_expense_id PK
        uuid template_id FK
        uuid category_id FK
        bigint amount
        varchar nominal
        timestamptz payment_date
        text description
        timestamptz created_at
        timestamptz updated_at
    }
```
