terraform {
  required_version = ">= 1.5"
  
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = var.aws_region
  
  default_tags {
    tags = local.common_tags
  }
}

locals {
  common_tags = {
    Environment = "dev"
    Project     = var.project_name
    ManagedBy   = "Terraform"
    Owner       = "DevOps-Team"
  }
}

# Cognito モジュールの使用
module "cognito" {
  source = "../../modules/cognito"

  environment  = "dev"
  project_name = var.project_name
  common_tags  = local.common_tags

  # 開発環境用の設定（セキュリティ要件を緩和）
  password_policy = {
    minimum_length    = 6  # 開発用に短く
    require_lowercase = true
    require_numbers   = true
    require_symbols   = false  # 開発用にシンボル不要
    require_uppercase = false  # 開発用に大文字不要
  }

  admin_create_user_only = false  # 開発者が自由にユーザー作成可能

  token_validity = {
    access_token  = 24   # 開発用に長め
    id_token     = 24   # 開発用に長め
    refresh_token = 30
  }

  allow_unauthenticated_identities = true  # 開発用にゲストアクセス許可

  # 開発環境用のIAMポリシー（より緩い権限）
  authenticated_user_policy_statements = [
    {
      Effect = "Allow"
      Action = [
        "mobileanalytics:PutEvents",
        "cognito-sync:*",
        "cognito-identity:*"
      ]
      Resource = ["*"]
    },
    {
      Effect = "Allow"
      Action = [
        "s3:GetObject",
        "s3:PutObject"
      ]
      Resource = ["arn:aws:s3:::dev-${var.project_name}-*/*"]
    }
  ]
}

module "ssm" {
  source = "../../modules/ssm"

  environment  = "dev"
  project_name = var.project_name
  common_tags  = local.common_tags

  db_user     = var.db_user
  db_password = var.db_password
  db_port     = var.db_port
  db_host     = var.db_host
  db_name     = var.db_name
}
