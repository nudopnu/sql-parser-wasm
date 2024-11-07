# sql-parser-wasm 🌐

This repository provides a WebAssembly (WASM) module for parsing [Goose](https://pressly.github.io/goose/) SQL migrations in the browser. It exposes two main functions:
- `parseMigrations(fileContent: string)`: Parses the content of a migration file, returning the `UP` and `DOWN` migration statements as plain text.
- `parseSql(sql: string)`: Parses individual SQL statements detected within a migration.

## Sample usage

```js
await parseMigrations(`-- +goose Up
CREATE TABLE refresh_tokens (
token TEXT PRIMARY KEY,
	user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	expires_at TIMESTAMP NOT NULL,
 	revoked_at TIMESTAMP,
 	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
 	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
 );
 -- +goose Down
 DROP TABLE refresh_tokens;`);
// {UpStatements: Array(1), DownStatements: Array(1), DisableTransactionUp: false, DisableTransactionDown: false}

await parseSQL(`
CREATE TABLE refresh_tokens (
token TEXT PRIMARY KEY,
	user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	expires_at TIMESTAMP NOT NULL,
 	revoked_at TIMESTAMP,
 	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
 	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
 );`);
// {Action: 'create', Table: {…}, NewName: {…}, IfExists: false, Columns: Array(6), …}

await parseSQL("DROP TABLE refresh_tokens;");
// {Action: 'drop', Table: {…}, NewName: {…}, IfExists: false}
```

## Setup

Make sure you have [Go installed](https://go.dev/dl/) (version 1.16 or later) on your machine.

#### Building the WebAssembly Binary

To compile the Go code into WebAssembly, run:

```bash
GOOS=js GOARCH=wasm go build -o example/main.wasm
```

This will produce `main.wasm` in the example directory.

> Note: The `wasm_exec.js` file in the `example` folder is provided by Go and is required to run Webssembly modules in the browser. For the latset version, see the Go source [here](https://github.com/golang/go/blob/master/lib/wasm/wasm_exec.js).
