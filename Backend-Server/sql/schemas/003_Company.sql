-- +goose up
CREATE TABLE companies (
    companyId SERIAL PRIMARY KEY,
    companyName VARCHAR(255) NOT NULL,
    companyLocation VARCHAR(255) NOT NULL,
    companyWebsite VARCHAR(255) NOT NULL,
    companyEmail VARCHAR(255) NOT NULL,
    HiringManagerName VARCHAR(255) DEFAULT 'N/A'

);





-- +goose down
DROP TABLE companies;