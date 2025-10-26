# DB接続情報 - ユーザー名
resource "aws_ssm_parameter" "db_user" {
  name        = "/lovers/dev/DB_USER"
  description = "データベースユーザー名"
  type        = "String"
  value       = var.db_user

  tags = {
    Environment = "dev"
    Project     = "lovers"
    Name        = "lovers-dev-db-user"
  }
}

# DB接続情報 - パスワード（暗号化）
resource "aws_ssm_parameter" "db_password" {
  name        = "/lovers/dev/DB_PASSWORD"
  description = "データベースパスワード"
  type        = "SecureString"
  value       = var.db_password

  tags = {
    Environment = "dev"
    Project     = "lovers"
    Name        = "lovers-dev-db-password"
  }
}

# DB接続情報 - ポート番号
resource "aws_ssm_parameter" "db_port" {
  name        = "/lovers/dev/DB_PORT"
  description = "データベースポート番号"
  type        = "String"
  value       = var.db_port

  tags = {
    Environment = "dev"
    Project     = "lovers"
    Name        = "lovers-dev-db-port"
  }
}

# DB接続情報 - ホスト
resource "aws_ssm_parameter" "db_host" {
  name        = "/lovers/dev/DB_HOST"
  description = "データベースホスト"
  type        = "String"
  value       = var.db_host

  tags = {
    Environment = "dev"
    Project     = "lovers"
    Name        = "lovers-dev-db-host"
  }
}

# DB接続情報 - データベース名
resource "aws_ssm_parameter" "db_name" {
  name        = "/lovers/dev/DB_NAME"
  description = "データベース名"
  type        = "String"
  value       = var.db_name

  tags = {
    Environment = "dev"
    Project     = "lovers"
    Name        = "lovers-dev-db-name"
  }
}