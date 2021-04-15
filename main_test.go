// package
package main

import (
	"context"
	"testing"

	appsv1aplha1 "github.com/mritunjaysharma394/policy-report-prototype/pkg/apis/wgpolicyk8s.io/v1alpha1"
	testclient "github.com/mritunjaysharma394/policy-report-prototype/pkg/generated/clientset/versioned/fake"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestCreatePolicyReport(t *testing.T) {
	policyTests := []struct {
		name         string
		policyreport *appsv1aplha1.PolicyReport
		ns           string
	}{
		{"demo-test-1", &appsv1aplha1.PolicyReport{
			ObjectMeta: metav1.ObjectMeta{
				Name: "demo",
			},
			Summary: appsv1aplha1.PolicyReportSummary{
				Pass: 10,
				Fail: 4,
				Warn: 0,
			},
		}, "default"},
		{"demo-test-2", &appsv1aplha1.PolicyReport{
			ObjectMeta: metav1.ObjectMeta{
				Name: "demo",
			},
			Summary: appsv1aplha1.PolicyReportSummary{
				Pass: 5,
				Fail: 4,
				Warn: 0,
			},
			Results: []*appsv1aplha1.PolicyReportResult{
				{
					Policy:      "test-policy",
					Rule:        "test-rule",
					Category:    "test-category",
					Result:      "pass",
					Scored:      true,
					Description: "test-description",
					Properties: map[string]string{
						"index":           "1",
						"audit":           "",
						"AuditEnv":        "",
						"AuditConfig":     "",
						"type":            "test-type",
						"remediation":     "test-remediation",
						"test_info":       "test",
						"actual_value":    "test-actual-value",
						"IsMultiple":      "true",
						"expected_result": "test-exp-result",
						"reason":          "test-reason",
					},
				},
			},
		}, "default"},
	}

	for _, pr := range policyTests {
		_, err := testclient.NewSimpleClientset().Wgpolicyk8sV1alpha1().PolicyReports(pr.ns).Create(context.TODO(), pr.policyreport, metav1.CreateOptions{})
		if err != nil {
			t.Fatalf("error creating policy report: %v", err)
		}
	}
}
