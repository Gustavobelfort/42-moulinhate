package cli

type ExamStatus struct {
	CurrentLevel   int
	RemainingLives int
	Started        bool
}

type LivePrefixState struct {
	LivePrefix string
	IsEnable   bool
}
