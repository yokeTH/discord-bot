-- CreateTable
CREATE TABLE "watchlist" (
    "channel_id" INT8 NOT NULL,
    "symbol" STRING NOT NULL,
    "created_at" TIMESTAMPTZ(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT "watchlist_pkey" PRIMARY KEY ("channel_id","symbol")
);

-- CreateTable
CREATE TABLE "pending_delete" (
    "req_id" STRING NOT NULL,
    "symbol" STRING NOT NULL,
    "expires_at" TIMESTAMPTZ(6) NOT NULL,

    CONSTRAINT "pending_delete_pkey" PRIMARY KEY ("req_id","symbol")
);
