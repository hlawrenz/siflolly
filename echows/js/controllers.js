'use strict';

/* Controllers */

angular.module('siflOlly.controllers', [])
  .controller('EchoCtrl', ['$scope', function($scope) {
      $scope.messages = [];
      $scope.echo = '';
      var sockHost = location.hostname+(location.port ? ':'+location.port: '')
      $scope._socket = new WebSocket("ws://"+sockHost+"/echosock");
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
      var sockHost = location.hostname+(location.port ? ':'+location.port: '')
      $scope._socket = new WebSocket("ws://"+sockHost+"/noise");
      $scope.status = 'Not Connected';

      $scope._socket.onopen = function (event) {
          $scope.status = 'Connected';
		  $scope.$apply();
      };

      $scope._socket.onmessage = function (event) {
          $scope.messages.unshift(event.data);
		  $scope.$apply();
      };

      $scope.sendMessage = function () {
          $scope._socket.send($scope.echo);
          $scope.echo = '';
      };

  }]);
