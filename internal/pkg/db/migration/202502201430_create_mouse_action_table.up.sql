CREATE TABLE IF NOT EXISTS mouse_action(
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR (50) UNIQUE NOT NULL,
    created_at BIGINT  NOT NULL DEFAULT 0,
    updated_at BIGINT  NOT NULL DEFAULT 0
    );

INSERT INTO mouse_action (name, created_at, updated_at) VALUES ('move', 0, 0);
INSERT INTO mouse_action (name, created_at, updated_at) VALUES ('click', 0, 0);
INSERT INTO mouse_action (name, created_at, updated_at) VALUES ('scroll', 0, 0);