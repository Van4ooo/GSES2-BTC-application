down: ## зупиняємо запущені контейнери
	@docker-compose down

up: ## піднімаємо сервіс
	@docker-compose up -d --build app-service --remove-orphans

test: ## тестування правильної роботи сервісу
	@docker-compose up -d --build test-api --remove-orphans
