--- TABLES
CREATE TABLE url (
	id VARCHAR PRIMARY KEY,
	original_url VARCHAR NOT NULL,
	code VARCHAR NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

--- INDEX
CREATE INDEX idx_url_id ON url (id);
CREATE INDEX idx_url_code ON url (code);
