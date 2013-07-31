'use strict';

/* Controllers */

angular.module('todo.controllers', []).
	controller('mainCtrl', ['$scope', '$http', 'Restangular', function($scope, $http, Restangular) {
		delete $http.defaults.headers.common['X-Requested-With'];

		var baseItems = Restangular.all('items');
		baseItems.getList().then(function(items) {
			$scope.items = items;
		});

		$scope.addItem = function(item) {
			if (item) {
				baseItems.post(item).then(function(item) {
					$scope.items.push(item);
					$("#item-content").val("");
				});
			}
		}

		$scope.removeItem = function(item) {
			if (item) {
				item.remove();
			}
		}

	}]
);