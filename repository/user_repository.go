package repository

import (
	"context"
	"errors"
	"latest/database"
	"latest/dto"
	"log"
	"time"
)

type UserRepository struct {
}

func (u *UserRepository) Store(user dto.User, hash string) (int64, error) {

	db := database.GetDatabase()

	sqlStatment := `
	INSERT INTO users (username, password, email, created_at)
	VALUES ($1, $2, $3, $4)
	`

	t, err := db.Exec(context.Background(),
		sqlStatment,
		user.Username,
		hash,
		user.Email,
		time.Now(),
	)

	if err != nil {
		log.Println(err)
		return 0, err
	}

	return t.RowsAffected(), err
}

/*
	Funções utilizadas para atualizar um usuário no banco de dados
*/

func (u *UserRepository) ChangePassword(password, email string) error {

	db := database.GetDatabase()
	ctx := context.Background()

	str := `UPDATE users
			SET password = $1
			where email = $2`

	tx, err := db.Begin(ctx)

	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	_, err = tx.Exec(ctx, str, password, email)

	if err != nil {
		return err
	}

	err = tx.Commit(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) ChangeEmail(newEmail, oldEmail string) error {

	db := database.GetDatabase()
	ctx := context.Background()

	str := `UPDATE users
			SET email = $1
			where email = $2`

	tx, err := db.Begin(ctx)

	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	_, err = tx.Exec(ctx, str, newEmail, oldEmail)

	if err != nil {
		return err
	}

	err = tx.Commit(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) DeleteUser(email string) error {

	db := database.GetDatabase()
	ctx := context.Background()

	str := `DELETE from users u
			WHERE u.email = $1`

	tx, err := db.Begin(ctx)

	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	_, err = tx.Exec(ctx, str, email)

	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	tx.Commit(ctx)

	return nil
}

/*
	Essas funções são utilizadas para checar os códigos enviados pelo usuário
	e alternar o status de verificado entre falso e verdadeiro
*/

func (u *UserRepository) SetVerifiedTrue(email string) (int64, error) {

	db := database.GetDatabase()
	ctx := context.Background()

	sqlStatement := `UPDATE users 
					 SET activated = true,
					 	 updated_at = $2
	   				 where email = $1`

	row, err := db.Exec(ctx, sqlStatement, email, time.Now())

	if err != nil {
		return 0, err
	}

	return row.RowsAffected(), nil
}

func (u *UserRepository) SetVerifiedFalse(email string) (int64, error) {

	db := database.GetDatabase()
	ctx := context.Background()

	sqlStatement := `UPDATE users 
					 SET activated = false,
					 	 updated_at = $2
	   				 where email = $1`

	row, err := db.Exec(ctx, sqlStatement, email, time.Now())

	if err != nil {
		return 0, err
	}

	return row.RowsAffected(), nil
}

func (u *UserRepository) CheckVerifyStatus(email string) error {

	var activated bool

	db := database.GetDatabase()

	ctx := context.Background()

	str := `SELECT activated from users u
			WHERE u.email = $1`

	err := db.QueryRow(ctx, str, email).Scan(&activated)

	if err != nil {
		return err
	}

	if activated {
		return errors.New("user already verified")
	}

	return nil

}

/*
	Essas funções são utilizadas para caso o usuário perca seu acesso a sua conta, seja
	não lembrando o seu e-mail ou a sua senha
*/

func (u *UserRepository) LostAccount(email string) (*dto.User, error) {

	db := database.GetDatabase()

	ctx := context.Background()

	var user dto.User

	str := `SELECT username from users u
			WHERE u.email = $1`

	err := db.QueryRow(ctx, str, email).Scan(&user.Username)

	if err != nil {
		return nil, errors.New("user not registred")
	}

	return &user, nil
}

func (u *UserRepository) LostPassword(email, password string) error {

	db := database.GetDatabase()
	ctx := context.Background()

	str := `UPDATE users
			SET password = $2
			where email = $1`

	tx, err := db.Begin(ctx)

	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	_, err = tx.Exec(ctx, str, email, password)

	if err != nil {
		return err
	}

	err = tx.Commit(ctx)

	if err != nil {
		return err
	}

	return nil
}

/*
	Funções auxiliares para os serviços na hora de cadastrar ou atualizar o usuário
*/

func (u *UserRepository) GetByUsername(username string) error {

	db := database.GetDatabase()

	ctx := context.Background()

	str := `SELECT username from users u
			WHERE u.username = $1`

	err := db.QueryRow(ctx, str, username).Scan(&username)

	if err == nil {
		return errors.New("user already registred")
	}

	return nil
}

func (u *UserRepository) GetByUserEmail(email string) (*dto.User, error) {

	db := database.GetDatabase()

	var user dto.User

	ctx := context.Background()

	str := `SELECT email, username, password from users u
			WHERE u.email = $1`

	err := db.QueryRow(ctx, str, email).Scan(&user.Email, &user.Username, &user.Password)

	if err != nil {
		return nil, errors.New("user not registred")
	}

	return &user, nil
}

/*
	Função utilizada pelo serviço para realizar a autenticação do usuário
*/

func (u *UserRepository) LoginWithUserAndPassword(username, password string) (string, error) {

	db := database.GetDatabase()

	var u_username string
	var u_password string

	sqlStatement := `
	SELECT username,
		   password
	FROM users
	WHERE username=$1 
	`

	err := db.QueryRow(context.Background(), sqlStatement, username).Scan(&u_username, &u_password)

	if err != nil {
		return "", errors.New("incorrect password or username")
	}

	return u_password, nil
}
