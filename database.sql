CREATE TABLE IF NOT EXISTS users (
     id SERIAL PRIMARY KEY,
     full_name VARCHAR(60) NOT NULL,
     password VARCHAR(64) NOT NULL,
     phone VARCHAR(15) UNIQUE NOT NULL,
     created_at TIMESTAMP NOT NULL,
     updated_at TIMESTAMP,
     deleted_at TIMESTAMP,
     login_success INT NOT NULL DEFAULT(0)
);

CREATE INDEX users_phone_desc ON users(phone DESC);
CREATE INDEX users_createdat_desc ON users(created_at DESC);
CREATE INDEX users_deletedat_desc ON users(deleted_at DESC);