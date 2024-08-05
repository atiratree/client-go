// Code generated by applyconfiguration-gen. DO NOT EDIT.

package internal

import (
	"fmt"
	"sync"

	typed "sigs.k8s.io/structured-merge-diff/v4/typed"
)

func Parser() *typed.Parser {
	parserOnce.Do(func() {
		var err error
		parser, err = typed.NewParser(schemaYAML)
		if err != nil {
			panic(fmt.Sprintf("Failed to parse schema: %v", err))
		}
	})
	return parser
}

var parserOnce sync.Once
var parser *typed.Parser
var schemaYAML = typed.YAMLObject(`types:
- name: com.github.openshift.api.insights.v1alpha1.DataGather
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: com.github.openshift.api.insights.v1alpha1.DataGatherSpec
      default: {}
    - name: status
      type:
        namedType: com.github.openshift.api.insights.v1alpha1.DataGatherStatus
      default: {}
- name: com.github.openshift.api.insights.v1alpha1.DataGatherSpec
  map:
    fields:
    - name: dataPolicy
      type:
        scalar: string
      default: ""
    - name: gatherers
      type:
        list:
          elementType:
            namedType: com.github.openshift.api.insights.v1alpha1.GathererConfig
          elementRelationship: atomic
- name: com.github.openshift.api.insights.v1alpha1.DataGatherStatus
  map:
    fields:
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Condition
          elementRelationship: associative
          keys:
          - type
    - name: dataGatherState
      type:
        scalar: string
    - name: finishTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
    - name: gatherers
      type:
        list:
          elementType:
            namedType: com.github.openshift.api.insights.v1alpha1.GathererStatus
          elementRelationship: associative
          keys:
          - name
    - name: insightsReport
      type:
        namedType: com.github.openshift.api.insights.v1alpha1.InsightsReport
      default: {}
    - name: insightsRequestID
      type:
        scalar: string
    - name: relatedObjects
      type:
        list:
          elementType:
            namedType: com.github.openshift.api.insights.v1alpha1.ObjectReference
          elementRelationship: atomic
    - name: startTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
- name: com.github.openshift.api.insights.v1alpha1.GathererConfig
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
    - name: state
      type:
        scalar: string
      default: ""
- name: com.github.openshift.api.insights.v1alpha1.GathererStatus
  map:
    fields:
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Condition
          elementRelationship: associative
          keys:
          - type
    - name: lastGatherDuration
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Duration
    - name: name
      type:
        scalar: string
      default: ""
- name: com.github.openshift.api.insights.v1alpha1.HealthCheck
  map:
    fields:
    - name: advisorURI
      type:
        scalar: string
      default: ""
    - name: description
      type:
        scalar: string
      default: ""
    - name: state
      type:
        scalar: string
      default: ""
    - name: totalRisk
      type:
        scalar: numeric
      default: 0
- name: com.github.openshift.api.insights.v1alpha1.InsightsReport
  map:
    fields:
    - name: downloadedAt
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
    - name: healthChecks
      type:
        list:
          elementType:
            namedType: com.github.openshift.api.insights.v1alpha1.HealthCheck
          elementRelationship: atomic
    - name: uri
      type:
        scalar: string
- name: com.github.openshift.api.insights.v1alpha1.ObjectReference
  map:
    fields:
    - name: group
      type:
        scalar: string
      default: ""
    - name: name
      type:
        scalar: string
      default: ""
    - name: namespace
      type:
        scalar: string
    - name: resource
      type:
        scalar: string
      default: ""
- name: io.k8s.apimachinery.pkg.apis.meta.v1.Condition
  map:
    fields:
    - name: lastTransitionTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
    - name: message
      type:
        scalar: string
      default: ""
    - name: observedGeneration
      type:
        scalar: numeric
    - name: reason
      type:
        scalar: string
      default: ""
    - name: status
      type:
        scalar: string
      default: ""
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.apimachinery.pkg.apis.meta.v1.Duration
  scalar: string
- name: io.k8s.apimachinery.pkg.apis.meta.v1.FieldsV1
  map:
    elementType:
      scalar: untyped
      list:
        elementType:
          namedType: __untyped_atomic_
        elementRelationship: atomic
      map:
        elementType:
          namedType: __untyped_deduced_
        elementRelationship: separable
- name: io.k8s.apimachinery.pkg.apis.meta.v1.ManagedFieldsEntry
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: fieldsType
      type:
        scalar: string
    - name: fieldsV1
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.FieldsV1
    - name: manager
      type:
        scalar: string
    - name: operation
      type:
        scalar: string
    - name: subresource
      type:
        scalar: string
    - name: time
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
- name: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
  map:
    fields:
    - name: annotations
      type:
        map:
          elementType:
            scalar: string
    - name: creationTimestamp
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
    - name: deletionGracePeriodSeconds
      type:
        scalar: numeric
    - name: deletionTimestamp
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
    - name: finalizers
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
    - name: generateName
      type:
        scalar: string
    - name: generation
      type:
        scalar: numeric
    - name: labels
      type:
        map:
          elementType:
            scalar: string
    - name: managedFields
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ManagedFieldsEntry
          elementRelationship: atomic
    - name: name
      type:
        scalar: string
    - name: namespace
      type:
        scalar: string
    - name: ownerReferences
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.OwnerReference
          elementRelationship: associative
          keys:
          - uid
    - name: resourceVersion
      type:
        scalar: string
    - name: selfLink
      type:
        scalar: string
    - name: uid
      type:
        scalar: string
- name: io.k8s.apimachinery.pkg.apis.meta.v1.OwnerReference
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
      default: ""
    - name: blockOwnerDeletion
      type:
        scalar: boolean
    - name: controller
      type:
        scalar: boolean
    - name: kind
      type:
        scalar: string
      default: ""
    - name: name
      type:
        scalar: string
      default: ""
    - name: uid
      type:
        scalar: string
      default: ""
    elementRelationship: atomic
- name: io.k8s.apimachinery.pkg.apis.meta.v1.Time
  scalar: untyped
- name: __untyped_atomic_
  scalar: untyped
  list:
    elementType:
      namedType: __untyped_atomic_
    elementRelationship: atomic
  map:
    elementType:
      namedType: __untyped_atomic_
    elementRelationship: atomic
- name: __untyped_deduced_
  scalar: untyped
  list:
    elementType:
      namedType: __untyped_atomic_
    elementRelationship: atomic
  map:
    elementType:
      namedType: __untyped_deduced_
    elementRelationship: separable
`)
