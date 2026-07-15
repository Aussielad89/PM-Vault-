package snapstock

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

type Snapstock struct {
	pvDir string
}

func New(pvDir string) *Snapstock {
	return &Snapstock{pvDir: pvDir}
}

func (s *Snapstock) Commit(message string) error {
	historyDir := filepath.Join(s.pvDir, "history")
	if err := os.MkdirAll(historyDir, 0755); err != nil {
		return err
	}

	timestamp := time.Now().Format("20060102-150405")
	snapshotDir := filepath.Join(historyDir, timestamp)
	if err := os.MkdirAll(snapshotDir, 0755); err != nil {
		return err
	}

	files := []string{"ledger.json", "snapshots.jsonl"}
	for _, f := range files {
		src := filepath.Join(s.pvDir, f)
		dst := filepath.Join(snapshotDir, f)
		if _, err := os.Stat(src); err == nil {
			data, err := os.ReadFile(src)
			if err != nil {
				continue
			}
			os.WriteFile(dst, data, 0644)
		}
	}

	meta := map[string]string{
		"timestamp": time.Now().Format(time.RFC3339),
		"message":   message,
	}
	metaPath := filepath.Join(snapshotDir, "meta.json")
	metaData, _ := os.ReadFile(metaPath)
	if metaData == nil {
		metaBytes, _ := json.MarshalIndent(meta, "", "  ")
		os.WriteFile(metaPath, metaBytes, 0644)
	}

	cmd := exec.Command("git", "-C", s.pvDir, "add", "-A")
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("git add failed: %s — %v", string(out), err)
	}

	cmd = exec.Command("git", "-C", s.pvDir, "commit", "-m", message)
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("git commit failed: %s — %v", string(out), err)
	}

	return nil
}
