app.factory('PDV', ['$http', '$q', '$timeout', function($http, $q, $timeout) {

	var factory = {

		pdv: false,

		find: function(options) {
			var deferred = $q.defer();

			if (factory.pdvs) {
				deferred.resolve(factory.pdvs);	
			} else {
				$http.get('api/pdv').success(function(data, status) {
					factory.pdvs = data;
					deferred.resolve(factory.pdvs);
				}).error(function(data, status) {
					deferred.reject('Impossible de récupérer les pdvs');
				});
			}
			return deferred.promise;
		},

		get: function(id) {
			var pdv = {};
			var deferred = $q.defer();
			var pdvs = factory.find().then(function(pdvs) {
				angular.forEach(factory.pdvs, function(value, key) {
					if (value.id == id) {
						pdv = value;
					};
				});
				deferred.resolve(pdv);
			}, function(msg) {
				deferred.reject(msg);
			});

			return deferred.promise;
		},
	};

	return factory;
}]);