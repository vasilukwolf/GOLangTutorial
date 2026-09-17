package main

import "fmt"

type AccessParams struct {
	Age        int
	IsPremium  bool
	IsBanned   bool
	AgeRaiting int
}

func accessControl(params AccessParams) {
	if params.IsBanned {
		fmt.Println("Вы забанены, если мы забанили вас по ошибке, обратитесь в поддержку.")
		return
	}
	if params.Age < params.AgeRaiting {
		fmt.Println("Извините, вы слишеком молоды для просмотра этого контента.")
		return
	}
	if !params.IsPremium {
		fmt.Println("Доступен с рекламой")
		return
	} else if params.IsPremium {
		fmt.Println("Приятного просмотра")
	}
	fmt.Println("Access Granted: User has access.")
}

func main() {
	accessControl(AccessParams{Age: 20, IsPremium: false, IsBanned: false, AgeRaiting: 18})
	accessControl(AccessParams{Age: 15, IsPremium: true, IsBanned: false, AgeRaiting: 16})
	accessControl(AccessParams{Age: 25, IsPremium: true, IsBanned: true, AgeRaiting: 18})
}
