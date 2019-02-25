DROP function IF EXISTS fill_rooms();

CREATE FUNCTION fill_rooms()
  RETURNS INTEGER AS $$
BEGIN
  FOR temp_hotel IN 1..18 LOOP
    FOR persons_count IN 1..2 LOOP
      FOR i IN 1..4 LOOP
        insert into rooms (room, persons, hotel_id) values (i, persons_count, temp_hotel);
      END LOOP;
    END LOOP;
  END LOOP;
  RETURN 1;
END;

$$
LANGUAGE plpgsql;

SELECT fill_rooms();


