package server

import (
	"reflect"
	"testing"

	bffconfig "github.com/teamgram/teamgram-server/v2/app/bff/bff/internal/config"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
)

func TestBuildBizBackedConfigSetsConcreteKitexClients(t *testing.T) {
	src := bffconfig.Config{
		RpcServerConf: kitex.RpcServerConf{
			ListenOn: "127.0.0.1:20010",
		},
		BizServiceClient: kitex.RpcClientConf{
			DestService: "service.biz_service",
			ServiceName: "RPCBizservice",
		},
		MsgClient: kitex.RpcClientConf{
			DestService: "messenger.msg",
			ServiceName: "RPCMsg",
		},
		SyncClient: kitex.RpcClientConf{
			DestService: "messenger.sync",
			ServiceName: "RPCSync",
		},
		TypingMinIntervalSeconds: 7,
	}

	chatInvites := buildChatInvitesConfig(src)
	if chatInvites.ListenOn != src.ListenOn {
		t.Fatalf("expected listen_on %q, got %q", src.ListenOn, chatInvites.ListenOn)
	}
	if chatInvites.ChatClient.DestService != "service.biz_service" {
		t.Fatalf("expected chat client dest service to be forwarded, got %#v", chatInvites.ChatClient)
	}
	if chatInvites.ChatClient.ServiceName != "RPCChat" {
		t.Fatalf("expected chat client service name RPCChat, got %#v", chatInvites.ChatClient)
	}
	if chatInvites.UserClient.DestService != "service.biz_service" {
		t.Fatalf("expected user client dest service to be forwarded, got %#v", chatInvites.UserClient)
	}
	if chatInvites.UserClient.ServiceName != "RPCUser" {
		t.Fatalf("expected user client service name RPCUser, got %#v", chatInvites.UserClient)
	}

	messages := buildMessagesConfig(src)
	if messages.ChatClient.DestService != "service.biz_service" {
		t.Fatalf("expected messages chat client dest service to be forwarded, got %#v", messages.ChatClient)
	}
	if messages.ChatClient.ServiceName != "RPCChat" {
		t.Fatalf("expected messages chat client service name RPCChat, got %#v", messages.ChatClient)
	}
	if messages.MsgClient.DestService != "messenger.msg" {
		t.Fatalf("expected msg client dest service to be forwarded, got %#v", messages.MsgClient)
	}
	if messages.MsgClient.ServiceName != "RPCMsg" {
		t.Fatalf("expected msg client service name RPCMsg, got %#v", messages.MsgClient)
	}
	if messages.UserClient.DestService != "service.biz_service" {
		t.Fatalf("expected messages user client dest service to be forwarded, got %#v", messages.UserClient)
	}
	if messages.UserClient.ServiceName != "RPCUser" {
		t.Fatalf("expected messages user client service name RPCUser, got %#v", messages.UserClient)
	}

	dialogs := buildDialogsConfig(src)
	if dialogs.DialogClient.DestService != "service.biz_service" {
		t.Fatalf("expected dialogs dialog client dest service to be forwarded, got %#v", dialogs.DialogClient)
	}
	if dialogs.DialogClient.ServiceName != "RPCDialog" {
		t.Fatalf("expected dialogs dialog client service name RPCDialog, got %#v", dialogs.DialogClient)
	}
	if dialogs.UserClient.DestService != "service.biz_service" {
		t.Fatalf("expected dialogs user client dest service to be forwarded, got %#v", dialogs.UserClient)
	}
	if dialogs.UserClient.ServiceName != "RPCUser" {
		t.Fatalf("expected dialogs user client service name RPCUser, got %#v", dialogs.UserClient)
	}
	if _, ok := reflect.TypeOf(dialogs).FieldByName("MessageClient"); ok {
		t.Fatal("dialogs config must not expose legacy MessageClient")
	}
	if dialogs.MsgClient.DestService != "messenger.msg" {
		t.Fatalf("expected dialogs msg client dest service to be forwarded, got %#v", dialogs.MsgClient)
	}
	if dialogs.MsgClient.ServiceName != "RPCMsg" {
		t.Fatalf("expected dialogs msg client service name RPCMsg, got %#v", dialogs.MsgClient)
	}
	if dialogs.SyncClient.DestService != "messenger.sync" {
		t.Fatalf("expected dialogs sync client dest service to be forwarded, got %#v", dialogs.SyncClient)
	}
	if dialogs.SyncClient.ServiceName != "RPCSync" {
		t.Fatalf("expected dialogs sync client service name RPCSync, got %#v", dialogs.SyncClient)
	}
	if dialogs.TypingMinIntervalSeconds != 7 {
		t.Fatalf("expected typing min interval 7, got %d", dialogs.TypingMinIntervalSeconds)
	}

	savedMessageDialogs := buildSavedMessageDialogsConfig(src)
	if savedMessageDialogs.DialogClient.DestService != "service.biz_service" {
		t.Fatalf("expected saved message dialogs dialog client dest service to be forwarded, got %#v", savedMessageDialogs.DialogClient)
	}
	if savedMessageDialogs.DialogClient.ServiceName != "RPCDialog" {
		t.Fatalf("expected saved message dialogs dialog client service name RPCDialog, got %#v", savedMessageDialogs.DialogClient)
	}

	drafts := buildDraftsConfig(src)
	if drafts.DialogClient.DestService != "service.biz_service" {
		t.Fatalf("expected drafts dialog client dest service to be forwarded, got %#v", drafts.DialogClient)
	}
	if drafts.DialogClient.ServiceName != "RPCDialog" {
		t.Fatalf("expected drafts dialog client service name RPCDialog, got %#v", drafts.DialogClient)
	}
	if drafts.UserClient.ServiceName != "RPCUser" {
		t.Fatalf("expected drafts user client service name RPCUser, got %#v", drafts.UserClient)
	}
	if drafts.ChatClient.ServiceName != "RPCChat" {
		t.Fatalf("expected drafts chat client service name RPCChat, got %#v", drafts.ChatClient)
	}
}

func TestBuildAuthorizationConfigUsesUnifiedBFFDependencies(t *testing.T) {
	c := bffconfig.Config{
		RpcServerConf: kitex.RpcServerConf{
			ListenOn: "127.0.0.1:0",
		},
		AuthSessionClient: kitex.RpcClientConf{
			DestService: "service.authsession",
			ServiceName: "RPCAuthsession",
		},
		BizServiceClient: kitex.RpcClientConf{
			DestService: "service.biz_service",
			ServiceName: "RPCBizservice",
		},
	}

	got := buildAuthorizationConfig(c)
	if got.AuthsessionClient.ServiceName != "RPCAuthsession" {
		t.Fatalf("AuthsessionClient.ServiceName = %q, want RPCAuthsession", got.AuthsessionClient.ServiceName)
	}
	if got.UserClient.ServiceName != "RPCUser" {
		t.Fatalf("UserClient.ServiceName = %q, want RPCUser", got.UserClient.ServiceName)
	}
	if got.UserClient.DestService != "service.biz_service" {
		t.Fatalf("UserClient.DestService = %q, want service.biz_service", got.UserClient.DestService)
	}
}

func TestBuildUpdatesConfigUsesUserupdatesClient(t *testing.T) {
	c := bffconfig.Config{
		RpcServerConf: kitex.RpcServerConf{
			ListenOn: "127.0.0.1:0",
		},
		BizServiceClient: kitex.RpcClientConf{
			DestService: "service.biz_service",
			ServiceName: "RPCBizservice",
		},
		UserupdatesClient: kitex.RpcClientConf{
			DestService: "messenger.userupdates",
			ServiceName: "RPCUserupdates",
		},
	}

	got := buildUpdatesConfig(c)
	if got.UserupdatesClient.ServiceName != "RPCUserupdates" {
		t.Fatalf("UserupdatesClient.ServiceName = %q, want RPCUserupdates", got.UserupdatesClient.ServiceName)
	}
	if got.UserupdatesClient.DestService != "messenger.userupdates" {
		t.Fatalf("UserupdatesClient.DestService = %q, want messenger.userupdates", got.UserupdatesClient.DestService)
	}
	if got.UserClient.ServiceName != "RPCUser" {
		t.Fatalf("UserClient.ServiceName = %q, want RPCUser", got.UserClient.ServiceName)
	}
	if got.UserClient.DestService != "service.biz_service" {
		t.Fatalf("UserClient.DestService = %q, want service.biz_service", got.UserClient.DestService)
	}
}
