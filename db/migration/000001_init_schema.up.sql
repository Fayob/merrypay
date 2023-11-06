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
  used_by VARCHAR UNIQUE REFERENCES users(username) ON DELETE CASCADE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT (now()) 
);

CREATE TABLE bank_details (
  id SERIAL PRIMARY KEY NOT NULL,
  bank_name VARCHAR NOT NULL,
  account_number VARCHAR NOT NULL,
  account_name VARCHAR NOT NULL,
  owner VARCHAR UNIQUE REFERENCES users(username) ON DELETE CASCADE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE withdrawal (
  id SERIAL PRIMARY KEY NOT NULL,
  amount BIGINT NOT NULL,
  withdraw_by VARCHAR NOT NULL REFERENCES users(username) ON DELETE CASCADE,
  status status_type DEFAULT 'pending' NOT NULL,
  initiated_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
  completed_at TIMESTAMPTZ NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE TABLE earnings (
  id SERIAL PRIMARY KEY NOT NULL,
  referrals INTEGER DEFAULT 0 NOT NULL,
  referral_balance BIGINT DEFAULT 0 NOT NULL,
  referral_total_earning BIGINT DEFAULT 0 NOT NULL,
  total_withdrawal BIGINT DEFAULT 0 NOT NULL,
  media_earning BIGINT DEFAULT 0 NOT NULL,
  owner VARCHAR NOT NULL UNIQUE REFERENCES users(username) ON DELETE CASCADE
);

CREATE INDEX ON coupon (created_by);
CREATE INDEX ON coupon (used_by);
CREATE INDEX ON bank_details (owner);
CREATE INDEX ON withdrawal (withdraw_by);
CREATE INDEX ON earnings (owner);