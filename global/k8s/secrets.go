package global

import (
	"context"
	"time"

	"github.com/Sirupsen/logrus"
	corev1 "github.com/ericchiang/k8s/apis/core/v1"
	"github.com/valentin2105/Nemo/global"
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
		logrus.Warn("Error " + err.Error())
	}

	var secrets corev1.SecretList
	if err := client.List(context.Background(), "", &secrets); err != nil {
		logrus.Warn("Error " + err.Error())
	}
	for _, secrets := range secrets.Items {
		//Name
		n := *secrets.Metadata.Name
		nc := global.TrimQuotes(n)
		// Namespace
		ns := *secrets.Metadata.Namespace
		nsc := global.TrimQuotes(ns)
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
