# Skip Platform Take Home Challenge

## Requirments

- Must have a AWS credentials to configure the AWS CLI
- Should create the appropriate S3 bucket and dynamodb table that host the terraform states and locks. see infra\backend.tf
- Configure the following github action secrets:
  - *AWS_ACCESS_KEY_ID*: self explanatory
  - *AWS_SECRET_ACCESS_KEY*: self explanatory
  - *POSTGRES_PASSWORD*: RDS PostgreSQL password
  - *REPO_NAME*: ECR repository name

## Initial Run

Execute the terraform.yml workflow to create all the necessary infrastruture. Once the workflow is complete, you must add the following environment variables to your github actions:

Variable of type secrets:

- POSTGRES_DSN: This should be as follow : `postgresql://postgres:POSTGRES_PASSWORD@postgres:5432/postgres`. The terraform output command should give you the necessary informations.
  
Normal variable:

- AWS_REGION: Please checkout the default region set in the terraform.tfvars file
- EKS_CLUSTER_NAME: The name of the EKS cluster created by the terraform script

Disclaimer: this could be set automatically, however due to lack of time , I had to do it manually.

## CICD

Once a push is triggered in the main branch, the CICD should trigger the following workflow:

- Lint test the code
- Execute go test
- Generate PROTO
- Build and push docker image
- Deploy to k8s

## Local development

To ease the setup of a local development environment, developers should use the docker compose file created.
