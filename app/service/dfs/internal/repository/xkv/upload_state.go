package xkv

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/teamgram/marmota/pkg/stores/kv"
)

const (
	uploadStateTTLSeconds         = 3 * 60 * 60
	uploadStateCacheRefTTLSeconds = 2 * 60 * 60
)

var ErrUploadStateNotFound = errors.New("dfs upload state not found")

type DfsFileInfo struct {
	Creator           int64
	FileID            int64
	Big               bool
	FileName          string
	FileTotalParts    int
	FirstFilePartSize int
	FilePartSize      int
	LastFilePartSize  int
	MimeType          string
	Mtime             int64
}

type UploadStateModel interface {
	SaveUploadPart(ctx context.Context, creator, fileID int64, partIndex int32, data []byte) error
	SaveUploadFileInfo(ctx context.Context, info *DfsFileInfo) error
	LoadUploadFileInfo(ctx context.Context, creator, fileID int64) (*DfsFileInfo, error)
	OpenUploadFileReader(ctx context.Context, info *DfsFileInfo) (io.ReadSeeker, error)
	SaveObjectCacheRef(ctx context.Context, objectID, creator, fileID int64) error
	LoadObjectCacheRef(ctx context.Context, objectID int64) (*DfsFileInfo, error)
}

type uploadStateModel struct {
	kv kv.ExtStore
}

func NewUploadStateModel(store kv.ExtStore) UploadStateModel {
	return &uploadStateModel{kv: store}
}

func uploadPartKey(creator, fileID int64) string {
	return fmt.Sprintf("file_%d_%d", creator, fileID)
}

func uploadInfoKey(creator, fileID int64) string {
	return fmt.Sprintf("file_info_%d_%d", creator, fileID)
}

func uploadCacheRefKey(objectID int64) string {
	return fmt.Sprintf("cache_file_info_%d", objectID)
}

func (m *uploadStateModel) SaveUploadPart(ctx context.Context, creator, fileID int64, partIndex int32, data []byte) error {
	key := uploadPartKey(creator, fileID)
	if err := m.kv.HsetCtx(ctx, key, strconv.FormatInt(int64(partIndex), 10), string(data)); err != nil {
		return fmt.Errorf("upload_state.SaveUploadPart hset: %w", err)
	}
	if err := m.kv.ExpireCtx(ctx, key, uploadStateTTLSeconds); err != nil {
		return fmt.Errorf("upload_state.SaveUploadPart expire: %w", err)
	}
	return nil
}

func (m *uploadStateModel) SaveUploadFileInfo(ctx context.Context, info *DfsFileInfo) error {
	if info == nil {
		return nil
	}
	key := uploadInfoKey(info.Creator, info.FileID)
	fields := fileInfoFields(info)
	for field, value := range fields {
		if err := m.kv.HsetCtx(ctx, key, field, value); err != nil {
			return fmt.Errorf("upload_state.SaveUploadFileInfo hset %s: %w", field, err)
		}
	}
	if err := m.kv.ExpireCtx(ctx, key, uploadStateTTLSeconds); err != nil {
		return fmt.Errorf("upload_state.SaveUploadFileInfo expire: %w", err)
	}
	return nil
}

func (m *uploadStateModel) LoadUploadFileInfo(ctx context.Context, creator, fileID int64) (*DfsFileInfo, error) {
	key := uploadInfoKey(creator, fileID)
	fields, err := m.kv.HgetallCtx(ctx, key)
	if err != nil {
		return nil, fmt.Errorf("upload_state.LoadUploadFileInfo hgetall: %w", err)
	}
	if len(fields) == 0 {
		return nil, ErrUploadStateNotFound
	}
	info, err := parseFileInfoFields(creator, fileID, fields)
	if err != nil {
		return nil, fmt.Errorf("upload_state.LoadUploadFileInfo parse: %w", err)
	}
	if info.FileTotalParts == 0 {
		totalParts, lastPartSize, err := m.loadUploadFilePartStats(ctx, creator, fileID)
		if err != nil {
			return nil, err
		}
		info.FileTotalParts = totalParts
		info.LastFilePartSize = lastPartSize
	}
	return info, nil
}

func (m *uploadStateModel) OpenUploadFileReader(ctx context.Context, info *DfsFileInfo) (io.ReadSeeker, error) {
	if info == nil || info.FileTotalParts <= 0 {
		return nil, ErrUploadStateNotFound
	}
	key := uploadPartKey(info.Creator, info.FileID)
	var buf bytes.Buffer
	for i := int32(0); i < int32(info.FileTotalParts); i++ {
		part, err := m.kv.HgetCtx(ctx, key, strconv.FormatInt(int64(i), 10))
		if err != nil {
			return nil, fmt.Errorf("upload_state.OpenUploadFileReader hget: %w", err)
		}
		if part == "" {
			return nil, ErrUploadStateNotFound
		}
		buf.WriteString(part)
	}
	return bytes.NewReader(buf.Bytes()), nil
}

func (m *uploadStateModel) loadUploadFilePartStats(ctx context.Context, creator, fileID int64) (int, int, error) {
	key := uploadPartKey(creator, fileID)
	totalParts, err := m.kv.HlenCtx(ctx, key)
	if err != nil {
		return 0, 0, fmt.Errorf("upload_state.loadUploadFilePartStats hlen: %w", err)
	}
	if totalParts == 0 {
		return 0, 0, ErrUploadStateNotFound
	}
	lastPart, err := m.kv.HgetCtx(ctx, key, strconv.Itoa(totalParts-1))
	if err != nil {
		return 0, 0, fmt.Errorf("upload_state.loadUploadFilePartStats hget: %w", err)
	}
	if lastPart == "" {
		return 0, 0, ErrUploadStateNotFound
	}
	return totalParts, len([]byte(lastPart)), nil
}

func (m *uploadStateModel) SaveObjectCacheRef(ctx context.Context, objectID, creator, fileID int64) error {
	value := fmt.Sprintf("%d_%d", creator, fileID)
	if err := m.kv.SetexCtx(ctx, uploadCacheRefKey(objectID), value, uploadStateCacheRefTTLSeconds); err != nil {
		return fmt.Errorf("upload_state.SaveObjectCacheRef setex: %w", err)
	}
	return nil
}

func (m *uploadStateModel) LoadObjectCacheRef(ctx context.Context, objectID int64) (*DfsFileInfo, error) {
	value, err := m.kv.GetCtx(ctx, uploadCacheRefKey(objectID))
	if err != nil {
		return nil, fmt.Errorf("upload_state.LoadObjectCacheRef get: %w", err)
	}
	if value == "" {
		return nil, ErrUploadStateNotFound
	}
	parts := strings.Split(value, "_")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid cache file info ref %q", value)
	}
	creator, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid cache file info creator %q: %w", parts[0], err)
	}
	fileID, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid cache file info file_id %q: %w", parts[1], err)
	}
	return m.LoadUploadFileInfo(ctx, creator, fileID)
}

func fileInfoFields(info *DfsFileInfo) map[string]string {
	fields := make(map[string]string, 7)
	if info.Big {
		fields["big"] = "true"
	}
	if info.FileName != "" {
		fields["file_name"] = info.FileName
	}
	if info.FileTotalParts != 0 {
		fields["file_total_parts"] = strconv.Itoa(info.FileTotalParts)
	}
	if info.FirstFilePartSize != 0 {
		fields["first_file_part_size"] = strconv.Itoa(info.FirstFilePartSize)
	}
	if info.FilePartSize != 0 {
		fields["file_part_size"] = strconv.Itoa(info.FilePartSize)
	}
	if info.LastFilePartSize != 0 {
		fields["last_file_part_size"] = strconv.Itoa(info.LastFilePartSize)
	}
	if info.Mtime != 0 {
		fields["mtime"] = strconv.FormatInt(info.Mtime, 10)
	}
	return fields
}

func parseFileInfoFields(creator, fileID int64, fields map[string]string) (*DfsFileInfo, error) {
	info := &DfsFileInfo{
		Creator: creator,
		FileID:  fileID,
	}
	var err error
	if raw := fields["big"]; raw != "" {
		info.Big = raw == "1" || strings.EqualFold(raw, "true")
	}
	info.FileName = fields["file_name"]
	if info.FileTotalParts, err = parseIntField(fields, "file_total_parts"); err != nil {
		return nil, err
	}
	if info.FirstFilePartSize, err = parseIntField(fields, "first_file_part_size"); err != nil {
		return nil, err
	}
	if info.FilePartSize, err = parseIntField(fields, "file_part_size"); err != nil {
		return nil, err
	}
	if info.LastFilePartSize, err = parseIntField(fields, "last_file_part_size"); err != nil {
		return nil, err
	}
	if info.Mtime, err = parseInt64Field(fields, "mtime"); err != nil {
		return nil, err
	}
	return info, nil
}

func parseIntField(fields map[string]string, name string) (int, error) {
	if fields[name] == "" {
		return 0, nil
	}
	n, err := strconv.Atoi(fields[name])
	if err != nil {
		return 0, fmt.Errorf("%s: %w", name, err)
	}
	return n, nil
}

func parseInt64Field(fields map[string]string, name string) (int64, error) {
	if fields[name] == "" {
		return 0, nil
	}
	n, err := strconv.ParseInt(fields[name], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", name, err)
	}
	return n, nil
}
