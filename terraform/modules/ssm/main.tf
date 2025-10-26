# DB接続情報 - ユーザー名
resource "aws_ssm_parameter" "db_user" {
  name        = "/lovers/dev/db_user"
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
  name        = "/lovers/dev/db_password"
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
  name        = "/lovers/dev/db_port"
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
  name        = "/lovers/dev/db_host"
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
  name        = "/lovers/dev/db_name"
  description = "データベース名"
  type        = "String"
  value       = var.db_name

  tags = {
    Environment = "dev"
    Project     = "lovers"
    Name        = "lovers-dev-db-name"
  }
}