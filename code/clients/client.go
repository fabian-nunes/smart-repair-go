package clients

type Client struct {
	name  string
	email string
	phone string
}

func newClient(name string, email string, phone string) Client {
	return Client{
		name:  name,
		email: email,
		phone: phone,
	}
}
