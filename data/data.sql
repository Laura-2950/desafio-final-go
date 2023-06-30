-- DROP DATABASE IF EXISTS dental_clinic;
-- CREATE DATABASE dental_clinic;

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

CREATE TABLE shifts (
    id INT NOT NULL AUTO_INCREMENT,
    patient_id INT,
    dentist_id INT,
    date_hour VARCHAR(50) NOT NULL,
	description VARCHAR(255) DEFAULT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (patient_id) REFERENCES patients(id),
    FOREIGN KEY (dentist_id) REFERENCES dentists(id)
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
	('John', 'Doe', '123 Main St', '123456789', '06/01/2023'),
	('Jane', 'Smith', '456 Elm St', '987654321', '06/02/2023'),
	('Michael', 'Johnson', '789 Oak St', '456789123', '06/03/2023'),
	('Emily', 'Williams', '321 Pine St', '789123456', '06/04/2023'),
	('Daniel', 'Brown', '654 Maple St', '321654987', '06/05/2023'),
	('Olivia', 'Jones', '987 Cedar St', '654987321', '06/06/2023'),
	('William', 'Davis', '147 Birch St', '987321654', '06/06/2023'),
	('Sophia', 'Miller', '258 Spruce St', '741852963', '06/05/2023'),
	('James', 'Wilson', '369 Pine St', '852963741', '06/05/2023'),
	('Emma', 'Anderson', '951 Oak St', '963741852', '06/01/2023'),
	('Benjamin', 'Clark', '753 Maple St', '852741963', '06/01/2023'),
	('Ava', 'Lewis', '456 Elm St', '741963852', '06/01/2023'),
	('Alexander', 'Hall', '852 Pine St', '963852741', '06/05/2023'),
	('Mia', 'Moore', '369 Oak St', '741852963', '06/04/2023'),
	('Ethan', 'Young', '753 Birch St', '852963741', '06/02/2023'),
	('Charlotte', 'King', '147 Cedar St', '963741852', '06/05/2023'),
	('Daniel', 'Lee', '951 Maple St', '741852963', '06/04/2023'),
	('Amelia', 'Gonzalez', '258 Pine St', '852963741', '06/06/2023'),
	('Henry', 'Harris', '753 Oak St', '963741852', '06/04/2023'),
	('Harper', 'Walker', '369 Birch St', '741852963', '06/02/2023');


INSERT INTO shifts (patient_id, dentist_id, date_hour, description) VALUES 
	(9, 20, '04/07/2023 11:30', "consulta anual"),
	(13, 3, '10/08/2023 18:45', "consulta anual"),
	(1, 15, '04/08/2023 15:00', "consulta anual"),
	(10, 5, '25/07/2023 18:00', "consulta anual"),
	(13, 20,'10/08/2023 18:30', "consulta anual"),
	(1, 5,'20/07/2023 11:30', "consulta anual"),
	(1, 18,'14/07/2023 08:00', "consulta anual");