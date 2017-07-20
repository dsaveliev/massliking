package models

import (
	"strconv"
	"time"

	"massliking/backend/config"
	. "massliking/backend/errors"
	"massliking/backend/instabot"
)

var INSTAGRAM_REGISTRY map[int64]*WorkersPool = map[int64]*WorkersPool{}

const INSTAGRAM_STATE_STOP = "stop"
const INSTAGRAM_STATE_START = "start"
const INSTAGRAM_STATE_SUSPECTED = "suspected"

type WorkersPool struct {
	WriteOpCh          chan *InstagramWriteOp
	InstagramCommandCh chan string
	StatsCommandCh     chan string
	ChannelCommandChs  map[string]chan string
}

type InstagramReadOp struct {
	Error error
}

type InstagramWriteOp struct {
	Instagram   *Instagram
	Callback    func(i *Instagram)
	ReadChannel chan *InstagramReadOp
}

type Limits struct {
	Like     int
	Comment  int
	Follow   int
	Unfollow int
}

type Speed struct {
	Like     int `json:"like"`
	Comment  int `json:"comment"`
	Follow   int `json:"follow"`
	Unfollow int `json:"unfollow"`
}

type Counters struct {
	Like      int
	Comment   int
	Follow    int
	Unfollow  int
	StartedAt time.Time
}

type Hours struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

//TODO: Add FollowersCount field for summary over the channels
type Instagram struct {
	Id     int64 `json:"id" xorm:"pk autoincr index"`
	UserId int64 `json:"-" xorm:"index"`

	Info *instabot.UserInfo `json:"info" xorm:"jsonb"`

	State string `json:"state" xorm:"varchar(250) notnull default 'stop'"`

	Username string `json:"username" xorm:"notnull index"`
	Password string `json:"password" xorm:"varchar(250) notnull"`
	Trusted  bool   `json:"trusted" xorm:"bool notnull default false"`
	*Hours   `json:"hours" xorm:"jsonb"`
	*Speed   `json:"speed" xorm:"jsonb"`

	*Counters `json:"-" xorm:"jsonb"`

	Version   int       `json:"-" xorm:"version"`
	CreatedAt time.Time `json:"created_at" xorm:"created"`
	UpdatedAt time.Time `json:"updated_at" xorm:"updated"`
}

func (i *Instagram) IdString() string {
	return strconv.FormatInt(i.Id, 10)
}

func (i *Instagram) Save(callback func(i *Instagram)) error {
	// Check account in the registry
	if INSTAGRAM_REGISTRY[i.Id] == nil {
		return MODEL_INSTAGRAM_NOT_FOUND
	}

	// Create message for async processing
	writeOp := &InstagramWriteOp{
		Instagram:   i,
		Callback:    callback,
		ReadChannel: make(chan *InstagramReadOp),
	}

	// Send message to goroutine
	INSTAGRAM_REGISTRY[i.Id].WriteOpCh <- writeOp
	// Call SyncSave inside goroutine and read response
	readOp := <-writeOp.ReadChannel

	// Return response
	return readOp.Error
}

func (i *Instagram) SyncSave(callback func(i *Instagram)) error {
	var err error

	user, err := i.User()
	if err != nil {
		return err
	}

	instagram, err := GetInstagram(user, i.Id)
	if err != nil {
		return err
	}

	// Look ma, no hands!
	*i = *instagram
	callback(i)

	instagram, err = UpdateInstagramById(i)
	if err != nil {
		return err
	}

	return nil
}

func (i *Instagram) FindLimits() *Limits {
	key := ""

	if i.Trusted {
		key = "trusted"
	} else {
		key = "non_trusted"
	}

	return &Limits{
		Like:     config.GetInt("limits." + key + ".likes"),
		Comment:  config.GetInt("limits." + key + ".comments"),
		Follow:   config.GetInt("limits." + key + ".follows"),
		Unfollow: config.GetInt("limits." + key + ".unfollows"),
	}
}

func (i *Instagram) CheckLimits(action string) bool {
	var err error

	duration := time.Since(i.Counters.StartedAt)

	if duration >= time.Hour*time.Duration(24) {
		err = i.Save(func(instagram *Instagram) {
			i.Counters = &Counters{
				Like:      0,
				Comment:   0,
				Follow:    0,
				Unfollow:  0,
				StartedAt: time.Now(),
			}
		})
		if err != nil {
			return false
		}
		return true
	}

	limits := i.FindLimits()

	switch action {
	case "like":
		if i.Counters.Like >= limits.Like {
			return false
		}
	case "comment":
		if i.Counters.Comment >= limits.Comment {
			return false
		}
	case "follow":
		if i.Counters.Follow >= limits.Follow {
			return false
		}
	case "unfollow":
		if i.Counters.Unfollow >= limits.Unfollow {
			return false
		}
	default:
		return false
	}

	return true
}

func (i *Instagram) ActionSpeed(action string) int {
	var speed int

	switch action {
	case "like":
		speed = i.Speed.Like
	case "comment":
		speed = i.Speed.Comment
	case "follow":
		speed = i.Speed.Follow
	case "unfollow":
		speed = i.Speed.Unfollow
	default:
		speed = 1
	}

	if speed <= 0 {
		speed = 1
	}

	return speed
}

func (i *Instagram) UpdateLimits(action string) error {
	var err error

	err = i.Save(func(i *Instagram) {
		switch action {
		case "like":
			i.Counters.Like += 1
		case "comment":
			i.Counters.Comment += 1
		case "follow":
			i.Counters.Follow += 1
		case "unfollow":
			i.Counters.Unfollow += 1
		default:
			//TODO: Process undefined action
		}
	})
	if err != nil {
		return err
	}

	return nil
}

func (i *Instagram) UpdateInfo(client *instabot.Client) error {
	info, err := client.SearchUsername(i.Username)
	if err != nil {
		return On(err, MODEL_INSTAGRAM_INFO_ERROR)
	}

	err = i.Save(func(i *Instagram) {
		i.Info = info.UserInfo
	})
	if err != nil {
		return On(err, MODEL_INSTAGRAM_INFO_ERROR)
	}

	return nil
}

func (i *Instagram) User() (*User, error) {
	var err error
	user := &User{}

	_, err = Engine.Id(i.UserId).Get(user)

	return user, On(err, MODEL_USER_NOT_FOUND)
}

func (i *Instagram) ValidateState(client *instabot.Client) error {
	var err error

	user, err := i.User()
	if err != nil {
		return err
	}

	// Check account presence
	instagram, err := GetInstagram(user, i.Id)
	if err != nil {
		return On(err, MODEL_INSTAGRAM_NOT_FOUND)
	}

	// Look ma, no hands!
	*i = *instagram

	// Check client validity
	if client.Suspected {
		err = i.Save(func(i *Instagram) {
			i.State = INSTAGRAM_STATE_SUSPECTED
		})
		if err != nil {
			return On(err, MODEL_INSTAGRAM_NOT_UPDATED)
		}
	}

	// Check account state
	if i.State != INSTAGRAM_STATE_START {
		return MODEL_INSTAGRAM_INACTIVE_ERROR
	}

	return nil
}
