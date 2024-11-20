package label

import (
	"fmt"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"testing"
)

func TestLabel(t *testing.T) {
	//labelSet := labels.SelectorFromSet(map[string]string{
	//	"app": "nginx",
	//})
	//listopts := metav1.ListOptions{
	//	LabelSelector: labelSet.String(),
	//	FieldSelector: fmt.Sprintf("spec.ports[0].nodePort=%s", "8080"),
	//	Limit:         10,
	//}

	mylabels := labels.Set{
		"app": "aaa",
		"bbb": "bbb",
	}

	sel := labels.NewSelector()
	req, err := labels.NewRequirement("bbb", selection.Equals, []string{"bbb"})
	if err != nil {
		panic(err.Error())
	}
	sel.Add(*req)
	if sel.Matches(mylabels) {
		fmt.Printf("Selector %v matched field set %v\n", sel, mylabels)
	} else {
		panic("Selector should have matched field set")
	}

	// Selector from string expression.
	sel, err = labels.Parse("foo==bar")
	if err != nil {
		panic(err.Error())
	}
	if sel.Matches(mylabels) {
		fmt.Printf("Selector %v matched label set %v\n", sel, mylabels)
	} else {
		panic("Selector should have matched labels")
	}

}
