package models

func (u *User) CreateInstagram(c *InstagramCredentials) (*Instagram, error) {
	return CreateInstagram(u, c)
}

func (u *User) GetInstagram(id int64) (*Instagram, error) {
	return GetInstagram(u, id)
}

func (u *User) FindInstagrams() ([]*Instagram, error) {
	return FindInstagrams(u)
}

func (u *User) UpdateInstagram(instagram *Instagram, c *InstagramCredentials) (*Instagram, error) {
	return UpdateInstagram(instagram, c)
}

func (u *User) DeleteInstagram(instagram *Instagram) error {
	return DeleteInstagram(instagram)
}
