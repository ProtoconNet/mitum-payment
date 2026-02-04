package digest

import (
	"net/http"

	cdigest "github.com/ProtoconNet/mitum-currency/v3/digest"
	ctypes "github.com/ProtoconNet/mitum-currency/v3/types"
	"github.com/ProtoconNet/mitum-payment/types"
	"github.com/ProtoconNet/mitum2/base"
)

var (
	HandlerPathPaymentDesign      = `/payment/{contract:(?i)` + ctypes.REStringAddressString + `}`
	HandlerPathPaymentAccountInfo = `/payment/{contract:(?i)` + ctypes.REStringAddressString + `}/account/{address:(?i)` + ctypes.REStringAddressString + `}`
)

func SetHandlers(hd *cdigest.Handlers) {
	get := 1000
	_ = hd.SetHandler(HandlerPathPaymentAccountInfo, HandlePaymentAccountInfo, true, get, get).
		Methods(http.MethodOptions, "GET")
	_ = hd.SetHandler(HandlerPathPaymentDesign, HandlePaymentDesign, true, get, get).
		Methods(http.MethodOptions, "GET")
}

func HandlePaymentDesign(hd *cdigest.Handlers, w http.ResponseWriter, r *http.Request) {
	cacheKey := cdigest.CacheKeyPath(r)
	if err := cdigest.LoadFromCache(hd.Cache(), cacheKey, w); err == nil {
		return
	}

	contract, err, status := cdigest.ParseRequest(w, r, "contract")
	if err != nil {
		cdigest.HTTP2ProblemWithError(w, err, status)

		return
	}

	if v, err, shared := hd.RG().Do(cacheKey, func() (interface{}, error) {
		return handlePaymentDesignInGroup(hd, contract)
	}); err != nil {
		cdigest.HTTP2HandleError(w, err)
	} else {
		cdigest.HTTP2WriteHalBytes(hd.Encoder(), w, v.([]byte), http.StatusOK)

		if !shared {
			cdigest.HTTP2WriteCache(w, cacheKey, hd.ExpireShortLived())
		}
	}
}

func handlePaymentDesignInGroup(hd *cdigest.Handlers, contract string) ([]byte, error) {
	var de *types.Design
	var st base.State

	de, st, err := PaymentDesign(hd.Database(), contract)
	if err != nil {
		return nil, err
	}

	i, err := buildPaymentDesign(hd, contract, *de, st)
	if err != nil {
		return nil, err
	}
	return hd.Encoder().Marshal(i)
}

func buildPaymentDesign(hd *cdigest.Handlers, contract string, de types.Design, st base.State) (cdigest.Hal, error) {
	h, err := hd.CombineURL(HandlerPathPaymentDesign, "contract", contract)
	if err != nil {
		return nil, err
	}

	var hal cdigest.Hal
	hal = cdigest.NewBaseHal(de, cdigest.NewHalLink(h, nil))

	h, err = hd.CombineURL(cdigest.HandlerPathBlockByHeight, "height", st.Height().String())
	if err != nil {
		return nil, err
	}
	hal = hal.AddLink("block", cdigest.NewHalLink(h, nil))

	for i := range st.Operations() {
		h, err := hd.CombineURL(cdigest.HandlerPathOperation, "hash", st.Operations()[i].String())
		if err != nil {
			return nil, err
		}
		hal = hal.AddLink("operations", cdigest.NewHalLink(h, nil))
	}

	return hal, nil
}

func HandlePaymentAccountInfo(hd *cdigest.Handlers, w http.ResponseWriter, r *http.Request) {
	cachekey := cdigest.CacheKeyPath(r)
	if err := cdigest.LoadFromCache(hd.Cache(), cachekey, w); err == nil {
		return
	}

	contract, err, status := cdigest.ParseRequest(w, r, "contract")
	if err != nil {
		cdigest.HTTP2ProblemWithError(w, err, status)

		return
	}

	account, err, status := cdigest.ParseRequest(w, r, "address")
	if err != nil {
		cdigest.HTTP2ProblemWithError(w, err, status)

		return
	}

	if v, err, shared := hd.RG().Do(cachekey, func() (interface{}, error) {
		return handlePaymentAccountInfoInGroup(hd, contract, account)
	}); err != nil {
		cdigest.HTTP2HandleError(w, err)
	} else {
		cdigest.HTTP2WriteHalBytes(hd.Encoder(), w, v.([]byte), http.StatusOK)

		if !shared {
			cdigest.HTTP2WriteCache(w, cachekey, hd.ExpireShortLived())
		}
	}
}

func handlePaymentAccountInfoInGroup(hd *cdigest.Handlers, contract, account string) ([]byte, error) {
	var accountInfoValue *AccountInfoValue

	accountInfoValue, err := AccountInfo(hd.Database(), contract, account)
	if err != nil {
		return nil, err
	}

	i, err := buildAccountInfoValue(hd, contract, *accountInfoValue)
	if err != nil {
		return nil, err
	}
	return hd.Encoder().Marshal(i)
}

func buildAccountInfoValue(hd *cdigest.Handlers, contract string, it AccountInfoValue) (cdigest.Hal, error) {
	h, err := hd.CombineURL(
		HandlerPathPaymentAccountInfo,
		"contract", contract, "address", it.setting.Address().String(),
	)
	if err != nil {
		return nil, err
	}

	var hal cdigest.Hal
	hal = cdigest.NewBaseHal(it, cdigest.NewHalLink(h, nil))

	return hal, nil
}
