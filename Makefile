DBFILE = test.db
FILEEXISTS = $(shell ls | grep ${DBFILE})

setupdb:
ifeq (${FILEEXISTS}, ${DBFILE})
	@echo 'DB file exists.'
else
	@echo 'Create DB file.'

	@sqlite3 ${DBFILE} < scripts/create_db.sql
endif

cleanupdb:
ifeq (${FILEEXISTS}, ${DBFILE})
	rm ${DBFILE}
endif

addtestuser: setupdb
	@sqlite3 ${DBFILE} < scripts/insert_testuser.sql