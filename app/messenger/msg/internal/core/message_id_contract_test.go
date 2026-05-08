package core

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/internal/svc"
	msgpb "github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
)

func TestNoClientMessageIDDirectlyFromPeerSeq(t *testing.T) {
	repoRoot := filepath.Clean("../../../../..")
	scanRoots := []string{
		filepath.Join(repoRoot, "app/messenger/msg"),
		filepath.Join(repoRoot, "app/messenger/userupdates"),
		filepath.Join(repoRoot, "app/bff/dialogs"),
		filepath.Join(repoRoot, "app/service/biz/dialog"),
	}
	forbidden := []*regexp.Regexp{
		regexp.MustCompile(`MessageID:\s*op\.PeerSeq\b`),
		regexp.MustCompile(`\bId:\s*int32\([^)]*PeerSeq[^)]*\)`),
		regexp.MustCompile(`\bMaxId:\s*int32\([^)]*PeerSeq[^)]*\)`),
		regexp.MustCompile(`\bMaxId:\s*messageEvent\.PeerSeq\b`),
		regexp.MustCompile(`\bTopMessage:\s*int32\([^)]*PeerSeq[^)]*\)`),
		regexp.MustCompile(`\bReadInboxMaxId:\s*int32\([^)]*PeerSeq[^)]*\)`),
		regexp.MustCompile(`\bReadOutboxMaxId:\s*int32\([^)]*PeerSeq[^)]*\)`),
		regexp.MustCompile(`\bPinnedMsgID:\s*int32\([^)]*PeerSeq[^)]*\)`),
	}

	var offenders []string
	for _, root := range scanRoots {
		err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
			if err != nil || d == nil || d.IsDir() || !strings.HasSuffix(path, ".go") {
				return err
			}
			if strings.HasSuffix(path, "message_id_contract_test.go") {
				return nil
			}
			body, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			text := string(body)
			for _, re := range forbidden {
				if match := re.FindString(text); match != "" {
					offenders = append(offenders, path+": "+match)
				}
			}
			return nil
		})
		if err != nil {
			t.Fatal(err)
		}
	}
	if len(offenders) != 0 {
		t.Fatalf("client-visible ids derived from peer_seq:\n%s", strings.Join(offenders, "\n"))
	}
}

func TestReadHistoryMaxIDUsesResolvedPublicIDAndPeerSeqSeparately(t *testing.T) {
	updatesClient := &fakeUserUpdatesClient{
		processWithEffectsResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:      1001,
			OperationId: readHistoryOperationID(1001, 1002, 55, 9001),
			Status:      1,
			Pts:         23,
			PtsCount:    1,
			CurrentPts:  23,
		}),
	}
	repo := &fakeMsgRepository{
		resolveByUserMessageID: map[resolveMessageKey]*repository.ResolvedMessageID{
			{userID: 1001, peerType: payload.PeerTypeUser, peerID: 1002, userMessageID: 55}: {
				UserID:             1001,
				PeerType:           payload.PeerTypeUser,
				PeerID:             1002,
				UserMessageID:      55,
				PeerSeq:            7,
				CanonicalMessageID: 7007,
			},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo, UserUpdates: updatesClient})

	if _, err := core.MsgReadHistoryV2(&msgpb.TLMsgReadHistoryV2{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  payload.PeerTypeUser,
		PeerId:    1002,
		MaxId:     55,
	}); err != nil {
		t.Fatalf("MsgReadHistoryV2() error = %v", err)
	}
	if repo.resolveInput.UserMessageID != 55 {
		t.Fatalf("resolver input user_message_id = %d, want requester public id 55", repo.resolveInput.UserMessageID)
	}
	var op payload.MessageOperationV1
	if err := json.Unmarshal(updatesClient.processWithEffects.Operation.Payload, &op); err != nil {
		t.Fatalf("decode read history payload: %v", err)
	}
	if op.ReadInboxMaxPeerSeq != 7 || op.ReadMaxUserMessageID != 55 {
		t.Fatalf("read history ids = peer_seq:%d public:%d, want peer_seq 7 public 55", op.ReadInboxMaxPeerSeq, op.ReadMaxUserMessageID)
	}
}
