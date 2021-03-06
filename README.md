# Boardinator == Board Coordinator (get it?!?!)

Boardinator is task management software used by [Santa Barbara
Hackerspace](http://sbhackerspace.com) to help our Board coordinate
more goodly.

Boardinator tracks which tasks our Board members have _been_ assigned
and _have_ assigned to other members of the 'space, then does useful
things with that information (e.g., sends reminder emails to the
assignee, emails the Board when a task is not completed on time, etc).


## Quickstart

### Postres Setup

The following instructions are for Postgres 8.4 and 9.1, and should be
almost identical for newer versions.

    sudo su postgres -c psql < db_setup.sql


### Create New Task

```
curl -X POST -d \
'{"name": "Boardinator MVP", "due_date": "2014-03-22T17:30:00-07:00", "assignee": "elimisteve@gmail.com", "description": "Finish API Task creation"}' \
http://localhost:6060/api/tasks
```


### Update Task

TODO: PUT should include the full Task, and the PATCH HTTP verb should
be used for partial updates such as this one:

`curl -X PUT -d '{"completed":true}' http://localhost:6060/api/tasks/49ebc56f-dfdb-4a11-4d9c-d83d488f987a`


### Get Task

`curl -X GET http://localhost:6060/api/tasks/49ebc56f-dfdb-4a11-4d9c-d83d488f987a`

or simply

`curl http://localhost:6060/api/tasks/49ebc56f-dfdb-4a11-4d9c-d83d488f987a`


### Delete Task

`curl -X DELETE http://localhost:6060/api/tasks/49ebc56f-dfdb-4a11-4d9c-d83d488f987a`


## AngularJS Dependencies installation

### Install Yeoman:

Visit http://yeoman.io/ for detailed documentation:

    npm install -g yo

or

    sudo npm install -g yo


### Use Bower to install your App Dependencies:

Once yeoman is installed...

    cd ui/
    bower install


## Run the App

When bower install is completed, simply type:

    grunt serve

and that will start your Angular Application on localhost!  More
likely you will want to run the backend.  To do so, type

    MANDRILL_KEY="(put it here)" go run boardinator.go

then visit <http://localhost:6060/>.
