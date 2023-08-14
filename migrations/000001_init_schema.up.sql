-- Create the user table
CREATE TABLE IF NOT EXISTS "user" (
  id BIGSERIAL PRIMARY KEY,
  username VARCHAR,
  password VARCHAR,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create the article table
CREATE TABLE IF NOT EXISTS article (
  id BIGSERIAL PRIMARY KEY,
  title VARCHAR,
  description VARCHAR,
  image_exam TEXT,
  article_type VARCHAR,
  data JSON,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE article ADD COLUMN user_id BIGSERIAL REFERENCES "user" (id);

