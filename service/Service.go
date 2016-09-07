package service

import (
	"net/http"
	"github.com/dghubble/sling"
	"github.com/emt/jira-helper/dto"
)

const baseUrl = "http://etiyajira.cloudapp.net/rest/"

type Client struct {
	AuthService  *AuthService
	IssueService *IssueService
}

func (s *AuthService) Login(user *dto.LoginRequest) (error) {
	response := &dto.LoginResponse{}
	_, err := s.sling.New().Post("auth/1/session").BodyJSON(user).ReceiveSuccess(response)
	if (err == nil) {
		s.user.StoreToken(response.Session.Value)
	}
	return err;
}

func NewClient(httpClient *http.Client) *Client {
	user := new(dto.User)
	client := &Client{
		AuthService: newAuthService(httpClient, user),
		IssueService : newIssueService(httpClient, user)}
	return client;
}

func (s *IssueService) WorkLog(req *dto.WorkLogRequest, issue string) (*http.Response, bool) {
	res, _ := s.sling.New().Post(issue + "/worklog").Add("Cookie", "JSESSIONID=" + s.user.GetToken()).BodyJSON(req).Receive(nil, nil);
	return res, res.StatusCode == 201;
}

type IssueService struct {
	user  *dto.User
	sling *sling.Sling
}

type AuthService struct {
	user  *dto.User
	sling *sling.Sling
}

func newAuthService(httpClient *http.Client, user *dto.User) *AuthService {
	return &AuthService{
		sling: sling.New().Client(httpClient).Base(baseUrl),
		user :user,
	}
}

func newIssueService(httpClient *http.Client, user *dto.User) *IssueService {
	return &IssueService{
		sling: sling.New().Client(httpClient).Base(baseUrl + "api/2/issue/"),
		user :user,
	}
}



