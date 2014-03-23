# Boardinator == Board Coordinator (get it?!?!)

Boardinator is task management software used by [Santa Barbara
Hackerspace](http://sbhackerspace.com) to help our Board coordinate
more goodly.

Boardinator tracks which tasks our Board members have _been_ assigned
and _have_ assigned to other members of the 'space, then does useful
things with that information (e.g., sends reminder emails to the
assignee, emails the Board when a task is not completed on time, etc).


## Tutorial

### Create New Task

```
curl -X POST -d '{"name": "Test Task", "due_date": "2014-03-22T19:30:00-07:00", "username": "Steve", "description": "Finish API Task creation"}' http://localhost:6060/api/tasks
```
