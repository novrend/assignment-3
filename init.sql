CREATE TABLE IF NOT EXISTS data (
    id   SERIAL PRIMARY KEY,
    water INTEGER,
    wind INTEGER
);

INSERT INTO data (water, wind) VALUES (0,0);