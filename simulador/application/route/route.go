package route

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Route struct {
	ID        string     `json:"routeID"`
	ClientID  string     `json:"clientID"`
	Positions []Position `json:"position"`
}

type Position struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type PartialRoutePosition struct {
	ID       string    `json:"routeID"`
	ClientID string    `json:"clientID"`
	Position []float64 `json:"position"`
	Finished bool      `json:"finished"`
}

// Método para criar uma nova rota
func (r *Route) LoadPositions() error {
	// Carrega as posições da rota
	if r.ID == "" {
		return errors.New("Route ID is required")
	}

	f, err := os.Open("destinations/" + r.ID + ".txt")
	if err != nil {
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		lat, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			return err
		}
		long, err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			return err
		}

		r.Positions = append(r.Positions, Position{
			Lat:  lat,
			Long: long,
		})

	}

	return nil
}

// Add Json em uma lista de String
// ExportJsonPositions generates a slice of string in Json using PartialRoutePosition struct
func (r *Route) ExportJsonPositions() ([]string, error) {
	var route PartialRoutePosition
	var result []string
	total := len(r.Positions)
	for k, v := range r.Positions {
		route.ID = r.ID
		route.ClientID = r.ClientID
		route.Position = []float64{v.Lat, v.Long}
		route.Finished = false
		if total-1 == k {
			route.Finished = true
		}
		jsonRoute, err := json.Marshal(route)
		if err != nil {
			return nil, err
		}
		result = append(result, string(jsonRoute))
	}
	return result, nil
}
