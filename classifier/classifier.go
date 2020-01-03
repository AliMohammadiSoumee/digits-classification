package classifier

import (
	"github.com/alidadar7676/digits-classification/class"
	"github.com/alidadar7676/digits-classification/matrix"
	"github.com/sirupsen/logrus"
)

var classNames = []string{"num0", "num1", "num2", "num3", "num4", "num5", "num6", "num7", "num8", "num9"}

type Classifier struct {
	baseDir string
	classes map[string]class.Class
}

func (c *Classifier) Classify(vec matrix.Vector) string {
	min := 10000000.0
	ans := ""
	for name, class := range c.classes {
		dis := class.Distance(vec)
		//logrus.Info("Classify: ", name, " ", dis)
		if dis < min {
			ans = name
			min = dis
		}
	}
	return ans
}

func NewClassifier(baseDir string) (Classifier, error) {
	classes := make(map[string]class.Class)

	for _, name := range classNames {
		logrus.Info("Classifier: Create class ", name)

		class, err := class.NewClass(baseDir, name)
		if err != nil {
			logrus.Error(err)
		} else {
			classes[name] = class
		}
	}

	return Classifier{
		classes: classes,
		baseDir: baseDir,
	}, nil
}
