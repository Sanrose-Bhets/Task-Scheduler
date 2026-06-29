-- +goose Up
CREATE TABLE applications (
    id SERIAL PRIMARY KEY,
    department VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL,
    company INTEGER NOT NULL,
    position VARCHAR(255) NOT NULL,
    interviewId INT,
    interviewId REFERENCES interviews(id),
    company REFERENCES companies(companyId),
    position REFERENCES jobs(positionName)
);



-- +goose Down
DROP TABLE applications;
