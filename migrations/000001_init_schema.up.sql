create table "user"(
  "id" bigserial primary key,
  "username" varchar not null,
  "password" varchar not null,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
create table "article"(
  "id" bigserial primary key,
  "title" varchar not null,
  "description" varchar not null,
  "data" json,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
