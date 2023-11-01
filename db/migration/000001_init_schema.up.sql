DROP TYPE IF EXISTS levels;
CREATE TYPE levels AS ENUM ('member', 'vendor', 'staff', 'admin');
DROP TYPE IF EXISTS status_type;
CREATE TYPE status_type AS ENUM ('pending', 'failed', 'successful');

CREATE TABLE users (
  username VARCHAR PRIMARY KEY NOT NULL,
  first_name VARCHAR NOT NULL,
  last_name VARCHAR NOT NULL,
  email VARCHAR UNIQUE NOT NULL,
  membership levels DEFAULT 'member' NOT NULL,
  password VARCHAR NOT NULL,
  updated_password_at TIMESTAMPTZ NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  created_at TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE coupon (
  digit VARCHAR PRIMARY KEY NOT NULL,
  created_by VARCHAR REFERENCES users(username) ON DELETE CASCADE,
  used_by VARCHAR REFERENCES users(username) ON DELETE CASCADE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT (now()) 
);

CREATE TABLE bank_details (
  id SERIAL PRIMARY KEY,
  bank_name VARCHAR,
  account_number INTEGER,
  account_name VARCHAR,
  owner VARCHAR REFERENCES users(username) ON DELETE CASCADE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE withdrawal (
  id SERIAL PRIMARY KEY,
  amount BIGINT,
  withdraw_by VARCHAR REFERENCES users(username) ON DELETE CASCADE,
  status status_type DEFAULT 'pending' NOT NULL,
  initiated_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
  completed_at TIMESTAMPTZ NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE TABLE earnings (
  id SERIAL PRIMARY KEY,
  referrals INTEGER,
  referral_balance BIGINT,
  referral_total_earning BIGINT,
  total_withdrawal BIGINT,
  media_earning BIGINT,
  owner VARCHAR REFERENCES users(username) ON DELETE CASCADE
);

CREATE INDEX ON coupon (created_by);
CREATE INDEX ON coupon (used_by);
CREATE INDEX ON bank_details (owner);
CREATE INDEX ON withdrawal (withdraw_by);
CREATE INDEX ON earnings (owner);