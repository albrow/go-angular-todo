'use strict';

/* Services */

angular.module('todo.services', ['ngResource']).
	factory('Item', function($resource) {		
		return $resource('http://127.0.0.1\\:6060/items/:itemId', {}, {
			update: {method: 'PUT'}
		});
	});