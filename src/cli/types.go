package cli

type ExamStatus struct {
	CurrentLevel    int
	RemainingLives  int
	Started         bool
	CurrentExercise string
	Retries         int
}

type LivePrefixState struct {
	LivePrefix string
	IsEnable   bool
}
