-- Select rows from a Table or View 'TableOrViewName' in schema 'SchemaName'
CREATE TABLE "user" (
                        "id" bigserial PRIMARY KEY,
                        "user_id" bigint NOT NULL,
                        "username" varchar(64) COLLATE "default" NOT NULL,
                        "password" varchar(64) COLLATE "default" NOT NULL,
                        "email" varchar(64) COLLATE "default",
                        "gender" smallint NOT NULL DEFAULT '0',
                        "create_time" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                        "update_time" timestamp NULL DEFAULT CURRENT_TIMESTAMP
);


DROP TABLE IF EXISTS community;
CREATE TABLE "community" (
                             "id" serial PRIMARY KEY,
                             "community_id" integer NOT NULL,
                             "community_name" varchar(128) COLLATE "default" NOT NULL,
                             "introduction" varchar(256) COLLATE "default" NOT NULL,
                             "create_time" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                             "update_time" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);


INSERT INTO community VALUES ('1', '1', 'Go', 'Golang', '2016-11-01 08:10:10', '2016-11-01 08:10:10');
INSERT INTO community VALUES ('2', '2', 'leetcode', '刷题刷题刷题', '2020-01-01 08:00:00', '2020-01-01 08:00:00');
INSERT INTO community VALUES ('3', '3', 'CS:GO', 'Rush B。。。', '2018-08-07 08:30:00', '2018-08-07 08:30:00');
INSERT INTO community VALUES ('4', '4', 'LOL', '欢迎来到英雄联盟!', '2016-01-01 08:00:00', '2016-01-01 08:00:00');


DROP TABLE IF EXISTS post;
CREATE TABLE "post" (
                        "id" BIGSERIAL PRIMARY KEY,
                        "post_id" BIGINT NOT NULL,
                        "title" VARCHAR(128) NOT NULL,
                        "content" VARCHAR(8192) NOT NULL,
                        "author_id" BIGINT NOT NULL,
                        "community_id" BIGINT NOT NULL,
                        "status" SMALLINT NOT NULL DEFAULT 1,
                        "create_time" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                        "update_time" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX "idx_post_id" ON "post" ("post_id");
CREATE INDEX "idx_author_id" ON "post" ("author_id");
CREATE INDEX "idx_community_id" ON "post" ("community_id");

DROP TABLE IF EXISTS "comment";
CREATE TABLE "comment" (
                           "id" BIGSERIAL PRIMARY KEY,
                           "comment_id" BIGINT NOT NULL,
                           "content" TEXT NOT NULL,
                           "post_id" BIGINT NOT NULL,
                           "author_id" BIGINT NOT NULL,
                           "parent_id" BIGINT NOT NULL DEFAULT '0',
                           "status" SMALLINT NOT NULL DEFAULT '1',
                           "create_time" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                           "update_time" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX "idx_comment_id" ON "comment" ("comment_id");
CREATE INDEX "idx_author_Id" ON "comment" ("author_id");