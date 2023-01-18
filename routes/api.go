package routes

func (client *Client) Hosts() []string {
	return client.hosts
}

func (client *Client) SetHosts(hosts ...string) *Client {
	client.hosts = hosts
	return client
}

func (client *Client) Pattern() string {
	return client.pattern
}

func (client *Client) SetContentType(contentType string) *Client {
	client.contentType = contentType
	return client
}

func (client *Client) ContentType() string {
	return client.contentType
}
