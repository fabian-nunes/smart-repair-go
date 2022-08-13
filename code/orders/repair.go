package orders

import "time"

type Repair struct {
	code        int
	name        string
	date        string
	cost        float64
	description string
	client      string
	status      int
}

func newRepair(code int, name string, cost float64, description string, client string) Repair {
	return Repair{
		code:        code,
		name:        name,
		date:        time.Now().Format("2006-01-02"),
		cost:        cost,
		description: description,
		client:      client,
		status:      0,
	}
}

func findRepair(code int) bool {
	for _, repair := range repairs {
		if repair.code == code {
			return true
		}
	}
	return false
}
