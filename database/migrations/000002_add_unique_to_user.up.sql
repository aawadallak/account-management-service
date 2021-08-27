ALTER TABLE users
ADD CONSTRAINT UC_users UNIQUE (username, email);