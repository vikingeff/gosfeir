"use strict";

shareApp.controller("homeController" ,function ($scope) {

    $scope.user = '42';

});

shareApp.controller("objectsController" ,function ($scope, Obj) {

	$scope.tri = 'Title';

    Obj.fetch().success(function(resp){
        $scope.objects = resp;
        console.log($scope.objects);
    });


    $scope.deleteObject = function(index, id) {
        Obj.remove(id)
            .success(function(resp){
                $scope.objects.splice(index, 1);
            }
        );
    };

	$scope.trier = function (tri) {
		if ($scope.tri === tri) {
			$scope.reverse = !$scope.reverse
		}
		$scope.tri = tri;
	}

});

shareApp.controller('editObjectController', function($scope, Obj, $routeParams, $location){

    var objectId = $routeParams.id;

    Obj.fetchOne(objectId).success(function(object){
        $scope.obj = object[0];
        console.log($scope.obj);
    });

    $scope.updateObject = function(object){
        Obj.update(object, objectId)
            .success(function() {
                $location.path('/objects');
            })
            .error(function(resp){
                console.log(resp);
            });
    };
});

shareApp.controller("objectFormController" ,function ($scope, Obj, $location) {

    $scope.showAlert = false;

    $scope.addObject = function(object){
        Obj.create(object)
            .success(function(){
                var newobject = {};
                angular.copy(object, newobject)
                $scope.objects = [];
                $scope.objects.push(newobject);
                $scope.obj = {};
                $location.path('/objects');
            })
    };
});
