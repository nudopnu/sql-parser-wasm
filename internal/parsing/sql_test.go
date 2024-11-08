package parsing

import "testing"

var shouldParse = `-- +goose Up
CREATE TABLE refresh_tokens (
token TEXT PRIMARY KEY,
	user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	expires_at TIMESTAMP NOT NULL,
 	revoked_at TIMESTAMP,
 	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
 	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
 );
 -- +goose Down
`

func Test(t *testing.T) {
	_, err := ParseSQL(shouldParse)
	if err != nil {
		t.Error(err)
	}
}
