DBFILE = test.db
DBDIR = ./internal/repository/sqlite
#DBDIR = .
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
	@sqlite3 ${DBFILE} < scripts/insert_testuser.sql
