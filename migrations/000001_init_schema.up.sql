-- Create the user table
CREATE TABLE "user" (
  id SERIAL PRIMARY KEY,
  username VARCHAR,
  password VARCHAR,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create the article table
CREATE TABLE article (
  id SERIAL PRIMARY KEY,
  title VARCHAR,
  description VARCHAR,
  imageExam TEXT,
  article_type VARCHAR,
  data JSON,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE article ADD COLUMN userId SERIAL REFERENCES "user" (id);
ALTER TABLE "user" ADD COLUMN article_id SERIAL REFERENCES article (id);

