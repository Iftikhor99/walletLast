package wallet

import (
	"fmt"
	"github.com/google/uuid"
	"reflect"
	"github.com/Iftikhor99/wallet/v1/pkg/types"
	"testing"
)

var defaultTestAccount = testAccount{
	phone: "+992000000001",
	balance: 10_000_00,
	payments: []struct {
		amount types.Money
		category types.PaymentCategory
	}{
		{amount: 1_000_00, category: "auto"},
	},
}

type testAccount struct {
	phone types.Phone
	balance types.Money
	payments []struct {
		amount types.Money
		category types.PaymentCategory
	}
}

type testService struct {
	*Service
}

func newTestService() *testService {
	return &testService{Service: &Service{}}
}


func TestService_FindPaymentByID_success(t *testing.T) {

	s := newTestService()
	_, payments, err := s.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
	}
		
	// Tpo6yem HavtTu nnaTéx
		
	payment := payments[0]
		
	got, err := s.FindPaymentByID(payment. ID)
		
	if err != nil {
		t.Errorf("FindPaymentByID(): error = %v", err)
		return
	}
		
	// CpaBHMBaem nnaTexu
	if !reflect.DeepEqual(payment, got) {
		t.Errorf("FindPaymentByID(): wrong payment returned = %v", err)
		return
	}
}

func TestService_FindPaymentByID_fail(t *testing.T) {
	// co3paém cepsuc
	s := newTestService()
	_, _, err := s.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
	}
	
	// TipoOyem HaWTM HeECyWeECTByWuMA nnaTéex
	_, err = s.FindPaymentByID(uuid.New().String())
	if err == nil {
		t.Error("FindPaymentByID(): must return error, returned nil")
		return
	}
	
	if err != ErrPaymentNotFound {
		t.Errorf("FindPaymentByID(): must return ErrPaymentNotFound, returned = %v", err)
		return
	}
	
}

func TestService_Reject_success(t *testing.T) {

// co3paém cepsuc
s := newTestService()

_, payments, err := s.addAccount(defaultTestAccount)

if err != nil {
	t.Error(err)
	return
}

// TipoOyem OTMeHMTb nnaTéx

payment := payments[0]

err = s.Reject(payment. ID)

if err != nil {
	t.Errorf("Reject(): error = %v", err)
	return
}

savedPayment, err := s.FindPaymentByID(payment. ID)
if err != nil {
	t.Errorf("Reject(): can't find payment by id, error = %v", err)
	return
}
if savedPayment.Status != types.PaymentStatusFail {
	t.Errorf("Reject(): status didn't changed, payment = %v", savedPayment)
	return
}

savedAccount, err := s.FindAccountByID(payment.AccountID)

if err != nil {
	t.Errorf("Reject(): can't find account by id, error = %v", err)
	return
}

if savedAccount.Balance != defaultTestAccount.balance {
	t.Errorf("Reject(): balance didn't changed, account = %v", savedAccount)
	return
}

}

func TestService_Repeat_success(t *testing.T) {

	s := newTestService()
	_, payments, err := s.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
	}
		
	// Tpo6yem HavtTu nnaTéx
		
	payment := payments[0]
		
	payment, err = s.Repeat(payment.ID)
		
	if err != nil {
		t.Errorf("Repeat(): error = %v", err)
		return
	}
		
}

func TestService_PayFromFavorite_success(t *testing.T) {

	s := newTestService()
	_, payments, err := s.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
	}
		
	// Tpo6yem HavtTu nnaTéx
		
	payment := payments[0]

	favorite, err := s.FavoritePayment(payment.ID, "Tcell")
		
	payment, err = s.PayFromFavorite(favorite.ID)
		
	if err != nil {
		t.Errorf("PayFromFavorite(): error = %v", err)
		return
	}
		
}


func (s *testService) addAccount(data testAccount) (*types.Account, []*types.Payment, error) {
	// perucTpupyemM TaM nonb30BaTena
	account, err := s.RegisterAccount (data.phone)
	if err != nil {
		return nil, nil, fmt.Errorf("can't register account, error = %v", err)
	}
	
	// MononHsem ero cyéT
	err = s.Deposit(account.ID, data.balance)
	if err != nil {
		return nil, nil, fmt.Errorf("can't deposity account, error = %v", err)
	}
	
	// BeinonHaem nnaTexu
	// MOKeM CO3MaTb CNavc Cpa3y HYKHOM ONMHbI, NOCKONbKy 3HAaeM Ppa3sMep
	payments := make([]*types.Payment, len(data.payments) )
	for i, payment := range data.payments {
		// Torga 30€Ccb padoTaem npocto yepes index, a He Yepe3 append
		payments[i], err = s.Pay(account.ID, payment.amount, payment.category)
		if err != nil {
		return nil, nil, fmt.Errorf("can't make payment, error = %v", err)
		}
	}
	
	return account, payments, nil
}