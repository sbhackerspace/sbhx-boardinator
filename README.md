# Boardinator == Board Coordinator (get it?!?!)

Boardinator is task management software used by [Santa Barbara
Hackerspace](http://sbhackerspace.com) to help our Board coordinate
more goodly.

Boardinator tracks which tasks our Board members have _been_ assigned
and _have_ assigned to other members of the 'space, then does useful
things with that information (e.g., sends reminder emails to the
assignee, emails the Board when a task is not completed on time, etc).


## Tutorial

### Postres Setup

In command line shell (probably Bash):

```
$ sudo su postgres
$ createuser boardinator
```

In Postgres shell (transition with `psql` command):

```
CREATE DATABASE boardinator;
CREATE TABLE tasks (
    Id          varchar(36) NOT NULL,
    Name        varchar(100) NOT NULL,
    Description varchar(4096),
    DueDate     timestamp with time zone,
    Assignee    varchar(100)
);
GRANT ALL PRIVILEGES ON tasks TO boardinator;
ALTER ROLE boardinator WITH PASSWORD 'boardinator';
```


### Create New Task

```
curl -X POST -d '{"name": "Boardinator MVP", "due_date": "2014-03-22T17:30:00-07:00", "assignee": "elimisteve@gmail.com", "description": "Finish API Task creation"}' http://localhost:6060/api/tasks
```
