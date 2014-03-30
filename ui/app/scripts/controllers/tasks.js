'use strict';

app.controller('TasksCtrl', ['$scope', '$http', '$filter', function ($scope, $http, $filter) {

    $scope.showTaskList = true;
    $scope.tasks = [
        {assignee: 'Jay Kan', name: 'Task Name', dueDate: '12/31/2015', description: 'Sample Task Description Sample Task Description Sample Task Description'},
        {assignee: 'AJ', name: 'Task Name', dueDate: '12/31/2015', description: 'Test Sample Task Description Description'},
        {assignee: 'Steve Phillips', name: 'Task Name', dueDate: '12/31/2015', description: 'Test Task Description Sample Task Description'},
        {assignee: 'Jim', name: 'Task Name', dueDate: '12/31/2015', description: 'Test Task Description Sample Task Description'},
        {assignee: 'Garry', name: 'Task Name', dueDate: '12/31/2015', description: 'Test Task Description Sample Task Description'},
        {assignee: 'Whatever', name: 'Task Name', dueDate: '12/31/2015', description: 'Test Task DescriptionSample Task Description'},
        {assignee: 'ABC', name: 'Task Name', dueDate: '12/31/2015', description: 'Test Task Sample Task Description Description'},
    ];
    $scope.submitTaskForm = function(formData) {         
        
        var dateFilter = $filter('date'),
            formattedDate = dateFilter(formData.dueDate, 'yyyy/MM/dd');   
                    
        var params = {            
            name: formData.name,
            description: formData.description,
            due_date: formData.dueDate,
            assignee: formData.assignee,
        };
        
        var url = "http://localhost:6060/api/tasks";
        
        $http.post('/api/tasks', params)
            .then(function(e) {

            })                 
        $scope.tasks.push(formData);
    };
   
    $scope.dateOptions = {
        'year-format': "'yyyy'",
        'starting-day': 1
    };    
    $scope.formats = ['MM/dd/yyyy', 'yyyy/MM/dd', 'shortDate'];
    $scope.format = $scope.formats[1];

  }]);
