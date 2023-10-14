-- Extensions
CREATE EXTENSION postgis;

-- Tables
-- офисы
CREATE TABLE public."office" (
    id bigserial NOT NULL,
    office_id bigserial NOT NULL,
    longitude numeric(9, 6) NOT NULL,
    latitude numeric(9, 6) NOT NULL,
    location point NOT NULL,
    geom geometry(Point, 4326) NULL,
    address text NOT NULL,
    office_name text NOT NULL,
    is_active bool NOT NULL DEFAULT true,
    timetable_individual jsonb NULL,
    timetable_enterprise jsonb NULL,
    metro_station text NOT NULL,
    handling_ids int8[] NULL,
    client_types int8[] NULL,
    max_people_on_window int8 NOT NULL DEFAULT 0,
    has_ramp bool NOT NULL DEFAULT false,
    created timestamp NOT NULL DEFAULT now(),
    updated timestamp NOT NULL DEFAULT now(),
    CONSTRAINT office_pkey PRIMARY KEY (id)
);
CREATE UNIQUE INDEX office_office_id_uindex ON public."office" USING btree (office_id);
CREATE INDEX ON "office" USING GIST(location);
CREATE INDEX ON "office" USING GIST(geom);

-- очередь с талонами
CREATE TABLE public."queue_tickets" (
    id bigserial NOT NULL,
    ticket_id bigserial NOT NULL,
    office_id int8 NOT NULL,
    handling_id int8 NOT NULL,
    created timestamp NOT NULL DEFAULT now(),
    updated timestamp NOT NULL DEFAULT now(),
    CONSTRAINT queue_tickets_pkey PRIMARY KEY (id)
);

-- список услуг
CREATE TABLE public."handling" (
    id bigserial NOT NULL,
    handling_id bigserial NOT NULL,
    title text NOT NULL,
    client_type int8 NOT NULL,
    handling_duration interval NOT NULL,
    created timestamp NOT NULL DEFAULT now(),
    updated timestamp NOT NULL DEFAULT now(),
    CONSTRAINT services_pkey PRIMARY KEY (id)
);
CREATE UNIQUE INDEX handling_handling_id_client_type_uindex ON public."handling" USING btree (handling_id, client_type);

-- рейтинг офисов
CREATE TABLE public."office_rating" (
    id bigserial NOT NULL,
    office_id int8 NOT NULL,
    rating float8 NOT NULL,
    created timestamp NOT NULL DEFAULT now(),
    updated timestamp NOT NULL DEFAULT now(),
    CONSTRAINT office_rating_pkey PRIMARY KEY (id)
);
CREATE UNIQUE INDEX office_rating_office_id_uindex ON public."office_rating" USING btree (office_id);

-- Insert Data
-- Вставка данных для физических лиц
INSERT INTO public."handling" (title, client_type, handling_duration)
VALUES
    ('Счета и платежи', 1, '1 min'::interval),
    ('Кредитование', 1, '2 min'::interval),
    ('Инвестиции', 1, '3 min'::interval),
    ('Пластиковые карты', 1, '1 min'::interval),
    ('Интернет-банкинг', 1, '2 min'::interval),
    ('Доверительное управление', 1, '3 min'::interval),
    ('Валютные операции', 1, '1 min'::interval),
    ('Сейфовое хранение', 1, '2 min'::interval);

-- Вставка данных для юридических лиц
INSERT INTO public."handling" (title, client_type, handling_duration)
VALUES
    ('Открытие корпоративного счета', 2, '1 min'::interval),
    ('Кредиты и финансирование', 2, '2 min'::interval),
    ('Управление денежными потоками', 2, '3 min'::interval),
    ('Обслуживание пластиковых карт', 2, '1 min'::interval),
    ('Управление рисками', 2, '2 min'::interval);

INSERT INTO public."office"
(longitude, latitude, "location", address, office_name, metro_station, has_ramp, geom)
VALUES(36.984314, 56.184479, '(36.984314, 56.184479)', 'address', 'Address 1', 'Metro', false, ST_SetSRID(ST_MakePoint('36.984314', '56.184479'), 4326));
