'use strict';

angular.module('uiApp', [
  'ngCookies',
  'ngResource',
  'ngSanitize',
  'ngRoute'
])
  .config(function ($routeProvider) {
    $routeProvider
      .when('/', {
        templateUrl: '/ui/app/views/main.html',
        controller: 'MainCtrl'
      })
      .when('/tasks', {
        templateUrl: '/ui/app/views/tasks.html',
        controller: 'TasksCtrl'
      })
      .when('/calendar', {
        templateUrl: '/ui/app/views/calendar.html',
        controller: 'CalendarCtrl'
      })
      .otherwise({
        redirectTo: '/'
      });
  });
