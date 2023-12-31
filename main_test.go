package main

import (
	"reflect"
	"testing"

	v1alpha1 "github.com/project-copacetic/copacetic/pkg/types/v1alpha1"
)

func Test_parseFakeReport(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		want    *FakeReport
		wantErr bool
	}{
		{
			name: "valid report",
			args: args{file: "testdata/fake_report.json"},
			want: &FakeReport{
				OSType:    "FakeOS",
				OSVersion: "42",
				Arch:      "amd64",
				Packages: []FakePackage{
					{
						Name:             "foo",
						InstalledVersion: "1.0.0",
						FixedVersion:     "1.0.1",
						VulnerabilityID:  "VULN001",
					},
					{
						Name:             "bar",
						InstalledVersion: "2.0.0",
						FixedVersion:     "2.0.1",
						VulnerabilityID:  "VULN002",
					},
					{
						Name:             "baz",
						InstalledVersion: "3.0.0",
						FixedVersion:     "",
						VulnerabilityID:  "VULN003",
					},
				},
			},
			wantErr: false,
		},
		{
			name:    "invalid file",
			args:    args{file: "testdata/nonexistent_file.json"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid json",
			args:    args{file: "testdata/invalid_report.json"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseFakeReport(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseFakeReport() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseFakeReport() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFakeParser(t *testing.T) {
	tests := []struct {
		name string
		want *FakeParser
	}{
		{
			name: "valid parser",
			want: &FakeParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newFakeParser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFakeParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFakeParser_Parse(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		k       *FakeParser
		args    args
		want    *v1alpha1.UpdateManifest
		wantErr bool
	}{
		{
			name: "valid report",
			k:    &FakeParser{},
			args: args{file: "testdata/fake_report.json"},
			want: &v1alpha1.UpdateManifest{
				APIVersion: v1alpha1.APIVersion,
				Metadata: v1alpha1.Metadata{
					OS: v1alpha1.OS{
						Type:    "FakeOS",
						Version: "42",
					},
					Config: v1alpha1.Config{
						Arch: "amd64",
					},
				},
				Updates: []v1alpha1.UpdatePackage{
					{
						Name:             "foo",
						InstalledVersion: "1.0.0",
						FixedVersion:     "1.0.1",
						VulnerabilityID:  "VULN001",
					},
					{
						Name:             "bar",
						InstalledVersion: "2.0.0",
						FixedVersion:     "2.0.1",
						VulnerabilityID:  "VULN002",
					},
				},
			},
			wantErr: false,
		},
		{
			name:    "invalid file",
			k:       &FakeParser{},
			args:    args{file: "testdata/nonexistent_file.json"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid json",
			k:       &FakeParser{},
			args:    args{file: "testdata/invalid_report.json"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.k.parse(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("FakeParser.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FakeParser.Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
