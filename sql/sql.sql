CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS seguidores;
DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(100) not null,
    criadoEm timestamp default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE seguidores(
   usuario_id int not null,
   FOREIGN KEY (usuario_id)
   REFERENCES usuarios(id)
   ON DELETE CASCADE,

   seguidor_id int not null,
   FOREIGN KEY (seguidor_id)
   REFERENCES usuarios(id)
   ON DELETE CASCADE,

   PRIMARY KEY (usuario_id, seguidor_id)
) ENGINE=INNODB;


USE devbook;

DROP TABLE IF EXISTS devices;

CREATE TABLE devices(
    id int auto_increment primary key,
    name varchar(50) not null,
    address varchar(50) not null unique,
    latitude float not null,
    longitude float not null
) ENGINE=INNODB;

USE devbook;

DROP TABLE IF EXISTS records;

CREATE TABLE records(
    id bigint auto_increment primary key,
    value DOUBLE not null,
    createdAt timestamp default current_timestamp(),
    device_id int not null,
    KEY FK_device_id (device_id),
    CONSTRAINT FK_device_id FOREIGN KEY (device_id)
    REFERENCES devices(id)
) ENGINE=INNODB;
