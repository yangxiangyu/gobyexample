package cmd

import ("testing")

func Test_DoSometing(t *testing.T) {
	result := DoSomething("haha")
	t.Log(result)
	t.Log("Done")
}

