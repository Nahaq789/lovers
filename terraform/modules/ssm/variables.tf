variable "environment" {
  description = "Environment name (dev, prod, etc.)"
  type        = string
}

variable "project_name" {
  description = "Project name for resource naming"
  type        = string
}

variable "common_tags" {
  description = "Common tags to apply to all resources"
  type        = map(string)
  default     = {}
}

variable "db_user" {
  description = "データベースユーザー名"
  type        = string
  sensitive   = true
}

variable "db_password" {
  description = "データベースパスワード"
  type        = string
  sensitive   = true
}

variable "db_port" {
  description = "データベースポート番号"
  type        = string
  default     = "5432"
}

variable "db_host" {
  description = "データベースホスト"
  type        = string
}

variable "db_name" {
  description = "データベース名"
  type        = string
}