'use strict';

app.controller('TasksCtrl', ['$scope', '$http', '$filter', function ($scope, $http, $filter) {

    $scope.showTaskList = true;
    $scope.tasks = [
        {assignor: 'Jay Kan', name: 'Task Name', dueDate: '12/31/2015', description: 'Sample Task Description Sample Task Description Sample Task Description', priority: 'High'},
        {assignor: 'AJ', name: 'Task Name', dueDate: '12/31/2015', description: 'Test Sample Task Description Description', priority: 'Medium'},
        {assignor: 'Steve Phillips', name: 'Task Name', dueDate: '12/31/2015', description: 'Test Task Description Sample Task Description', priority: 'Low'},
        {assignor: 'Jim', name: 'Task Name', dueDate: '12/31/2015', description: 'Test Task Description Sample Task Description', priority: 'Medium'},
        {assignor: 'Garry', name: 'Task Name', dueDate: '12/31/2015', description: 'Test Task Description Sample Task Description', priority: 'Low'},
        {assignor: 'Whatever', name: 'Task Name', dueDate: '12/31/2015', description: 'Test Task DescriptionSample Task Description', priority: 'High'},
        {assignor: 'ABC', name: 'Task Name', dueDate: '12/31/2015', description: 'Test Task Sample Task Description Description', priority: 'High'},
    ];
    $scope.submitTaskForm = function(formData) {         
        // Format Date 
        var datefilter = $filter('date'),
            formattedDate = datefilter(formData.dueDate, 'yyyy/MM/dd');   
        // API call            
        // var params = {            
        //     Name: formData.name,
        //     Description: formData.description,
        //     DueDate: formData.dueDate,
        //     Assignee: formData.assignee,
        // };

        // $http.post('url', params)
        //     .then(function(e) {

        //     })                 
        $scope.tasks.push(formData);
    };
   
    $scope.dateOptions = {
        'year-format': "'yyyy'",
        'starting-day': 1
    };    
    $scope.formats = ['MM/dd/yyyy', 'yyyy/MM/dd', 'shortDate'];
    $scope.format = $scope.formats[1];

  }]);
