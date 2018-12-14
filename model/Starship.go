package model

type Starship struct {
	Name                   string   `json:"name"`
	Model                  string   `json:"model"`
	Manufacturer           string   `json:"manufacturer"`
	Cost_in_credits        string   `json:"cost_in_credits"`
	Length                 string   `json:"length"`
	Max_atmosphering_speed string   `json:"max_atmosphering_speed"`
	Crew                   string   `json:"crew"`
	Passengers             string   `json:"passengers"`
	Cargo_capacity         string   `json:"Cargo_capacity"`
	Consumables            string   `json:"consumables"`
	Hyperdrive_rating      string   `json:"hyperdrive_rating"`
	MGLT_                  string   `json:"MGLT"`
	Starship_class         string   `json:"starship_class"`
	Pilots                 []string `json:"pilots"`
	Films                  []string `json:"films"`
	Created                string   `json:"created"`
	Edited                 string   `json:"edited"`
	Url                    string   `json:"url"`
}
