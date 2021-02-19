package integration_test

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	sdk "github.com/irisnet/service-sdk-go"
	"github.com/irisnet/service-sdk-go/service"
	"github.com/irisnet/service-sdk-go/types"
	"github.com/irisnet/service-sdk-go/types/store"
)

const (
	nodeURI  = "tcp://127.0.0.1:26657"
	grpcAddr = "127.0.0.1:9090"
	chainID  = "test"
	charset  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	addr     = "iaa1s7cmgyu6xqergksdwq63fjhpekwrt7h2k4j4hu"
)

var (
	path = os.ExpandEnv("$HOME/.iriscli")
)

type IntegrationTestSuite struct {
	suite.Suite

	serviceClient service.ServiceClient

	r            *rand.Rand
	rootAccount  MockAccount
	randAccounts []MockAccount
}

// MockAccount define a account for test
type MockAccount struct {
	Name, Password string
	Address        types.AccAddress
}

func TestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(IntegrationTestSuite))
}

func (s *IntegrationTestSuite) SetupSuite() {
	options := []types.Option{
		types.KeyDAOOption(store.NewMemory(nil)),
		types.TimeoutOption(6),
		types.CachedOption(true),
	}
	cfg, err := types.NewClientConfig(nodeURI, grpcAddr, chainID, options...)
	if err != nil {
		panic(err)
	}

	s.serviceClient = sdk.NewServiceClient(cfg)

	s.r = rand.New(rand.NewSource(time.Now().UnixNano()))
	s.rootAccount = MockAccount{
		Name:     "v1",
		Password: "1234567890",
		Address:  types.MustAccAddressFromBech32(addr),
	}

	s.initAccount()
}

func (s *IntegrationTestSuite) TearDownSuite() {
	_ = os.Remove(path)
}

func (s *IntegrationTestSuite) initAccount() {
	_, err := s.serviceClient.Import(s.Account().Name,
		s.Account().Password,
		string(getPrivKeyArmor()))
	if err != nil {
		panic(err)
	}

	//var receipts bank.Receipts
	for i := 0; i < 5; i++ {
		name := s.RandStringOfLength(10)
		pwd := s.RandStringOfLength(16)
		address, _, err := s.serviceClient.Insert(name, "11111111")
		if err != nil {
			panic(fmt.Sprintf("generate test account failed, err: %s", err.Error()))
		}

		s.randAccounts = append(s.randAccounts, MockAccount{
			Name:     name,
			Password: pwd,
			Address:  types.MustAccAddressFromBech32(address),
		})
	}
}

// RandStringOfLength return a random string
func (s *IntegrationTestSuite) RandStringOfLength(l int) string {
	var result []byte
	bytes := []byte(charset)
	for i := 0; i < l; i++ {
		result = append(result, bytes[s.r.Intn(len(bytes))])
	}
	return string(result)
}

// GetRandAccount return a random test account
func (s *IntegrationTestSuite) GetRandAccount() MockAccount {
	return s.randAccounts[s.r.Intn(len(s.randAccounts))]
}

// Account return a test account
func (s *IntegrationTestSuite) Account() MockAccount {
	return s.rootAccount
}

func getPrivKeyArmor() []byte {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path = filepath.Dir(path)
	path = filepath.Join(path, "integration_test/scripts/priv.key")
	bz, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return bz
}
