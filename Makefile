db-up:
	 sql-migrate up -config=dbconfig.yml -env="development"
db-down:
	 sql-migrate down -config=dbconfig.yml -env="development"
start-db:
	docker run --name timewise -d -p 5432:5432 -e PGDATA=/var/lib/postgresql/data/pgdata -v /Users/ryan/postgres:/var/lib/postgresql/data -e POSTGRES_PASSWORD=123456 -e POSTGRES_USER=guru -e POSTGRES_DB=timewise postgres