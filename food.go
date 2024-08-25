package foods

type Food struct {
	Name        string
	Categories  []string
	ServingSize int
	ServingUOM  string
	Calories    int
	Nutrients   []Nutrient
}
