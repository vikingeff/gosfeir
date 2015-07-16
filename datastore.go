package main

import (
  "net/http"
	"appengine"
	"appengine/datastore"
)

func addData(d Data, r *http.Request) string {
	c := appengine.NewContext(r)
	key := datastore.NewIncompleteKey(c, "object", nil)
	a, _ := datastore.Put(c, key, &d)
	return (a.String())
}

func putData(d Data, r *http.Request, id int64) string {
	c := appengine.NewContext(r)
	key := datastore.NewKey(c, "object", "", id, nil)
	a, _ := datastore.Put(c, key, &d)
	return (a.String())
}

func getData(r *http.Request, id int64) ([]Data, []*datastore.Key, string) {
  var listKeys []Data
  var q *(datastore.Query)

  c := appengine.NewContext(r)
	if id == 0 {
		q = datastore.NewQuery("object")
	} else {
		key := datastore.NewKey(c, "object", "", id, nil)
		q = datastore.NewQuery("object").Filter("__key__ =", key)
	}
  if keys, err := q.GetAll(c, &listKeys); err != nil {
    return nil, nil, err.Error()
  } else {
    return listKeys, keys, ""
  }
}
