# sql-parser-wasm 🌐

I needed a way to parse [goose](https://pressly.github.io/goose/) sql migrations in the browser. This repository is meant to be compiled to WebAssembly (WASM) and exposes two functions:
- `parseMigrations(fileContent: string)` should be called to parse a migration file's content and returns the containing up and down migrations as plain text
- `parseSql(sql: string)` should be called to parse each detected statement

## Usage

```js
const migrations = await parseMigrations(`-- +migrate Up
CREATE TABLE refresh_tokens (
token TEXT PRIMARY KEY,
	user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	expires_at TIMESTAMP NOT NULL,
 	revoked_at TIMESTAMP,
 	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
 	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
 );
 -- +migrate Down
 DROP TABLE refresh_tokens;`);
```

## Setup

Compile with
```bash
GOOS=js GOARCH=wasm go build -o example/main.wasm
```
