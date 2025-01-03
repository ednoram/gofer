-- schema.sql

-- Create user table
CREATE TABLE IF NOT EXISTS user (
  user_id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL
);

-- Create api key table
CREATE TABLE IF NOT EXISTS api_key (
  key_id INTEGER PRIMARY KEY AUTOINCREMENT,
  api_key VARCHAR(50) NOT NULL,
  user_id INTEGER NOT NULL,
  FOREIGN KEY (user_id) REFERENCES user(user_id)
);

-- Create task table
CREATE TABLE IF NOT EXISTS task (
  task_id INTEGER PRIMARY KEY AUTOINCREMENT,
  title VARCHAR(30) NOT NULL,
  description TEXT,
  completed BOOLEAN DEFAULT False
);

