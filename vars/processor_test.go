package vars

import (
	"testing"

	"strconv"

	"github.com/jmartin82/mmock/definition"
)

//DummyDataFaker is used in tests
type DummyDataFaker struct {
	Dummy string
}

func NewDummyDataFaker(dummyString string) DummyDataFaker {
	result := DummyDataFaker{Dummy: dummyString}
	return result
}

func (ddf DummyDataFaker) Brand() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) Character() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) Characters() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) CharactersN(n int) string {
	return ddf.Dummy + strconv.Itoa(n)
}
func (ddf DummyDataFaker) City() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) Color() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) Company() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) Continent() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) Country() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) CreditCardVisa() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) CreditCardMasterCard() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) CreditCardAmericanExpress() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) Currency() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) CurrencyCode() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) Day() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) Digits() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) DigitsN(n int) string {
	return ddf.Dummy + strconv.Itoa(n)
}
func (ddf DummyDataFaker) EmailAddress() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) FirstName() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) FullName() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) LastName() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) Gender() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) IPv4() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) Language() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) Model() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) Month() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) Year() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) MonthShort() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) Paragraph() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) Paragraphs() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) ParagraphsN(n int) string {
	return ddf.Dummy + strconv.Itoa(n)
}
func (ddf DummyDataFaker) Phone() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) Product() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) Sentence() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) Sentences() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) SentencesN(n int) string {
	return ddf.Dummy + strconv.Itoa(n)
}
func (ddf DummyDataFaker) SimplePassword() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) State() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) StateAbbrev() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) Street() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) StreetAddress() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) UserName() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) WeekDay() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) Word() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) Words() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) WordsN(n int) string {
	return ddf.Dummy + strconv.Itoa(n)
}
func (ddf DummyDataFaker) Zip() string {
	return ddf.Dummy
}
func (ddf DummyDataFaker) Int(n int) string {
	return ddf.Dummy + strconv.Itoa(n)
}
func (ddf DummyDataFaker) Float(n int) string {
	return ddf.Dummy + strconv.Itoa(n)
}
func (ddf DummyDataFaker) UUID() string {
	return "00000000-0000-0000-0000-000000000000"
}

func getProcessor() Processor {
	return Processor{FillerFactory: MockFillerFactory{FakeAdapter: NewDummyDataFaker("AleixMG")}}
}

func TestReplaceTags(t *testing.T) {

	req := definition.Request{}
	req.Body = "hi!"
	val := make(definition.Values)
	val["param1"] = []string{"valParam"}
	req.QueryStringParameters = val

	cookie := make(definition.Cookies)
	cookie["cookie1"] = "valCookie"
	req.Cookies = cookie

	res := definition.Response{}
	res.Body = "Request Body {{request.body}}. Query {{request.query.param1}}. Cookie: {{request.cookie.cookie1}}. Random: {{fake.UserName}}"

	cookie = make(definition.Cookies)
	cookie["cookie1"] = "valCookie"
	cookie["cookie2"] = "{{fake.UserName}}"
	res.Cookies = cookie

	val = make(definition.Values)
	val["header1"] = []string{"valHeader"}
	val["header2"] = []string{"valHeader", "{{request.query.param1}}"}

	res.Headers = val

	mock := definition.Mock{Request: req, Response: res}
	varsProcessor := getProcessor()
	varsProcessor.Eval(&req, &mock)

	if mock.Response.Body != "Request Body hi!. Query valParam. Cookie: valCookie. Random: AleixMG" {
		t.Error("Replaced tags in body not match", res.Body)
	}

	if mock.Response.Cookies["cookie2"] != "AleixMG" {
		t.Error("Replaced tags in cookie match", mock.Response.Cookies["cookie2"])
	}

	if mock.Response.Headers["header2"][1] != "valParam" {
		t.Error("Replaced tags in headers match", mock.Response.Headers["header2"][1])
	}
}

func TestReplaceUndefinedFakeTag(t *testing.T) {
	req := definition.Request{}
	val := make(definition.Values)
	val["param1"] = []string{"valParam"}
	req.QueryStringParameters = val

	cookie := make(definition.Cookies)
	cookie["cookie1"] = "valCookie"
	req.Cookies = cookie

	res := definition.Response{}
	res.Body = "Request {{request.query.param2}}. Cookie: {{request.cookie.cookie2}}. Random: {{fake.otherOption}}"

	mock := definition.Mock{Request: req, Response: res}
	varsProcessor := getProcessor()
	varsProcessor.Eval(&req, &mock)

	if mock.Response.Body != "Request {{request.query.param2}}. Cookie: {{request.cookie.cookie2}}. Random: {{fake.otherOption}}" {
		t.Error("Replaced tags in body not match", mock.Response.Body)
	}

}

func TestReplaceTagWithSpace(t *testing.T) {
	req := definition.Request{}
	val := make(definition.Values)
	val["param1"] = []string{"valParam"}
	req.QueryStringParameters = val

	cookie := make(definition.Cookies)
	cookie["cookie1"] = "valCookie"
	req.Cookies = cookie

	res := definition.Response{}
	res.Body = "Request {{ request.query.param1}}. Cookie: {{request.cookie.cookie1 }}. Random: {{fake.UserName }}"

	mock := definition.Mock{Request: req, Response: res}
	varsProcessor := getProcessor()
	varsProcessor.Eval(&req, &mock)

	if mock.Response.Body != "Request valParam. Cookie: valCookie. Random: AleixMG" {
		t.Error("Replaced tags in body not match", mock.Response.Body)
	}
}

func TestReplaceUrlPathVars(t *testing.T) {

	mockReq := definition.Request{}
	mockReq.Path = "/users/:id"
	res := definition.Response{}
	res.Body = "{ \"id\": {{request.path.id}} }"

	mock := definition.Mock{Request: mockReq, Response: res}
	varsProcessor := getProcessor()

	req := definition.Request{}
	req.Path = "/users/15"
	varsProcessor.Eval(&req, &mock)

	if mock.Response.Body != "{ \"id\": 15 }" {
		t.Error("Replaced url param in body not match", mock.Response.Body)
	}
}

func TestReplaceTagWithParameter(t *testing.T) {
	req := definition.Request{}

	res := definition.Response{}
	res.Body = "Random: {{fake.CharactersN(15)}}"

	mock := definition.Mock{Request: req, Response: res}
	varsProcessor := getProcessor()
	varsProcessor.Eval(&req, &mock)

	if mock.Response.Body != "Random: AleixMG15" {
		t.Error("Replaced tags in body not match", mock.Response.Body)
	}
}

func TestReplaceTagWithParameterNoParameterPassed(t *testing.T) {
	req := definition.Request{}

	res := definition.Response{}
	res.Body = "Random: {{fake.CharactersN}}"

	mock := definition.Mock{Request: req, Response: res}
	varsProcessor := getProcessor()
	varsProcessor.Eval(&req, &mock)

	if mock.Response.Body != "Random: {{fake.CharactersN}}" {
		t.Error("Replaced tags in body not match", mock.Response.Body)
	}
}

func TestReplaceMissingTags(t *testing.T) {
	req := definition.Request{}

	res := definition.Response{}
	res.Body = "Request Body {{request.body}}. Query {{request.query.param1}}. Cookie: {{request.cookie.cookie1}}."

	mock := definition.Mock{Request: req, Response: res}
	varsProcessor := getProcessor()
	varsProcessor.Eval(&req, &mock)

	if mock.Response.Body != "Request Body {{request.body}}. Query {{request.query.param1}}. Cookie: {{request.cookie.cookie1}}." {
		t.Error("Replaced missing tags not match", mock.Response.Body)
	}
}

func TestReplaceExistingBodyTag(t *testing.T) {
	req := definition.Request{}
	req.Body = "{\"persons\":{\"bodies\":[{\"body\":\"is my body\"}, {\"body\":\"is your body\"}]}}"

	res := definition.Response{}
	res.Body = "Request Body {{request.body.persons.bodies.1.body}}"

	mock := definition.Mock{Request: req, Response: res}
	varsProcessor := getProcessor()
	varsProcessor.Eval(&req, &mock)

	if mock.Response.Body != "Request Body \"is your body\"" {
		t.Error("Replaced missing tags not match", mock.Response.Body)
	}
}

func TestReplaceUnExistingBodyTag(t *testing.T) {
	req := definition.Request{}
	req.Body = "{\"persons\":{\"bodies\":[{\"body\":\"is my body\"}, {\"body\":\"is your body\"}]}}"

	res := definition.Response{}
	res.Body = "Request Body {{request.body.whatever}}"

	mock := definition.Mock{Request: req, Response: res}
	varsProcessor := getProcessor()
	varsProcessor.Eval(&req, &mock)

	if mock.Response.Body != "Request Body {{request.body.whatever}}" {
		t.Error("Unexisting tags in body match", mock.Response.Body)
	}
}
