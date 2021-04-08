package utils

import "context"

//DrainAccumulator drains an accumulator
func DrainAccumulator(ctx context.Context, acc <-chan (<-chan struct{})) <-chan struct{} {
	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			select {
			case <-ctx.Done():
				return
			case d, ok := <-acc:
				if ok {
					return
				}
				select {
				case <-ctx.Done():
					return
				case <-d:
				}
			}
		}

	}()

	return done
}
