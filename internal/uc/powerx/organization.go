package powerx

import (
	"PowerX/internal/model"
	"PowerX/internal/types"
	"PowerX/internal/types/errorx"
	"context"
	"fmt"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strings"
)

type OrganizationUseCase struct {
	db *gorm.DB
}

func NewOrganizationUseCase(db *gorm.DB) *OrganizationUseCase {
	return &OrganizationUseCase{
		db: db,
	}
}

func (uc *OrganizationUseCase) Init() {
	var count int64
	if err := uc.db.Model(&Department{}).Count(&count).Error; err != nil {
		panic(errors.Wrap(err, "init root dep failed"))
	}
	if count == 0 {
		dep := defaultDepartment()
		if err := uc.db.Model(&Department{}).Create(&dep).Error; err != nil {
			panic(errors.Wrap(err, "init root dep failed"))
		}
	}
}

const (
	GenderMale   = "male"
	GenderFeMale = "female"
	GenderUnKnow = "un_know"
)

const (
	EmployeeStatusDisabled = "disabled"
	EmployeeStatusEnabled  = "enabled"
)

type Employee struct {
	model.Model
	Account       string `gorm:"unique"`
	Name          string
	NickName      string
	Desc          string
	Position      string
	JobTitle      string
	DepartmentId  int64
	Department    *Department
	MobilePhone   string
	Gender        string
	Email         string
	ExternalEmail string
	Avatar        string
	Password      string
	Status        string `gorm:"index"`
	IsReserved    bool
	IsActivated   bool
}

const defaultCost = bcrypt.MinCost

// 生成哈希密码
func hashPassword(password string) (hashedPwd string, err error) {
	newPassword, err := bcrypt.GenerateFromPassword([]byte(password), defaultCost)
	if err != nil {
		return "", errors.Wrap(err, "gen pwd failed")
	}
	return string(newPassword), nil
}

// 校验密码
func verifyPassword(hashedPwd string, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd))
	return err == nil
}

func (e *Employee) HashPassword() (err error) {
	if e.Password != "" {
		e.Password, err = hashPassword(e.Password)
	}
	return nil
}

type Department struct {
	model.Model
	Name        string
	PId         int64
	PDep        *Department `gorm:"foreignKey:PId"`
	LeaderId    int64
	Leader      *Employee     `gorm:"foreignKey:LeaderId"`
	Ancestors   []*Department `gorm:"many2many:department_ancestors;"`
	Desc        string
	PhoneNumber string
	Email       string
	Remark      string
	IsReserved  bool
}

func defaultDepartment() *Department {
	return &Department{
		Name:       "组织架构",
		PId:        0,
		Desc:       "根节点, 别删除",
		IsReserved: true,
	}
}

func (uc *OrganizationUseCase) VerifyPassword(hashedPwd string, pwd string) bool {
	return verifyPassword(hashedPwd, pwd)
}

func (uc *OrganizationUseCase) CreateEmployee(ctx context.Context, employee *Employee) (err error) {
	if err := uc.db.WithContext(ctx).Create(&employee).Error; err != nil {
		// todo use errors.Is() when gorm update ErrDuplicatedKey
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return errorx.WithCause(errorx.ErrBadRequest, "账号已存在")
		}
		panic(err)
	}
	return nil
}

func (uc *OrganizationUseCase) FindAccountsByIds(ctx context.Context, employeeIds []int64) (accounts []string) {
	err := uc.db.WithContext(ctx).Model(&Employee{}).Where("id in ?", employeeIds).Pluck("account", &accounts).Error
	if err != nil {
		panic(errors.Wrap(err, "find accounts by ids failed"))
	}
	return accounts
}

func (uc *OrganizationUseCase) PatchEmployeeByUserId(ctx context.Context, employee *Employee, employeeId int64) error {
	result := uc.db.WithContext(ctx).Model(&Employee{}).Where(employee.ID).Updates(&employee)
	if result.Error != nil {
		panic(result.Error)
	}
	if result.RowsAffected == 0 {
		return errorx.WithCause(errorx.ErrBadRequest, "未找到员工")
	}
	return nil
}

type FindManyEmployeesOption struct {
	Ids             []int64
	Accounts        []string
	Names           []string
	LikeName        string
	Emails          []string
	LikeEmail       string
	DepIds          []int64
	Positions       []string
	PhoneNumbers    []string
	LikePhoneNumber string
	Statuses        []string
	PageIndex       int
	PageSize        int
}

func buildFindManyEmployeesQueryNoPage(query *gorm.DB, opt *FindManyEmployeesOption) *gorm.DB {
	if len(opt.Ids) > 0 {
		query.Where("id in ?", opt.Ids)
	}
	if len(opt.Names) > 0 {
		query.Where("name in ?", opt.Names)
	} else if opt.LikeName != "" {
		query.Where("name like ?", fmt.Sprintf("%s%%", opt.LikeName))
	}
	if len(opt.Emails) > 0 {
		query.Where("email in ?", opt.Emails)
	} else if opt.LikeEmail != "" {
		query.Where("email like ?", fmt.Sprintf("%s%%", opt.LikeEmail))
	}
	if len(opt.PhoneNumbers) > 0 {
		query.Where("mobile_phone in ?")
	} else if opt.LikePhoneNumber != "" {
		query.Where("mobile_phone like ?", fmt.Sprintf("%s%%", opt.LikePhoneNumber))
	}
	if len(opt.Positions) > 0 {
		query.Where("position in ?", opt.Positions)
	}
	if len(opt.Accounts) > 0 {
		query.Where("account in ?", opt.Accounts)
	}
	if len(opt.DepIds) > 0 {
		query.Where("? && department_ids", pq.Int64Array(opt.DepIds))
	}
	if len(opt.Statuses) > 0 {
		query.Where("status in ?", opt.Statuses)
	}
	return query
}

func (uc *OrganizationUseCase) FindManyEmployeesPage(ctx context.Context, opt *FindManyEmployeesOption) types.Page[*Employee] {
	var employees []*Employee
	var count int64
	query := uc.db.WithContext(ctx).Model(&Employee{})

	if opt.PageIndex == 0 {
		opt.PageIndex = 1
	}
	if opt.PageSize == 0 {
		opt.PageSize = 20
	}

	if opt.PageIndex != 0 && opt.PageSize != 0 {
		query.Offset((opt.PageIndex - 1) * opt.PageSize).Limit(opt.PageSize)
	}
	query = buildFindManyEmployeesQueryNoPage(query, opt)
	if err := query.Count(&count).Error; err != nil {
		panic(errors.Wrap(err, "find employees failed"))
	}
	if err := query.Find(&employees).Error; err != nil {
		panic(errors.Wrap(err, "find employees failed"))
	}
	return types.Page[*Employee]{
		List:      employees,
		PageIndex: opt.PageIndex,
		PageSize:  opt.PageSize,
		Total:     count,
	}
}

type EmployeeLoginOption struct {
	Account     string
	PhoneNumber string
	Email       string
}

func (uc *OrganizationUseCase) FindOneEmployeeByLoginOption(ctx context.Context, option *EmployeeLoginOption) (employee *Employee, err error) {
	if *option == (EmployeeLoginOption{}) {
		panic(errors.New("option empty"))
	}

	var queryEmployee Employee
	if option.Account != "" {
		queryEmployee.Account = option.Account
	}
	if option.Email != "" {
		queryEmployee.Email = option.Email
	}
	if option.PhoneNumber != "" {
		queryEmployee.MobilePhone = option.PhoneNumber
	}

	if err = uc.db.WithContext(ctx).Model(&Employee{}).Where(&queryEmployee).First(&employee).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.WithCause(errorx.ErrBadRequest, "用户不存在, 请检查登录信息")
		}
		panic(err)
	}
	return
}

func (uc *OrganizationUseCase) FindOneEmployeeById(ctx context.Context, id int64) (employee *Employee, err error) {
	if err = uc.db.WithContext(ctx).Where(id).Preload("Department").First(&employee).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.WithCause(errorx.ErrBadRequest, "用户不存在")
		}
		panic(err)
	}
	return
}

func (uc *OrganizationUseCase) UpdateEmployeeById(ctx context.Context, employee *Employee, employeeId int64) error {
	whereCase := Employee{
		Model: model.Model{
			ID: employeeId,
		},
		IsReserved: false,
	}
	result := uc.db.WithContext(ctx).Where(whereCase, "is_reserved").Updates(employee)
	if result.RowsAffected == 0 {
		return errorx.WithCause(errorx.ErrBadRequest, "更新失败, 用户保留或不存在")
	}
	err := result.Error
	if err != nil {
		panic(errors.Wrap(err, "delete employee failed"))
	}
	return nil
}

func (uc *OrganizationUseCase) DeleteEmployeeById(ctx context.Context, id int64) error {
	result := uc.db.WithContext(ctx).Where(Employee{IsReserved: false}, "is_reserved").Delete(&Employee{}, id)
	err := result.Error
	if err != nil {
		panic(errors.Wrap(err, "delete employee failed"))
	}
	if result.RowsAffected == 0 {
		return errorx.WithCause(errorx.ErrBadRequest, "删除失败")
	}
	return nil
}

func (uc *OrganizationUseCase) FindAllPositions(ctx context.Context) (positions []string) {
	err := uc.db.WithContext(ctx).Model(Employee{}).Pluck("position", &positions).Error
	if err != nil {
		panic(err)
	}
	return positions
}

func (uc *OrganizationUseCase) FindOneDepartment(ctx context.Context, id int64) (department *Department, err error) {
	department = &Department{}
	if err := uc.db.WithContext(ctx).Preload("Leader").First(department, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.WithCause(errorx.ErrBadRequest, "部门未找到")
		}
		panic(err)
	}
	return department, nil
}

func (uc *OrganizationUseCase) CreateDepartment(ctx context.Context, dep *Department) error {
	if dep.PId == 0 {
		return errorx.WithCause(errorx.ErrBadRequest, "必须指定父部门Id")
	}
	db := uc.db.WithContext(ctx)
	// 查询父节点
	var pDep *Department
	if err := db.Preload("Ancestors").Where(dep.PId).First(&pDep).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errorx.WithCause(errorx.ErrBadRequest, "父部门不存在")
		}
		panic(errors.Wrap(err, "query parent Dep failed"))
	}
	for _, ancestor := range pDep.Ancestors {
		dep.Ancestors = append(dep.Ancestors, ancestor)
	}
	dep.Ancestors = append(dep.Ancestors, pDep)

	if err := db.Create(dep).Error; err != nil {
		panic(errors.Wrap(err, "create dep failed"))
	}
	return nil
}

type FindManyDepartmentsOption struct {
	DepIds   []int64
	LikeName string
}

func (uc *OrganizationUseCase) FindManyDepartmentsPage(ctx context.Context, option *types.PageOption[FindManyDepartmentsOption]) *types.Page[*Department] {
	var deps []*Department
	var count int64
	query := uc.db.WithContext(ctx).Model(Department{})

	if len(option.Option.DepIds) > 0 {
		query.Where(option.Option.DepIds)
	}

	if option.Option.LikeName != "" {
		query.Where("name like ?", "%"+option.Option.LikeName+"%")
	}

	if err := query.Count(&count).Error; err != nil {
		panic(err)
	}
	if option.PageIndex != 0 && option.PageSize != 0 {
		query.Offset((option.PageIndex - 1) * option.PageSize).Limit(option.PageSize)
	}
	if err := query.Find(&deps).Error; err != nil {
		panic(errors.Wrap(err, "query deps failed"))
	}
	return &types.Page[*Department]{
		List:      deps,
		PageIndex: option.PageIndex,
		PageSize:  option.PageSize,
		Total:     count,
	}
}

func (uc *OrganizationUseCase) FindManyDepartmentsByRootId(ctx context.Context, rootId int64) (departments []*Department, err error) {
	departments = []*Department{}
	if err := uc.db.WithContext(ctx).Model(Department{}).Preload("Leader").Preload("Ancestors").
		Joins("LEFT JOIN department_ancestors ON departments.id = department_ancestors.department_id").
		Where("department_ancestors.ancestor_id = ?", rootId).Or("departments.id = ?", rootId).
		Find(&departments).Error; err != nil {
		panic(err)
	}
	if len(departments) == 0 {
		return nil, errorx.WithCause(errorx.ErrBadRequest, "根部门不存在")
	}
	return
}

func (uc *OrganizationUseCase) FindAllDepartments(ctx context.Context) (departments []*Department) {
	if err := uc.db.WithContext(ctx).Preload("Leader").Find(&departments).Error; err != nil {
		panic(err)
	}
	return
}

func (uc *OrganizationUseCase) CountEmployeeInDepartmentByIds(ctx context.Context, depIds []int64) (count int64) {
	if err := uc.db.WithContext(ctx).Model(Employee{}).Where("department_id in ?", depIds).Count(&count).Error; err != nil {
		panic(err)
	}
	return count
}

func (uc *OrganizationUseCase) PatchDepartmentById(ctx context.Context, id int64, dep *Department) error {
	result := uc.db.WithContext(ctx).Model(Department{}).Where(Department{IsReserved: false}, "is_reserved").
		Where("id = ?", id).Updates(&dep)
	if result.Error != nil {
		panic(result.Error)
	}
	if result.RowsAffected == 0 {
		return errorx.WithCause(errorx.ErrBadRequest, "更新失败")
	}
	return nil
}

func (uc *OrganizationUseCase) DeleteDepartmentById(ctx context.Context, id int64) error {
	db := uc.db.WithContext(ctx)
	queryWhere := db.Model(Department{}).
		Joins("LEFT JOIN department_ancestors ON departments.id = department_ancestors.department_id").
		Where("department_ancestors.ancestor_id = ?", id).Or("departments.id = ?", id).
		Select("id")
	result := db.Model(Department{}).Where(Department{IsReserved: false}, "is_reserved").
		Where("id in (?)", queryWhere).Delete(&Department{})
	if result.Error != nil {
		panic(result.Error)
	}
	if result.RowsAffected == 0 {
		return errorx.WithCause(errorx.ErrBadRequest, "删除失败")
	}
	return nil
}
