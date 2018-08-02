package server

import (
	"path"
	"net/http"
	"net/url"
	"encoding/json"
	"gopkg.in/mgo.v2"
//	"gopkg.in/mgo.v2/bson"
	//"github.com/unrolled/render"
	"github.com/markllama/hexgame/types/db"
	"github.com/markllama/hexgame/types/api"
)

func GameListHandleFunc(s *mgo.Session) (func(http.ResponseWriter, *http.Request)) {
	
	f := func(w http.ResponseWriter, r *http.Request) {
		//var g db.Game

		sc := s.Copy()
		defer sc.Close()
		c := sc.DB("hexgame").C("games")
		
		w.Header().Add("Content-Type", "application/json")

		//_, name := path.Split(r.URL.Path)

		var hg []db.Game
		
		c.Find(nil).All(&hg)
	
		gamerefs := make([]api.GameRef, len(hg))

		gurl := url.URL{Scheme: "http", Host: r.Host}
		
		for index, game := range hg {
			gurl.Path = path.Join(r.URL.Path, game.Name)
			gamerefs[index].Name = game.Name
			gamerefs[index].URL = gurl.String()
		}
		
		jgames, _ := json.Marshal(gamerefs)
		//			jgames, _ := json.Marshal(hg)
		
		w.Write([]byte(jgames))
	}

	return f
}