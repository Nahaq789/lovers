

---

## システム設計の最終まとめ

### アーキテクチャ

```
クライアント → Lambda（Go） → Cognito（認証）
                             → Supabase PostgreSQL（データ）
```

すべてのAPIはLambda経由で統一

---

### データベース設計

```sql
-- ユーザーテーブル
CREATE TABLE "user" (
    user_id UUID PRIMARY KEY,           -- Cognito sub
    email VARCHAR(50) NOT NULL UNIQUE,
    name VARCHAR(20) NOT NULL
);

-- RLS有効化（Supabase API経由のアクセスを全拒否）
ALTER TABLE "user" ENABLE ROW LEVEL SECURITY;
```

**アクセス制御:**
- Supabase API直接: RLSで拒否
- Lambda（PostgreSQL接続）: RLSバイパス（全アクセス可能）

---

### セキュリティ設計

1. **認証**: Cognito JWT
2. **認可**: Lambda内でJWT検証 + 本人確認（JWT内subとリクエストパラメータを照合）
3. **多層防御**: RLS（1層目）+ Lambda認証（2層目）

---

### エンドポイント一覧

| エンドポイント | 認証 | 説明 |
|-------------|------|------|
| POST /signup | 不要 | サインアップ（Cognito登録のみ） |
| POST /confirm-email | 不要 | メール認証 + 自動ログイン |
| POST /complete-profile | 必須 | プロフィール登録（userテーブル登録） |
| POST /login | 不要 | ログイン |
| POST /refresh-token | 不要 | トークン更新 |
| POST /forgot-password | 不要 | パスワードリセット開始 |
| POST /confirm-forgot-password | 不要 | パスワードリセット確認 |
| DELETE /users/me | 必須 | アカウント削除 |

---

### サインアップ〜登録完了フロー

```
【ステップ1: サインアップ】
POST /signup
・Cognito登録
・6桁コード送信
・この時点でuserテーブルは未登録

【ステップ2: メール認証】
POST /confirm-email
・Cognito認証確認
・自動ログイン
・JWTトークン返却

【ステップ3: プロフィール登録】
POST /complete-profile (JWT必須)
・JWT検証 → cognitoSub, email取得
・user_id存在チェック（二重登録防止）
・email存在チェック（孤立レコード検出）
・userテーブル登録
  成功: 登録完了
  失敗: userテーブル削除 → Cognito削除
```

---

### エラーハンドリング方針

**サインアップ:**
- Cognito失敗 → userテーブル登録しない

**プロフィール登録:**
- user_id重複 → "既に登録済み"エラー
- email重複 → "メールアドレス使用済み"エラー + 管理者通知
- DB登録失敗 → userテーブル削除 → Cognito削除

**ユーザー削除（パターンB）:**
- ① userテーブル削除
- ② Cognito削除
- ②失敗時の復元: 後で検討

---

### 保留事項（今後検討）

- [ ] メール認証コード再送機能
- [ ] プロフィール更新機能
- [ ] パスワード変更機能
- [ ] メール認証タイムアウト対応
- [ ] ユーザー削除失敗時の復元処理

---
