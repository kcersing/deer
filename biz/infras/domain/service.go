package domain

type Service struct {
	repo       Repository
	eventBus   EventBus
	txProvider TransactionProvider
}

func (s Service) CreateOrder(cmd CreateOrderCommand) (interface{}, error) {

	return s.txProvider.Transactional(func() (interface{}, error) {
		aggregate := &Aggregate{}
		// 加载聚合根
		if err := aggregate.CreateOrder(cmd); err != nil {
			return "", err
		}
		// 保存聚合根
		if err := s.repo.Save(aggregate); err != nil {
			return "", err
		}
		// 发布领域事件
		s.eventBus.Publish(aggregate.Events...)
		return aggregate.Id, nil
	})
}
