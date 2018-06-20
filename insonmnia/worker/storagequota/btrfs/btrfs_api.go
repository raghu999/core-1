package btrfs

import (
	"context"
	"fmt"
	"os/exec"
	"strconv"

	log "github.com/noxiouz/zapctx/ctxlog"
	"go.uber.org/zap"
)

type btrfsIf interface {
	QuotaEnable(ctx context.Context, path string) error
	QuotaExists(ctx context.Context, qgroupID string, path string) (bool, error)
	QuotaCreate(ctx context.Context, qgroupID string, path string) error
	QuotaDestroy(ctx context.Context, qgroupID string, path string) error
	QuotaLimit(ctx context.Context, sizeInBytes uint64, qgroupID string, path string) error
	QuotaAssign(ctx context.Context, src string, dst string, path string) error
	QuotaRemove(ctx context.Context, src string, dst string, path string) error
	GetQuotaID(ctx context.Context, path string) (string, error)
}

var API btrfsIf = btrfsCLI{}

type btrfsCLI struct{}

func (btrfsCLI) QuotaEnable(ctx context.Context, path string) error {
	output, err := exec.CommandContext(ctx, "btrfs", "quota", "enable", path).Output()
	if err != nil {
		log.G(ctx).Error("failed to create to enable btrfs quota", zap.String("path", path), zap.Error(err), zap.ByteString("output", output))
		return err
	}
	return nil
}

func (btrfsCLI) QuotaExists(ctx context.Context, qgroupID string, path string) (bool, error) {
	// TODO: add implementation
	return false, fmt.Errorf("QuotaExists not implemented")
}

func (btrfsCLI) GetQuotaID(ctx context.Context, path string) (string, error) {
	return "", fmt.Errorf("GetQuotaID not implemented")
}

func (btrfsCLI) QuotaCreate(ctx context.Context, qgroupID string, path string) error {
	output, err := exec.CommandContext(ctx, "btrfs", "qgroup", "create", qgroupID, path).Output()
	if err != nil {
		log.G(ctx).Error("failed to create btrfs qgroup", zap.String("path", path), zap.String("qgroupid", qgroupID), zap.Error(err), zap.ByteString("output", output))
		return err
	}
	return nil
}

func (btrfsCLI) QuotaDestroy(ctx context.Context, qgroupID string, path string) error {
	output, err := exec.CommandContext(ctx, "btrfs", "qgroup", "destroy", qgroupID, path).Output()
	if err != nil {
		log.G(ctx).Error("failed to destroy btrfs qgroup", zap.String("path", path), zap.String("qgroupid", qgroupID), zap.Error(err), zap.ByteString("output", output))
		return err
	}
	return nil
}

func (btrfsCLI) QuotaLimit(ctx context.Context, sizeInBytes uint64, qgroupID string, path string) error {
	output, err := exec.CommandContext(ctx, "btrfs", "qgroup", "limit", strconv.FormatUint(sizeInBytes, 10), qgroupID, path).Output()
	if err != nil {
		log.G(ctx).Error("failed to limit btrfs qgroup", zap.String("qgroupid", qgroupID), zap.Error(err), zap.ByteString("output", output))
		return err
	}
	return nil
}

func (btrfsCLI) QuotaAssign(ctx context.Context, src string, dst string, path string) error {
	output, err := exec.CommandContext(ctx, "btrfs", "qgroup", "assign", src, dst, path).Output()
	if err != nil {
		log.G(ctx).Error("failed to assign btrfs qgroup", zap.String("src", src), zap.String("dst", dst), zap.Error(err), zap.ByteString("output", output))
		return err
	}
	return nil
}

func (btrfsCLI) QuotaRemove(ctx context.Context, src string, dst string, path string) error {
	output, err := exec.CommandContext(ctx, "btrfs", "qgroup", "remove", src, dst, path).Output()
	if err != nil {
		log.G(ctx).Error("failed to remove btrfs qgroup", zap.String("src", src), zap.String("dst", dst), zap.Error(err), zap.ByteString("output", output))
		return err
	}
	return nil
}
