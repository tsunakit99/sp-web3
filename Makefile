# Makefile

.PHONY: dev-up dev-down supabase-up supabase-down

dev-up:
	@echo "Starting Supabase..."
	supabase start &
	@echo "Starting frontend/backend..."
	docker compose -f docker-compose.dev.yml up --build

dev-down:
	docker compose -f docker-compose.dev.yml down
	supabase stop

supabase-up:
	supabase start

supabase-down:
	supabase stop
	
