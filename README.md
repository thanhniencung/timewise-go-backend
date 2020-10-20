#start postgres server
docker run --restart=always --name timewise -d -p 5432:5432 -e PGDATA=/var/lib/postgresql/data/pgdata -v /Users/ryan/postgres:/var/lib/postgresql/data -e POSTGRES_PASSWORD=123456 -e POSTGRES_USER=guru -e POSTGRES_DB=timewise postgres

ref: https://hub.docker.com/_/postgres

DB Diagram
https://dbdiagram.io/d/5db91e4cfa792a62f50da62e