-- Create "ent_types" table
CREATE TABLE "ent_types" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "type" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "ent_types_type_key" to table: "ent_types"
CREATE UNIQUE INDEX "ent_types_type_key" ON "ent_types" ("type");
-- Create "users" table
CREATE TABLE "users" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY (START WITH 4294967296), "email" character varying NOT NULL, "password" character varying NOT NULL, "name" character varying NOT NULL, "ip" character varying NOT NULL, "device" character varying NOT NULL, "verified" boolean NOT NULL DEFAULT false, "blocked" boolean NOT NULL DEFAULT false, "username" character varying NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, PRIMARY KEY ("id"));
-- Create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX "users_email_key" ON "users" ("email");
-- Create index "users_username_key" to table: "users"
CREATE UNIQUE INDEX "users_username_key" ON "users" ("username");
-- Create "logins" table
CREATE TABLE "logins" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "ip" character varying NOT NULL, "device" character varying NOT NULL, "access" character varying NOT NULL, "refresh" character varying NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "user_logins" bigint NULL, PRIMARY KEY ("id"), CONSTRAINT "logins_users_logins" FOREIGN KEY ("user_logins") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Add pk ranges for ('logins'),('users') tables
INSERT INTO "ent_types" ("type") VALUES ('logins'), ('users');
