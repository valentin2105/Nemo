package global

import (
	"context"
	"fmt"
	"log"

	corev1 "github.com/ericchiang/k8s/apis/core/v1"
)

type Secret struct {
	Name      string
	Namespace string
	Age       string
}

type SecretList []Secret

func ListSecrets() SecretList {
	pl := make(SecretList, 0)
	client, err := LoadClient(Kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	var secrets corev1.SecretList
	if err := client.List(context.Background(), "", &secrets); err != nil {
		log.Fatal(err)
	}
	for _, secrets := range secrets.Items {
		//Name
		n := fmt.Sprintf("%q", *secrets.Metadata.Name)
		nc := TrimQuotes(n)
		// Namespace
		ns := fmt.Sprintf("%q", *secrets.Metadata.Namespace)
		nsc := TrimQuotes(ns)
		//Age
		a := fmt.Sprintf("%q", *secrets.Metadata.CreationTimestamp)
		ac := TrimQuotes(a)
		// Put in slice
		p := Secret{Name: nc, Namespace: nsc, Age: ac}
		pl = append(pl, p)
	}
	return pl
}
