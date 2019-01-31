package authy

import (
	"io/ioutil"
	"os"
	"testing"
)

func Test_findAuthyId(t *testing.T) {

	a := authyClient{}
	tmpfile, err := ioutil.TempFile("", "authyid")

	if err != nil {
		t.Fatalf("cannot create temp file %v", err)
	}

	defer os.Remove(tmpfile.Name())
	a.Config.File = tmpfile.Name()

	ioutil.WriteFile(tmpfile.Name(), []byte(`
piper 123
hook 456
hook 789
`), os.ModePerm)

	{
		id, err := a.findAuthyId("piper")
		if err != nil {
			t.Fatalf("findId failed %v", err)
		}

		if id != "123" {
			t.Error("find id return wrong value")
		}

	}
	{
		id, err := a.findAuthyId("hook")
		if err != nil {
			t.Fatalf("findId failed %v", err)
		}

		if id != "456" {
			t.Error("find id return wrong value")
		}

	}

}
