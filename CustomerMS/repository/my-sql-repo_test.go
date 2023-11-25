package repository_test

import (
	dbinit "BillingGo/db"
	"BillingGo/models"
	"BillingGo/repository"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	customer     models.Customer
	mockCustomer = models.Customer{
		CustomerId:    "2",
		CustomerName:  "Diku",
		ContactNumber: "345-345-345",
		Address:       "Ahmednagar",
		Priority:      "high",
	}
	cust *models.Customer
)

func TestCreateService(t *testing.T) {

	db, mock := dbinit.NewMockDB()

	repo := &repository.MySQLRepository{db}
	defer func() {
		logrus.Print("Im closing db connection for CREATE")
		repo.Close()
	}()
	//rows := sqlmock.NewRows([]string{"CustomerId", "CustomerName", "ContactNumber", "Address", "Priority"}).AddRow(mockCustomer.CustomerId, mockCustomer.CustomerName, mockCustomer.ContactNumber, mockCustomer.Address, mockCustomer.Priority)
	query := `INSERT\s.*INTO\s.*customers.*`
	mock.ExpectBegin()
	//mock.ExpectExec(query).WithArgs(mockCustomer.CustomerId, mockCustomer.CustomerName, mockCustomer.ContactNumber, mockCustomer.Address, mockCustomer.Priority).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
	//mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(mockCustomer.CustomerId, mockCustomer.CustomerName, mockCustomer.ContactNumber, mockCustomer.Address, mockCustomer.Priority, sqlmock.AnyArg, sqlmock.AnyArg, sqlmock.AnyArg).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	resultcustomer, err := repo.CreateCutomer(&mockCustomer)

	//assertions for failing
	assert.NoError(t, err)
	assert.Nil(t, err)
	assert.Equal(t, mockCustomer.CustomerId, resultcustomer.CustomerId)
	assert.Equal(t, mockCustomer.CustomerName, resultcustomer.CustomerName)
	assert.Equal(t, mockCustomer.ContactNumber, resultcustomer.ContactNumber)
	assert.Equal(t, mockCustomer.Address, resultcustomer.Address)
	assert.Equal(t, mockCustomer.Priority, resultcustomer.Priority)

}

func TestGetById(t *testing.T) {

	db, mock := dbinit.NewMockDB()

	repo := &repository.MySQLRepository{db}

	rows := sqlmock.NewRows([]string{"CustomerId", "CustomerName", "ContactNumber", "Address", "Priority"}).AddRow(mockCustomer.CustomerId, mockCustomer.CustomerName, mockCustomer.ContactNumber, mockCustomer.Address, mockCustomer.Priority)

	query := `SELECT\s.*FROM\s.customers. WHERE customer_id= \\?`
	mock.ExpectQuery(query).WillReturnRows(rows)

	resultcustomer, err := repo.GetCutomerById(mockCustomer.CustomerId)

	defer func() {
		repo.Close()
	}()
	assert.NotNil(t, resultcustomer)
	assert.NoError(t, err)
	assert.Equal(t, mockCustomer.CustomerId, resultcustomer.CustomerId)
	assert.Equal(t, mockCustomer.CustomerName, resultcustomer.CustomerName)
	assert.Equal(t, mockCustomer.ContactNumber, resultcustomer.ContactNumber)
	assert.Equal(t, mockCustomer.Address, resultcustomer.Address)
	assert.Equal(t, mockCustomer.Priority, resultcustomer.Priority)
	assert.Nil(t, err)
}

func TestGetByIdFail(t *testing.T) {

	db, mock := dbinit.NewMockDB()

	repo := &repository.MySQLRepository{db}

	rows := sqlmock.NewRows([]string{"CustomerId", "CustomerName", "ContactNumber", "Address", "Priority"})

	query := `SELECT\s.*FROM\s.customers. WHERE customer_id= \\?`
	mock.ExpectQuery(query).WillReturnRows(rows)

	result, err := repo.GetCutomerById("1")
	logrus.Printf("REsulet: %v", result)
	defer func() {
		repo.Close()
	}()
	assert.Error(t, err)

	assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
	// assert.Empty(t, result.CustomerId)
	// assert.Empty(t, result.CustomerName)
	// assert.Empty(t, result.ContactNumber)
	// assert.Empty(t, result.Address)
	// assert.Empty(t, result.Priority)
	// assert.Empty(t, result.CreatedAt)
	// assert.Empty(t, result.UpdatedAt)
	// assert.Empty(t, result.DeletedAt)
	// assert.Empty(t, result)
}
func TestGetAll(t *testing.T) {
	var customers []*models.Customer
	db, mock := dbinit.NewMockDB()

	repo := &repository.MySQLRepository{db}
	defer func() {
		repo.Close()
	}()
	rows := sqlmock.NewRows([]string{"CustomerId", "CustomerName", "ContactNumber", "Address", "Priority"}).AddRow(mockCustomer.CustomerId, mockCustomer.CustomerName, mockCustomer.ContactNumber, mockCustomer.Address, mockCustomer.Priority)

	query := `SELECT\s.*FROM\s.*customers.*`
	mock.ExpectQuery(query).WillReturnRows(rows)

	customers, err := repo.GetAllCutomer()

	assert.NotNil(t, customers)
	assert.NoError(t, err)
	assert.Nil(t, err)
	assert.Len(t, customers, 1)
	assert.Equal(t, mockCustomer.CustomerId, customers[0].CustomerId)
	assert.Equal(t, mockCustomer.CustomerName, customers[0].CustomerName)
	assert.Equal(t, mockCustomer.ContactNumber, customers[0].ContactNumber)
	assert.Equal(t, mockCustomer.Address, customers[0].Address)
	assert.Equal(t, mockCustomer.Priority, customers[0].Priority)
}

func TestUpdate(t *testing.T) {

	db, mock := dbinit.NewMockDB()

	repo := &repository.MySQLRepository{db}

	defer func() {
		repo.Close()
	}()

	//rows := sqlmock.NewRows([]string{"CustomerId", "CustomerName", "ContactNumber", "Address", "Priority"}).AddRow(mockCustomer.CustomerId, mockCustomer.CustomerName, mockCustomer.ContactNumber, mockCustomer.Address, mockCustomer.Priority)
	query := `UPDATE\s.* customers SET.*`

	mock.ExpectBegin()
	//mock.ExpectExec(query).WithArgs(mockCustomer.CustomerId, mockCustomer.CustomerName, mockCustomer.ContactNumber, mockCustomer.Address, mockCustomer.Priority).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(0, 1))
	//mock.ExpectQuery(query).WithArgs(mockCustomer.CustomerId, mockCustomer.CustomerName, mockCustomer.ContactNumber, mockCustomer.Address, mockCustomer.Priority).WillReturnRows(rows)
	//mock.ExpectQuery(query).WithArgs(mockCustomer.CustomerId, mockCustomer.CustomerName, mockCustomer.ContactNumber, mockCustomer.Address, mockCustomer.Priority).WillReturnRows(rows)
	mock.ExpectCommit()

	resultcustomer, err := repo.UpdateCutomer(&mockCustomer)

	//assertions for failing
	logrus.Print(resultcustomer)
	logrus.Print(err)
	//assert.NoError(t, err)
	//assert.Nil(t, err)

}
func TestDelete(t *testing.T) {

	db, mock := dbinit.NewMockDB()

	repo := &repository.MySQLRepository{db}

	defer func() {
		repo.Close()
	}()

	query := `DELETE\s.*FROM\s.customers. WHERE customer_id = \?`
	mock.ExpectBegin()
	// mock.ExpectQuery(query1).WillReturnRows(rows)
	mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	resultcustomer, err := repo.DeleteCutomer(mockCustomer.CustomerId)

	//assertions for failing
	logrus.Print(resultcustomer)
	logrus.Print(err)
	//assert.NoError(t, err)
	//assert.Nil(t, err)

}
func TestDeleteFail(t *testing.T) {

	db, mock := dbinit.NewMockDB()

	repo := &repository.MySQLRepository{db}

	defer func() {
		repo.Close()
	}()

	query := `DELETE\s.*FROM\s.customers. WHERE customer_id = \?`
	mock.ExpectBegin()
	// mock.ExpectQuery(query1).WillReturnRows(rows)
	mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	resultcustomer, err := repo.DeleteCutomer("5")

	//assertions for failing
	assert.Error(t, err)

	//assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
	assert.Empty(t, resultcustomer) //-> not empty

}

func TestMySQLRepository_UpdateCutomer(t *testing.T) {
	//db, _ := dbinit.NewMockDB()

	type args struct {
		model *models.Customer
	}
	tests := []struct {
		name    string
		m       *repository.MySQLRepository
		args    args
		want    *models.Customer
		wantErr bool
	}{
		// TODO: Add test cases.
		// {

		// 	name:    "Update Case1",
		// 	m:       &repository.MySQLRepository{db},
		// 	args:    args{model: &mockCustomer},
		// 	want:    cust,
		// 	wantErr: false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.UpdateCutomer(tt.args.model)
			if (err != nil) != tt.wantErr {
				t.Errorf("MySQLRepository.UpdateCutomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MySQLRepository.UpdateCutomer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMySQLRepository_DeleteCutomer(t *testing.T) {
	db, _ := dbinit.NewMockDB()
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		m       *repository.MySQLRepository
		args    args
		want    *models.Customer
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "DELETE Fails",
			m:       &repository.MySQLRepository{db},
			args:    args{id: "5"},
			want:    cust,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.DeleteCutomer(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("MySQLRepository.DeleteCutomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MySQLRepository.DeleteCutomer() = %v, want %v", got, tt.want)
			}
		})
	}
}
