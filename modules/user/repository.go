package user

import (
	"database/sql"
	"errors"
	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
	"go-barbershop/utils/constant"
	"go-barbershop/utils/logger"
	"go-barbershop/utils/permission"
)

type Repository interface {
	Login(ctx *gin.Context, user LoginRequest) (result User, err error)
	SignUp(ctx *gin.Context, user User) (err error)
	GetListPermissionByRoleId(ctx *gin.Context, user User) (result []permission.Permission, err error)
}

type userRepository struct {
	db *sql.DB
}

func NewRepository(database *sql.DB) Repository {
	return &userRepository{
		db: database,
	}
}

func (r *userRepository) Login(ctx *gin.Context, user LoginRequest) (result User, err error) {
	conn := goqu.New(constant.Postgres.Dialect(), r.db)
	dialect := conn.From(constant.User.TableName()).
		Select(
			goqu.C("id"),
			goqu.C("username"),
			goqu.C("password"),
			goqu.C("role_id"),
		).
		Where(
			goqu.I("username").Eq(user.Username),
		)

	_, err = dialect.ScanStruct(&result)
	if err != nil {
		logger.ErrorWithCtx(ctx, nil, err)
		err = errors.New("failed login")
		return
	}

	return
}

func (r *userRepository) SignUp(ctx *gin.Context, user User) (err error) {
	conn := goqu.New(constant.Postgres.Dialect(), r.db)
	dataset := conn.Insert(constant.User.TableName()).Rows(
		goqu.Record{
			"username":  user.Username,
			"full_name": user.Username,
			"password":  user.Password,
		},
	)

	_, err = dataset.Executor().Exec()
	if err != nil {
		err = errors.New("failed sign up user")
		return err
	}

	return nil
}

func (r *userRepository) GetListPermissionByRoleId(ctx *gin.Context, user User) (result []permission.Permission, err error) {
	conn := goqu.New(constant.Postgres.Dialect(), r.db)
	dataset := conn.From(goqu.T(constant.RolePermission.TableName()).As("rp")).
		Join(goqu.T(constant.Permission.TableName()).As("p"), goqu.On(
			goqu.I("rp.permission_id").Eq(goqu.I("p.id")),
		)).
		Select(
			goqu.I("p.access_code"),
			goqu.I("p.grant_code"),
		).
		Where(
			goqu.I("rp.role_id").Eq(user.RoleId),
		)

	err = dataset.ScanStructs(&result)
	if err != nil {
		logger.ErrorWithCtx(ctx, nil, err)
		return
	}

	return
}
