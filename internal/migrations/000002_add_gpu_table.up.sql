CREATE TABLE gpu
(
    id           SERIAL PRIMARY KEY,
    model        TEXT,
    memory       INT,
    clock_speed  INT,
    memory_type  TEXT,
    tdp          INT,
    architecture TEXT,
    rank         INT,
    release_date TIMESTAMPTZ,
    inserted_at  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);