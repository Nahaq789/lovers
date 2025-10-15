@run:
    go run ./...

@test:
    go test ./... -cover

@deploy:
    cd ./terraform/environments/dev/ && \
    terraform init && \
    terraform plan -var-file="terraform.tfvars" && \
    terraform apply -var-file="terraform.tfvars"
