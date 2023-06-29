DROP DATABASE IF EXISTS dental_clinic;
CREATE DATABASE dental_clinic;

CREATE TABLE IF NOT EXISTS dentists (
	id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    registration_number VARCHAR(255) NOT NULL
);

CREATE TABLE patients (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    address VARCHAR(255),
    dni VARCHAR(20) NOT NULL,
    registration_date VARCHAR(20) NOT NULL
);

INSERT INTO dentists (name, last_name, registration_number) VALUES
	('John', 'Doe', '12345'),
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


INSERT INTO patients (name, last_name, address, dni, registration_date) VALUES
	('John', 'Doe', '123 Main St', '123456789', '2023-06-01'),
	('Jane', 'Smith', '456 Elm St', '987654321', '2023-06-02'),
	('Michael', 'Johnson', '789 Oak St', '456789123', '2023-06-03'),
	('Emily', 'Williams', '321 Pine St', '789123456', '2023-06-04'),
	('Daniel', 'Brown', '654 Maple St', '321654987', '2023-06-05'),
	('Olivia', 'Jones', '987 Cedar St', '654987321', '2023-06-06'),
	('William', 'Davis', '147 Birch St', '987321654', '2023-06-07'),
	('Sophia', 'Miller', '258 Spruce St', '741852963', '2023-06-08'),
	('James', 'Wilson', '369 Pine St', '852963741', '2023-06-09'),
	('Emma', 'Anderson', '951 Oak St', '963741852', '2023-06-10'),
	('Benjamin', 'Clark', '753 Maple St', '852741963', '2023-06-11'),
	('Ava', 'Lewis', '456 Elm St', '741963852', '2023-06-12'),
	('Alexander', 'Hall', '852 Pine St', '963852741', '2023-06-13'),
	('Mia', 'Moore', '369 Oak St', '741852963', '2023-06-14'),
	('Ethan', 'Young', '753 Birch St', '852963741', '2023-06-15'),
	('Charlotte', 'King', '147 Cedar St', '963741852', '2023-06-16'),
	('Daniel', 'Lee', '951 Maple St', '741852963', '2023-06-17'),
	('Amelia', 'Gonzalez', '258 Pine St', '852963741', '2023-06-18'),
	('Henry', 'Harris', '753 Oak St', '963741852', '2023-06-19'),
	('Harper', 'Walker', '369 Birch St', '741852963', '2023-06-20');