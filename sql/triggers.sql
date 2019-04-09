--rooms
----create rooms
CREATE OR REPLACE FUNCTION new_room() RETURNS TRIGGER
AS $$
BEGIN
  UPDATE rooms SET city_id = (select hotels.city_id from hotels where hotels.id =  new.hotel_id )
  where rooms.id = new.id;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS tr_new_room ON rooms;

CREATE TRIGGER tr_new_room AFTER INSERT ON rooms
  FOR EACH ROW EXECUTE PROCEDURE new_room();








------------users
----create user
-- CREATE OR REPLACE FUNCTION new_user() RETURNS TRIGGER
--   AS $$
--     BEGIN
--
--       INSERT INTO profiles(user_id) values (new.id);
--       RETURN NEW;
--
--     END;
--   $$ LANGUAGE plpgsql;
--
-- DROP TRIGGER IF EXISTS tr_new_user ON users;
-- CREATE TRIGGER tr_new_user AFTER INSERT ON users
-- FOR EACH ROW EXECUTE PROCEDURE new_user();
----update user
-- CREATE OR REPLACE FUNCTION update_user() RETURNS TRIGGER AS $$
--   BEGIN
--
--     UPDATE users set updated_at = NOW() WHERE id = OLD.id;
--     RETURN OLD;
--
--   END;
-- $$ LANGUAGE plpgsql;
--
-- DROP TRIGGER IF EXISTS  tr_update_user ON users;
-- CREATE TRIGGER tr_update_user AFTER UPDATE ON users FOR EACH ROW EXECUTE  PROCEDURE  update_user();
--
-- ----delete user
-- CREATE OR REPLACE FUNCTION delete_user() RETURNS TRIGGER AS $$
--
--   BEGIN
--
--       DELETE FROM bookings WHERE customer_id = OLD.id;
--       RETURN OLD;
--
--   END;
--   $$ LANGUAGE plpgsql;
--
-- DROP TRIGGER IF EXISTS  tr_delete_user ON users;
-- CREATE TRIGGER tr_delete_user BEFORE DELETE ON users FOR EACH ROW EXECUTE  PROCEDURE  delete_user();
--
-- ------------cities
-- ----update city
-- CREATE OR REPLACE FUNCTION update_city() RETURNS TRIGGER AS $$
-- BEGIN
--
--   UPDATE cities set updated_at = NOW() WHERE id = OLD.id;
--   RETURN OLD;
--
-- END;
-- $$ LANGUAGE plpgsql;
--
-- DROP TRIGGER IF EXISTS  tr_update_city ON cities;
-- CREATE TRIGGER tr_update_city AFTER UPDATE ON cities FOR EACH ROW EXECUTE  PROCEDURE  update_city();
--
-- --hotels
-- ----update hotels
-- CREATE OR REPLACE FUNCTION update_hotel() RETURNS TRIGGER AS $$
-- BEGIN
--
--   UPDATE hotels set updated_at = NOW() WHERE id = OLD.id;
--   RETURN OLD;
--
-- END;
-- $$ LANGUAGE plpgsql;
--
-- DROP TRIGGER IF EXISTS  tr_update_hotel ON hotels;
-- CREATE TRIGGER tr_update_hotel AFTER UPDATE ON hotels FOR EACH ROW EXECUTE  PROCEDURE  update_hotel();


----update rooms
-- CREATE OR REPLACE FUNCTION update_room() RETURNS TRIGGER AS $$
-- BEGIN
--
--   UPDATE rooms set updated_at = NOW() WHERE id = OLD.id;
--   RETURN OLD;
--
-- END;
-- $$ LANGUAGE plpgsql;
--
-- DROP TRIGGER IF EXISTS  tr_update_room ON rooms;
-- CREATE TRIGGER tr_update_room AFTER UPDATE ON rooms FOR EACH ROW EXECUTE  PROCEDURE  update_room();

--bookings
----update bookings
-- CREATE OR REPLACE FUNCTION update_booking() RETURNS TRIGGER AS $$
-- BEGIN
--
--   UPDATE bookings set updated_at = NOW() WHERE id = OLD.id;
--   RETURN OLD;
--
-- END;
-- $$ LANGUAGE plpgsql;
--
-- DROP TRIGGER IF EXISTS  tr_update_booking ON bookings;
-- CREATE TRIGGER tr_update_booking AFTER UPDATE ON bookings FOR EACH ROW EXECUTE  PROCEDURE  update_booking();

-- TODO: триггер на добавление предложений и номеров по занятии и добавлении номеров