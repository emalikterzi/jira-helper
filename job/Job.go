package job

import (
	"time"
	"regexp"
	"strings"
)

type Job struct {
	TaskName   string
	JobValue   string
	Comment    string
	JobType    string
	JobDate    time.Time
	JobDateStr string
}

func IsJobExist(Comment string) (string, error) {
	task, err1 := regexp.Compile("\\[[\\w]+-[\\w\\d]+\\]")
	task2, err2 := regexp.Compile("(###([\\w]+=[\\d]+[mh]?))")

	foundTask := task.FindAllString(Comment, 1)
	foundJobs := task2.FindAllString(Comment, 1)
	if (err1 != nil || err2 != nil || len(foundJobs) == 0 || len(foundTask) == 0) {
		return "", err1;
	}
	return foundTask[0], nil;
}

func FindWorkLogJob(Comment string) (*Job) {
	task, _ := regexp.Compile("(log=[\\d]+[mh]?)");
	logs := task.FindAllString(Comment, 1);
	if (len(logs) == 0) {
		return nil;
	}
	log := logs[0];
	splittedLog := strings.Split(log, "=");
	return &Job{JobType:splittedLog[0], JobValue:splittedLog[1]};
}