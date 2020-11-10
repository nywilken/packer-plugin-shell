// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package shell

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName     *string           `mapstructure:"packer_build_name" cty:"packer_build_name" hcl:"packer_build_name"`
	PackerBuilderType   *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type" hcl:"packer_builder_type"`
	PackerCoreVersion   *string           `mapstructure:"packer_core_version" cty:"packer_core_version" hcl:"packer_core_version"`
	PackerDebug         *bool             `mapstructure:"packer_debug" cty:"packer_debug" hcl:"packer_debug"`
	PackerForce         *bool             `mapstructure:"packer_force" cty:"packer_force" hcl:"packer_force"`
	PackerOnError       *string           `mapstructure:"packer_on_error" cty:"packer_on_error" hcl:"packer_on_error"`
	PackerUserVars      map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables" hcl:"packer_user_variables"`
	PackerSensitiveVars []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables" hcl:"packer_sensitive_variables"`
	Inline              []string          `cty:"inline" hcl:"inline"`
	Script              *string           `cty:"script" hcl:"script"`
	Scripts             []string          `cty:"scripts" hcl:"scripts"`
	ValidExitCodes      []int             `mapstructure:"valid_exit_codes" cty:"valid_exit_codes" hcl:"valid_exit_codes"`
	Vars                []string          `mapstructure:"environment_vars" cty:"environment_vars" hcl:"environment_vars"`
	EnvVarFormat        *string           `mapstructure:"env_var_format" cty:"env_var_format" hcl:"env_var_format"`
	Binary              *bool             `cty:"binary" hcl:"binary"`
	RemotePath          *string           `mapstructure:"remote_path" cty:"remote_path" hcl:"remote_path"`
	ExecuteCommand      *string           `mapstructure:"execute_command" cty:"execute_command" hcl:"execute_command"`
	InlineShebang       *string           `mapstructure:"inline_shebang" cty:"inline_shebang" hcl:"inline_shebang"`
	PauseAfter          *string           `mapstructure:"pause_after" cty:"pause_after" hcl:"pause_after"`
	UseEnvVarFile       *bool             `mapstructure:"use_env_var_file" cty:"use_env_var_file" hcl:"use_env_var_file"`
	RemoteFolder        *string           `mapstructure:"remote_folder" cty:"remote_folder" hcl:"remote_folder"`
	RemoteFile          *string           `mapstructure:"remote_file" cty:"remote_file" hcl:"remote_file"`
	StartRetryTimeout   *string           `mapstructure:"start_retry_timeout" cty:"start_retry_timeout" hcl:"start_retry_timeout"`
	SkipClean           *bool             `mapstructure:"skip_clean" cty:"skip_clean" hcl:"skip_clean"`
	ExpectDisconnect    *bool             `mapstructure:"expect_disconnect" cty:"expect_disconnect" hcl:"expect_disconnect"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":          &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":        &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_core_version":        &hcldec.AttrSpec{Name: "packer_core_version", Type: cty.String, Required: false},
		"packer_debug":               &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":               &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":            &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":      &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables": &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"inline":                     &hcldec.AttrSpec{Name: "inline", Type: cty.List(cty.String), Required: false},
		"script":                     &hcldec.AttrSpec{Name: "script", Type: cty.String, Required: false},
		"scripts":                    &hcldec.AttrSpec{Name: "scripts", Type: cty.List(cty.String), Required: false},
		"valid_exit_codes":           &hcldec.AttrSpec{Name: "valid_exit_codes", Type: cty.List(cty.Number), Required: false},
		"environment_vars":           &hcldec.AttrSpec{Name: "environment_vars", Type: cty.List(cty.String), Required: false},
		"env_var_format":             &hcldec.AttrSpec{Name: "env_var_format", Type: cty.String, Required: false},
		"binary":                     &hcldec.AttrSpec{Name: "binary", Type: cty.Bool, Required: false},
		"remote_path":                &hcldec.AttrSpec{Name: "remote_path", Type: cty.String, Required: false},
		"execute_command":            &hcldec.AttrSpec{Name: "execute_command", Type: cty.String, Required: false},
		"inline_shebang":             &hcldec.AttrSpec{Name: "inline_shebang", Type: cty.String, Required: false},
		"pause_after":                &hcldec.AttrSpec{Name: "pause_after", Type: cty.String, Required: false},
		"use_env_var_file":           &hcldec.AttrSpec{Name: "use_env_var_file", Type: cty.Bool, Required: false},
		"remote_folder":              &hcldec.AttrSpec{Name: "remote_folder", Type: cty.String, Required: false},
		"remote_file":                &hcldec.AttrSpec{Name: "remote_file", Type: cty.String, Required: false},
		"start_retry_timeout":        &hcldec.AttrSpec{Name: "start_retry_timeout", Type: cty.String, Required: false},
		"skip_clean":                 &hcldec.AttrSpec{Name: "skip_clean", Type: cty.Bool, Required: false},
		"expect_disconnect":          &hcldec.AttrSpec{Name: "expect_disconnect", Type: cty.Bool, Required: false},
	}
	return s
}
