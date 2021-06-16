package accountsvc

var (
	service Service
)

type Service interface {
}

func init() {
	service = newServiceImpl()
}

func Inject() Service {
	return service
}
