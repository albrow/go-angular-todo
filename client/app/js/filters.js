"use strict";

angular.module('todo.filters', []).filter('done', function() {
  return function(input) {
    return _.where(input, {done: true});
  };
}).filter('undone', function() {
	return function(input) {
		return _.where(input, {done: false});
	}
});