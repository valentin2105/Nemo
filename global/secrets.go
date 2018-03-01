package global

import (
	"context"
	"fmt"
	"log"
	"time"

	corev1 "github.com/ericchiang/k8s/apis/core/v1"
)

// Secret - kubectl get secret
type Secret struct {
	Name      string
	Namespace string
	Created   string
}

// SecretList - list of secret
type SecretList []Secret

// ListSecrets - return a list of secret
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
		//Created
		c := secrets.Metadata.GetCreationTimestamp()
		cs := c.GetSeconds()
		csc := time.Unix(cs, 0)
		cscf := csc.Format(DefaultDateFormat)
		// Put in slice
		p := Secret{Name: nc, Namespace: nsc, Created: cscf}
		pl = append(pl, p)
	}
	return pl
}
