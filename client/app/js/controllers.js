'use strict';

/* Controllers */

angular.module('todo.controllers', []).
	controller('mainCtrl', ['$scope', '$http', 'Item', function($scope, $http, Item) {
		delete $http.defaults.headers.common['X-Requested-With'];
		$scope.items = Item.query();

	}]
);