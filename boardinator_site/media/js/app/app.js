'use strict';

define(
    [
        'route_resolver'
    ],
    function () {
        var app = angular.module('app',
            [
                'ngRoute',
                'ngResource',
                'routeResolverModule',
                'ui.calendar',
                'ui.bootstrap'
            ]
        );

        app.config(['$routeProvider', '$locationProvider', 'routeResolverProvider', '$controllerProvider', '$compileProvider', '$filterProvider', '$provide',
            function ($routeProvider, $locationProvider, routeResolverProvider, $controllerProvider, $compileProvider, $filterProvider, $provide) {

                app.register =
                {
                    controller: $controllerProvider.register,
                    directive: $compileProvider.directive,
                    filter: $filterProvider.register,
                    factory: $provide.factory,
                    service: $provide.service,
                    constant: $provide.constant
                };

                //Define routes - controllers will be loaded dynamically
                var route = routeResolverProvider.route,
                    hasController = true;

                /**
                *  Resolve URLs and associated controllers
                *  @param (1)   URL path name
                *  @param (2)   (route.resolve) the name of the associated view/controller file
                *  @param (3)   hasController (boolean) load the controller for the view
                */
                $routeProvider
                    .when('/',                          route.resolve('dashboard'))
                    .when('/task',                      route.resolve('task'))
                    .when('/calendar',                  route.resolve('calendar'))
                    .otherwise({ redirectTo: '/' });

                    $locationProvider.html5Mode(true).hashPrefix('!');
            }
        ]);
        return app;
    }
);

