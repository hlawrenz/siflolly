'use strict';

/* Controllers */

angular.module('siflOlly.controllers', [])
  .controller('EchoCtrl', ['$scope', function($scope) {
      $scope.messages = [];
      $scope.echo = '';
      $scope._socket = new WebSocket("ws://localhost:8888/echosock");
      $scope.status = 'Not Connected';

      $scope._socket.onopen = function (event) {
          $scope.status = 'Connected';
		  $scope.$apply();
      };

      $scope._socket.onmessage = function (event) {
          $scope.messages.push(event.data);
		  $scope.$apply();
      };

      $scope.sendMessage = function () {
          $scope._socket.send($scope.echo);
      };
  }])
  .controller('NoiseCtrl', ['$scope', function($scope) {
      $scope.messages = [];
      $scope.echo = '';
      $scope._socket = new WebSocket("ws://localhost:8888/noise");
      $scope.status = 'Not Connected';

      $scope._socket.onopen = function (event) {
          $scope.status = 'Connected';
		  $scope.$apply();
      };

      $scope._socket.onmessage = function (event) {
          $scope.messages.unshift(event.data);
          $scope.echo = '';
		  $scope.$apply();
      };

      $scope.sendMessage = function () {
          $scope._socket.send($scope.echo);
      };

  }]);
