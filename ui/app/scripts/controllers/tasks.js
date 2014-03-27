'use strict';

app.controller('TasksCtrl', ['$scope', '$http', function ($scope, $http) {
    $scope.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];

    $scope.showTaskList = true;

    $scope.submitTaskForm = function (fromData) {
        var params = {        
            Name: formData.name,
            Description: formData.description,
            DueDate: formData.dueDate,
            Assignee: formData.assignee,
        };

        $http.post('url', params)
            .then(function(e) {                
            })      
    };
   
    $scope.dateOptions = {
        'year-format': "'yy'",
        'starting-day': 1
    };

  }]);
