package paths

import "fmt"

type Paths struct {
	StatePath    string
	CachePath    string
	IndexPath    string
	SettingsPath string
	SessionPath  string
	ChatsDir     string
}

func ResolvePaths() (Paths, error) {
	p := Paths{}
	statePath, err := StatePath()
	if err != nil {
		return Paths{}, fmt.Errorf("resolving state path: %w", err)
	}
	cachePath, err := CachePath()
	if err != nil {
		return Paths{}, fmt.Errorf("resolving cache path: %w", err)
	}
	p.CachePath = cachePath

	indexPath, err := IndexPath()
	if err != nil {
		return Paths{}, fmt.Errorf("resolving index path: %w", err)
	}
	p.IndexPath = indexPath

	settingsPath, err := SettingsPath()
	if err != nil {
		return Paths{}, fmt.Errorf("resolving settings path: %w", err)
	}
	p.SettingsPath = settingsPath

	sessionPath, err := SessionPath()
	if err != nil {
		return Paths{}, fmt.Errorf("resolving session path: %w", err)
	}
	p.SessionPath = sessionPath

	chatsDir, err := ChatsDir()
	if err != nil {
		return Paths{}, fmt.Errorf("resolving chats dir: %w", err)
	}
	p.ChatsDir = chatsDir
	return p, nil
}
