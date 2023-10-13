-- Tables
-- офисы
CREATE TABLE public."office" (
    id bigserial NOT NULL,
    office_id int8 NOT NULL,
    longitude numeric(9, 6) NOT NULL,
    latitude numeric(9, 6) NOT NULL,
    location point NOT NULL,
    address text NOT NULL,
    office_name text NOT NULL,
    is_active bool NOT NULL DEFAULT true,
    timetable jsonb NULL,
    metro_station text NOT NULL,
    has_ramp bool NOT NULL DEFAULT false,
    created timestamp NOT NULL DEFAULT now(),
    updated timestamp NOT NULL DEFAULT now(),
    CONSTRAINT office_pkey PRIMARY KEY (id)
);
CREATE UNIQUE INDEX office_office_id_uindex ON public."office" USING btree (office_id);
CREATE INDEX ON "office" USING GIST(location);

-- очередь с талонами
CREATE TABLE public."queue_tickets" (
    id bigserial NOT NULL,
    ticket_number text NOT NULL,
    service_id int8 NOT NULL,
    created timestamp NOT NULL DEFAULT now(),
    updated timestamp NOT NULL DEFAULT now(),
    CONSTRAINT queue_tickets_pkey PRIMARY KEY (id)
);

-- рейтинг офисов
CREATE TABLE public."office_rating" (
    id bigserial NOT NULL,
    office_id int8 NOT NULL,
    rating float8 NOT NULL,
    ddate timestamp NOT NULL DEFAULT now(),
    created timestamp NOT NULL DEFAULT now(),
    updated timestamp NOT NULL DEFAULT now(),
    CONSTRAINT office_rating_pkey PRIMARY KEY (id)
);
CREATE UNIQUE INDEX office_rating__office_id_ddate_uindex ON public."office_rating" USING btree (office_id, ddate);

-- Extensions
CREATE EXTENSION postgis;