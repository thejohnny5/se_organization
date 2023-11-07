.PHONY: db clean run

db:
	psql -U postgres -d test -f $(PWD)/pkg/models/init.sql

clean:
	psql -U postgres -d test -f $(PWD)/pkg/models/clean.sql

# run:
# 	npm run dev
# 	air
