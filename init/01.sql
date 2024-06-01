-- sensordb.sensor_data definition

CREATE TABLE IF NOT EXISTS `sensor_data` (
  `id` int NOT NULL AUTO_INCREMENT,
  `value` float DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL,
  `id1` varchar(255) DEFAULT NULL,
  `id2` int DEFAULT NULL,
  `timestamp` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
);