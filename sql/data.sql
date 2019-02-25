--fill users
insert into users (phone, name, password)
values ('89000000001', 'user1', 'user1');
insert into users (phone, name, password)
values ('89000000002', 'user2', 'user2');
insert into users (phone, name, password)
values ('89000000003', 'user3', 'user3');
insert into users (phone, name, password)
values ('89000000004', 'user4', 'user4');
insert into users (phone, name, password)
values ('89000000005', 'user5', 'user5');
insert into users (phone, name, password)
values ('89000000006', 'user6', 'user6');
insert into users (phone, name, password)
values ('89000000007', 'user7', 'user7');
insert into users (phone, name, password)
values ('89000000008', 'user8', 'user8');
insert into users (phone, name, password)
values ('89000000009', 'user9', 'user9');
insert into users (phone, name, password)
values ('89000000010', 'user10', 'user10');

--fill cities
insert into cities (name, photo)--1
values ('Москва', 'https://i1.wp.com/gubdaily.ru/wp-content/uploads/2018/01/moskva.jpg?resize=720%2C481&ssl=1'),
       ('Санкт-Петербург', 'https://i0.wp.com/gubdaily.ru/wp-content/uploads/2018/01/piter.jpg?resize=720%2C480&ssl=1'),
       ('Тюмень', 'https://i2.wp.com/gubdaily.ru/wp-content/uploads/2018/01/tumen.jpg?resize=720%2C480&ssl=1'),
       ('Грозный', 'https://i2.wp.com/gubdaily.ru/wp-content/uploads/2018/01/groznii.jpg?resize=720%2C479&ssl=1'),
       ('Казань', 'https://i0.wp.com/gubdaily.ru/wp-content/uploads/2018/01/kazan1.jpg?resize=720%2C471&ssl=1'),
       ('Краснодар', 'https://i0.wp.com/gubdaily.ru/wp-content/uploads/2018/01/krasnodar.jpg?resize=720%2C477&ssl=1'),
       ('Уфа', 'https://i0.wp.com/gubdaily.ru/wp-content/uploads/2018/01/ufa1.jpg?resize=720%2C480&ssl=1'),
       ('Новосибирск',
        'https://i1.wp.com/gubdaily.ru/wp-content/uploads/2018/01/novosibirskqq.jpg?resize=720%2C480&ssl=1'),
       ('Красноярск',
        'https://i0.wp.com/gubdaily.ru/wp-content/uploads/2018/01/krasnoyarsk.jpg?resize=720%2C480&ssl=1'),
       ('Кемерово', 'https://i1.wp.com/gubdaily.ru/wp-content/uploads/2018/01/kemerovo.jpg?resize=720%2C404&ssl=1');

--fill hotels
---- Мoscow
insert into hotels (name, city_id, photo)
values ('Four Seasons', 1, 'http://rating.msk.ru/img/hotels/four_seasons/thumb/four_seasons_thumb.jpg'),
       ('Лотте Москва', 1, 'http://rating.msk.ru/img/hotels/lotte/thumb/lotte_thumb.jpg'),
       ('Националь', 1, 'http://rating.msk.ru/img/hotels/nacional/thumb/nacional_thumb.jpg'),
       ('Свиссотель Красные Холмы', 1, 'http://rating.msk.ru/img/hotels/jazz_hotel/thumb/jazz_hotel_thumb.jpg'),
       ('Арарат Парк Хаятт', 1, 'http://rating.msk.ru/img/hotels/mirit/thumb/mirit_thumb.jpg'),
       ('Baltschug Kempinski', 1, 'http://rating.msk.ru/img/hotels/medeya/thumb/medeya_thumb.jpg'),
       ('The Ritz-Carlton Moscow', 1, 'http://rating.msk.ru/img/hotels/savoy/thumb/savoy_thumb.jpg'),
       ('Метрополь', 1, 'http://rating.msk.ru/img/hotels/metropol/thumb/metropol_thumb.jpg'),
       ('Moscow Marriott Royal Aurora', 1, 'http://rating.msk.ru/img/hotels/domino_hotel/thumb/domino_hotel_thumb.jpg'),
----SPB
       ('Коринтия Санкт-Петербург',
        2,
        'http://where.ru/media/pic/ano/9/corinthia_hotel_st_petersburg__korintiya_sankt-peterburg.jpg'),
       ('Ренессанс Санкт-Петербург Балтик Отель',
        2,
        'http://where.ru/media/pic/ano/4/renaissance_baltic_exterior_prevminijpg.jpg'),
       ('Belmond Grand Hotel Europe', 2, 'http://where.ru/media/pic/ano/7/grand_otel_evropa.jpg'),
       ('Талион Империал Отель', 2, 'http://where.ru/media/pic/ano/5/taleon_imperial_hotel__talion_imperial_otel.jpg'),
       ('Отель “Кемпински Мойка 22”', 2, 'http://where.ru/media/pic/ano/4/otel_kempinski_moyka_22_.jpg'),
       ('Park Inn by Radisson Невский, Санкт-Петербург',
        2,
        'http://where.ru/media/pic/ano/8/park_inn_by_radisson__nevskiy_sankt-peterburg.jpg'),
       ('Отель “Астория”', 2, 'http://where.ru/media/pic/ano/9/rfh_hotel_astoria.jpg'),
       ('Crowne Plaza St. Petersburg Airport',
        2,
        'http://where.ru/media/pic/ano/10/crowne_plaza_st_petersburg_airport.jpg'),
       ('W St. Petersburg', 2, 'http://where.ru/media/pic/ano/5/who3270gr111762minijpgjpg.jpg');


DROP function IF EXISTS fill_rooms();

CREATE FUNCTION fill_rooms()
  RETURNS INTEGER AS $$
BEGIN
  FOR temp_hotel IN 1..18 LOOP
    FOR persons_count IN 1..5 LOOP
      FOR i IN 1..4 LOOP
        insert into rooms (room, persons, hotel_id) values (i, persons_count, temp_hotel);
      END LOOP;
    END LOOP;
  END LOOP;

  RETURN (SELECT COUNT(id) from rooms);
END;

$$
LANGUAGE plpgsql;

SELECT fill_rooms();
