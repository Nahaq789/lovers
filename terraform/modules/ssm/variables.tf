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