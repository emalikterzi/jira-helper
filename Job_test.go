package main

import (
	"testing"
	"github.com/emt/jira-helper/job"
)

func TestJobExist(t *testing.T) {
	comment := "[TEPP-481] test Commits ###log=1";
	s, _ := job.IsJobExist(comment);
	if (s != "[TEPP-481]") {
		t.Errorf("Job Not Found in %s", comment);
		return;
	}
}

func TestFindWorkLogJob(t *testing.T) {
	comment := "[TEPP-481] test Commits ###log=1m";
	task := job.FindWorkLogJob(comment);
	if (task != nil) {
		t.Logf("Log found %s", task);
		return;
	}
	t.Errorf("Log work not found in %s", comment);
}