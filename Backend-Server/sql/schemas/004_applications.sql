-- +goose Up
CREATE TABLE applications (
    id SERIAL PRIMARY KEY,
    department VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL,
    company INTEGER NOT NULL REFERENCES companies(companyId),
    position VARCHAR(255) NOT NULL REFERENCES jobs(positionName),
    interviewId INT REFERENCES Interviews(id)
);



-- +goose Down
DROP TABLE applications;
