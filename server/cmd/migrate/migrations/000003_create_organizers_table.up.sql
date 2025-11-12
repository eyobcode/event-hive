CREATE TABLE IF NOT EXISTS organizers (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    organization_name VARCHAR(255),
    contact_email VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_organizer_user
        FOREIGN KEY (user_id) REFERENCES users(id)
        ON DELETE CASCADE
);
