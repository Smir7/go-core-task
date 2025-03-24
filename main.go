package main

type Semaphore chan struct{}

func (s Semaphore) Add(n int) {
	for i := 0; i < n; i++ {
		s <- struct{}{}
	}
}

func (s Semaphore) Delete(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

func main() {
	const count = 10

	var s = make(Semaphore, count)

	for i := 0; i < count; i++ {
		go func(n int) {
			defer s.Add(1)

			print(n, " ")
		}(i)
	}

	s.Delete(count)
}
