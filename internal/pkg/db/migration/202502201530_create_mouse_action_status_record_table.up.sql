CREATE TABLE IF NOT EXISTS mouse_action_status_record_raw(
            id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
            user_id BIGINT UNSIGNED NOT NULL,
            mouse_action_id BIGINT UNSIGNED NOT NULL,
            x BIGINT NOT NULL,
            y BIGINT NOT NULL,
            time BIGINT NOT NULL,
            created_at BIGINT  NOT NULL DEFAULT 0,
            updated_at BIGINT  NOT NULL DEFAULT 0,
            FOREIGN KEY(user_id) REFERENCES user(id),
            FOREIGN KEY(mouse_action_id) REFERENCES mouse_action(id)
    );