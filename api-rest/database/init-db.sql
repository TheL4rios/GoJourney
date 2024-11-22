DROP TABLE IF EXISTS users;

CREATE TABLE users (
	id VARCHAR(32) PRIMARY KEY,
	email VARCHAR(255) NOT NULL UNIQUE,
	password VARCHAR(255) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

INSERT INTO users (id, email, password) VALUES 
('1', 'test@test.com', 'test'),
('2', 'test2@test.com', 'test2'),
('3', 'test3@test.com', 'test3'),
('4', 'test4@test.com', 'test4'),
('5', 'test5@test.com', 'test5');

DROP TABLE IF EXISTS posts;

CREATE TABLE posts (
	id VARCHAR(32) PRIMARY KEY,
	post_content VARCHAR(255) NOT NULL,
	user_id VARCHAR(32) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),

	FOREIGN KEY (user_id) REFERENCES users(id)
)
