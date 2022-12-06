build: ## サービスの構築
	docker compose build

up: ## サービス立ち上げ
	docker compose up

stop: ## サービスを停止
	docker compose stop

kill: ## サービスを強制停止
	docker compose kill

down: ## サービスの停止とコンテナの削除
	docker compose down

restart: ## サービスの再起動
	docker compose restart

help: ## ヘルプを表示
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'