output "ecr_repository_url" {
  value       = aws_ecr_repository.my_ecr_repository.repository_url
  description = "The URL of the ECR repository."
}

output "db_instance_endpoint" {
  value       = aws_db_instance.my_db_instance.endpoint
  description = "The endpoint of the RDS PostgreSQL database."
}

output "db_instance_username" {
  value       = aws_db_instance.my_db_instance.username
  description = "The username for the RDS PostgreSQL database."
}