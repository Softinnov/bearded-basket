var app = angular.module('app', ['ngRoute']);

app.config(['$routeProvider', function($routeProvider) {
	$routeProvider
		.when('/', {
			templateUrl: 'static/pdvs.html',
			controller: 'PDVsCtrl'
		})
		.when('/pdvs/:id', {
			templateUrl: 'static/pdv.html',
			controller: 'PDVsCtrl'
		})
		.otherwise({
			redirectTo: '/'
		})
}])