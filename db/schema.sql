-- schema.sql

-- Create user table
CREATE TABLE user (
  user_id INTEGER PRIMARY KEY AUTOINCREMENT,
  username TEXT NOT NULL UNIQUE
);

-- Create api key table
CREATE TABLE api_key (
  key_id INTEGER PRIMARY KEY AUTOINCREMENT,
  api_key VARCHAR(50) NOT NULL UNIQUE,
  user_id INTEGER NOT NULL,
  FOREIGN KEY (user_id) REFERENCES user(user_id)
);

-- Create task table
CREATE TABLE task (
  task_id INTEGER PRIMARY KEY AUTOINCREMENT,
  title VARCHAR(30) NOT NULL,
  description TEXT,
  completed BOOLEAN DEFAULT False,
  created_by INTEGER NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (created_by) REFERENCES user(user_id)
);

