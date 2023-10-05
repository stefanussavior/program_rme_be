CREATE TABLE IF NOT EXISTS users(
    id VARCHAR (100) PRIMARY KEY AUTO_INCREMENT,
    nama VARCHAR (100) NULL,
    username VARCHAR (100) NULL,
    email VARCHAR (100) NOT NULL, 
    password VARCHAR (50) NOT NULL,
    alamat TEXT NULL,
    no_telp INT (20) NULL
);