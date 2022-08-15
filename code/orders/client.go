package orders

type Client struct {
	name    string
	email   string
	phone   string
	repairs int
}

func newClient(name string, email string, phone string) Client {
	return Client{
		name:    name,
		email:   email,
		phone:   phone,
		repairs: 0,
	}
}
