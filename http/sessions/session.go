package sessions

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis/v3"
	"github.com/goquiz/api/helpers"
	"time"
)

type Sess struct {
	Session *session.Session
}

var Global Sess

func New(c *fiber.Ctx, global bool) (*Sess, error) {
	var s *Sess
	if global {
		s = &Global
	}
	storage := redis.New(redis.Config{
		Host:     helpers.Env.Redis.Host,
		Port:     helpers.Env.Redis.Port,
		Username: helpers.Env.Redis.Username,
		Password: helpers.Env.Redis.Password,
	})
	store := session.New(session.Config{
		Expiration:     time.Duration(helpers.Env.Session.ExpiresIn) * time.Minute,
		KeyLookup:      helpers.Env.Session.KeyLookup,
		CookieSecure:   helpers.Env.Session.CookieSecure,
		CookieHTTPOnly: helpers.Env.Session.CookieHttpOnly,
		CookieSameSite: helpers.Env.Session.CookieSameSite,
		Storage:        storage,
		CookieDomain:   helpers.Env.Session.CookieDomain,
	})
	sess, err := store.Get(c)
	if err != nil {
		s.Session = nil
		return s, err
	}
	s.Session = sess
	return s, nil
}

func (s *Sess) Get(key string) interface{} {
	return s.Session.Get(key)
}

func (s *Sess) Set(key string, value interface{}) {
	s.Session.Set(key, value)
}

func (s *Sess) Delete(key string) {
	s.Session.Delete(key)
}

func (s *Sess) Destroy() error {
	err := s.Session.Destroy()
	if err != nil {
		return err
	}
	s.Session = nil
	return nil
}

// NewGlobalSessionHandler a global middleware to enable sessions in the application
func NewGlobalSessionHandler(c *fiber.Ctx) error {
	_, _ = New(c, true)
	return c.Next()
}

func (s *Sess) Save() {
	err := s.Session.Save()
	if err != nil {
		fmt.Println("Failed to save session", err)
	}
}

func (s *Sess) Id() string {
	return s.Session.ID()
}
