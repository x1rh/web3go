package tonx

import (
	"context"

	"github.com/pkg/errors"
	"github.com/tonkeeper/tongo/liteapi"
	"github.com/tonkeeper/tongo/ton"
	"github.com/tonkeeper/tongo/tonconnect"
)

type ServerConfig struct {
	Secret          string `json:"Secret"`          //  "klafdsldssskdsssdd"
	LifeTimePayload int64  `json:"LifeTimePayload"` // 3000
	LifeTimeProof   int64  `json:"LifeTimeProof"`   // 3000
}

type Server struct {
	server *tonconnect.Server
}

type Params struct {
	Timestamp         int64
	DomainLengthBytes int64
	Address           string
	DomainValue       string
	Signature         string
	Payload           string
	StateInit         string
}

func NewServer(c ServerConfig) (*Server, error) {
	tonClient, err := liteapi.NewClientWithDefaultMainnet()
	if err != nil {
		return nil, errors.Wrap(err, "fail to new a ton client")
	}
	srv, err := tonconnect.NewTonConnect(tonClient, c.Secret, tonconnect.WithLifeTimePayload(c.LifeTimePayload), tonconnect.WithLifeTimeProof(c.LifeTimeProof))
	if err != nil {
		return nil, errors.Wrap(err, "fail to new a tonconnect")
	}
	return &Server{
		server: srv,
	}, nil
}

func (s *Server) VerifySignature(ctx context.Context, p *Params) (bool, error) {
	proof := &tonconnect.Proof{
		Address: p.Address,
		Proof: tonconnect.ProofData{
			Timestamp: p.Timestamp,
			Signature: p.Signature,
			Domain:    p.DomainValue,
			Payload:   p.Payload,
			StateInit: p.StateInit,
		},
	}
	ok, _, err := s.server.CheckProof(ctx, proof, s.server.CheckPayload, tonconnect.StaticDomain(p.DomainValue))
	if err != nil {
		return false, errors.Wrap(err, "fail to verify signature")
	}
	return ok, nil
}

func (s *Server) GeneratePayload() (string, error) {
	return s.server.GeneratePayload()
}

// GetNonBounceableAddressByHex
// example: 根据 HEX：0:2c477e26fd20dde1f6e4faf242c1458d2b5ec056e 格式地址获取 UQAsR34m_SDd4fbk-vJCwUW 格式地址
func (s *Server) GetNonBounceableAddressByHex(hex string) (string, error) {
	accountId, err := ton.AccountIDFromRaw(hex)
	if err != nil {
		return "", errors.Wrap(err, "fail to get ton address account id")
	}
	return accountId.ToHuman(false, false), nil
}
