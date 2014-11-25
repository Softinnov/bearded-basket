app.controller('UsersCtrl', ['$scope', '$http', '$window', function($scope, $http, $window) {
	$scope.userEdit = null;
	$scope.userNew = null;
	$scope.userCurrent = null;
	$scope.users = [];
	$scope.roles = [];
	$scope.showCreateUser = false;
	$scope.showEditUser = false;
	$scope.error = null;
	$scope.load = 0;

	$http({method: 'GET', url: 'api/users'}).success(function(data) {
		$scope.users = data;
		$scope.load++;
	}).error(function(){
		$window.location.href = '../';
	});

	$http({method: 'GET', url: 'api/roles'}).success(function(data) {
		$scope.roles = data;
		$scope.load++;
	}).error(function(e) {
		$scope.errors = e;
	});

	$http({method: 'GET', url: 'api/user'}).success(function(data) {
		$scope.userCurrent = data;
		$scope.load++;
	}).error(function() {
		$window.location.href = '../';
	})

	$scope.newUser = function() {
		$scope.showCreateUser = true;
		$scope.showEditUser = false;
	};

	$scope.modifyUser = function(user) {
		$scope.userEdit = angular.copy(user);
		$scope.showCreateUser = false;
		$scope.showEditUser = true;;
	};

	$scope.createUser = function() {
		$http({
			method: 'POST',
			url: 'api/users',
			data: $scope.userNew,
		}).success(function(data){
			$scope.userNew.u_id = data;
			$scope.users.push($scope.userNew);
			$scope.showCreateUser = false;
			$scope.userNew = {u_role:1};
		}).error(function(e){
			$scope.error = e;
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
			$scope.userEdit = {};
			$scope.showEditUser = false;
		}).error(function(e){
			$scope.error = e;
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
			$scope.showEditUser = false;
		}).error(function(e){
			$scope.error = e;
		});
	};
}]);
