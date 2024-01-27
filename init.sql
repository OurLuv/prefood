
-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS category_id_seq;

-- Table Definition
CREATE TABLE "public"."category" (
    "id" int4 NOT NULL DEFAULT nextval('category_id_seq'::regclass),
    "name" varchar,
    "restaurant_id" int4 NOT NULL,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS food_id_seq;

-- Table Definition
CREATE TABLE "public"."food" (
    "id" int4 NOT NULL DEFAULT nextval('food_id_seq'::regclass),
    "name" varchar,
    "description" varchar,
    "category_id" int4,
    "price" int4,
    "in_stock" bool,
    "created_at" timestamp,
    "image" varchar,
    "restaurant_id" int4,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "public"."order_food" (
    "order_id" int4,
    "food_id" int4,
    "quantity" int4
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS orders_id_seq;

-- Table Definition
CREATE TABLE "public"."orders" (
    "id" int4 NOT NULL DEFAULT nextval('orders_id_seq'::regclass),
    "restaurant_id" int4,
    "name" varchar(255),
    "phone" varchar(20),
    "total" int4,
    "status" varchar(255) DEFAULT 'IN_PROCCESS'::character varying,
    "channel" varchar(255),
    "additive" varchar(255),
    "discount" varchar,
    "ordered" timestamp,
    "arrive" timestamp,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS restaurants_id_seq;

-- Table Definition
CREATE TABLE "public"."restaurants" (
    "id" int4 NOT NULL DEFAULT nextval('restaurants_id_seq'::regclass),
    "name" varchar,
    "client_id" int4,
    "phone" varchar,
    "country" varchar,
    "state" varchar,
    "city" varchar,
    "street" varchar,
    "email" varchar,
    "created_at" timestamp,
    "open" bool DEFAULT false,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS users_id_seq;

-- Table Definition
CREATE TABLE "public"."users" (
    "id" int4 NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    "firstname" varchar,
    "lastname" varchar,
    "password" varchar,
    "email" varchar,
    PRIMARY KEY ("id")
);

