app.controller('PDVsCtrl', ['$scope', '$http', function($scope, $http) {

	$http({method: 'GET', url: 'api/pdv'}).success(function(data){
		$scope.pdvs = data;
	}).error(function(){
		alert("Cannot load pdvs");
	});

}]);