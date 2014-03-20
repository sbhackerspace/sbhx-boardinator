'use strict';

define(function () {

    var routeResolverModule = angular.module('routeResolverModule', []);

    //Must be a provider since it will be injected into module.config()
    routeResolverModule.provider('routeResolver', function () {

        this.$get = function () {
            return this;
        };

        this.routeConfig = function () {
            var viewsDirectory = '/media/js/app/views/',
                controllersDirectory = '/media/js/app/controllers/',

            setBaseDirectories = function (viewsDir, controllersDir) {
                viewsDirectory = viewsDir;
                controllersDirectory = controllersDir;
            },

            getViewsDirectory = function (folder) {
                if(folder) { return viewsDirectory + folder + '/'; }
                else { return viewsDirectory; }
            },

            getControllersDirectory = function (folder) {
                if(folder) { return controllersDirectory + folder  + '/'; }
                else { return controllersDirectory; }
            };

            return {
                setBaseDirectories: setBaseDirectories,
                getControllersDirectory: getControllersDirectory,
                getViewsDirectory: getViewsDirectory
            };
        }();

        String.prototype.toCamel = function(){
	        return this.replace(/(\-[a-z])/g, function($1){return $1.toUpperCase().replace('-','');});
        };

        this.route = function (routeConfig) {

            var resolve = function (baseName, hasController, folder) {
                var routeDef = {};
                routeDef.templateUrl = routeConfig.getViewsDirectory(folder) + baseName + '.html';
                if(hasController) {
                    routeDef.controller = baseName.toCamel() + 'Controller';
                    routeDef.resolve = {
                        load: ['$q', '$rootScope', function ($q, $rootScope) {
                            var dependencies = [routeConfig.getControllersDirectory(folder) + baseName + '.js'];
                            return resolveDependencies($q, $rootScope, dependencies);
                        }]
                    };
                }
                return routeDef;
            },

            resolveDependencies = function ($q, $rootScope, dependencies) {
                var defer = $q.defer();
                require(dependencies, function () {
                    defer.resolve();
                    $rootScope.$apply()
                });

                return defer.promise;
            };

            return {
                resolve: resolve
            }
        }(this.routeConfig);
    });
});