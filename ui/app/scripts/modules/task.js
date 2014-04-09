'use strict';

    angular.module('taskModule', [])

        .provider('taskService', function() {

            this.$get = function($http, $q, $rootScope, $filter, BoardinatorTask) {

                var dateFilter = $filter('date');

                var getTasks = function() {
                    var deferred = $q.defer(),
                        dateFilter = $filter('date'),                       
                        data = {};

                    $http.get('/api/tasks', data).then(function(e) {
                        if(e.status === 200) {
                            var p,
                                newArr = [],
                                tasks = e.data;
                                        
                            for (var i=0; i<tasks.length; i++) {
                                var task = tasks[i];
                                for (p in task) {
                                    if (p === 'due_date') {
                                        task['due_date'] = dateFilter(task.due_date, 'yyyy/MM/dd');
                                        newArr.push(task);
                                    }
                                }
                            }
                            deferred.resolve(newArr);                                                          
                        }                          
                    });
                    return deferred.promise;
                };  

                var createNewTask = function(params) {
                    var deferred = $q.defer();

                    var data = {
                        name: params.name,
                        description: params.description.replace(/\n/g, " "),
                        assignee: params.assignee,
                        due_date: params.dueDate
                    };

                    $http.post('/api/tasks', data).then(function(e) {
                        if(e.status === 200) {
                            var p,
                                task = e.data;

                            for (p in task) {
                                if (p === 'due_date') {
                                    task['due_date'] = dateFilter(task.due_date, 'yyyy/MM/dd');
                                }
                            }
                            deferred.resolve(task);
                        }
                    });
                    return deferred.promise;
                };

                var editTask = function(params) {
                    var url = '/api/tasks/' + params.id,                        
                        date = new Date(params.dudDate),
                        deferred = $q.defer();
                    
                    var data = {
                        name: params.name,
                        description: params.description.replace(/\n/g, " "),
                        assignee: params.assignee,
                        due_date: date
                    };

                    $http.put(url, data).then(function(e) {
                        if(e.status === 200) {
                            var p,
                                editedTask = e.data;

                            for (p in editedTask) {
                                if (p === 'due_date') {
                                    editedTask['due_date'] = dateFilter(editedTask.due_date, 'yyyy/MM/dd');
                                }
                            }
                            deferred.resolve(editedTask);
                        }
                    });
                    return deferred.promise;
                };

                var deleteTask = function(params) {
                    var url = '/api/tasks/' + params.id,
                        deferred = $q.defer();

                    $http.delete(url, {}).then(function(e) {
                        if(e.status === 200) {
                            deferred.resolve(e.data.response);    
                        }                        
                    });
                    return deferred.promise;
                };                  

                return {
                    getTasks:      getTasks,
                    createNewTask: createNewTask,
                    editTask:      editTask,
                    deleteTask:    deleteTask
                }
           };
        })

        .factory('BoardinatorTask', function() {

            var BoardinatorTask = function(JsonData) {

                for(var key in JsonData) {
                    this[key] = JsonData[key];
                }
            };
            return BoardinatorTask;
        })
        