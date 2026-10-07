package document

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"

	svcapitypes "github.com/aws-controllers-k8s/ssm-controller/apis/v1alpha1"
)

func TestDocumentTagOnlyDelta(t *testing.T) {
	documentName := "ack-document-tag-test"
	tag := func(key, value string) *svcapitypes.Tag {
		return &svcapitypes.Tag{Key: aws.String(key), Value: aws.String(value)}
	}
	desired := &resource{ko: &svcapitypes.Document{
		Spec: svcapitypes.DocumentSpec{
			Name: &documentName,
			Tags: []*svcapitypes.Tag{tag("environment", "production")},
		},
	}}
	latest := &resource{ko: &svcapitypes.Document{
		Spec: svcapitypes.DocumentSpec{
			Name: &documentName,
			Tags: []*svcapitypes.Tag{tag("environment", "development")},
		},
	}}

	delta := newResourceDelta(desired, latest)
	if !delta.DifferentAt("Spec.Tags") {
		t.Fatal("expected tag drift to be detected")
	}
	if delta.DifferentExcept("Spec.Tags") {
		t.Fatal("expected tag drift to be the only delta")
	}
}
