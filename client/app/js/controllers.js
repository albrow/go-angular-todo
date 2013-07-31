'use strict';

/* Controllers */

angular.module('todo.controllers', []).
	controller('mainCtrl', ['$scope', '$http', 'Restangular', function($scope, $http, Restangular) {
		delete $http.defaults.headers.common['X-Requested-With'];
		$scope.currentNewItem = null;

		var baseItems = Restangular.all('items');
		baseItems.getList().then(function(items) {
			$scope.items = items;
		});

		$scope.addItem = function(item) {
			if (item) {
				baseItems.post(item).then(function(item) {
					$scope.currentNewItem = item;
					$scope.items.unshift(item);
					$("#item-content-input").val("");
					setTimeout(function() {
						$scope.currentNewItem = null;
						$scope.$digest();
					}, 10);
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