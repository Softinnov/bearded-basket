var app = angular.module('app', ['ngRoute']);

app.config(['$routeProvider', function($routeProvider) {
	$routeProvider.when('/', {
		templateUrl: 'html/users.html',
		controller: 'UsersCtrl'
	})
}]);
