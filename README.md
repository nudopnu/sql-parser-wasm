# sql-parser-wasm

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