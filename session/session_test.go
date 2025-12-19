package session

import (
	"cli-chat/chat"
	"cli-chat/fileatomic"
	"cli-chat/index"
	"cli-chat/paths"
	"cli-chat/providers"
	"cli-chat/settings"
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type fixture struct {
	sessionPath  string
	settingsPath string
	indexPath    string

	wroteState    State
	wroteSettings settings.Settings
	created1      time.Time
	updated1      time.Time
	created2      time.Time
	updated2      time.Time
}
type Opts struct {
	writeChat1 bool
	writeChat2 bool
}

func initFirstRun(t *testing.T) fixture {
	t.Helper()

	t.Setenv("XDG_CONFIG_HOME", t.TempDir())
	t.Setenv("XDG_DATA_HOME", t.TempDir())

	sessionPath, err := paths.SessionPath()
	require.NoError(t, err)
	settingsPath, err := paths.SettingsPath()
	require.NoError(t, err)
	indexPath, err := paths.IndexPath()
	require.NoError(t, err)

	return fixture{
		sessionPath:  sessionPath,
		settingsPath: settingsPath,
		indexPath:    indexPath,
	}
}

func initWithExistingFiles(t *testing.T, opt Opts) fixture {
	t.Helper()

	fx := initFirstRun(t)

	set := settings.Settings{
		CommandPrefix: "This",
		Timeout:       15,
		Columns:       4,
		ApiKey:        "TEST",
	}
	data, err := json.MarshalIndent(set, "", "  ")
	require.NoError(t, err)
	require.NoError(t, fileatomic.Write(fx.settingsPath, data, 0o600))

	st := State{
		LastProvider: providers.Provider{
			Name:    "yet",
			Key:     "ANOTHER",
			BaseURL: "test",
			Model:   "succeeds",
		},
	}
	data, err = json.MarshalIndent(st, "", "  ")
	require.NoError(t, err)
	require.NoError(t, fileatomic.Write(fx.sessionPath, data, 0o600))

	chatsDir, err := paths.ChatsDir()
	require.NoError(t, err)

	created1 := time.Date(2025, 1, 2, 4, 4, 5, 6, time.UTC)
	updated1 := created1.Add(10 * time.Minute)
	created2 := time.Date(2025, 1, 2, 3, 4, 5, 6, time.UTC)
	updated2 := created2.Add(10 * time.Minute)

	db := index.DB{
		Chats: []index.ChatMeta{
			{ID: "1", CreatedAt: created1, UpdatedAt: updated1},
			{ID: "2", CreatedAt: created2, UpdatedAt: updated2},
		},
	}
	require.NoError(t, db.Save(fx.indexPath))

	chat1 := chat.Chat{
		ID:        "1",
		CreatedAt: created1,
		UpdatedAt: updated1,
		Conversation: []chat.Message{
			{Role: "user", Content: "dumb question"},
			{Role: "assistant", Content: "assertive answer"},
		},
	}
	chat2 := chat.Chat{
		ID:        "2",
		CreatedAt: created2,
		UpdatedAt: updated2,
		Conversation: []chat.Message{
			{Role: "user", Content: "dumber question"},
			{Role: "assistant", Content: "eagerly assertive answer"},
		},
	}

	if opt.writeChat1 {
		require.NoError(t, chat1.Write(chatsDir))
	}
	if opt.writeChat2 {
		require.NoError(t, chat2.Write(chatsDir))
	}

	fx.wroteSettings = set
	fx.wroteState = st
	fx.created1 = created1
	fx.updated1 = updated1
	fx.created2 = created2
	fx.updated2 = updated2
	return fx
}

func TestLoadOrCreate_FirstRun(t *testing.T) {
	fx := initFirstRun(t)

	s, err := LoadOrCreate(fx.sessionPath)
	require.NoError(t, err)

	require.NotNil(t, s.Client)
	require.NotNil(t, s.Chat)
	require.NotNil(t, s.DB)

	_, err = os.Stat(fx.settingsPath)
	require.NoError(t, err)

	_, err = os.Stat(fx.sessionPath)
	require.ErrorIs(t, err, os.ErrNotExist)

	_, err = os.Stat(fx.indexPath)
	require.ErrorIs(t, err, os.ErrNotExist)
}

func TestLoadOrCreate_LoadsExistingFiles(t *testing.T) {
	fx := initWithExistingFiles(t, Opts{writeChat1: true, writeChat2: true})

	s, err := LoadOrCreate(fx.sessionPath)
	require.NoError(t, err)

	require.NotNil(t, s.Client)
	require.NotNil(t, s.Chat)
	require.NotNil(t, s.DB)

	require.Equal(t, fx.wroteState.LastProvider, s.Provider)

	require.Equal(t, fx.wroteSettings, s.AppSettings)

	require.Len(t, s.Chat.Conversation, 2)
	require.Equal(t, "user", s.Chat.Conversation[0].Role)
	require.Equal(t, "dumb question", s.Chat.Conversation[0].Content)
	require.Equal(t, "assistant", s.Chat.Conversation[1].Role)
	require.Equal(t, "assertive answer", s.Chat.Conversation[1].Content)
	require.True(t, s.Chat.CreatedAt.Equal(fx.created1))
	require.True(t, s.Chat.UpdatedAt.Equal(fx.updated1))

	require.Len(t, s.DB.Chats, 2)
	var have1, have2 bool
	for _, m := range s.DB.Chats {
		switch m.ID {
		case "1":
			have1 = true
			require.True(t, m.CreatedAt.Equal(fx.created1))
			require.True(t, m.UpdatedAt.Equal(fx.updated1))
		case "2":
			have2 = true
			require.True(t, m.CreatedAt.Equal(fx.created2))
			require.True(t, m.UpdatedAt.Equal(fx.updated2))
		}
	}
	require.True(t, have1, "db should contain chat meta id=1")
	require.True(t, have2, "db should contain chat meta id=2")
}

func TestLoadOrCreate_ChatMissing(t *testing.T) {
	fx := initWithExistingFiles(t, Opts{writeChat1: false, writeChat2: true})
	_, err := LoadOrCreate(fx.sessionPath)
	require.Error(t, err)
}
