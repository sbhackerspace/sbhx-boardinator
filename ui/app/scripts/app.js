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
        templateUrl: '/media/ui/app/views/main.html',
        controller: 'MainCtrl'
      })
      .when('/tasks', {
        templateUrl: '/media/ui/app/views/tasks.html',
        controller: 'TasksCtrl'
      })
      .when('/calendar', {
        templateUrl: '/media/ui/app/views/calendar.html',
        controller: 'CalendarCtrl'
      })
      .otherwise({
        redirectTo: '/'
      });
  });
