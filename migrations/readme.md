# Мои скрипты для работе с миграциями

использую:  
https://github.com/golang-migrate/migrate

	migrate create -ext sql -dir migrations create_extension
	migrate create -ext sql -dir migrations create_events
	migrate create -ext sql -dir migrations create_index

**для основной базы**  

	createdb -h localhost -p 5432 -U postgres pg_calendar
	migrate -path migrations -database postgres://postgres:12345@localhost:5432/pg_calendar?sslmode=disable up
	migrate -path migrations -database postgres://postgres:12345@localhost:5432/pg_calendar?sslmode=disable down


**для тестовой базы**  

	createdb -h localhost -p 5432 -U postgres pg_calendar_test
	migrate -path migrations -database postgres://postgres:12345@localhost:5432/pg_calendar_test?sslmode=disable up
	migrate -path migrations -database postgres://postgres:12345@localhost:5432/pg_calendar_test?sslmode=disable down


CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
DROP TABLE IF EXISTS events;
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


INSERT INTO public.events("header", "text", "start_time", "end_time", "user_name", "reminder_before") VALUES 
('event 1', 'описание события','2020-04-01 09:00:00'::timestamp , '2020-04-01 10:30:00'::timestamp, 'User1', interval '1 hour' ),
('event 2', 'описание события','2020-04-01 09:00:00'::timestamp , '2020-04-01 10:30:00'::timestamp, 'User1', interval '0' ),
('event 3', 'описание события','2020-04-01 09:00:00'::timestamp , '2020-04-01 10:30:00'::timestamp, 'User1', interval '0' ),
('event 4', 'описание события','2020-04-01 09:00:00'::timestamp , '2020-04-01 10:30:00'::timestamp, 'User1', interval '30 minutes' ),
('event 5', 'описание события','2020-04-01 09:00:00'::timestamp , '2020-04-01 10:30:00'::timestamp, 'User1', interval '30 minutes' ),
('event 6', 'описание события','2020-04-01 09:00:00'::timestamp , '2020-04-01 10:30:00'::timestamp, 'User1', interval '2 hour' ),
('event 7', 'описание события','2020-04-01 09:00:00'::timestamp , '2020-04-01 10:30:00'::timestamp, 'User1', interval '1 day' )
;


INSERT INTO public.events("id", "header", "text", "start_time", "end_time", "user_name", "reminder_before") VALUES 
(uuid_generate_v4(), 'event 1', 'описание события','2020-04-01 09:00:00'::timestamp , '2020-04-01 10:30:00'::timestamp, 'User1', interval '1 hour' ),
(uuid_generate_v4(), 'event 2', 'описание события','2020-04-01 09:00:00'::timestamp , '2020-04-01 10:30:00'::timestamp, 'User1', interval '0' ),
(uuid_generate_v4(), 'event 3', 'описание события','2020-04-01 09:00:00'::timestamp , '2020-04-01 10:30:00'::timestamp, 'User1', interval '0' ),
(uuid_generate_v4(), 'event 4', 'описание события','2020-04-01 09:00:00'::timestamp , '2020-04-01 10:30:00'::timestamp, 'User1', interval '30 minutes' ),
(uuid_generate_v4(), 'event 5', 'описание события','2020-04-01 09:00:00'::timestamp , '2020-04-01 10:30:00'::timestamp, 'User1', interval '30 minutes' ),
(uuid_generate_v4(), 'event 6', 'описание события','2020-04-01 09:00:00'::timestamp , '2020-04-01 10:30:00'::timestamp, 'User1', interval '2 hour' ),
(uuid_generate_v4(), 'event 7', 'описание события','2020-04-01 09:00:00'::timestamp , '2020-04-01 10:30:00'::timestamp, 'User1', interval '1 day' )
;
