'use strict';

app.controller('TasksCtrl', ['$scope', '$http', '$filter', function ($scope, $http, $filter) {

    $scope.showTaskList = true;
    $scope.tasks = [
        {assignee: 'Jay Kan', name: 'Task Name', due_date: '12/31/2015', description: 'Sample Task Description Sample Task Description Sample Task Description'},
        {assignee: 'AJ', name: 'Task Name', due_date: '12/31/2015', description: 'Test Sample Task Description Description'},
        {assignee: 'Steve Phillips', name: 'Task Name', due_date: '12/31/2015', description: 'Test Task Description Sample Task Description'},
        {assignee: 'Jim', name: 'Task Name', due_date: '12/31/2015', description: 'Test Task Description Sample Task Description'},
        {assignee: 'Garry', name: 'Task Name', due_date: '12/31/2015', description: 'Test Task Description Sample Task Description'},
        {assignee: 'Whatever', name: 'Task Name', due_date: '12/31/2015', description: 'Test Task DescriptionSample Task Description'},
        {assignee: 'ABC', name: 'Task Name', due_date: '12/31/2015', description: 'Test Task Sample Task Description Description'},
    ];
    $scope.submitTaskForm = function(formData) {         
        
        var dateFilter = $filter('date'),
            formattedDate = dateFilter(formData.dueDate, 'yyyy/MM/dd');   
                    
        var params = {            
            name: formData.name,
            description: formData.description.replace(/\n/g, " "),
            due_date: formData.dueDate,
            assignee: formData.assignee,
        };
                
        $http.post('/api/tasks', params)
            .then(function(e) {
                if(e.status === 200) {
                    var p,
                        dateFilter = $filter('date'),                    
                        task = e.data;                                        
                    for (p in task) {
                        if (p == 'due_date') {
                            task['due_date'] = dateFilter(task.due_date, 'yyyy/MM/dd');
                        }                            
                    }                                                                                  
                    $scope.tasks.push(task);
                }                            
            })                         
    };  
   
    $scope.dateOptions = {
        'year-format': "'yyyy'",
        'starting-day': 1
    };    
    $scope.formats = ['MM/dd/yyyy', 'yyyy/MM/dd', 'shortDate'];
    $scope.format = $scope.formats[1];

  }]);
