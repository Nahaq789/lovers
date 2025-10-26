output "db_user_parameter_name" {
  description = "DB_USERパラメータ名"
  value       = aws_ssm_parameter.db_user.name
}

output "db_password_parameter_name" {
  description = "DB_PASSWORDパラメータ名"
  value       = aws_ssm_parameter.db_password.name
}

output "db_port_parameter_name" {
  description = "DB_PORTパラメータ名"
  value       = aws_ssm_parameter.db_port.name
}

output "db_host_parameter_name" {
  description = "DB_HOSTパラメータ名"
  value       = aws_ssm_parameter.db_host.name
}

output "db_name_parameter_name" {
  description = "DB_NAMEパラメータ名"
  value       = aws_ssm_parameter.db_name.name
}

output "db_parameter_arns" {
  description = "全DBパラメータのARN"
  value = {
    db_user     = aws_ssm_parameter.db_user.arn
    db_password = aws_ssm_parameter.db_password.arn
    db_port     = aws_ssm_parameter.db_port.arn
    db_host     = aws_ssm_parameter.db_host.arn
    db_name     = aws_ssm_parameter.db_name.arn
  }
}