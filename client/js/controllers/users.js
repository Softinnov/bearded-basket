app.controller('UsersCtrl', ['$scope', '$http', '$window', function($scope, $http, $window) {
	$scope.userEdit = null;
	$scope.userNew = null;
	$scope.userCurrent = null;
	$scope.users = [];
	$scope.roles = [];

	$http({method: 'GET', url: 'api/users'}).success(function(data) {
		$scope.users = data;
	}).error(function(){
		alert("cannot get users");
		//$window.location.href = '../';
	});

	$http({method: 'GET', url: 'api/roles'}).success(function(data) {
		$scope.roles = data;
	}).error(function() {
		alert("cannot get roles");
	});

	$http({method: 'GET', url: 'api/user'}).success(function(data) {
		$scope.userCurrent = data;
	}).error(function() {
		alert("cannot get currentUser");
	})

	$scope.newUser = function() {
		$scope.userNew = {};
	};

	$scope.modifyUser = function(user) {
		$scope.userEdit = angular.copy(user);
	};

	$scope.createUser = function() {
		$http({
			method: 'POST',
			url: 'api/users',
			data: $scope.userNew,
		}).success(function(data){
			$scope.users.push($scope.userNew);
			$scope.userNew = false;
		}).error(function(){
			alert('error');
		});
	};

	$scope.editUser = function() {
		$http({
			method: 'PUT',
			url: 'api/users/' + $scope.userEdit.u_id,
			data: $scope.userEdit,
		}).success(function(data){
			angular.forEach($scope.users, function(user, k) {
				if (user.u_id == $scope.userEdit.u_id) {
					$scope.users[k] = angular.copy($scope.userEdit);
					$scope.users[k].u_pass = "";
				}
			});
			$scope.userEdit = false;
		}).error(function(){
			alert('error');
		});
	};

	$scope.deleteUser = function() {
		$http({
			method: 'DELETE',
			url: 'api/users/' + $scope.userEdit.u_id,
		}).success(function(data){
			_.remove($scope.users, function(user) {
				return user.u_id == $scope.userEdit.u_id
			});
			$scope.userEdit = false;
		}).error(function(){
			alert('error');
		});
	};
}]);
