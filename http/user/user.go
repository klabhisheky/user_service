package user

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"developer.zopsmart.com/go/backend/zs/types"
	"github.com/klabhisheky/user_service/model"
	"github.com/klabhisheky/user_service/services"
)

type userServer struct {
	svc services.User
}

func New(usrsvc services.User) *userServer {
	return &userServer{svc: usrsvc}
}

func (usrsvr *userServer) Index(r *http.Request) (interface{}, error) {

	usrId, err := strconv.Atoi(r.FormValue("UserId"))
	if err != nil {
		return nil, types.ErrInvalidParam{[]string{"UserId"}}
	}

	//service call
	svcres, err := usrsvr.svc.Find(r.Context(), usrId)
	if err != nil {
		return nil, err
	}

	return svcres, err
}

func (usrsvr *userServer) Create(r *http.Request) (interface{}, error) {
	var usr model.User
	b, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(b, &usr)
	if err != nil {
		return nil, err
	}
	if (usr.Address == "") || (usr.Name == "") || (usr.Phone == "") {
		return nil, types.ErrMissingParam{"Name/Address/Phone"}
	}

	svcres, err := usrsvr.svc.Create(r.Context(), &usr)
	if err != nil {
		return nil, err
	}

	return svcres, nil

}
