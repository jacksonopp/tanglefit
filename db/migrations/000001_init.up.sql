CREATE TYPE "role" AS ENUM (
  'member',
  'owner',
  'admin',
  'instructor'
);

CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamp,
  "role" role
);

CREATE TABLE "gyms" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "owner_id" int,
  "member_id" int
);

CREATE TABLE "classes" (
  "id" SERIAL PRIMARY KEY,
  "gym_id" int,
  "attendee_id" int,
  "instructor_id" int,
  "date" timestamp NOT NULL,
  "max_attendees" int NOT NULL,
  "class_type" varchar
);

ALTER TABLE "gyms" ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("id");

ALTER TABLE "gyms" ADD FOREIGN KEY ("member_id") REFERENCES "users" ("id");

ALTER TABLE "classes" ADD FOREIGN KEY ("gym_id") REFERENCES "gyms" ("id");

ALTER TABLE "classes" ADD FOREIGN KEY ("attendee_id") REFERENCES "users" ("id");

ALTER TABLE "classes" ADD FOREIGN KEY ("instructor_id") REFERENCES "users" ("id");

CREATE INDEX ON "users" ("email");
