-- +goose up
CREATE TABLE jobs (
    positionName VARCHAR(255) PRIMARY KEY,
    companyId INTEGER NOT NULL,
    experienceRequired VARCHAR(255) NOT NULL,
    expectedCompensation VARCHAR(255) NOT NULL,
    employmentType VARCHAR(255) NOT NULL,
    jobLocationType VARCHAR(255) NOT NULL,
    companyId REFERENCES companies(companyId),
);



-- +goose down
DROP TABLE jobs;