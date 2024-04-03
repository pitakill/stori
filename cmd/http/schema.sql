CREATE TABLE IF NOT EXISTS users (
  id         TEXT PRIMARY KEY,
  first_name TEXT NOT NULL,
  last_name  TEXT NOT NULL,
  email      TEXT UNIQUE NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS accounts (
  id         TEXT PRIMARY KEY,
  user_id    TEXT NOT NULL,
  bank       TEXT NOT NULL,
  number     TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY(user_id) REFERENCES users(user_id)
);

CREATE TABLE IF NOT EXISTS transactions (
  id         TEXT PRIMARY KEY,
  account_id TEXT NOT NULL,
  date       TIMESTAMP NOT NULL,
  credit     BOOLEAN NOT NULL,
  amount     TEXT NOT NULL, -- for store precision we save it as string
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY(account_id) REFERENCES accounts(account_id)
);
