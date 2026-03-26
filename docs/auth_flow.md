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

1. **認証**: Cognito JWT（`ADMIN_USER_PASSWORD_AUTH` フロー）
2. **認可**: Lambda内でJWT検証 + 本人確認（JWT内subとリクエストパラメータを照合）
3. **多層防御**: RLS（1層目）+ Lambda認証（2層目）

**JWT検証:**
- CognitoのJWKSエンドポイントから公開鍵を取得しキャッシュ
- Lambdaローカルで検証（Cognitoへの問い合わせなし）

**トークン管理:**
- AccessToken / RefreshToken はいずれも HttpOnly Cookie で管理
- リフレッシュは401レスポンスをトリガーにフロントが自動実行し、成功後に元のリクエストをリトライ
- ログアウトは `RevokeToken`（当該デバイスのみ無効化）

---

### エンドポイント一覧

| エンドポイント | 認証 | Cognito API | 説明 |
|-------------|------|------------|------|
| POST /signup | 不要 | SignUp | サインアップ（Cognito登録のみ） |
| POST /confirm-email | 不要 | ConfirmSignUp + AdminInitiateAuth | メール認証 + 自動ログイン |
| POST /complete-profile | 必須 | - | プロフィール登録（userテーブル登録） |
| POST /login | 不要 | AdminInitiateAuth | ログイン |
| POST /refresh-token | 不要 | AdminInitiateAuth (REFRESH_TOKEN_AUTH) | トークン更新 |
| POST /logout | 必須 | RevokeToken | ログアウト（当該デバイスのみ） |
| POST /forgot-password | 不要 | ForgotPassword | パスワードリセット開始 |
| POST /confirm-forgot-password | 不要 | ConfirmForgotPassword | パスワードリセット確認 |
| DELETE /users/me | 必須 | AdminDeleteUser | アカウント削除 |

---

### 認証フロー

#### ログイン
1. `POST /login { email, password }`
2. Go → Cognito `AdminInitiateAuth`
3. AccessToken / RefreshToken を HttpOnly Cookie にセット

#### APIリクエスト
1. Cookie の AccessToken をGoがJWKS公開鍵でローカル検証
2. 期限切れの場合 → 401を返す
3. フロントが `POST /refresh-token` を自動実行 → AccessToken を更新
4. 元のリクエストをリトライ

#### ログアウト
1. `POST /logout`
2. Go → Cognito `RevokeToken`（RefreshToken を無効化）
3. Cookie（AccessToken / RefreshToken）をクリア
4. ※発行済みAccessTokenはJWTの有効期限まで技術的には有効（許容リスク）

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
・自動ログイン（AdminInitiateAuth）
・AccessToken / RefreshToken を Cookie にセット

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

**ユーザー削除:**
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
