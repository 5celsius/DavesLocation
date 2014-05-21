package daveloc

import (
	"fmt"
	"net/http"
  // "encoding/json"
  "time"
  // "log"
)

type DaveCoord struct {
    lat     float32
    lon     float32
}

type DaveLoc struct {
  coord   DaveCoord
	heading int16
}

func init() {
  http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request)  {
  w.Header().Set("Content-Type", "application/json")
  loc := getDaveLoc()
  fmt.Fprintf(w, "{\"lat\": %f, \"lon\" : %f,\"heading\": %d}", loc.coord.lat, loc.coord.lon, loc.heading)
}

func getCurrentCoord(start DaveCoord, finish DaveCoord, percent float32) DaveCoord {
  lat := ((finish.lat - start.lat) * percent) + start.lat
  lon := ((finish.lon - start.lon) * percent) + start.lon
  return DaveCoord{lat, lon}
}

func getDaveLoc() *DaveLoc {
  topRight     := DaveCoord{47.631568, -122.132382}
  topLeft      := DaveCoord{47.631503, -122.143143}
  bottomLeft   := DaveCoord{47.627852, -122.143068}
  bottomRight  := DaveCoord{47.627801, -122.132479}

  currentTime      := time.Now()
  currentSec       := currentTime.Second()
  currentPos       := currentSec / 15
  percentComplete  := float32(currentSec) / 60.0
  var returnCoord DaveCoord;
  var heading int16;

  switch currentPos {
    case 0:
      returnCoord = getCurrentCoord(topRight, topLeft, percentComplete)
      heading = 270
    case 1:
      returnCoord = getCurrentCoord(topLeft, bottomLeft, percentComplete)
      heading = 180
    case 2:
      returnCoord = getCurrentCoord(bottomLeft, bottomRight, percentComplete)
      heading = 90
    case 3:
      returnCoord = getCurrentCoord(bottomRight, topRight, percentComplete)
      heading = 0
  }
  return &DaveLoc{ returnCoord, heading }
}
