'use strict';


// Declare app level module which depends on filters, and services
angular.module('siflOlly', [
  'ngRoute',
  'siflOlly.filters',
  'siflOlly.services',
  'siflOlly.directives',
  'siflOlly.controllers'
]).
config(['$routeProvider', function($routeProvider) {
  $routeProvider.when('/echo', {templateUrl: 'partials/echo.html', controller: 'EchoCtrl'});
  $routeProvider.when('/noise', {templateUrl: 'partials/noise.html', controller: 'NoiseCtrl'});
  $routeProvider.otherwise({redirectTo: '/echo'});
}]);
