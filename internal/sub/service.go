package sub

type SubscriptionService struct {
	SubRepo *SubscriptionRepository
}

func NewSubscriptionService(subscriptionRepository *SubscriptionRepository) *SubscriptionService {
	return &SubscriptionService{
		SubRepo: subscriptionRepository,
	}
}

func(s *SubscriptionService) Create(sub *Subscription) (*Subscription,error) {
	exists, err := s.SubRepo.Exists(sub.ServiceName, sub.UserID)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrSubExists
	}
	return s.SubRepo.Create(sub)
}