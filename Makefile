.PHONY: help
help: ## Affiche cette aide
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY:dev
dev: public/assets/main.js ## Lance le serveur de développement
	APP_ENV=dev gow run .

.PHONY: build
build:
	bun run build
	go build

public/assets/main.js:
	touch public/assets/main.js
