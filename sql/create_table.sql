-- 1. user テーブル
CREATE TABLE IF NOT EXISTS "user" (
    user_id UUID PRIMARY KEY,
    email VARCHAR(254) NOT NULL UNIQUE,
    user_name VARCHAR(20) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);

COMMENT ON TABLE "user" IS 'ユーザー情報';
COMMENT ON COLUMN "user".user_id IS 'ユーザーID（主キー）';
COMMENT ON COLUMN "user".email IS 'メールアドレス';
COMMENT ON COLUMN "user".user_name IS 'ユーザー名';

-- 2. group テーブル
CREATE TABLE IF NOT EXISTS "group" (
    group_id UUID PRIMARY KEY,
    created_by UUID NOT NULL,
    group_name VARCHAR(20) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    CONSTRAINT fk_group_created_by FOREIGN KEY (created_by) REFERENCES "user"(user_id)
);

COMMENT ON TABLE "group" IS 'グループ情報';
COMMENT ON COLUMN "group".group_id IS 'グループID（主キー）';
COMMENT ON COLUMN "group".created_by IS 'グループ作成者のユーザーID';
COMMENT ON COLUMN "group".group_name IS 'グループ名';

-- 3. group_member テーブル
CREATE TABLE IF NOT EXISTS group_member (
    group_member_id UUID PRIMARY KEY,
    group_id UUID NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    CONSTRAINT fk_group_member_group FOREIGN KEY (group_id) REFERENCES "group"(group_id),
    CONSTRAINT fk_group_member_user FOREIGN KEY (user_id) REFERENCES "user"(user_id),
    CONSTRAINT uq_group_member UNIQUE (group_id, user_id)
);

COMMENT ON TABLE group_member IS 'グループメンバー（多対多の中間テーブル）';
COMMENT ON COLUMN group_member.group_member_id IS 'グループメンバーID（主キー）';
COMMENT ON COLUMN group_member.group_id IS '所属グループID';
COMMENT ON COLUMN group_member.user_id IS 'メンバーのユーザーID';

-- 4. category テーブル
CREATE TABLE IF NOT EXISTS category (
    category_id UUID PRIMARY KEY,
    group_id UUID NOT NULL,
    created_by UUID NOT NULL,
    category_name VARCHAR(10) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    CONSTRAINT fk_category_group FOREIGN KEY (group_id) REFERENCES "group"(group_id),
    CONSTRAINT fk_category_created_by FOREIGN KEY (created_by) REFERENCES "user"(user_id)
);

COMMENT ON TABLE category IS 'カテゴリマスタ（グループごと）';
COMMENT ON COLUMN category.category_id IS 'カテゴリID（主キー）';
COMMENT ON COLUMN category.group_id IS '所属グループID';
COMMENT ON COLUMN category.created_by IS 'カテゴリ作成者のユーザーID';
COMMENT ON COLUMN category.category_name IS 'カテゴリ名';

-- 5. expense テーブル
CREATE TABLE IF NOT EXISTS expense (
    expense_id UUID NOT NULL,
    user_id UUID NOT NULL,
    group_id UUID NOT NULL,
    category_id UUID NOT NULL,
    amount BIGINT NOT NULL,
    nominal VARCHAR(15) NOT NULL,
    payment_date TIMESTAMPTZ NOT NULL,
    description TEXT,
    deleted_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    CONSTRAINT expense_pkey PRIMARY KEY (expense_id, user_id),
    CONSTRAINT fk_expense_group FOREIGN KEY (group_id) REFERENCES "group"(group_id),
    CONSTRAINT fk_expense_user FOREIGN KEY (user_id) REFERENCES "user"(user_id),
    CONSTRAINT fk_expense_category FOREIGN KEY (category_id) REFERENCES category(category_id)
);

COMMENT ON TABLE expense IS '支出明細';
COMMENT ON COLUMN expense.expense_id IS '支出ID（複合主キー）';
COMMENT ON COLUMN expense.user_id IS '支払者のユーザーID（複合主キー）';
COMMENT ON COLUMN expense.group_id IS '所属グループID';
COMMENT ON COLUMN expense.category_id IS 'カテゴリID';
COMMENT ON COLUMN expense.amount IS '支払金額';
COMMENT ON COLUMN expense.nominal IS '名目';
COMMENT ON COLUMN expense.payment_date IS '支払日';
COMMENT ON COLUMN expense.description IS '説明';
COMMENT ON COLUMN expense.deleted_at IS '削除日時（論理削除）';

-- 6. expense_log テーブル
CREATE TABLE IF NOT EXISTS expense_log (
    expense_log_id UUID PRIMARY KEY,
    expense_id UUID NOT NULL,
    group_id UUID NOT NULL,
    user_id UUID NOT NULL,
    operation VARCHAR(10) NOT NULL,
    before_data JSONB,
    after_data JSONB,
    created_at TIMESTAMPTZ NOT NULL,
    CONSTRAINT fk_expense_log_expense FOREIGN KEY (expense_id, user_id) REFERENCES expense(expense_id, user_id),
    CONSTRAINT fk_expense_log_group FOREIGN KEY (group_id) REFERENCES "group"(group_id),
    CONSTRAINT chk_operation CHECK (operation IN ('add', 'edit', 'delete'))
);

COMMENT ON TABLE expense_log IS '支出変更履歴ログ';
COMMENT ON COLUMN expense_log.expense_log_id IS 'ログID（主キー）';
COMMENT ON COLUMN expense_log.expense_id IS '対象の支出ID';
COMMENT ON COLUMN expense_log.group_id IS '所属グループID';
COMMENT ON COLUMN expense_log.user_id IS '操作対象の支払者ユーザーID';
COMMENT ON COLUMN expense_log.operation IS '操作種別（add/edit/delete）';
COMMENT ON COLUMN expense_log.before_data IS '変更前データ（JSONB形式）';
COMMENT ON COLUMN expense_log.after_data IS '変更後データ（JSONB形式）';

-- 7. template テーブル
CREATE TABLE IF NOT EXISTS template (
    template_id UUID PRIMARY KEY,
    group_id UUID NOT NULL,
    created_by UUID NOT NULL,
    template_name VARCHAR(20) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    CONSTRAINT fk_template_group FOREIGN KEY (group_id) REFERENCES "group"(group_id),
    CONSTRAINT fk_template_created_by FOREIGN KEY (created_by) REFERENCES "user"(user_id)
);

COMMENT ON TABLE template IS '支出テンプレート';
COMMENT ON COLUMN template.template_id IS 'テンプレートID（主キー）';
COMMENT ON COLUMN template.group_id IS '所属グループID';
COMMENT ON COLUMN template.created_by IS 'テンプレート作成者のユーザーID';
COMMENT ON COLUMN template.template_name IS 'テンプレート名';

-- 8. template_expense テーブル
CREATE TABLE IF NOT EXISTS template_expense (
    template_expense_id UUID PRIMARY KEY,
    template_id UUID NOT NULL,
    category_id UUID NOT NULL,
    amount BIGINT NOT NULL,
    nominal VARCHAR(15) NOT NULL,
    payment_date TIMESTAMPTZ NOT NULL,
    description TEXT,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    CONSTRAINT fk_template_expense_template FOREIGN KEY (template_id) REFERENCES template(template_id),
    CONSTRAINT fk_template_expense_category FOREIGN KEY (category_id) REFERENCES category(category_id)
);

COMMENT ON TABLE template_expense IS 'テンプレート支出';
COMMENT ON COLUMN template_expense.template_expense_id IS 'テンプレート支出ID（主キー）';
COMMENT ON COLUMN template_expense.template_id IS '所属テンプレートID';
COMMENT ON COLUMN template_expense.category_id IS 'カテゴリID';
COMMENT ON COLUMN template_expense.amount IS '金額';
COMMENT ON COLUMN template_expense.nominal IS '名目';
COMMENT ON COLUMN template_expense.payment_date IS '支払日';
COMMENT ON COLUMN template_expense.description IS '説明';

-- インデックス作成
CREATE INDEX IF NOT EXISTS idx_group_member_group_id ON group_member(group_id);
CREATE INDEX IF NOT EXISTS idx_group_member_user_id ON group_member(user_id);
CREATE INDEX IF NOT EXISTS idx_expense_group_id ON expense(group_id);
CREATE INDEX IF NOT EXISTS idx_expense_user_id ON expense(user_id);
CREATE INDEX IF NOT EXISTS idx_expense_payment_date ON expense(payment_date);
CREATE INDEX IF NOT EXISTS idx_expense_deleted_at ON expense(deleted_at);
CREATE INDEX IF NOT EXISTS idx_expense_log_expense_id ON expense_log(expense_id);
CREATE INDEX IF NOT EXISTS idx_expense_log_group_id ON expense_log(group_id);
CREATE INDEX IF NOT EXISTS idx_expense_log_created_at ON expense_log(created_at);
CREATE INDEX IF NOT EXISTS idx_category_group_id ON category(group_id);
CREATE INDEX IF NOT EXISTS idx_template_group_id ON template(group_id);
CREATE INDEX IF NOT EXISTS idx_template_expense_template_id ON template_expense(template_id);
