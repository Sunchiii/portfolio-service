ALTER TABLE article
ADD image_exam TEXT;
ALTER TABLE article
ADD article_type VARCHAR;
ALTER TABLE article
ADD user_id INTEGER REFERENCES "user" (id);

ALTER TABLE "user" 
ADD article_id INTEGER REFERENCES article (id);
