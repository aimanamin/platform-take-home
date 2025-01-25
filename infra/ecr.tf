# Create an Elastic Container Registry (ECR)
resource "aws_ecr_repository" "my_ecr_repository" {
  name = var.ecr_repository_name

  image_tag_mutability = "IMMUTABLE"
}