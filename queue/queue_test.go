package queue

import (
  "testing"
)

func TestStart (t *testing.T) {
  if(Start()) {
    t.Log("PASSED: start ran successfully")
  } else {
    t.Error("FAILED: Start went wrong")
  }
}

func TestPullEmpty(t *testing.T){
  var expectedmessage = "EMPTY";
  var message string = Pull()
  if (message == expectedmessage) {
    t.Logf("PASSED: pulled %s from queue", message)
  } else {
    t.Errorf("FAILED: expected to pull %s, but pulled %s", expectedmessage, message)
  }
}

func TestPush(t *testing.T) {
  Push("string1")
  Push("string2")
}

func TestPullFirst(t *testing.T){
  var expectedmessage = "string2";
  var message string = Pull()
  if (message == expectedmessage) {
    t.Logf("PASSED: pulled %s from queue", message)
  } else {
    t.Errorf("FAILED: expected to pull %s, but pulled %s", expectedmessage, message)
  }
}

func TestPullSecond(t *testing.T){
  var expectedmessage = "string1";
  var message string = Pull()
  if (message == expectedmessage) {
    t.Logf("PASSED: pulled %s from queue", message)
  } else {
    t.Errorf("FAILED: expected to pull %s, but pulled %s", expectedmessage, message)
  }
}

func TestPushFULL(t *testing.T) {
  for i := 0; i < 10; i++ {
    Push("fillingQueue")
	}
  if (Push("QueueAlreadyFull")) {
    t.Error("FAILED: we managed to push to an already full queue")
  } else {
    t.Log("PASSED: we failed to add to the queue as its full")
  }
}
