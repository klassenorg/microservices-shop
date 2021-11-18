package grpc

import (
	pb "cartsvc/gen/proto"
	"cartsvc/internal/service"
	mock_service "cartsvc/internal/service/mocks"
	log "cartsvc/pkg/logger"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"net"
	"strconv"
	"testing"
)

func dialer(h *Handler) func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	pb.RegisterCartServiceServer(server, h)

	go func() {
		if err := server.Serve(listener); err != nil {
			panic(err)
		}
	}()

	return func(ctx context.Context, s string) (net.Conn, error) {
		return listener.Dial()
	}
}

const (
	userID    = "1234-test"
	productID = "123456"
	count     = int64(10)
)

var (
	countString = strconv.FormatInt(count, 10)
	ctx         = context.Background()

	testError = errors.New("test error")
)

func TestHandler_GetCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cart := mock_service.NewMockCart(ctrl)
	cart.EXPECT().GetCart(gomock.Any(), userID).Return(map[string]string{productID: countString}, nil)

	conn, err := getConn(cart)
	assert.NoError(t, err)
	defer conn.Close() //nolint:errcheck

	client := pb.NewCartServiceClient(conn)

	req := &pb.CartRequest{Id: userID}
	res, err := client.GetCart(ctx, req)
	assert.NoError(t, err)

	want := map[string]string{productID: countString}

	assert.Equal(t, want, res.Cart)
}

func TestHandler_GetCartErr(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cart := mock_service.NewMockCart(ctrl)
	cart.EXPECT().GetCart(gomock.Any(), userID).Return(map[string]string{}, testError)

	conn, err := getConn(cart)
	assert.NoError(t, err)
	defer conn.Close() //nolint:errcheck

	client := pb.NewCartServiceClient(conn)

	req := &pb.CartRequest{Id: userID}
	_, err = client.GetCart(ctx, req)
	assert.Error(t, err)
}

func TestHandler_AddToCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cart := mock_service.NewMockCart(ctrl)
	cart.EXPECT().AddToCart(gomock.Any(), userID, productID, count).Return(nil)

	conn, err := getConn(cart)
	assert.NoError(t, err)
	defer conn.Close() //nolint:errcheck

	client := pb.NewCartServiceClient(conn)

	req := &pb.CartUpdateRequest{
		UserId:    userID,
		ProductId: productID,
		Count:     uint32(count),
	}

	_, err = client.AddToCart(ctx, req)
	assert.NoError(t, err)
}

func TestHandler_AddToCartErr(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cart := mock_service.NewMockCart(ctrl)
	cart.EXPECT().AddToCart(gomock.Any(), userID, productID, count).Return(testError)

	conn, err := getConn(cart)
	assert.NoError(t, err)
	defer conn.Close() //nolint:errcheck

	client := pb.NewCartServiceClient(conn)

	req := &pb.CartUpdateRequest{
		UserId:    userID,
		ProductId: productID,
		Count:     uint32(count),
	}

	_, err = client.AddToCart(ctx, req)
	assert.Error(t, err)
}

func TestHandler_RemoveFromCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cart := mock_service.NewMockCart(ctrl)
	cart.EXPECT().RemoveFromCart(gomock.Any(), userID, productID, count).Return(nil)

	conn, err := getConn(cart)
	assert.NoError(t, err)
	defer conn.Close() //nolint:errcheck

	client := pb.NewCartServiceClient(conn)

	req := &pb.CartUpdateRequest{
		UserId:    userID,
		ProductId: productID,
		Count:     uint32(count),
	}

	_, err = client.RemoveFromCart(ctx, req)
	assert.NoError(t, err)
}

func TestHandler_RemoveFromCartErr(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cart := mock_service.NewMockCart(ctrl)
	cart.EXPECT().RemoveFromCart(gomock.Any(), userID, productID, count).Return(testError)

	conn, err := getConn(cart)
	assert.NoError(t, err)
	defer conn.Close() //nolint:errcheck

	client := pb.NewCartServiceClient(conn)

	req := &pb.CartUpdateRequest{
		UserId:    userID,
		ProductId: productID,
		Count:     uint32(count),
	}

	_, err = client.RemoveFromCart(ctx, req)
	assert.Error(t, err)
}

func TestHandler_RemoveAllFromCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cart := mock_service.NewMockCart(ctrl)
	cart.EXPECT().RemoveAllFromCart(gomock.Any(), userID).Return(nil)

	conn, err := getConn(cart)
	assert.NoError(t, err)
	defer conn.Close() //nolint:errcheck

	client := pb.NewCartServiceClient(conn)

	req := &pb.CartRequest{Id: userID}

	_, err = client.RemoveAllFromCart(ctx, req)
	assert.NoError(t, err)
}

func TestHandler_RemoveAllFromCartErr(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cart := mock_service.NewMockCart(ctrl)
	cart.EXPECT().RemoveAllFromCart(gomock.Any(), userID).Return(testError)

	conn, err := getConn(cart)
	assert.NoError(t, err)
	defer conn.Close() //nolint:errcheck

	client := pb.NewCartServiceClient(conn)

	req := &pb.CartRequest{Id: userID}

	_, err = client.RemoveAllFromCart(ctx, req)
	assert.Error(t, err)
}

func getConn(cart *mock_service.MockCart) (*grpc.ClientConn, error) {
	l := zap.New(nil).Sugar()
	logger := &log.Logger{SugaredLogger: l}
	services := &service.Services{Cart: cart}
	h := NewHandler(services, logger)

	return grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer(h)))
}
