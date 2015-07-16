"use strict";

shareApp.factory("Obj", function ($http) {
    var API_URI = '/api/objects/';

    return {

        fetch : function() {
            return $http.get(API_URI);
        },

        create : function(object) {
            return  $http.post(API_URI, object);
        },

        remove  : function(id) {
            return $http.delete(API_URI + id);
        },

        fetchOne : function(id) {
            return $http.get(API_URI + id);
        },

        update : function(object, id) {
             return $http.put(API_URI + id, object);
        }

    };

});
