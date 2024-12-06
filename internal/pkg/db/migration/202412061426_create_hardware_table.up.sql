CREATE TABLE IF NOT EXISTS hardware(
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR (50) UNIQUE NOT NULL,
    created_at BIGINT  NOT NULL DEFAULT 0,
    updated_at BIGINT  NOT NULL DEFAULT 0
    );

INSERT INTO hardware (name, created_at, updated_at) VALUES ('cpu', 0, 0);
INSERT INTO hardware (name, created_at, updated_at) VALUES ('disk', 0, 0);
INSERT INTO hardware (name, created_at, updated_at) VALUES ('memory', 0, 0);