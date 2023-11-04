
db:
	"psql -U postgres -d test -f ${PWD}/pkg/models/init.sql"

clean:
	"psql -U postgres -d test -f ${PWD}pkg/models/clean.sql"