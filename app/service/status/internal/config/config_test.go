package config

import "testing"

func TestConfig_Validate(t *testing.T) {
	tests := []struct {
		name         string
		statusExpire int
		wantErr      bool
	}{
		{"positive expire", 300, false},
		{"zero expire", 0, true},
		{"negative expire", -1, true},
		{"one second", 1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{StatusExpire: tt.statusExpire}
			err := c.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
