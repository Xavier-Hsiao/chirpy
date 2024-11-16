-- +goose UP
ALTER TABLE users
ADD COLUMN hashed_password TEXT DEFAULT 'unset';

-- +goose Down
ALTER TABLE users
DROP COLUMN hashed_password;