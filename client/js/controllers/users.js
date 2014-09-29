app.controller('UsersCtrl', ['$scope', '$http', '$window', function($scope, $http, $window) {
  $scope.userEdit = false;

  $http({method: 'GET', url: 'api/users'}).success(function(data){
    $scope.users = data;
  }).error(function(){
    $window.location.href = '../';
  });

  $scope.modifyUser = function(user) {
    $scope.userEdit = angular.copy(user)
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
        }
      });
      $scope.userEdit = false;
    }).error(function(){
      alert('error');
    });
  }

}]);

