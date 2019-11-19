# bbs api

## Getting Started

Start the db container.

```bash
docker-compose up -d
```

Run migrations.

```bash
go get -v github.com/rubenv/sql-migrate/...
sql-migrate up --config=config/database/dbconfig.yml
```
