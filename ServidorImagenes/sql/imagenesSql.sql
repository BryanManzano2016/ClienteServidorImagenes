DROP TABLE IF EXISTS imagenes;
CREATE TABLE imagenes(
	nombre VARCHAR PRIMARY KEY,
	datos BYTEA,
	ancho INTEGER,
	largo INTEGER
);
INSERT INTO imagenes(nombre, datos) VALUES('crisis.jpg', pg_read_binary_file('C:\crisis.jpg'));

SELECT pg_read_binary_file('C:\crisis.jpg');
SELECT pg_stat_file('C:\crisis.jpg');