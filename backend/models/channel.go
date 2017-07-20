package models

import (
	"strconv"
	"time"

	. "massliking/backend/errors"
)

const CHANNEL_STATE_STOP = "stop"
const CHANNEL_STATE_START = "start"
const CHANNEL_STATE_EMPTY = "empty"

var ACTIONS = [4]string{
	"like",
	"comment",
	"follow",
	"unfollow",
}

var TARGETS = [5]string{
	"followers",
	"subscriptions",
	"hashtag",
	"likes",
	"comments",
}

type Queue struct {
	MaxId   string
	Leads   []int
	Targets []int
}

type Channel struct {
	Id             int64     `json:"id"              xorm:"pk autoincr index"`
	InstagramId    int64     `json:"-"               xorm:"index"`
	Action         string    `json:"action"          xorm:"varchar(250) notnull"`
	Target         string    `json:"target"          xorm:"varchar(250) notnull"`
	Value          string    `json:"value"           xorm:"varchar(250) notnull"`
	Comments       []string  `json:"comments"        xorm:"jsonb"`
	State          string    `json:"state"           xorm:"varchar(250) notnull default 'stop'"`
	LeadsCount     int       `json:"leads_count"     xorm:"integer"`
	TargetsCount   int       `json:"targets_count"   xorm:"integer"`
	FollowersCount int       `json:"followers_count" xorm:"integer default 0"`
	Version        int       `json:"-"               xorm:"version"`
	CreatedAt      time.Time `json:"created_at"      xorm:"created"`
	UpdatedAt      time.Time `json:"updated_at"      xorm:"updated"`
	*Queue         `json:"-" xorm:"jsonb"`
}

func (c *Channel) IdString() string {
	return strconv.FormatInt(c.Id, 10)
}

func (c *Channel) Leads() map[int]bool {
	leads := map[int]bool{}

	for _, l := range c.Queue.Leads {
		leads[l] = true
	}

	return leads
}

func (c *Channel) Save(update func(c *Channel)) error {
	var err error

	channel := Channel{}
	_, err = Engine.Id(c.Id).Get(&channel)
	if err != nil {
		return On(err, MODEL_CHANNEL_NOT_UPDATED)
	}

	*c = channel
	update(c)

	c.LeadsCount = len(c.Queue.Leads)
	c.TargetsCount = len(c.Queue.Targets)

	_, err = Engine.Id(c.Id).Update(c)
	if err != nil {
		return On(err, MODEL_CHANNEL_NOT_UPDATED)
	}

	return nil
}

func (c *Channel) Instagram() (*Instagram, error) {
	var err error

	instagram := &Instagram{}

	_, err = Engine.Id(c.InstagramId).Get(instagram)

	return instagram, On(err, MODEL_INSTAGRAM_NOT_FOUND)
}
