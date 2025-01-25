# Create an Amazon RDS PostgreSQL database
resource "aws_db_instance" "my_db_instance" {
  identifier             = var.db_instance_identifier
  engine                 = "postgres"
  instance_class         = "db.t3.micro"  # Change to a supported instance class
  allocated_storage       = 20
  username               = var.db_instance_username
  password               = var.db_instance_password
  publicly_accessible     = false
  skip_final_snapshot     = true
}