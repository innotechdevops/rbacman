
swaggen:
	swag init -g cmd/rbacman/main.go -o docs/apispec

api-dev:
	go run cmd/rbacman/main.go -env development

api-prod:
	go run cmd/rbacman/main.go -env production
