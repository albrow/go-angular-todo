'use strict';


// Declare app level module which depends on filters, and services
angular.module('todo', ['todo.controllers', 'todo.filters','restangular']).
  config(['$routeProvider', 'RestangularProvider', function($routeProvider, RestangularProvider) {
	$routeProvider.when('/', {templateUrl: 'partials/main.html', controller: 'mainCtrl'});
    $routeProvider.otherwise({redirectTo: '/'});
    RestangularProvider.setBaseUrl("http://localhost:6060");
  }]);