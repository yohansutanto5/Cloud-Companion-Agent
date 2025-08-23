CREATE TABLE connects (
    id INT PRIMARY KEY,
    host VARCHAR(255),
	port INT
);

insert into connects (id, host, port) values (1, 'localhost', 5432);