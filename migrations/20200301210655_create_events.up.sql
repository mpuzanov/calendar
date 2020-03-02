CREATE TABLE events
(
	"id" uuid NOT NULL DEFAULT uuid_generate_v4(),
	"header" character varying(100) NOT NULL,
	"text" text,
	"start_time" timestamp without time zone NOT NULL,
	"end_time" timestamp without time zone NOT NULL,
	"user_name" character varying(50) NOT NULL,
	"reminder_before" interval(16),
	CONSTRAINT events_pkey PRIMARY KEY ("id")
);