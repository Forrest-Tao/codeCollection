package events

import (
	corev1 "k8s.io/api/core/v1"
)

func filterEvent(events []corev1.Event, eventtype string) []corev1.Event {
	if events == nil || len(events) == 0 {
		return events
	}
	res := make([]corev1.Event, 0)
	for _, event := range events {
		if event.Type == eventtype {
			res = append(res, event)
		}
	}
	return res
}

func isRunningOrSucceeded(pod corev1.Pod) bool {
	switch pod.Status.Phase {
	case corev1.PodSucceeded, corev1.PodRunning:
		return true
	}
	return false
}

func getWarningEvent(events []corev1.Event) []corev1.Event {
	if events == nil || len(events) == 0 {
		return nil
	}
	res := make([]corev1.Event, 0)
	for _, event := range events {
		if event.Type == corev1.EventTypeWarning {
			res = append(res, event)
		}
	}
	return res
}
