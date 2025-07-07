package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	http.HandleFunc("/run", runJobHandler)
	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func runJobHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	domain := r.URL.Query().Get("domain")
	if domain == "" {
		http.Error(w, "Missing domain", http.StatusBadRequest)
		return
	}

	config, err := rest.InClusterConfig()
	if err != nil {
		http.Error(w, "Failed to get cluster config", http.StatusInternalServerError)
		return
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		http.Error(w, "Failed to create clientset", http.StatusInternalServerError)
		return
	}

	jobName := fmt.Sprintf("pentest-%s", randString(6))
	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: jobName,
		},
		Spec: batchv1.JobSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					RestartPolicy: corev1.RestartPolicyNever,
					Containers: []corev1.Container{
						{
							Name:  "pen-test",
							Image: os.Getenv("PENTEST_IMAGE"),
							Args:  []string{domain},
						},
					},
				},
			},
		},
	}

	_, err = clientset.BatchV1().Jobs("default").Create(context.TODO(), job, metav1.CreateOptions{})
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create job: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Job %s created to scan %s\n", jobName, domain)
}

func randString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[randInt(len(letters))]
	}
	return string(b)
}

func randInt(n int) int {
	f, _ := os.Open("/dev/urandom")
	defer f.Close()
	b := make([]byte, 1)
	f.Read(b)
	return int(b[0]) % n
}
