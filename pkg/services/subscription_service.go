package services

import "github.com/ainurqa95/expect-artist/pkg/repositories"

type SubscriptionsService struct {
	subscriptionRepository repositories.SubscriptionRepository
}

func NewSubscriptionService(subscriptionRepository repositories.SubscriptionRepository) *SubscriptionsService {
	return &SubscriptionsService{
		subscriptionRepository: subscriptionRepository,
	}
}

func (service *SubscriptionsService) ExistsSubscriptionBy(artistId int, userId int) bool {
	return service.subscriptionRepository.ExistsSubscriptionBy(artistId, userId)
}

func (service *SubscriptionsService) Create(artistId int, userId int) error {
	return service.subscriptionRepository.Create(artistId, userId)
}
