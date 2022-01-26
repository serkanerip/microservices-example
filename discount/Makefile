postgres:
	docker rm -f pg_test pg_admin_test
	docker run -p 5432:5432 -e POSTGRES_PASSWORD=postgres --name pg_test -d postgres
	docker run -p 8080:80 -e PGADMIN_DEFAULT_EMAIL=admin@admin.com -e PGADMIN_DEFAULT_PASSWORD=root --name pg_admin_test -d dpage/pgadmin4