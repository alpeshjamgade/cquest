CREATE TABLE laptop
(
    id             SERIAL PRIMARY KEY,
    brand          TEXT,
    model          TEXT,
    cpu_id         INT,
    gpu_id         INT,
    ram            INT,
    ssd            INT,
    hdd            INT,
    usb_c          TEXT,
    usb_a          TEXT,
    hdmi           TEXT,
    ethernet       TEXT,
    headphone_jack TEXT,
    weight         float,
    price          float,
    rank           int,
    release_date   TIMESTAMPTZ,
    inserted_at    TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_cpu FOREIGN KEY (cpu_id) REFERENCES cpu (id),
    CONSTRAINT fk_gpu FOREIGN KEY (gpu_id) REFERENCES gpu (id)
);