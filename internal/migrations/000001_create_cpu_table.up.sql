CREATE TABLE cpu
(
    id           SERIAL PRIMARY KEY,
    model        TEXT,
    cores        INT,
    threads      INT,
    cache        INT,
    base_clock   INT,
    max_clock    INT,
    rank         INT,
    release_date TIMESTAMPTZ,
    inserted_at  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);