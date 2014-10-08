app.controller('UsersCtrl', ['$scope', '$http', '$window', function($scope, $http, $window) {
  $scope.userEdit = false;
  $scope.userNew = false;

  $http({method: 'GET', url: 'api/users'}).success(function(data){
    $scope.users = data;
  }).error(function(){
    alert("error");
    //$window.location.href = '../';
  });

  $scope.modifyUser = function(user) {
    $scope.userEdit = angular.copy(user);
  };

  $scope.newUser = function() {
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
      var i = 0;

      for (;i<$scope.users.length;i++) {
	if ($scope.users[i].u_id == $scope.userEdit.u_id) {
	  break;
	}
      }
      $scope.users.splice(i, 1);
      $scope.userEdit = false;
    }).error(function(){
      alert('error');
    });
  };

}]);

