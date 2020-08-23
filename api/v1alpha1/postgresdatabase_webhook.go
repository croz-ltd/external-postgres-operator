/*


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

package v1alpha1

import (
	"errors"
	"k8s.io/apimachinery/pkg/runtime"
	"regexp"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var postgresdatabaselog = logf.Log.WithName("postgresdatabase-resource")

func (r *PostgresDatabase) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-database-croz-net-v1alpha1-postgresdatabase,mutating=true,failurePolicy=fail,groups=database.croz.net,resources=postgresdatabases,verbs=create;update,versions=v1alpha1,name=mpostgresdatabase.kb.io

var _ webhook.Defaulter = &PostgresDatabase{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *PostgresDatabase) Default() {
	postgresdatabaselog.Info("default", "name", r.Name)

	if r.Spec.Definition.Encoding == "" {
		r.Spec.Definition.Encoding = UTF8
	}

	//if r.Spec.Definition.ConnectionLimit == nil {
	//	r.Spec.Definition.ConnectionLimit = -1
	//}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// +kubebuilder:webhook:verbs=create;update,path=/validate-database-croz-net-v1alpha1-postgresdatabase,mutating=false,failurePolicy=fail,groups=database.croz.net,resources=postgresdatabases,versions=v1alpha1,name=vpostgresdatabase.kb.io

var _ webhook.Validator = &PostgresDatabase{}

var identifiersRegex = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *PostgresDatabase) ValidateCreate() error {
	postgresdatabaselog.Info("validate create", "name", r.Name)

	if !identifiersRegex.MatchString(r.Spec.Name) {
		return errors.New("Requested database name does not compile by Postgres SQL indetifiers pattern. Required pattern: ^[a-zA-Z_][a-zA-Z0-9_]*$")
	}

	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *PostgresDatabase) ValidateUpdate(old runtime.Object) error {
	postgresdatabaselog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *PostgresDatabase) ValidateDelete() error {
	postgresdatabaselog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
