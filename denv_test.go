// Package denv manages whether the application is running in Development, Integration, Staging, or Production.
package denv

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewEnv(t *testing.T) {
	t.Run("when the environment variable is not set", func(t *testing.T) {
		t.Setenv("DENV_DEPLOYMENT_ENV", "")

		_, err := NewEnv()
		if errors.Is(err, ErrEnvParse) {
			t.Errorf("NewEnv() error = %v, wantErr %v", err, false)
		}
	})

	t.Run("when the environment variable is set to an invalid value", func(t *testing.T) {
		t.Setenv("DENV_DEPLOYMENT_ENV", "invalid")

		_, err := NewEnv()
		if errors.Is(err, ErrEnvParse) {
			t.Errorf("NewEnv() error = %v, wantErr %v", err, false)
		}
	})

	t.Run("when the environment variable is set to a development", func(t *testing.T) {
		t.Setenv("DENV_DEPLOYMENT_ENV", "development")

		got, err := NewEnv()
		if err != nil {
			t.Fatal(err)
		}
		want := &Env{
			Environment: "development",
		}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("value is mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("when the environment variable is set to an integration", func(t *testing.T) {
		t.Setenv("DENV_DEPLOYMENT_ENV", "integration")

		got, err := NewEnv()
		if err != nil {
			t.Fatal(err)
		}
		want := &Env{
			Environment: "integration",
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("value is mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("when the environment variable is set to a staging", func(t *testing.T) {
		t.Setenv("DENV_DEPLOYMENT_ENV", "staging")

		got, err := NewEnv()
		if err != nil {
			t.Fatal(err)
		}
		want := &Env{
			Environment: "staging",
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("value is mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("when the environment variable is set to a production", func(t *testing.T) {
		t.Setenv("DENV_DEPLOYMENT_ENV", "production")

		got, err := NewEnv()
		if err != nil {
			t.Fatal(err)
		}
		want := &Env{
			Environment: "production",
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("value is mismatch (-want +got):\n%s", diff)
		}
	})
}

func TestEnvString(t *testing.T) {
	t.Parallel()
	type fields struct {
		Environment string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "when the environment is development",
			fields: fields{
				Environment: "development",
			},
			want: "development",
		},
		{
			name: "when the environment is integration",
			fields: fields{
				Environment: "integration",
			},
			want: "integration",
		},
		{
			name: "when the environment is staging",
			fields: fields{
				Environment: "staging",
			},
			want: "staging",
		},
		{
			name: "when the environment is production",
			fields: fields{
				Environment: "production",
			},
			want: "production",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			e := &Env{
				Environment: tt.fields.Environment,
			}
			if got := e.String(); got != tt.want {
				t.Errorf("Env.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvIsDevelopment(t *testing.T) {
	t.Parallel()
	type fields struct {
		Environment string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "when the environment is development",
			fields: fields{
				Environment: "development",
			},
			want: true,
		},
		{
			name: "when the environment is integration",
			fields: fields{
				Environment: "integration",
			},
			want: false,
		},
		{
			name: "when the environment is staging",
			fields: fields{
				Environment: "staging",
			},
			want: false,
		},
		{
			name: "when the environment is production",
			fields: fields{
				Environment: "production",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			e := &Env{
				Environment: tt.fields.Environment,
			}
			if got := e.IsDevelopment(); got != tt.want {
				t.Errorf("Env.IsDevelopment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvIsIntegration(t *testing.T) {
	t.Parallel()
	type fields struct {
		Environment string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "when the environment is development",
			fields: fields{
				Environment: "development",
			},
			want: false,
		},
		{
			name: "when the environment is integration",
			fields: fields{
				Environment: "integration",
			},
			want: true,
		},
		{
			name: "when the environment is staging",
			fields: fields{
				Environment: "staging",
			},
			want: false,
		},
		{
			name: "when the environment is production",
			fields: fields{
				Environment: "production",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			e := &Env{
				Environment: tt.fields.Environment,
			}
			if got := e.IsIntegration(); got != tt.want {
				t.Errorf("Env.IsIntegration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvIsStaging(t *testing.T) {
	t.Parallel()
	type fields struct {
		Environment string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "when the environment is development",
			fields: fields{
				Environment: "development",
			},
			want: false,
		},
		{
			name: "when the environment is integration",
			fields: fields{
				Environment: "integration",
			},
			want: false,
		},
		{
			name: "when the environment is staging",
			fields: fields{
				Environment: "staging",
			},
			want: true,
		},
		{
			name: "when the environment is production",
			fields: fields{
				Environment: "production",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			e := &Env{
				Environment: tt.fields.Environment,
			}
			if got := e.IsStaging(); got != tt.want {
				t.Errorf("Env.IsStaging() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvIsProduction(t *testing.T) {
	t.Parallel()
	type fields struct {
		Environment string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "when the environment is development",
			fields: fields{
				Environment: "development",
			},
			want: false,
		},
		{
			name: "when the environment is integration",
			fields: fields{
				Environment: "integration",
			},
			want: false,
		},
		{
			name: "when the environment is staging",
			fields: fields{
				Environment: "staging",
			},
			want: false,
		},
		{
			name: "when the environment is production",
			fields: fields{
				Environment: "production",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			e := &Env{
				Environment: tt.fields.Environment,
			}
			if got := e.IsProduction(); got != tt.want {
				t.Errorf("Env.IsProduction() = %v, want %v", got, tt.want)
			}
		})
	}
}
