CREATE TABLE IF NOT EXISTS secret_message (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    message TEXT NOT NULL,
    user_id INT NOT NULL
);

CREATE TABLE IF NOT EXISTS public.user (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL
);