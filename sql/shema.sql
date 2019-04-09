CREATE TABLE users (
  id       SERIAL PRIMARY KEY,
  phone    VARCHAR(255) NOT NULL,
  name     VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  mail     VARCHAR(255) NOT NULL DEFAULT 'Нет почты',
  pasport  VARCHAR(255) NOT NULL DEFAULT 'Нет паспорта'
);

CREATE TABLE cities (
  id     SERIAL PRIMARY KEY,
  name   VARCHAR(255) NOT NULL,
  photo  VARCHAR(255) NOT NULL DEFAULT 'null',
  offers INTEGER      NOT NULL DEFAULT 0,
  rooms  INTEGER      NOT NULL DEFAULT 0
);

CREATE TABLE hotels (
  id          SERIAL PRIMARY KEY,
  name        VARCHAR(255) NOT NULL,
  country     VARCHAR(255) NOT NULL DEFAULT 'Россия',
  address     VARCHAR(255) NOT NULL DEFAULT 'Не указан',
  latitude    FLOAT        NOT NULL DEFAULT 46.2062966,
  longitude   FLOAT        NOT NULL DEFAULT 6.1466899,
  photo       VARCHAR(255) NOT NULL DEFAULT 'null',
  description TEXT         NOT NULL DEFAULT 'Нет описания',
  stars       INTEGER      NOT NULL DEFAULT 0,

  city_id     INTEGER      NOT NULL,
  FOREIGN KEY (city_id) references cities (id)
);

CREATE TABLE rooms (
  id          SERIAL PRIMARY KEY,
  room        INTEGER NOT NULL,
  persons     INTEGER NOT NULL,
  floor       INTEGER NOT NULL DEFAULT 1,
  housing     INTEGER NOT NULL DEFAULT 1,
  description TEXT    NOT NULL DEFAULT 'Нет описания',
  price       INTEGER NOT NULL DEFAULT 1,
  city_id     INTEGER NOT NULL DEFAULT 0,
  hotel_id    INTEGER NOT NULL,
  FOREIGN KEY (hotel_id) references hotels (id)
);

CREATE TABLE bookings (
  id             SERIAL PRIMARY KEY,
  start_datetime TIMESTAMP NOT NULL,
  end_datetime   TIMESTAMP NOT NULL,
  cost           INTEGER   NOT NULL,
  hotel_id       INTEGER   NOT NULL,
  FOREIGN KEY (hotel_id) references hotels (id),
  room_id        INTEGER   NOT NULL,
  FOREIGN KEY (hotel_id) references rooms (id),
  customer_id    INTEGER   NOT NULL,
  FOREIGN KEY (hotel_id) references users (id)
);
