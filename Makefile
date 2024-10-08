.PHONY: help
help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY:dev
dev: public/assets/main.js ## Start the development server
	parallel -j 3 --line-buffer ::: "APP_ENV=dev gow -r=false run ." "bun run dev" "templ generate --watch"

.PHONY: build
build: ## Generates executable
	bun run build
	go build

public/assets/main.js:
	mkdir -p public/assets
	touch public/assets/main.js
