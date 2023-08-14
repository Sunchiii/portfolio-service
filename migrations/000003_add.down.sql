ALTER TABLE article
DROP COLUMN image_exam;

ALTER TABLE article
DROP COLUMN article_type;

ALTER TABLE article 
DROP COLUMN user_id;

ALTER TABLE "user" 
DROP COLUMN article_id;

