package models

import (
	"time"

	. "massliking/backend/errors"
)

func CreateInstagram(user *User, c *InstagramCredentials) (*Instagram, error) {
	var err error
	instagram := &Instagram{}

	instagram.UserId = user.Id
	instagram.Username = c.Username
	instagram.Password = c.Password
	instagram.Trusted = c.Trusted
	instagram.State = INSTAGRAM_STATE_STOP
	instagram.Hours = &Hours{
		Max: c.Hours.Max,
		Min: c.Hours.Min,
	}
	instagram.Counters = &Counters{
		Like:      0,
		Comment:   0,
		Follow:    0,
		Unfollow:  0,
		StartedAt: time.Now(),
	}
	instagram.Speed = &Speed{
		Like:     c.Speed.Like,
		Comment:  c.Speed.Comment,
		Follow:   c.Speed.Follow,
		Unfollow: c.Speed.Unfollow,
	}

	_, err = Engine.Insert(instagram)

	return instagram, On(err, MODEL_INSTAGRAM_NOT_CREATED)
}

func GetInstagram(user *User, id int64) (*Instagram, error) {
	var err error
	instagram := &Instagram{}

	has, err := Engine.
		Id(id).
		And("user_id = ?", user.Id).
		Get(instagram)

	if err != nil || has != true {
		return instagram, MODEL_INSTAGRAM_NOT_FOUND
	}

	return instagram, nil
}

func FindInstagrams(user *User) ([]*Instagram, error) {
	var err error
	instagrams := []*Instagram{}

	err = Engine.
		Where("user_id = ?", user.Id).
		Find(&instagrams)

	return instagrams, On(err, MODEL_INSTAGRAM_COLLECTION_NOT_FOUND)
}

func FindAllInstagrams() ([]*Instagram, error) {
	var err error
	instagrams := []*Instagram{}

	err = Engine.Find(&instagrams)

	return instagrams, On(err, MODEL_INSTAGRAM_COLLECTION_NOT_FOUND)
}

func UpdateInstagramById(instagram *Instagram) (*Instagram, error) {
	var err error

	_, err = Engine.Id(instagram.Id).UseBool("trusted").Update(instagram)
	if err != nil {
		return instagram, On(err, MODEL_INSTAGRAM_NOT_UPDATED)
	}

	return instagram, nil
}

func UpdateInstagram(instagram *Instagram, c *InstagramCredentials) (*Instagram, error) {
	var err error

	err = instagram.Save(func(instagram *Instagram) {
		instagram.Username = c.Username
		instagram.Password = c.Password
		instagram.Trusted = c.Trusted
		instagram.Hours = &Hours{
			Max: c.Hours.Max,
			Min: c.Hours.Min,
		}
		instagram.Speed = &Speed{
			Like:     c.Speed.Like,
			Comment:  c.Speed.Comment,
			Follow:   c.Speed.Follow,
			Unfollow: c.Speed.Unfollow,
		}
	})

	return instagram, err
}

func DeleteInstagram(instagram *Instagram) error {
	var err error

	channels, err := instagram.FindChannels()
	if err != nil {
		return err
	}

	for _, c := range channels {
		err = instagram.DeleteChannel(c)
		if err != nil {
			return err
		}
	}

	_, err = Engine.
		Id(instagram.Id).
		Delete(&Instagram{})

	return On(err, MODEL_INSTAGRAM_NOT_DELETED)
}
