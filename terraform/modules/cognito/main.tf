resource "aws_cognito_user_pool" "main" {
  name = "${var.environment}-${var.project_name}-user-pool"

  # パスワードポリシー
  password_policy {
    minimum_length    = var.password_policy.minimum_length
    require_lowercase = var.password_policy.require_lowercase
    require_numbers   = var.password_policy.require_numbers
    require_symbols   = var.password_policy.require_symbols
    require_uppercase = var.password_policy.require_uppercase
  }

  # MFA設定
  mfa_configuration = var.mfa_configuration

  # ユーザー属性
  schema {
    attribute_data_type = "String"
    name                = "email"
    required            = true
    mutable             = true
  }

  schema {
    attribute_data_type = "String"
    name                = "name"
    required            = true
    mutable             = true
  }

  # アカウント復旧設定
  account_recovery_setting {
    recovery_mechanism {
      name     = "verified_email"
      priority = 1
    }
  }

  # メール設定
  email_configuration {
    email_sending_account = var.email_sending_account
    from_email_address    = var.from_email_address
  }

  # 検証設定
  auto_verified_attributes = ["email"]

  # ユーザープール作成時の設定
  admin_create_user_config {
    allow_admin_create_user_only = var.admin_create_user_only
  }

  tags = merge(var.common_tags, {
    Name        = "${var.environment}-${var.project_name}-user-pool"
    Environment = var.environment
  })
}

# User Pool Client
resource "aws_cognito_user_pool_client" "main" {
  name         = "${var.environment}-${var.project_name}-client"
  user_pool_id = aws_cognito_user_pool.main.id

  # クライアント設定
  generate_secret = false
  
  # 認証フロー
  explicit_auth_flows = [
    "ALLOW_USER_SRP_AUTH",
    "ALLOW_REFRESH_TOKEN_AUTH"
  ]

  # トークン設定
  access_token_validity  = var.token_validity.access_token
  id_token_validity     = var.token_validity.id_token  
  refresh_token_validity = var.token_validity.refresh_token

  # トークン有効期間単位
  token_validity_units {
    access_token  = "hours"
    id_token     = "hours"
    refresh_token = "days"
  }

  # OAuth設定
  supported_identity_providers = ["COGNITO"]
  
  callback_urls = var.callback_urls
  logout_urls   = var.logout_urls

  # OAuth スコープ
  allowed_oauth_flows = ["code"]
  allowed_oauth_scopes = [
    "email",
    "openid",
    "profile"
  ]
  allowed_oauth_flows_user_pool_client = true
}

# Identity Pool
resource "aws_cognito_identity_pool" "main" {
  identity_pool_name               = "${var.environment}-${var.project_name}-identity-pool"
  allow_unauthenticated_identities = var.allow_unauthenticated_identities

  cognito_identity_providers {
    client_id     = aws_cognito_user_pool_client.main.id
    provider_name = aws_cognito_user_pool.main.endpoint
  }

  tags = merge(var.common_tags, {
    Name        = "${var.environment}-${var.project_name}-identity-pool"
    Environment = var.environment
  })
}

# IAM Role for authenticated users
resource "aws_iam_role" "authenticated" {
  name = "${var.environment}-${var.project_name}-cognito-authenticated-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRoleWithWebIdentity"
        Effect = "Allow"
        Principal = {
          Federated = "cognito-identity.amazonaws.com"
        }
        Condition = {
          StringEquals = {
            "cognito-identity.amazonaws.com:aud" = aws_cognito_identity_pool.main.id
          }
          "ForAnyValue:StringLike" = {
            "cognito-identity.amazonaws.com:amr" = "authenticated"
          }
        }
      }
    ]
  })

  tags = merge(var.common_tags, {
    Name        = "${var.environment}-${var.project_name}-cognito-authenticated-role"
    Environment = var.environment
  })
}

# IAM Policy for authenticated users
resource "aws_iam_role_policy" "authenticated" {
  name = "${var.environment}-${var.project_name}-authenticated-policy"
  role = aws_iam_role.authenticated.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = var.authenticated_user_policy_statements
  })
}

# Identity Pool Role Attachment
resource "aws_cognito_identity_pool_roles_attachment" "main" {
  identity_pool_id = aws_cognito_identity_pool.main.id

  roles = {
    "authenticated" = aws_iam_role.authenticated.arn
  }
}
