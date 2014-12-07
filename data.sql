-- ----------------------------
--  Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS "public"."users";
CREATE TABLE "public"."users" (
	id serial primary key, -- Sequence structure is created automatically when using 'serial'
	name varchar(40) NOT NULL COLLATE "default",
	email varchar(40) NOT NULL COLLATE "default",
	password varchar(60)
)
WITH (OIDS=FALSE);