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

  }])
  .controller('CacophonyCtrl', ['$scope', function($scope) {
      $scope.messages = [];
      $scope.echo = '';
      $scope.nick = '';
      $scope.error = '';
      var sockHost = location.hostname+(location.port ? ':'+location.port: '')
      $scope._socket = new WebSocket("ws://"+sockHost+"/cacophony");
      $scope.status = 'Not Connected';

      $scope._socket.onopen = function (event) {
          $scope.status = 'Connected';
          $scope.$apply();
      };

      $scope._socket.onmessage = function (event) {
          console.log(event.data)
          var message = JSON.parse(event.data);
          if (message.Op == "say") {

              message.Payload.forEach(function(e) {
                  $scope.messages.unshift({
                      from: message.From,
                      msg: e
                  });
              });
          } else if (message.Op == "nick") {
              $scope.nick = message.To;
          } else if (message.Op == "error") {
              $scope.error = message.Payload[0];
          }
          
          $scope.$apply();
      };

      $scope.setNick = function () {
          var message = {
              Op: "nick",
              Payload: [$scope.nick]
          }
          $scope._socket.send(JSON.stringify(message));
          $scope.echo = '';
      };

      $scope.sendMessage = function () {
          var message = {
              Op: "say",
              Payload: [$scope.echo]
          }
          $scope._socket.send(JSON.stringify(message));
          $scope.echo = '';
      };

  }]);
