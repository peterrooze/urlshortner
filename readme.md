```DATABASE_URL="postgres://postgres:mysecretpassword@localhost:5432/urlshortener?sslmode=disable" go run ./cmd/server/```


```
#!/bin/bash

# Start PostgreSQL container if not running
docker start urlshortener-db || docker run --name urlshortener-db -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres

# Ensure database exists
docker exec -it urlshortener-db psql -U postgres -c "CREATE DATABASE urlshortener" || true

# Run the application
DATABASE_URL="postgres://postgres:mysecretpassword@localhost:5432/urlshortener?sslmode=disable" go run ./cmd/server/

``` 

list db entries:
docker exec -it urlshortener-db psql -U postgres -d urlshortener -c "SELECT * FROM urls;"



![URL Shortener Screenshot](https://i.imgur.com/SSwPS6j.png)