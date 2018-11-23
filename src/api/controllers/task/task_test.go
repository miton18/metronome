package taskctrl

import (
	"testing"
)

func Test_validateCrestion(t *testing.T) {
	type args struct {
		payload string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{{
		name: "12 h of period and 6 h of epsilon",
		args: args{
			payload: `{"name": "my-event", "urn": "https://toto", "schedule": "R/2018-11-22T10:12:01Z/PT12H/ET6H"}`,
		},
		want: nil,
	}, {
		name: "6 min of period and 3 min of epsilon",
		args: args{
			payload: `{"name": "my-event", "urn": "https://toto", "schedule": "R/2018-11-21T14:49:11Z/PT6M/ET3M"}`,
		},
		want: nil,
	}, {
		name: "2 min of period 1 m of epsilon",
		args: args{
			payload: `{"name": "my-event", "urn": "https://toto", "schedule": "R/2018-11-15T18:06:47Z/PT2M/ET1M"}`,
		},
		want: nil,
	}, {
		name: "1 day and 3 seconds of epsilon",
		args: args{
			payload: `{"name": "my-event", "urn": "https://toto", "schedule": "R/2016-12-31T10:00:00Z/P1DT/ET3S"}`,
		},
		want: nil,
	}, {
		name: "DHMS 1",
		args: args{
			payload: `{"name": "my-event", "urn": "https://toto", "schedule": "R/2016-12-31T10:00:00Z/P31DT1H5M32S/ET3S"}`,
		},
		want: nil,
	}, {
		name: "DHMS 2",
		args: args{
			payload: `{"name": "my-event", "urn": "https://toto", "schedule": "R/2016-12-31T10:00:00Z/P31DT1H5M/ET3S"}`,
		},
		want: nil,
	}, {
		name: "DHM 1",
		args: args{
			payload: `{"name": "my-event", "urn": "https://toto", "schedule": "R/2016-12-31T10:00:00Z/P31DT1H5M/ET3S"}`,
		},
		want: nil,
	}, {
		name: "DHM 2 (failed because no value before M)",
		args: args{
			payload: `{"name": "my-event", "urn": "https://toto", "schedule": "R/2016-12-31T10:00:00Z/P31DT1HM32S/ET3S"}`,
		},
		want: true,
	}, {
		name: "DMS 1",
		args: args{
			payload: `{"name": "my-event", "urn": "https://toto", "schedule": "R/2016-12-31T10:00:00Z/P31DT5M32S/ET3S"}`,
		},
		want: nil,
	}, {
		name: "DMS 2 (failed because no value before H)",
		args: args{
			payload: `{"name": "my-event", "urn": "https://toto", "schedule": "R/2016-12-31T10:00:00Z/P31DTH5M32S/ET3S"}`,
		},
		want: true,
	}, {
		name: "HMS 1",
		args: args{
			payload: `{"name": "my-event", "urn": "https://toto", "schedule": "R/2016-12-31T10:00:00Z/PT1H5M32S/ET3S"}`,
		},
		want: nil,
	}, {
		name: "HMS 2 (missing T)",
		args: args{
			payload: `{"name": "my-event", "urn": "https://toto", "schedule": "R/2016-12-31T10:00:00Z/P31T1H5M32S/ET3S"}`,
		},
		want: true,
	}, {
		name: "DH 1",
		args: args{
			payload: `{"name": "my-event", "urn": "https://toto", "schedule": "R/2016-12-31T10:00:00Z/P31DT1H/ET3S"}`,
		},
		want: nil,
	}, {
		name: "DH 2 (No unit after 5)",
		args: args{
			payload: `{"name": "my-event", "urn": "https://toto", "schedule": "R/2016-12-31T10:00:00Z/P31DT1H5/ET3S"}`,
		},
		want: true,
	}, {
		name: "DM 1",
		args: args{
			payload: `{"name": "my-event", "urn": "https://toto", "schedule": "R/2016-12-31T10:00:00Z/P31DT5M/ET3S"}`,
		},
		want: nil,
	}, {
		name: "DM 2 (No unit after 32)",
		args: args{
			payload: `{"name": "my-event", "urn": "https://toto", "schedule": "R/2016-12-31T10:00:00Z/P31DT5M32/ET3S"}`,
		},
		want: true,
	}, {
		name: "DS 1",
		args: args{
			payload: `{"name": "my-event", "urn": "https://toto", "schedule": "R/2016-12-31T10:00:00Z/P31DT32S/ET3S"}`,
		},
		want: nil,
	}, {
		name: "DS 2 (No H value)",
		args: args{
			payload: `{"name": "my-event", "urn": "https://toto", "schedule": "R/2016-12-31T10:00:00Z/P31DTH32S/ET3S"}`,
		},
		want: true,
	}, {
		name: "HM 1",
		args: args{
			payload: `{"name": "my-event", "urn": "https://toto", "schedule": "R/2016-12-31T10:00:00Z/PT1H5M/ET3S"}`,
		},
		want: nil,
	}, {
		name: "HM 2 (No unit after 32)",
		args: args{
			payload: `{"name": "my-event", "urn": "https://toto", "schedule": "R/2016-12-31T10:00:00Z/PT1H5M32/ET3S"}`,
		},
		want: true,
	}, {
		name: "HS 1",
		args: args{
			payload: `{"name": "my-event", "urn": "https://toto", "schedule": "R/2016-12-31T10:00:00Z/PT1H32S/ET3S"}`,
		},
		want: nil,
	}, {
		name: "HS 2 (No value before M)",
		args: args{
			payload: `{"name": "my-event", "urn": "https://toto", "schedule": "R/2016-12-31T10:00:00Z/PT1HM32S/ET3S"}`,
		},
		want: true,
	}, {
		name: "MS 1",
		args: args{
			payload: `{"name": "my-event", "urn": "https://toto", "schedule": "R/2016-12-31T10:00:00Z/PT5M32S/ET3S"}`,
		},
		want: nil,
	}, {
		name: "MS 2 (No value before H)",
		args: args{
			payload: `{"name": "my-event", "urn": "https://toto", "schedule": "R/2016-12-31T10:00:00Z/PTH5M32S/ET3S"}`,
		},
		want: true,
	}, {
		name: "D 1",
		args: args{
			payload: `{"name": "my-event", "urn": "https://toto", "schedule": "R/2016-12-31T10:00:00Z/P1DT/ET3S"}`,
		},
		want: nil,
	}, {
		name: "D 2 (No unit after 5)",
		args: args{
			payload: `{"name": "my-event", "urn": "https://toto", "schedule": "R/2016-12-31T10:00:00Z/PT1H5/ET3S"}`,
		},
		want: true,
	}, {
		name: "H 1",
		args: args{
			payload: `{"name": "my-event", "urn": "https://toto", "schedule": "R/2016-12-31T10:00:00Z/PT1H/ET3S"}`,
		},
		want: nil,
	}, {
		name: "H 2 (No unit after 5)",
		args: args{
			payload: `{"name": "my-event", "urn": "https://toto", "schedule": "R/2016-12-31T10:00:00Z/PT1H5/ET3S"}`,
		},
		want: true,
	}, {
		name: "M 1",
		args: args{
			payload: `{"name": "my-event", "urn": "https://toto", "schedule": "R/2016-12-31T10:00:00Z/PT5M/ET3S"}`,
		},
		want: nil,
	}, {
		name: "M 2 (No unit after 32)",
		args: args{
			payload: `{"name": "my-event", "urn": "https://toto", "schedule": "R/2016-12-31T10:00:00Z/PT5M32/ET3S"}`,
		},
		want: true,
	}, {
		name: "S 1",
		args: args{
			payload: `{"name": "my-event", "urn": "https://toto", "schedule": "R/2016-12-31T10:00:00Z/PT32S/ET3S"}`,
		},
		want: nil,
	}, {
		name: "S 2 (No unit after 32)",
		args: args{
			payload: `{"name": "my-event", "urn": "https://toto", "schedule": "R/2016-12-31T10:00:00Z/PT32/ETM3S"}`,
		},
		want: true,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validateCrestion(tt.args.payload)
			if got == nil && tt.want != nil || tt.want == nil && got != nil {
				t.Errorf("validateCrestion() = %v, want %v", got, tt.want)
			}
		})
	}
}
