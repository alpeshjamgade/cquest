CREATE TABLE cpu
(
    id          SERIAL PRIMARY KEY,
    Model       TEXT,
    Cores       INT,
    Threads     INT,
    Cache       INT,
    BaseClock   INT,
    MaxClock    INT,
    Rank        INT,
    ReleaseDate TIMESTAMPTZ,
    inserted_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);