.PHONY: db clean run

prod-compose-up:
	docker compose --env-file .env -f container_definitions/docker-compose.prod.yml up -d
prod-compose-down:
	docker compose -f container_definitions/docker-compose.prod.yml down

# Remove images
prod-clean:
	docker image rm se_org_prod

dev-compose-up:
	docker compose --env-file .env -f container_definitions/docker-compose.dev.yml up -d

dev-compose-down:
	docker compose -f container_definitions/docker-compose.dev.yml down

# Remove images
# Remove volumes
dev-clean:
	make dev-compose-down
	docker image rm se_org_dev_frontend
	docker image rm se_org_dev_backend

db:
	psql -U postgres -d test -f $(PWD)/db/init/init.sql

clean:
	psql -U postgres -d test -f $(PWD)/db/clean.sql

build_docker:
	docker compose --env-file .env up

connect:
	psql -h localhost -p 9003 -U postgres -d test