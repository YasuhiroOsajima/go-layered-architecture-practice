DBFILE = test.db
DBDIR = .
DBPATH = ${DBDIR}/${DBFILE}
FILEEXISTS = $(shell ls ${DBDIR} | grep ${DBFILE})

setupdb:
ifeq (${FILEEXISTS}, ${DBFILE})
	@echo 'DB file exists.'
else
	@echo 'Create DB file.'

	@sqlite3 ${DBPATH} < scripts/create_db.sql
endif

cleanupdb:
ifeq (${FILEEXISTS}, ${DBFILE})
	rm ${DBPATH}
endif

addtestuser: setupdb
	@sqlite3 ${DBPATH} < scripts/insert_testuser.sql

build:
	@wire cmd/wire.go
	@go build cmd/command.go cmd/wire_gen.go

runapp:
	@export SQLITE_PATH="`pwd`/test.db"; ./command.exe

unittest:
	@go test go-layered-architecture-practice/internal/repository/sqlite/user
