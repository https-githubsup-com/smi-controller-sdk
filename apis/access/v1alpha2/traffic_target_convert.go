package v1alpha2

import (
	"github.com/servicemeshinterface/smi-controller-sdk/apis/access/v1alpha3"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

/*
Our "spoke" versions need to implement the
[`Convertible`](https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/conversion?tab=doc#Convertible)
interface.  Namely, they'll need `ConvertTo` and `ConvertFrom` methods to convert to/from
the hub version.
*/

/*
ConvertTo is expected to modify its argument to contain the converted object.
Most of the conversion is straightforward copying, except for converting our changed field.
*/
// ConvertTo converts this TrafficTarget to the Hub version (v1alpha3).
func (src *TrafficTarget) ConvertTo(dstRaw conversion.Hub) error {
	traffictargetlog.Info("ConvertTo v1alpha3 from v1alpha2")

	dst := dstRaw.(*v1alpha3.TrafficTarget)
	dst.ObjectMeta = src.ObjectMeta

	dst.TypeMeta = src.TypeMeta
	dst.APIVersion = v1alpha3.GroupVersion.Identifier()

	dst.Spec.Destination = v1alpha3.IdentityBindingSubject{
		Kind:      src.Spec.Destination.Kind,
		Name:      src.Spec.Destination.Name,
		Namespace: src.Spec.Destination.Namespace,
	}

	dst.Spec.Sources = []v1alpha3.IdentityBindingSubject{}
	for _, ibs := range src.Spec.Sources {
		s := v1alpha3.IdentityBindingSubject{
			Kind:      ibs.Kind,
			Name:      ibs.Name,
			Namespace: ibs.Namespace,
		}

		dst.Spec.Sources = append(dst.Spec.Sources, s)
	}

	dst.Spec.Rules = []v1alpha3.TrafficTargetRule{}
	for _, ibs := range src.Spec.Rules {
		s := v1alpha3.TrafficTargetRule{
			Kind:    ibs.Kind,
			Name:    ibs.Name,
			Matches: ibs.Matches,
		}

		dst.Spec.Rules = append(dst.Spec.Rules, s)
	}

	return nil
}

/*
ConvertFrom is expected to modify its receiver to contain the converted object.
Most of the conversion is straightforward copying, except for converting our changed field.
*/

// ConvertFrom converts from the Hub version (v1alpha3) to this version.
func (dst *TrafficTarget) ConvertFrom(srcRaw conversion.Hub) error {
	traffictargetlog.Info("ConvertFrom v1alpha3 to v1alpha2")

	src := srcRaw.(*v1alpha3.TrafficTarget)
	dst.ObjectMeta = src.ObjectMeta

	dst.TypeMeta = src.TypeMeta
	dst.APIVersion = GroupVersion.Identifier()

	dst.Spec.Destination = IdentityBindingSubject{
		Kind:      src.Spec.Destination.Kind,
		Name:      src.Spec.Destination.Name,
		Namespace: src.Spec.Destination.Namespace,
	}

	dst.Spec.Sources = []IdentityBindingSubject{}
	for _, ibs := range src.Spec.Sources {
		s := IdentityBindingSubject{
			Kind:      ibs.Kind,
			Name:      ibs.Name,
			Namespace: ibs.Namespace,
		}

		dst.Spec.Sources = append(dst.Spec.Sources, s)
	}

	dst.Spec.Rules = []TrafficTargetRule{}
	for _, ibs := range src.Spec.Rules {
		s := TrafficTargetRule{
			Kind:    ibs.Kind,
			Name:    ibs.Name,
			Matches: ibs.Matches,
		}

		dst.Spec.Rules = append(dst.Spec.Rules, s)
	}

	return nil
}
