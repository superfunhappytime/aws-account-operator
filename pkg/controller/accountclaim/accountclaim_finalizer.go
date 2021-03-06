package accountclaim

import (
	"context"

	"github.com/go-logr/logr"
	awsv1alpha1 "github.com/openshift/aws-account-operator/pkg/apis/aws/v1alpha1"
	"github.com/openshift/aws-account-operator/pkg/controller/utils"
)

func (r *ReconcileAccountClaim) addFinalizer(reqLogger logr.Logger, accountClaim *awsv1alpha1.AccountClaim) error {
	reqLogger.Info("Adding Finalizer for the AccountClaim")
	accountClaim.SetFinalizers(append(accountClaim.GetFinalizers(), accountClaimFinalizer))

	// Update CR
	err := r.client.Update(context.TODO(), accountClaim)
	if err != nil {
		reqLogger.Error(err, "Failed to update AccountClaim with finalizer")
		return err
	}
	return nil
}

func (r *ReconcileAccountClaim) removeFinalizer(reqLogger logr.Logger, accountClaim *awsv1alpha1.AccountClaim, finalizerName string) error {
	reqLogger.Info("Removing Finalizer for the AccountClaim")
	accountClaim.SetFinalizers(utils.Remove(accountClaim.GetFinalizers(), finalizerName))

	// Update CR
	err := r.client.Update(context.TODO(), accountClaim)
	if err != nil {
		reqLogger.Error(err, "Failed to remove AccountClaim finalizer")
		return err
	}
	return nil
}
