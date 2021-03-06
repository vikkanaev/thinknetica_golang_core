package main

import (
	"context"
	"fmt"
	"os"
	"thinknetica_golang_core/Lesson_16-DB-App/pkg/storage"
	"thinknetica_golang_core/Lesson_16-DB-App/pkg/storage/postgres"
)

type app struct {
	storage storage.Interface
}

func new(conn string) (*app, error) {
	app := app{}
	var err error

	app.storage, err = postgres.New(conn)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	return &app, err
}

func main() {
	ctx := context.Background()
	pwd := os.Getenv("go_thinknetica_films_password")
	conn := "postgres://postgres:" + pwd + "@localhost/go_thinknetica_films"
	a, err := new(conn)
	if err != nil {
		fmt.Print(err)
		return
	}

	// Не забываем очищать ресурсы и закрывать соединения
	defer a.storage.Close()

	// Добавление фильмов
	films := []storage.Film{{Title: "Generation Pi", Year: 1999, StudioId: 1}}
	err = a.storage.AddFilms(ctx, films)
	if err != nil {
		fmt.Print(err)
		return
	}

	// Получение списка фильмов.
	data, err := a.storage.Films(ctx, 0)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("%+v\n", data)

	// Удаление фильма
	err = a.storage.DelFilm(ctx, 1)
	if err != nil {
		fmt.Print(err)
		return
	}

	// Изменения параметров фильма
	err = a.storage.UpdateFilm(ctx, 2, storage.Film{Title: "Okay", Year: 2015, StudioId: 1})
	if err != nil {
		fmt.Print(err)
		return
	}
}
