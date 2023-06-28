package gRPCHandler

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	"github.com/VlPakhomov/url_shortener/internal/config"
	"github.com/VlPakhomov/url_shortener/pkg/logger"
	"github.com/VlPakhomov/url_shortener/pkg/validator"
	api "github.com/VlPakhomov/url_shortener/proto"
)

type service interface {
	GetUrl(ctx context.Context, short string) (string, error)
	ShortenUrl(ctx context.Context, url string) (string, error)
}

type GrpcHandler struct {
	serv service
	api.UnimplementedGrpcHandlerServer
}

func (hl GrpcHandler) GetUrl(ctx context.Context, request *api.GetUrlRequest) (*api.GetUrlResponse, error) {

	p, _ := peer.FromContext(ctx)
	addr := p.Addr.String()

	logger.Debugf(ctx, "handle on %s | transportMode=gRPC", addr)

	rawShortUrl := request.RawShortUrl

	if !validator.IsShortUrl(ctx, string(config.Get(config.ShortUrlPattern)), rawShortUrl) {

		return nil, status.Error(codes.InvalidArgument, "| transportMode=gRPC")
	}

	rawUrl, err := hl.serv.GetUrl(ctx, rawShortUrl)
	if err != nil {

		//logger.Errorf(ctx, "bad response: internal server error %v | transportMode=gRPC", err)

		return nil, fmt.Errorf("%v | transportMode=gRPC", err)
	}

	logger.Info(ctx, "responded | transportMode=gRPC")

	return &api.GetUrlResponse{RawUrl: rawUrl}, nil

}

func (hl GrpcHandler) ShortenUrl(ctx context.Context, request *api.ShortenUrlRequest) (*api.ShortenUrlResponse, error) {

	p, _ := peer.FromContext(ctx)
	addr := p.Addr.String()
	logger.Debugf(ctx, "handle on %s | transportMode=gRPC", addr)

	rawUrl := request.RawUrl

	if !validator.IsUrl(ctx, rawUrl) {

		return nil, status.Error(codes.InvalidArgument, "| transportMode=gRPC")
	}

	rawShortUrl, err := hl.serv.ShortenUrl(ctx, rawUrl)
	if err != nil {

		//logger.Errorf(ctx, "bad response: internal server error %v | transportMode=gRPC", err)

		return nil, fmt.Errorf("%v | transportMode=gRPC", err)

	}

	logger.Info(ctx, "responded | transportMode=gRPC")

	return &api.ShortenUrlResponse{RawShortUrl: rawShortUrl}, nil

}

func NewHandler(s service) *GrpcHandler {

	return &GrpcHandler{
		serv: s,
	}
}
