package chats

import "context"

func (s *serv) Delete(ctx context.Context, id int64) error {
	err := s.chatsRepository.Delete(ctx, id)

	return err
}
