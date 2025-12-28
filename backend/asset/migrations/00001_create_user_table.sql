-- +goose Up
CREATE TABLE public.users (
   id SERIAL PRIMARY KEY,
   name VARCHAR(50) NOT NULL,
   login VARCHAR(50) NOT NULL UNIQUE,
   password VARCHAR(255) NOT NULL,
   created_at TIMESTAMP WITH TIME ZONE
);

-- +goose Down
DROP TABLE public.users;