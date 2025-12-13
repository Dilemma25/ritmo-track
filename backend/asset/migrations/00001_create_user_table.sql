-- +goose Up
CREATE TABLE users (
   id SERIAL PRIMARY KEY,
   name TEXT NOT NULL,
   login TEXT NOT NULL,
   password TEXT NOT NULL,
   created_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

-- +goose Down
DROP TABLE users;