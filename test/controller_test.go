package test

import (
	"bytes"
	"encoding/json"
	"fintechpractices/global"
	Init "fintechpractices/init"
	"fintechpractices/internal/schema"
	"fintechpractices/tools"
	"fmt"
	"io"
	"net/http"
	"testing"
)

var prefix string

func Start() {
	Init.Initialization()
	prefix = fmt.Sprintf("http://%s:%d", global.AppCfg.ServerCfg.Host, global.AppCfg.ServerCfg.Port)
	go func() {
		if err := global.Engine.Run(fmt.Sprintf("%s:%d", global.AppCfg.ServerCfg.Host, global.AppCfg.ServerCfg.Port)); err != nil {
			panic(err.Error())
		}
	}()
}

func TestPublickey(t *testing.T) {
	Start()

	req, err := http.NewRequest(
		http.MethodGet,
		prefix+"/pubkey",
		nil,
	)
	if err != nil {
		t.Error(err.Error())
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err.Error())
	}

	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	fmt.Println(string(b))

}

func TestRegister(t *testing.T) {
	Start()

	bSlice, err := tools.Encrypt("ccpractices")
	if err != nil {
		t.Error(err.Error())
	}

	body, err := json.Marshal(schema.RegisterReq{
		UserName:    "test_admin",
		UserAccount: "test_admin",
		DecryptData: fmt.Sprintf("%x", bSlice),
	})
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Printf("%s", string(body))
	req, err := http.NewRequest(
		http.MethodPost,
		prefix+"/register",
		bytes.NewReader(body),
	)
	if err != nil {
		t.Error(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err.Error())
	}

	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	fmt.Println(string(b))
}

func requestForLogin(t *testing.T) (token string) {
	bSlice, err := tools.Encrypt("ccpractices")
	if err != nil {
		t.Error(err.Error())
	}
	body, err := json.Marshal(schema.AuthReq{
		UserAccount: "test_admin",
		DecryptData: fmt.Sprintf("%x", bSlice),
	})

	if err != nil {
		t.Error(err.Error())
	}

	req, err := http.NewRequest(
		http.MethodPost,
		prefix+"/login",
		bytes.NewReader(body),
	)
	if err != nil {
		t.Error(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err.Error())
	}
	defer resp.Body.Close()

	b, _ := io.ReadAll(resp.Body)

	dict := map[string]interface{}{}
	_ = json.Unmarshal(b, &dict)
	return dict["token"].(string)
}

func TestLogin(t *testing.T) {
	Start()

	token := requestForLogin(t)

	claims, err := tools.ParseToken(token)
	if err != nil {
		t.Error(err.Error())
	}
	if claims.UserAccount != "test_admin" {
		t.Error("unexpect user account")
	}
	fmt.Printf("%#v\n", *claims)
}

func TestGetDp(t *testing.T) {
	Start()

	token := requestForLogin(t)
	req, err := http.NewRequest(
		http.MethodGet,
		prefix+"/dp?page_no=1&page_size=5",
		nil,
	)
	if err != nil {
		t.Error(err.Error())
	}
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err.Error())
	}
	defer resp.Body.Close()

	b, _ := io.ReadAll(resp.Body)

	res := schema.GetDpResp{}
	json.Unmarshal(b, &res)
	if res.Code != 2000 || len(res.Data) > 5 {
		t.Error("unexpected resp")
	}

	fmt.Printf("%v\n", res)
}

func TestGetResource(t *testing.T) {
	Start()

	token := requestForLogin(t)
	{
		req, err := http.NewRequest(
			http.MethodGet,
			prefix+"/resource/tone?page_no=1&page_size=2",
			nil,
		)
		if err != nil {
			t.Error(err.Error())
		}
		req.Header.Set("Authorization", "Bearer "+token)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Error(err.Error())
		}
		defer resp.Body.Close()

		b, _ := io.ReadAll(resp.Body)

		res := schema.GetResourceResp{}
		json.Unmarshal(b, &res)
		if res.Code != 2000 || len(res.Data) != 2 {
			t.Error("unexpected resp")
		}

		fmt.Printf("%#v\n", res)
	}

	{
		req, err := http.NewRequest(
			http.MethodGet,
			prefix+"/resource/image?page_no=1&page_size=1",
			nil,
		)
		if err != nil {
			t.Error(err.Error())
		}
		req.Header.Set("Authorization", "Bearer "+token)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Error(err.Error())
		}
		defer resp.Body.Close()

		b, _ := io.ReadAll(resp.Body)

		res := schema.GetResourceResp{}
		json.Unmarshal(b, &res)
		if res.Code != 2000 || len(res.Data) != 1 {
			t.Error("unexpected resp")
		}

		fmt.Printf("%#v\n", res)
	}

}

func TestDeleteDp(t *testing.T) {
	Start()

	token := requestForLogin(t)
	{
		req, err := http.NewRequest(
			http.MethodDelete,
			prefix+"/dp/test_dp_3.mp4",
			nil,
		)
		req.Header.Set("Authorization", "Bearer "+token)
		if err != nil {
			t.Error(err.Error())
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Error(err.Error())
		}
		defer resp.Body.Close()

		b, _ := io.ReadAll(resp.Body)
		res := schema.CommResp{}

		json.Unmarshal(b, &res)
		if res.Code != 2000 {
			t.Error("unexpect resp")
		}
		fmt.Printf("%#v\n", res)
	}

	{
		req, err := http.NewRequest(
			http.MethodDelete,
			prefix+"/dp/test_dp_4.mp4",
			nil,
		)
		req.Header.Set("Authorization", "Bearer "+token)
		if err != nil {
			t.Error(err.Error())
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Error(err.Error())
		}
		defer resp.Body.Close()

		b, _ := io.ReadAll(resp.Body)
		res := schema.CommResp{}

		json.Unmarshal(b, &res)
		if res.Code != 2000 {
			t.Error("unexpect resp")
		}
		fmt.Printf("%#v\n", res)
	}

	{
		req, err := http.NewRequest(
			http.MethodDelete,
			prefix+"/dp/test_dp_2.mp4",
			nil,
		)
		req.Header.Set("Authorization", "Bearer "+token)
		if err != nil {
			t.Error(err.Error())
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Error(err.Error())
		}
		defer resp.Body.Close()

		b, _ := io.ReadAll(resp.Body)
		res := schema.CommResp{}

		json.Unmarshal(b, &res)
		if res.Code != 2003 {
			t.Error("unexpect resp")
		}
		fmt.Printf("%#v\n", res)
	}
}

func TestDeleteResource(t *testing.T) {
	Start()

	token := requestForLogin(t)
	{
		req, err := http.NewRequest(
			http.MethodDelete,
			prefix+"/resource/test_image_1.png",
			nil,
		)
		req.Header.Set("Authorization", "Bearer "+token)
		if err != nil {
			t.Error(err.Error())
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Error(err.Error())
		}
		defer resp.Body.Close()

		b, _ := io.ReadAll(resp.Body)
		res := schema.CommResp{}

		json.Unmarshal(b, &res)
		if res.Code != 2000 {
			t.Error("unexpect resp")
		}
		fmt.Printf("%#v\n", res)
	}

	{
		req, err := http.NewRequest(
			http.MethodDelete,
			prefix+"/resource/test_tone_1.wav",
			nil,
		)
		req.Header.Set("Authorization", "Bearer "+token)
		if err != nil {
			t.Error(err.Error())
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Error(err.Error())
		}
		defer resp.Body.Close()

		b, _ := io.ReadAll(resp.Body)
		res := schema.CommResp{}

		json.Unmarshal(b, &res)
		if res.Code != 2000 {
			t.Error("unexpect resp")
		}
		fmt.Printf("%#v\n", res)
	}

	{
		req, err := http.NewRequest(
			http.MethodDelete,
			prefix+"/resource/test_image_111.png",
			nil,
		)
		req.Header.Set("Authorization", "Bearer "+token)
		if err != nil {
			t.Error(err.Error())
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Error(err.Error())
		}
		defer resp.Body.Close()

		b, _ := io.ReadAll(resp.Body)
		res := schema.CommResp{}

		json.Unmarshal(b, &res)
		if res.Code != 2003 {
			t.Error("unexpect resp")
		}
		fmt.Printf("%#v\n", res)
	}
}
