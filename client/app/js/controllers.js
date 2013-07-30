'use strict';

/* Controllers */

angular.module('todo.controllers', []).
	controller('mainCtrl', ['$scope', '$http', 'Restangular', function($scope, $http, Restangular) {
		delete $http.defaults.headers.common['X-Requested-With'];
		
		var baseItems = Restangular.all('items');
		$scope.items = baseItems.getList();

		$scope.addItem = function(item) {
			if (item) {
				$scope.items.push(angular.copy(item));
				baseItems.post(item);
				$("#item-content").val("");
			}
		}

	}]
);