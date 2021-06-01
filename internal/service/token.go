package service

import (
	"context"
	"fmt"
	"net/smtp"
	"time"

	"github.com/nguli-team/bakalo/internal/config"
	"github.com/nguli-team/bakalo/internal/domain"
	"github.com/nguli-team/bakalo/internal/util"
)

type tokenService struct {
	tokenRepository domain.TokenRepository
	smtpConfig      config.SMTPConfig
}

func NewTokenService(tokenRepository domain.TokenRepository, smtpConfig config.SMTPConfig) domain.TokenService {
	return &tokenService{
		tokenRepository: tokenRepository,
		smtpConfig:      smtpConfig,
	}
}

func (s tokenService) CheckVIPStatus(ctx context.Context, ip string) bool {
	serverToken, err := s.tokenRepository.FindByIP(ctx, ip)
	if err != nil {
		return false
	}
	return int64(serverToken.ValidUntil) > time.Now().Unix() && serverToken.IsValid
}

func (s tokenService) ValidateToken(ctx context.Context, token string, pin int) bool {
	serverToken, err := s.tokenRepository.FindByToken(ctx, token)
	if err != nil || pin != serverToken.PIN {
		return false
	}
	return int64(serverToken.ValidUntil) > time.Now().Unix() && serverToken.IsValid
}

func (s tokenService) CreateNewToken(ctx context.Context, ip string, email string) (*domain.VipToken, error) {
	newToken := &domain.VipToken{
		IP:         ip,
		Token:      util.RandomAlphaNumString(16),
		PIN:        util.RandomIntLength(6),
		Email:      email,
		ValidUntil: uint32(time.Now().AddDate(1, 0, 0).Unix()),
	}

	token, err := s.tokenRepository.Create(ctx, newToken)
	if err != nil {
		return nil, err
	}

	err = s.SendTokenViaEmail(ctx, token)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s tokenService) UpdateTokenIP(ctx context.Context, ip string, token string) (*domain.VipToken, error) {
	vipToken, err := s.tokenRepository.FindByToken(ctx, token)
	if err != nil {
		return nil, err
	}
	vipToken.IP = ip

	newToken, err := s.tokenRepository.Update(ctx, vipToken)

	if err != nil {
		return nil, err
	}

	return newToken, nil
}

func (s tokenService) SendTokenViaEmail(ctx context.Context, token *domain.VipToken) error {
	auth := smtp.PlainAuth("", s.smtpConfig.Email, s.smtpConfig.Password, s.smtpConfig.Host)
	smtpAddr := fmt.Sprintf("%s:%d", s.smtpConfig.Host, s.smtpConfig.Port)

	body := fmt.Sprintf(
		"From: %s <%s>\r\n"+
			"To: %s\r\n"+
			"Subject: Pembelian VIP Token Bakalo.li\r\n\r\n"+
			"Terimakasih atas pembeliannya! \r\n"+
			"VIP Token kamu adalah: %s\r\n"+
			"PIN kamu adalah: %d\r\n",
		s.smtpConfig.SenderName,
		s.smtpConfig.Email,
		token.Email,
		token.Token,
		token.PIN,
	)

	err := smtp.SendMail(smtpAddr, auth, s.smtpConfig.Email, []string{token.Email}, []byte(body))
	if err != nil {
		return err
	}

	return nil
}
