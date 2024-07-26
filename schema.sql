CREATE TABLE IF NOT EXISTS secret_message (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    message TEXT NOT NULL
);