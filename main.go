package main

// Сделать конвейер чисел Даны два канала. В первый пишутся числа типа uint8. Нужно, чтобы числа читались
//из первого канала по мере поступления,
//затем эти числа должны преобразовываться в float64 и возводиться в куб и результат записывался во второй канал.
//
//Напишите main функцию, в которой протестируете весь вышеописанный функционал. Выведите результаты на экран.
//
//Напишите unit тесты к созданным функциям

func main() {
	var in, out = make(chan uint8), make(chan float64)

	go func() {
		for i := uint8(0); i < 10; i++ {
			in <- i
		}
		close(in)
	}()

	go func() {
		defer close(out)

		for {
			num, isOpened := <-in
			if !isOpened {
				return
			}

			out <- float64(num * num * num)
		}
	}()

	for num := range out {
		println(num)
	}
}
