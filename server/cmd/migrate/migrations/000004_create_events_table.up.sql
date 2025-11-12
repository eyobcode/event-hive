CREATE TABLE IF NOT EXISTS events (
    id BIGSERIAL PRIMARY KEY,
    organizer_id BIGINT NOT NULL,
    category_id BIGINT,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    location VARCHAR(255),
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_event_organizer
        FOREIGN KEY (organizer_id) REFERENCES organizers(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_event_category
        FOREIGN KEY (category_id) REFERENCES categories(id)
        ON DELETE SET NULL
);
