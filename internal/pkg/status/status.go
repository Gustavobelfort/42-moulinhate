package status

type Exam struct {
	CurrentLevel    int
	RemainingLives  int
	Started         bool
	CurrentExercise string
	Retries         int
}

type Exercise struct {
	ExerciseName     string `json:"exerciseName"`
	ExpectedFile     string `json:"expectedFile"`
	AllowedFunctions string `json:"allowedFunctions"`
	ExerciseNumber   int64  `json:"exerciseNumber"`
	ExerciseLevel    int64  `json:"exerciseLevel"`
}
