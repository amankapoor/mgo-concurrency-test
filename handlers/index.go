package handlers

import (
	"html/template"
	"net/http"
	"runtime"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"fmt"
)

func IndexHandler(sessionFromMain *mgo.Session) http.Handler {
	fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("started index handler")

		//Tweak this upper limit
		//Note: No idea why, but none of the given three runtime.NumGoroutine() executes
		for i := 0; i < 1000; i++ {
			go openNewCloneSession(sessionFromMain)
			time.Sleep(1 * time.Nanosecond)
			fmt.Println("Entered Clone Record # ", i)
			fmt.Println("No. of Goroutines (CLONE) are: ", runtime.NumGoroutine())

			go openNewCopySession(sessionFromMain)
			time.Sleep(1 * time.Nanosecond)
			fmt.Println("Entered Copy Record # ", i)
			fmt.Println("No. of Goroutines (COPY) are: ", runtime.NumGoroutine())

			go openNewSession(sessionFromMain)
			time.Sleep(1 * time.Nanosecond)
			fmt.Println("Entered New Record # ", i)
			fmt.Println("No. of Goroutines (NEW) are: ", runtime.NumGoroutine())

		}
	})
	return fn
}

func openNewSession(session *mgo.Session) {
	title := "New"
	slug := "New"
	body := "New"

	id := bson.NewObjectId()
	//Inserting
	p := Posts{
		ID:    id,
		Title: title,
		Slug:  slug,
		Body:  template.HTML(body),
	}

	//We don't care about reading our writes here, so copy
	sess := session.New()
	defer sess.Close()
	collection := sess.DB("testdb").C("posts")

	insertionError := collection.Insert(p)
	if insertionError != nil {
		panic(insertionError)
	}

	fmt.Println(id.Hex())
	sess.Refresh()
}

func openNewCopySession(session *mgo.Session) {
	title := "Copy"
	slug := "Copy"
	body := "Copy"

	id := bson.NewObjectId()
	//Inserting
	p := Posts{
		ID:    id,
		Title: title,
		Slug:  slug,
		Body:  template.HTML(body),
	}

	//We don't care about reading our writes here, so copy
	sess := session.Copy()
	defer sess.Close()
	collection := sess.DB("testdb").C("posts")

	insertionError := collection.Insert(p)
	if insertionError != nil {
		panic(insertionError)
	}

	fmt.Println(id.Hex())
	sess.Refresh()
}

func openNewCloneSession(session *mgo.Session) {
	title := "Clone"
	slug := "Clone"
	body := "Clone"

	id := bson.NewObjectId()
	//Inserting
	p := Posts{
		ID:    id,
		Title: title,
		Slug:  slug,
		Body:  template.HTML(body),
	}

	//We don't care about reading our writes here, so copy
	sess := session.Clone()
	defer sess.Close()
	collection := sess.DB("testdb").C("posts")

	insertionError := collection.Insert(p)
	if insertionError != nil {
		panic(insertionError)
	}

	fmt.Println(id.Hex())
	sess.Refresh()
}
