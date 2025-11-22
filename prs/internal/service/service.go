package service

type PRService interface {

}

type prservice struct {

}

func NewPRService() PRService {
	return &prservice{}
}