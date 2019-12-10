CREATE TABLE IF NOT EXISTS users (
    	user_id         VARCHAR(100) PRIMARY KEY,
	email           VARCHAR(100) NOT NULL UNIQUE,
	user_name       VARCHAR(100) NOT NULL UNIQUE,
	password        VARCHAR(100) NOT NULL
) CHARSET=utf8;
