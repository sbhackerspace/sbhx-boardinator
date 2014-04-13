'use strict';

app.controller('TasksCtrl', 

    ['$scope', '$http', '$filter', 'taskService',

    function ($scope, $http, $filter, taskService) {

        $scope.tasks = [];
        $scope.task = {};      
        var dateFilter = $filter('date');        
       
        $scope.loadTasks = function() {           
            taskService.getTasks().then(function(tasks) {
                $scope.tasks = tasks;               
            })
        };
        
        $scope.submitTaskForm = function(formData) {                           
            taskService.createNewTask(formData).then(function(task) {                                        
                $scope.task = task;
                $scope.tasks.push(task);                
            })
        };

        $scope.complete = function(task) {            
            var status = $scope.task.completed;
            if(status === 'true') {
                console.log("Completed My Task");
            } else {
                console.log("Failed to complete my Task");
            }
            task.completed = true;            
            var params = { id: task.id, completed: task.completed };                       
        };

        $scope.editTask = function(selectedTask) {                                         
            $scope.task = selectedTask;                                     
            $scope.showTaskList = false;
            $scope.showTaskForm = true;
            $scope.editedTask = true;
            $scope.showCompletedTaskList = false;

            $scope.submitEdited = function(task) {     
                taskService.editTask(task).then(function(editedTask) {                                                
                    $scope.task.due_date= editedTask.due_date;
                    $scope.showCompletedTaskList = true;
                })                        
            }; 

            $scope.deleteTask = function(task) {                       
                taskService.deleteTask(task).then(function(response) {
                    if(response.indexOf('Task deleted successfully!') !== -1) {
                        $scope.tasks.splice($scope.tasks.indexOf(task), 1);
                        $scope.showTaskForm = false;
                        $scope.showTaskList = true;
                        $scope.showCompletedTaskList = true;
                    }
                })
            };       
        }; 
        $scope.dateOptions = {
            'year-format': "'yyyy'",
            'starting-day': 1
        };    
        $scope.formats = ['MM/dd/yyyy', 'yyyy/MM/dd', 'shortDate'];
        $scope.format = $scope.formats[1];
  }]);
