package models

import (
	. "massliking/backend/errors"
)

func CreateChannel(instagram *Instagram, c *ChannelCredentials) (*Channel, error) {
	var err error

	channel := &Channel{
		InstagramId:  instagram.Id,
		Action:       c.Action,
		Target:       c.Target,
		Value:        c.Value,
		Comments:     c.Comments,
		State:        CHANNEL_STATE_STOP,
		LeadsCount:   0,
		TargetsCount: 0,
		Queue: &Queue{
			MaxId:   "",
			Leads:   []int{},
			Targets: []int{},
		},
	}

	_, err = Engine.Insert(channel)

	return channel, On(err, MODEL_CHANNEL_NOT_CREATED)
}

func GetChannel(instagram *Instagram, id int64) (*Channel, error) {
	var err error
	channel := &Channel{}

	has, err := Engine.
		Id(id).
		And("instagram_id = ?", instagram.Id).
		Get(channel)

	if err != nil || has != true {
		return channel, On(err, MODEL_CHANNEL_NOT_FOUND)
	}

	return channel, nil
}

func FindChannels(instagram *Instagram) ([]*Channel, error) {
	var err error
	channels := []*Channel{}

	err = Engine.
		Where("instagram_id = ?", instagram.Id).
		Find(&channels)

	return channels, On(err, MODEL_CHANNEL_COLLECTION_NOT_FOUND)
}

func FindActiveChannels(instagram *Instagram, action string) ([]*Channel, error) {
	var err error
	channels := []*Channel{}

	err = Engine.
		Sql("SELECT * FROM channel WHERE instagram_id = ? AND action = ? AND ((state = ?) OR (state = ? AND targets_count > 0))",
			instagram.Id,
			action,
			CHANNEL_STATE_START,
			CHANNEL_STATE_EMPTY).
		Find(&channels)

	return channels, On(err, MODEL_CHANNEL_COLLECTION_NOT_FOUND)
}

func UpdateChannel(channel *Channel, c *ChannelCredentials) (*Channel, error) {
	var err error

	err = channel.Save(func(channel *Channel) {
		channel.Action = c.Action
		channel.Target = c.Target
		channel.Value = c.Value
		channel.Comments = c.Comments
	})

	return channel, err
}

func DeleteChannel(channel *Channel) error {
	var err error

	_, err = Engine.
		Id(channel.Id).
		Delete(&Channel{})

	return On(err, MODEL_CHANNEL_NOT_DELETED)
}
