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

- go-licenses

```bash
go install github.com/google/go-licenses@latest
```

## Make a new project

```bash
gonew github.com/WBear29/go-restapi-server-template example.com/<project>
```

## Run

```bash
go mod tidy
go mod download
docker compose up -d
go run main.go
```

## build

```bash
make build
```

## lint

```bash
make lint
make lint-fix
```

## Gen migration file

```bash
sql-migrate new <migration_name>
```

## Use packages Licenses

- show
    ```bash
    make licenses-show
    ```
- save
    ```bash
    make licenses-report
    ```

# References

- [gonew](https://go.dev/blog/gonew)
- [sql-migrate](https://github.com/rubenv/sql-migrate)
- [go-licenses](https://github.com/google/go-licenses)
- [go-clean-template](https://github.com/evrone/go-clean-template)