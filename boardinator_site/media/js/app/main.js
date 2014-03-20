require.config({
    baseUrl: '/media/js/app',
    urlArgs: 'v='+new Date().getTime(),
    paths: {
        app:                      'app',
        route_resolver:           'modules/route_resolver'
    }
});

require(
    [
        'app'
    ],
    function (){
        angular.bootstrap(document, ['app']);
    }
);

