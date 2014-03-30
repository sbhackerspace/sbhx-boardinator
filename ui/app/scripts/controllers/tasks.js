'use strict';

app.controller('TasksCtrl', ['$scope', '$http', '$filter', function ($scope, $http, $filter) {

    $scope.tasks = [];
    $scope.loadTasks = function() {
        var data = {};
        $http.get('/api/tasks', data)
            .then(function(e) {
                if(e.status === 200) {
                    var p,
                        dateFilter = $filter('date'),
                        tasks = e.data;                        
                    for (var i=0; i<tasks.length; i++) {
                        var task = tasks[i];
                        for (p in task) {
                            if (p === 'due_date') {
                                task['due_date'] = dateFilter(task.due_date, 'yyyy/MM/dd');
                            }
                        }
                    }
                    $scope.tasks = tasks;
                }
            })
    };
    
    $scope.submitTaskForm = function(formData) {         
                                    
        var params = {            
            name: formData.name,
            description: formData.description.replace(/\n/g, " "),            
            assignee: formData.assignee,
            due_date: formData.dueDate,
        };
                
        $http.post('/api/tasks', params)
            .then(function(e) {
                if(e.status === 200) {
                    var p,
                        dateFilter = $filter('date'),                    
                        task = e.data;                                        
                    for (p in task) {
                        if (p === 'due_date') {
                            task['due_date'] = dateFilter(task.due_date, 'yyyy/MM/dd');
                        }                            
                    }                                                                                  
                    $scope.tasks.push(task);
                }                            
            })                         
    };

    $scope.editTask = function (task) {                
        $scope.selectedTask = task;        
        $scope.task = {};
        console.log(task);
        $scope.task.name = task.name;
        $scope.task.description = task.description;
        $scope.task.assignee = task.assignee;  
        $scope.task.dueDate  = task.due_date;            
        $scope.showTaskList = false;
        $scope.showTaskForm = true;
        $scope.editedTask = true;

        $scope.submitEdited = function(task) {
            console.log(task);

        };        

    };
   
    $scope.dateOptions = {
        'year-format': "'yyyy'",
        'starting-day': 1
    };    
    $scope.formats = ['MM/dd/yyyy', 'yyyy/MM/dd', 'shortDate'];
    $scope.format = $scope.formats[1];

  }]);
