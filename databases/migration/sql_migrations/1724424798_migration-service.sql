-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE services (
       id              SERIAL PRIMARY KEY,
       name            VARCHAR(256) NOT NULL,
       description     VARCHAR(256),
       file_name       VARCHAR(256),
       created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
       created_by      VARCHAR(356) DEFAULT 'SYSTEM',
       modified_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
       modified_by     VARCHAR(356) DEFAULT 'SYSTEM'
);

-- +migrate StatementEnd