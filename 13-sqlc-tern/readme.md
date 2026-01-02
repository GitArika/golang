# Install necessary dependencies

## sqlc

```sh
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

once installed, create a sqlc.yaml running the following:

```sh
#!/bin/sh

cat << 'EOF' > sqlc.yaml
version: "2"
sql:
  - engine: "postgresql"
    queries: "query.sql"
    schema: "schema.sql"
    gen:
      go:
        package: "main"
        out: "."
        sql_package: "pgx/v5"
EOF
```

## tern

Installe o tern.

```sh
go install github.com/jackc/tern/v2@latest
```

Inicialize. 
> Note que um arquivo tern.conf será gerado na raiz do projeto.

_é necessário editar as configurações do arquivo tern.conf para prosseguir_

```sh
tern init
```

Crie sua primeira migration

```sh
tern new create_table_authors
```

Exemplo de migration
```sql
CREATE TABLE authors (
  id   BIGSERIAL PRIMARY KEY,
  name text      NOT NULL,
  bio  text
);

---- create above / drop below ----

DROP TABLE authors;
```

Aplicando a migration:

```sh
tern migrate
```

Crie um arquivo query.sql (para uso do sqlc)

```
#!/bin/sh

cat << 'EOF' > query.sql
-- name: GetAuthor :one
SELECT * FROM authors
WHERE id = $1 LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM authors
ORDER BY name;

-- name: CreateAuthor :one
INSERT INTO authors (
  name, bio
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateAuthor :exec
UPDATE authors
  set name = $2,
  bio = $3
WHERE id = $1;

-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = $1;
EOF
```

Generate code

```sh
sqlc generate -f ./sqlc.yaml
```

Por fim, atualize a lista de dependencias

```sh
go mod tidy
```
