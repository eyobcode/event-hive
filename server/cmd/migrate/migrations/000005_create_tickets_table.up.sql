CREATE TABLE IF NOT EXISTS tickets (
    id BIGSERIAL PRIMARY KEY,
    event_id BIGINT NOT NULL,
    ticket_type VARCHAR(50),
    price NUMERIC(10,2),
    quantity INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_ticket_event
        FOREIGN KEY (event_id) REFERENCES events(id)
        ON DELETE CASCADE
);
