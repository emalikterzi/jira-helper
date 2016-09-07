Jira Helper
-----------

---

### Intro

For Fun :)

##### Setup ( Windows )
---

Dont know  O_o


##### Setup ( Linux )
---

GO Binary : 1.7

Env GOROOT and GOPATH must be setted

```
!# ~/.bashrc

M2_HOME=/home/emt/Apache/apache-maven-3.3.9
JAVA_HOME=/home/emt/Jdks/jdk1.8.0_101
GRADLE_HOME=/home/emt/gradle-3.0
GOROOT=/home/emt/Binaries/Go/1.7
GOPATH=/home/emt/go_path_env

export M2_HOME
export JAVA_HOME
export GRADLE_HOME
export GOPATH
export GOROOT

export PATH=$PATH:$M2_HOME/bin:$JAVA_HOME/bin:$GRADLE_HOME/bin:$GOROOT/bin
```

Sample Properties file

```
!#  jira-helper.yml

repo: [/home/emt/Etiya-Projects-Git/Telaura-E-Commerce_UI]
db: /home/emt/test.db
author: Enes
jira:
  email: enesmalik.terzi@etiya.com
  password: :)
```

to overload  jira-helper location edit main.go in source file

```
!# main.go
const propertiesFileName = "jira-helper.yml"

func main() {

	data, err := ioutil.ReadFile("/home/emt/" + propertiesFileName)
	if (err != nil) {
		panic(err);
	}

	properties := container.JiraProperties{}

	prErr := yaml.Unmarshal(data, &properties)
```

### Usage
----

```
!# bash
go build # first time
go install # first time

$GOROOT/bin/jira-helper
```

After executed commits will be logged automatically


Sample commit

```
!# bash
git commit -m "git commit -m "[TEPP-604] testset #jira;log=50m"
git commit -m "git commit -m "[TEPP-604] testset #jira;log=1h"
```



### To Dos
---


- [x] WorkLog
- [ ] Close Issue
- [ ] Resolve Issue
- [ ] Create Branch From Issue