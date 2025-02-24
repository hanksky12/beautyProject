CREATE TABLE IF NOT EXISTS hardware_status_record(
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL,
    hardware_id BIGINT UNSIGNED NOT NULL,
    percent DOUBLE NOT NULL,
    time BIGINT NOT NULL,
    processed BOOLEAN NOT NULL,
    created_at BIGINT  NOT NULL DEFAULT 0,
    updated_at BIGINT  NOT NULL DEFAULT 0,
    FOREIGN KEY(user_id) REFERENCES user(id),
    FOREIGN KEY(hardware_id) REFERENCES hardware(id)
    );

CREATE TABLE IF NOT EXISTS hardware_status_record_raw(
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL,
    hardware_id BIGINT UNSIGNED NOT NULL,
    percent DOUBLE NOT NULL,
    time BIGINT NOT NULL,
    created_at BIGINT  NOT NULL DEFAULT 0,
    updated_at BIGINT  NOT NULL DEFAULT 0,
    FOREIGN KEY(user_id) REFERENCES user(id),
    FOREIGN KEY(hardware_id) REFERENCES hardware(id)
    );

CREATE TABLE IF NOT EXISTS hardware_status_record_average(
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL,
    hardware_id BIGINT UNSIGNED NOT NULL,
    percent DOUBLE NOT NULL,
    time BIGINT NOT NULL,
    created_at BIGINT  NOT NULL DEFAULT 0,
    updated_at BIGINT  NOT NULL DEFAULT 0,
    FOREIGN KEY(user_id) REFERENCES user(id),
    FOREIGN KEY(hardware_id) REFERENCES hardware(id)

    );