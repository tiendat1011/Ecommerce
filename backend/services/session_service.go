package services

import (
	"ecommerce-project/daos"
	"time"
)

type SessionService struct {
	sessionDAO daos.SessionDAO
}

func NewSessionService(sessionDAO daos.SessionDAO) *SessionService {
	return &SessionService{
		sessionDAO: sessionDAO,
	}
}

func (s *SessionService) SaveSession(token, userID string) error {
	ttl := 24 * time.Hour
	return s.sessionDAO.SaveSession(token, userID, ttl)
}

func (s *SessionService) ValidateSession(token string) (string, error) {
	return s.sessionDAO.GetSession(token)
}

func (s *SessionService) DeleteSession(token string) error {
	return s.sessionDAO.DeleteSession(token)
}