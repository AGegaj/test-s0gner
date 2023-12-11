CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE signatures (
    id UUID NOT NULL DEFAULT (uuid_generate_v4()),
    user_id VARCHAR(255) NOT NULL,
    signature VARCHAR(255) NOT NULL UNIQUE,
    answers TEXT[],
    questions TEXT[],
    timestamp TIMESTAMPTZ NOT NULL,

    CONSTRAINT "signatures_pkey" PRIMARY KEY ("id")
);