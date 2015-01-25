CREATE USER boardinator WITH PASSWORD 'boardinator';
CREATE DATABASE boardinator;
CREATE TABLE tasks (
    Id             varchar(36) NOT NULL,
    Name           varchar(100) NOT NULL,
    Description    varchar(4096),
    DueDate        timestamp with time zone,
    Assignee       varchar(100),
    Completed      boolean NOT NULL,
    CompletionDate timestamp with time zone
);
GRANT ALL PRIVILEGES ON tasks TO boardinator;
