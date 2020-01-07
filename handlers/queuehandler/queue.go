package queuehandler

func Enqueue(txt string) bool {
	if txt == "true" {
		return true
	}
	return false
}

func EnqueueTexts(txt1 string, txt2 string) bool {
	if txt1 == txt2 {
		return true
	}
	return false
}

func Dequeue() bool {
	return false
}
