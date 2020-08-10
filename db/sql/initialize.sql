DROP SCHEMA IF EXISTS location;
CREATE SCHEMA location;
DROP SCHEMA IF EXISTS event;
CREATE SCHEMA event;
USE location;

DROP TABLE IF EXISTS gps;
CREATE TABLE gps(
  id INT AUTO_INCREMENT PRIMARY KEY,
  event_id INT NOT NULL,
  date DATETIME,
  latitude DOUBLE,
  longitude DOUBLE,
  altitude DOUBLE,
  accuracy DOUBLE,
  altitudeAccuracy DOUBLE,
  heading DOUBLE,
  speed DOUBLE
);
