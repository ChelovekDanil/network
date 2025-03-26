CREATE DATABASE network;

CREATE TABLE users(
    ID char(36) PRIMARY KEY NOT NULL,
    FirstName varchar(30) NOT NULL,
    LastName varchar(30) NOT NULL
);