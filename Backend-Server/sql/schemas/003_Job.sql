-- +goose up
CREATE TABLE jobs (
    positionName VARCHAR(255) PRIMARY KEY,
    companyId INTEGER NOT NULL REFERENCES companies(companyId),
    experienceRequired VARCHAR(255) NOT NULL,
    expectedCompensation VARCHAR(255) NOT NULL,
    employmentType VARCHAR(255) NOT NULL,
    jobLocationType VARCHAR(255) NOT NULL
);



-- +goose down
DROP TABLE jobs;