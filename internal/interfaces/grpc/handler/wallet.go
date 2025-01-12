package grpc_handler

import (
	"context"
	"strings"

	pb "github.com/equitas-foundation/bamp-ocean/api-spec/protobuf/gen/go/ocean/v1"
	"github.com/equitas-foundation/bamp-ocean/internal/core/application"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type wallet struct {
	appSvc *application.WalletService
}

func NewWalletHandler(
	appSvc *application.WalletService,
) pb.WalletServiceServer {
	return &wallet{
		appSvc: appSvc,
	}
}

func (w *wallet) GenSeed(
	ctx context.Context, _ *pb.GenSeedRequest,
) (*pb.GenSeedResponse, error) {
	mnemonic, err := w.appSvc.GenSeed(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.GenSeedResponse{
		Mnemonic: strings.Join(mnemonic, " "),
	}, nil
}

func (w *wallet) CreateWallet(
	ctx context.Context, req *pb.CreateWalletRequest,
) (*pb.CreateWalletResponse, error) {
	mnemonic, err := parseMnemonic(req.GetMnemonic())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	password, err := parsePassword(req.GetPassword())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := w.appSvc.CreateWallet(
		ctx, strings.Split(mnemonic, " "), password,
	); err != nil {
		return nil, err
	}

	return &pb.CreateWalletResponse{}, nil
}

func (w *wallet) Unlock(
	ctx context.Context, req *pb.UnlockRequest,
) (*pb.UnlockResponse, error) {
	password, err := parsePassword(req.GetPassword())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := w.appSvc.Unlock(ctx, password); err != nil {
		return nil, err
	}

	return &pb.UnlockResponse{}, nil
}

func (w *wallet) Lock(
	ctx context.Context, req *pb.LockRequest,
) (*pb.LockResponse, error) {
	password, err := parsePassword(req.GetPassword())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := w.appSvc.Lock(ctx, password); err != nil {
		return nil, err
	}

	return &pb.LockResponse{}, nil
}

func (w *wallet) ChangePassword(
	ctx context.Context, req *pb.ChangePasswordRequest,
) (*pb.ChangePasswordResponse, error) {
	currentPwd, err := parsePassword(req.GetCurrentPassword())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	newPwd, err := parsePassword(req.GetNewPassword())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := w.appSvc.ChangePassword(
		ctx, currentPwd, newPwd,
	); err != nil {
		return nil, err
	}

	return &pb.ChangePasswordResponse{}, nil
}

func (w *wallet) RestoreWallet(
	req *pb.RestoreWalletRequest, stream pb.WalletService_RestoreWalletServer,
) error {
	mnemonic, err := parseMnemonic(req.GetMnemonic())
	if err != nil {
		return status.Error(codes.InvalidArgument, err.Error())
	}
	password, err := parsePassword(req.GetPassword())
	if err != nil {
		return status.Error(codes.InvalidArgument, err.Error())
	}
	birthdayBlock, err := parseBlockHeight(req.GetBirthdayBlockHeight())
	if err != nil {
		return status.Error(codes.InvalidArgument, err.Error())
	}
	rootPath, err := parseRootPath(req.GetRootPath())
	if err != nil {
		return status.Error(codes.InvalidArgument, err.Error())
	}

	chMessages := make(chan application.WalletRestoreMessage)
	go w.appSvc.RestoreWallet(
		stream.Context(), chMessages,
		strings.Split(mnemonic, " "), rootPath, password, birthdayBlock,
		req.GetEmptyAccountThreshold(), req.GetUnusedAddressThreshold(),
	)

	for msg := range chMessages {
		if msg.Err != nil {
			return msg.Err
		}
		if err := stream.Send(&pb.RestoreWalletResponse{
			Message: msg.Message,
		}); err != nil {
			return err
		}
	}

	return nil
}

func (w *wallet) Status(ctx context.Context, _ *pb.StatusRequest) (*pb.StatusResponse, error) {
	status := w.appSvc.GetStatus(ctx)
	return &pb.StatusResponse{
		Initialized: status.IsInitialized,
		Unlocked:    status.IsUnlocked,
		Synced:      status.IsSynced,
	}, nil
}

func (w *wallet) GetInfo(ctx context.Context, _ *pb.GetInfoRequest) (*pb.GetInfoResponse, error) {
	info, err := w.appSvc.GetInfo(ctx)
	if err != nil {
		return nil, err
	}
	network := parseNetwork(info.Network)
	accounts := parseAccounts(info.Accounts)
	return &pb.GetInfoResponse{
		Network:             network,
		NativeAsset:         info.NativeAsset,
		RootPath:            info.RootPath,
		BirthdayBlockHash:   info.BirthdayBlockHash,
		BirthdayBlockHeight: info.BirthdayBlockHeight,
		Accounts:            accounts,
		BuildInfo: &pb.BuildInfo{
			Version: info.BuildInfo.Version,
			Commit:  info.BuildInfo.Commit,
			Date:    info.BuildInfo.Date,
		},
	}, nil
}

func (w *wallet) Auth(
	ctx context.Context,
	req *pb.AuthRequest,
) (*pb.AuthResponse, error) {
	password, err := parsePassword(req.GetPassword())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	verified, err := w.appSvc.Auth(ctx, password)
	if err != nil {
		return nil, err
	}

	return &pb.AuthResponse{
		Verified: verified,
	}, nil
}
