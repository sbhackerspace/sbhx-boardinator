'use strict';

var app = angular.module('uiApp', [
  'ngCookies',
  'ngResource',
  'ngSanitize',
  'ngRoute'
]);


app.config(function ($routeProvider) {
    $routeProvider
      .when('/', {
        templateUrl: 'views/main.html',
        controller: 'MainCtrl'
      })
      .when('/tasks', {
        templateUrl: 'views/tasks.html',
        controller: 'TasksCtrl'
      })
      .when('/calendar', {
        templateUrl: 'views/calendar.html',
        controller: 'CalendarCtrl'
      })
      .otherwise({
        redirectTo: '/'
      });
  });


