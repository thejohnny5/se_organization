.PHONY: db clean run

db:
	psql -U postgres -d test -f $(PWD)/db/init.sql

clean:
	psql -U postgres -d test -f $(PWD)/db/clean.sql

# run:
# 	npm run dev
# 	air
build_docker:
	docker compose --env-file .env up