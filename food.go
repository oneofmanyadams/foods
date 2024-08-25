package foods

type Food struct {
	Name        string
	ServingSize int
	ServingUOM  string
	Calories    int
	Categories  []string
	Nutrients   []Nutrient
}
