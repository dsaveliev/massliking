package models

func (i *Instagram) CreateChannel(c *ChannelCredentials) (*Channel, error) {
	return CreateChannel(i, c)
}

func (i *Instagram) GetChannel(id int64) (*Channel, error) {
	return GetChannel(i, id)
}

func (i *Instagram) FindChannels() ([]*Channel, error) {
	return FindChannels(i)
}

func (i *Instagram) FindActiveChannels(action string) ([]*Channel, error) {
	return FindActiveChannels(i, action)
}

func (i *Instagram) UpdateChannel(channel *Channel, c *ChannelCredentials) (*Channel, error) {
	return UpdateChannel(channel, c)
}

func (i *Instagram) DeleteChannel(channel *Channel) error {
	return DeleteChannel(channel)
}
