CREATE SEQUENCE user_id_seq;

CREATE TABLE users (
    id int NOT NULL PRIMARY KEY DEFAULT nextval('user_id_seq'),
    name varchar(50) NOT NULL,
    username varchar(20) UNIQUE NOT NULL,
    email varchar(30) NOT NULL,
    password varchar(20) NOT NULL,
    gender varchar(10) NOT NULL,
    role varchar(10) NOT NULL,
    course varchar(20) NOT NULL,
    image varchar(225) ,
    about varchar(300),
    token varchar(225) ,
    joindate varchar(225) NOT NULL
); 

ALTER SEQUENCE user_id_seq OWNED BY users.id;