CREATE TABLE IF NOT EXISTS watchlist (
    channel_id INT8 NOT NULL,
    symbol     TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    PRIMARY KEY (channel_id, symbol)
);

CREATE TABLE IF NOT EXISTS pending_delete (
    req_id     TEXT NOT NULL,
    symbol     TEXT NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL,
    PRIMARY KEY (req_id, symbol)
);
