# How to Use

## Before

- gonew

```bash
go install golang.org/x/tools/cmd/gonew@latest
```

- sql-migrate

```bash
go install github.com/rubenv/sql-migrate/...@latest
```

## Make a new project

```bash
gonew github.com/WBear29/go-restapi-server-template example.com/<project>
```

## Run

```bash
go mod tidy
go mod download
go run main.go
```

## Gen migration file

```bash
sql-migrate new <migration_name>
```

# References

- [gonew](https://go.dev/blog/gonew) is required.
- [go-clean-template](https://github.com/evrone/go-clean-template)