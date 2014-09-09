app.controller('PDVsCtrl', ['$scope', 'PDV', '$routeParams', '$rootScope', function($scope, PDV, $routeParams, $rootScope) {

	$rootScope.loading = true;

	PDV.get($routeParams.id).then(function(pdv) {
		$rootScope.loading = false;
		$scope.id = pdv.pv_id;
		$scope.name = pdv.pv_nom;
		$scope.expire = pdv.pv_abo_expire;
	}, function(msg) {
		alert(msg);
	});

	$scope.pdvs = PDV.find().then(function(pdvs) {
		$rootScope.loading = false;
		$scope.pdvs = pdvs;
	}, function(msg) {
		alert(msg);
	});

}])