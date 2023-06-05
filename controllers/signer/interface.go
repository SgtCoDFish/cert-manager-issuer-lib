/*
Copyright 2023 The cert-manager Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package signer

import (
	"context"
	"crypto/x509"

	cmapi "github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/cert-manager/issuer-lib/api/v1alpha1"
)

type Sign func(ctx context.Context, cr CertificateRequestObject, issuerObject v1alpha1.Issuer) ([]byte, error)
type Check func(ctx context.Context, issuerObject v1alpha1.Issuer) error

type CertificateRequestObject interface {
	metav1.Object

	GetRequest() (template *x509.Certificate, csr []byte, err error)

	GetConditions() []cmapi.CertificateRequestCondition
}
