set dotenv-load := true

@run:
    go run ./...

@di_aws:
    wire ./cmd/di/aws

@di_auth:
    wire ./cmd/di/auth

@di_user:
    wire ./cmd/di/user

@di_group:
    wire ./cmd/di/group

@di_template:
    wire ./cmd/di/template

@di_expense:
    wire ./cmd/di/expense

@di_all:
    just di_aws
    just di_auth
    just di_user
    just di_group
    just di_template
    just di_expense

@test:
    go test ./... -cover

@deploy:
    cd ./terraform/environments/dev/ && \
    terraform init && \
    terraform plan -var-file="terraform.tfvars" && \
    terraform apply -var-file="terraform.tfvars"

@run_db:
    docker compose up -d

@stop_db:
    docker compose stop
