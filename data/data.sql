DROP DATABASE IF EXISTS dental_clinic;
CREATE DATABASE dental_clinic;

CREATE TABLE IF NOT EXISTS dentists (
	id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    registration_number VARCHAR(255) NOT NULL
);

INSERT INTO dentists (name, last_name, registration_number)
	VALUES ('John', 'Doe', '12345'),
	('Jane', 'Smith', '67890'),
	('Michael', 'Johnson', '54321'),
	('Emily', 'Williams', '09876'),
	('Daniel', 'Brown', '13579'),
	('Sophia', 'Jones', '24680'),
	('William', 'Davis', '98765'),
	('Olivia', 'Taylor', '43210'),
	('James', 'Miller', '56789'),
	('Ava', 'Anderson', '01234'),
	('Benjamin', 'Wilson', '98765'),
	('Mia', 'Martinez', '54321'),
	('Alexander', 'Thomas', '45678'),
	('Charlotte', 'Harris', '78901'),
	('Ethan', 'Clark', '23456'),
	('Amelia', 'Lewis', '90123'),
	('Jacob', 'Young', '56789'),
	('Lily', 'Walker', '34567'),
	('Matthew', 'Moore', '78901'),
	('Sofia', 'King', '89012');
