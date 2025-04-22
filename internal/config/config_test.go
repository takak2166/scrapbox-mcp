package config

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLoadConfig(t *testing.T) {
	tests := map[string][]struct {
		env     map[string]string
		want    *Config
		wantErr bool
	}{
		"ok: valid config with default port": {
			{
				env: map[string]string{
					"SCRAPBOX_SID":     "test_sid",
					"SCRAPBOX_PROJECT": "test_project",
				},
				want: &Config{
					ScrapboxSID: "test_sid",
					ProjectName: "test_project",
					Port:        8080,
				},
				wantErr: false,
			},
		},
		"ok: valid config with custom port": {
			{
				env: map[string]string{
					"SCRAPBOX_SID":     "test_sid",
					"SCRAPBOX_PROJECT": "test_project",
					"PORT":             "3000",
				},
				want: &Config{
					ScrapboxSID: "test_sid",
					ProjectName: "test_project",
					Port:        3000,
				},
				wantErr: false,
			},
		},
		"err: invalid PORT": {
			{
				env: map[string]string{
					"SCRAPBOX_SID":     "test_sid",
					"SCRAPBOX_PROJECT": "test_project",
					"PORT":             "invalid",
				},
				want:    nil,
				wantErr: true,
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			for _, tc := range tt {
				// Set environment variables
				for k, v := range tc.env {
					os.Setenv(k, v)
				}

				// Execute test
				got, err := LoadConfig()

				// Clear environment variables
				for k := range tc.env {
					os.Unsetenv(k)
				}

				if tc.wantErr {
					if err == nil {
						t.Error("LoadConfig() error = nil, wantErr true")
					}
					continue
				}

				if err != nil {
					t.Errorf("LoadConfig() error = %v, wantErr false", err)
					continue
				}

				if diff := cmp.Diff(tc.want, got); diff != "" {
					t.Errorf("LoadConfig() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
